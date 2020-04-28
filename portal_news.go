package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"portal_news/api_handler"
	"portal_news/db"
)

func main(){

	// *** Set Db
	dbConnector, err := getDbConnector()
	if err != nil{
		panic(err)
	}

	err = dbConnector.SetDbInstance()
	if err != nil{
		panic(err)
	}
	defer db.Instance.Close()

	db.Instance.LogMode(true)

	// Set Db ***

	// *** Set Router

	f, _ := os.Create("./gin.log")
	gin.DefaultWriter = io.MultiWriter(f,os.Stdout)

	router := gin.Default()


	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	v1 := router.Group("/")
	{
		v1.GET("/", api_handler.Home)
	}

	// Set Router ***

	// Server Start
	router.Run(":8080")
}

func getDbConnector() (*db.Connector, error){
	buf, err := ioutil.ReadFile(`C:\Users\SONG\Documents\study\go\src\portal_news\db_info.yaml`)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}

	var connector *db.Connector

	err = yaml.Unmarshal(buf, &connector)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}

	return connector, nil
}