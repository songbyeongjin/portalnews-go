package main

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"portal_news/api_handler"
	"portal_news/db"
	"portal_news/service"
)

func main(){
	// *** Set Db

	setDb()
	defer db.Instance.Close()

	// Set Db ***

	// *** Set Oauth

	service.SetOauthGoogleConfig()

	// Set Oauth ***

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

	store := cookie.NewStore([]byte("secret"))//ToDo encryption
	router.Use(sessions.Sessions("mysession", store))

	//temporary path for debug mode
	router.Static("/assets", `C:\Users\SONG\Documents\study\go\src\portal_news\assets`)

	//set render
	router.HTMLRender = createRender()

	//set auth router
	router.GET("/", api_handler.HomeGet)
	router.GET("/login", api_handler.LoginGet)
	router.POST("/login-auth", api_handler.LoginAuthPost)
	router.GET("/google-oauth", api_handler.GoogleOauthGet)
	router.GET("/google-oauth/callback", api_handler.GoogleOauthCallbackGet)
	router.GET("/logout", api_handler.LogoutGet)
	router.POST("/signup", api_handler.SignUpPost)

	//set news group router
	newsRouter := router.Group("/news")
	{
		newsRouter.GET("/", api_handler.PortalGet)
		newsRouter.GET("/naver", api_handler.Naverget)
		newsRouter.GET("/nate", api_handler.NateGet)
		newsRouter.GET("/daum", api_handler.Daumget)
	}


	//set  mypage group router
	myPageRouter := router.Group("/mypage").Use(service.LoginCheck())
	{
		myPageRouter.GET("/", api_handler.MyPageGet)
	}

	//set review group router
	reviewRouter := router.Group("/review").Use(service.LoginCheck())
	{
		reviewRouter.GET("/*queryUrl", api_handler.WriteReviewGET)
		reviewRouter.POST("/*queryUrl", api_handler.WriteReviewPOST)
	}

	return router
}

//Create render for using template block
func createRender() multitemplate.Renderer {
	rootPath := `C:\Users\SONG\Documents\study\go\src\portal_news\templates\`
	homePath := rootPath + `home.tmpl`
	newsPath := rootPath + `news.tmpl`
	loginPath := rootPath + `login.tmpl`
	notLoginPath := rootPath + `not_login.tmpl`
	writeReviewPath := rootPath + `write_review.tmpl`
	myPageath := rootPath + `mypage.tmpl`


	defineRootPath := `C:\Users\SONG\Documents\study\go\src\portal_news\templates\define\`
	defineHeaderPath := defineRootPath + `define_header.tmpl`
	defineNavigationPath := defineRootPath + `define_navigation.tmpl`
	defineLoginPath := defineRootPath + `define_login.tmpl`


	r := multitemplate.NewRenderer()
	r.AddFromFilesFuncs("news", template.FuncMap{
		"AddHttpsString": service.AddHttpsString,
	},newsPath, defineHeaderPath, defineNavigationPath)

	r.AddFromFiles("home", homePath, defineHeaderPath, defineNavigationPath)
	r.AddFromFiles("login", loginPath, defineHeaderPath,defineNavigationPath, defineLoginPath)
	r.AddFromFiles("notLogin", notLoginPath, defineHeaderPath,defineNavigationPath, defineLoginPath)
	r.AddFromFiles("writeReview", writeReviewPath, defineHeaderPath,defineNavigationPath)
	r.AddFromFiles("myPage", myPageath, defineHeaderPath,defineNavigationPath)
	//r.AddFromFiles("news", rootPath + `news.tmpl`, headerPath)

	return r
}