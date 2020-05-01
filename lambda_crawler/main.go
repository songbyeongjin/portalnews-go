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

	str, err :=crawlNews()
	if err != nil{
		return str, err
	}

	return "crawling succeeded", nil
}

func main() {
	lambda.Start(Crawling)
	//Crawling(context.Background())
}

func crawlNews()(string, error){
	//delete exist record rankin news
	var news model.RankingNews
	db.Instance.Delete(&news)

	// *** crawl news
	wg := sync.WaitGroup{}

	//crawl nate news
	var err error
	var str string

	wg.Add(1)
	go func(){
		str, err = crawlNateNewsAndSave()
		wg.Done()
	}()

	//crawl naver news
	wg.Add(1)
	go func(){
		str, err = crawlNaverNewsAndSave()
		wg.Done()
	}()

	//crawl daum news

	wg.Wait()
	// crawl news ***

	return str, err
}
//crawl nate news
func crawlNateNewsAndSave() (string, error){
	//get nate news
	nateNews := service.CrawlNateNews()

	if len(nateNews) != service.NewsCount{
		return "nate news crawling failed", fmt.Errorf("nate news item count error")
	}

	service.SaveNews(nateNews)

	return "nate news crawling succeed", nil
}

//crawl nate news
func crawlNaverNewsAndSave() (string, error){
	//get nate news
	naverNews := service.CrawlNaverNews()

	if len(naverNews) != service.NewsCount{
		return "nate news crawling failed", fmt.Errorf("nate news item count error")
	}

	service.SaveNews(naverNews)

	return "nate news crawling succeed", nil
}
