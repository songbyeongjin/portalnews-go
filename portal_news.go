package main

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
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

	setDb()
	defer db.Instance.Close()

	// Set Db ***

	// *** Set Router

	f, _ := os.Create("./server.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := setRouter()

	// Set Router ***

	// Server Start
	router.Run(":8080")
}

func setDb() {
	dbConnector, err := getDbConnector()
	if err != nil {
		panic(err)
	}

	err = dbConnector.SetDbInstance()
	if err != nil {
		panic(err)
	}

	db.Instance.LogMode(true)
}

//Set Db information from yaml
func getDbConnector() (*db.Connector, error){
	//temporary path for debug mode
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


// Set router
func setRouter() *gin.Engine{
	router := gin.Default()

	//temporary path for debug mode
	router.Static("/assets", `C:\Users\SONG\Documents\study\go\src\portal_news\assets`)

	//temporary path for debug mode
	router.LoadHTMLGlob(`C:\Users\SONG\Documents\study\go\src\portal_news\templates\*`)
	router.HTMLRender = createRender()

	//set default router
	router.GET("/", api_handler.Home)

	//set news group router
	newsRouter := router.Group("/news")
	{
		newsRouter.GET("/", api_handler.Portal)
		newsRouter.GET("/naver", api_handler.Naver)
		newsRouter.GET("/nate", api_handler.Nate)
		newsRouter.GET("/daum", api_handler.Daum)
	}

	return router
}

//Create render for using template block
func createRender() multitemplate.Renderer {
	rootPath := `C:\Users\SONG\Documents\study\go\src\portal_news\templates\`
	headerPath := rootPath + `header.tmpl`

	r := multitemplate.NewRenderer()
	r.AddFromFiles("home", rootPath + `home.tmpl`, headerPath)
	r.AddFromFiles("news", rootPath + `news.tmpl`, headerPath)

	return r
}