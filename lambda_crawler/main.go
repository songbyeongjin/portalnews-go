package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
	"os"
	"portal_news/lambda_crawler/db"
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

	wg := sync.WaitGroup{}
	//crawl nate news
	var err error
	var str string

	wg.Add(1)
	go func(){
		err, str = crawlNateNewsAndSave()
		wg.Done()
	}()

	//crawl naver news
	wg.Add(1)
	go func(){
		err, str = crawlNaverNewsAndSave()
		wg.Done()
	}()

	//crawl daum news

	wg.Wait()
	return "crawling succeeded", nil
}

func main() {
	lambda.Start(Crawling)
}

//crawl nate news
func crawlNateNewsAndSave() (error, string){
	//get nate news
	nateNews := service.CrawlNateNews()

	if len(nateNews) != service.NewsCount{
		return fmt.Errorf("nate news item count error"), "nate news crawling failed"
	}

	service.SaveNews(nateNews)

	return nil, "nate news crawling succeed"
}

//crawl nate news
func crawlNaverNewsAndSave() (error, string){
	//get nate news
	naverNews := service.CrawlNaverNews()

	if len(naverNews) != service.NewsCount{
		return fmt.Errorf("nate news item count error"), "nate news crawling failed"
	}

	service.SaveNews(naverNews)

	return nil, "nate news crawling succeed"
}
