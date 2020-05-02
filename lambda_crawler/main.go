package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
	"os"
	"portal_news/lambda_crawler/db"
	"portal_news/lambda_crawler/model"
	"portal_news/lambda_crawler/service"
	"strconv"
	"sync"
)

func Crawling(ctx context.Context) (string, error) {
	dbConnector := db.Connector{

		Dialect: os.Getenv("db_dialect"),
		Host:  os.Getenv("db_host"),
		Port:  os.Getenv("db_port"),
		Dbname:  os.Getenv("db_name"),
		User:  os.Getenv("db_user"),
		Password:  os.Getenv("db_password"),

		/*
		Dialect: "",
		Host:  "",
		Port:  "",
		Dbname:  "",
		User:  "",
		Password:  "",
		*/
	}

	dbErr := dbConnector.SetDbInstance()
	if dbErr != nil{
		return "db init failed", dbErr
	}

	err :=crawlNews()
	if err != nil{
		return "crawling error occurred", err
	}

	return "crawling succeeded", nil
}


type NewsCrawler interface {
	CrawlNews() []model.RankingNews
	GetNewsUrls(rootUrl string) []string
	GetNews(newsUrls []string) []model.RankingNews
}


func main() {
	lambda.Start(Crawling)
	//Crawling(context.Background())

}

func crawlNews() error {
	//delete exist record rankin news
	var news model.RankingNews
	db.Instance.Delete(&news)

	// *** crawl news

	newsCrawlers := []NewsCrawler{
		service.NateNewsCrawler{},
		service.NaverNewsCrawler{},
	}

	wg := sync.WaitGroup{}
	wg.Add(len(newsCrawlers))

	var err error
	for i, newsCrawler := range newsCrawlers {

		inNewsCrawler := newsCrawler
		go func(index int){
			defer wg.Done()

			news := inNewsCrawler.CrawlNews()

			if len(news) != service.NewsCount{
				err = fmt.Errorf("news item count error roop Counter : " + strconv.Itoa(index))
			}
			service.SaveNews(news)
		}(i)
	}

	wg.Wait()

	return err
}

