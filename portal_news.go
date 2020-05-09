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
	"portal_news/const_val"
	"portal_news/controller"
	"portal_news/db"
	"portal_news/service"
)

func main() {
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
func getDbConnector() (*db.Connector, error) {
	//temporary path for debug mode
	buf, err := ioutil.ReadFile(`C:\Users\SONG\Documents\study\go\src\portal_news\db_info.yaml`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var connector *db.Connector

	err = yaml.Unmarshal(buf, &connector)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return connector, nil
}

// Set router
func setRouter() *gin.Engine {
	router := gin.Default()

	store := cookie.NewStore([]byte(const_val.SessionKey))
	router.Use(sessions.Sessions("mySession", store))

	//temporary path for debug mode
	router.Static("/assets", `C:\Users\SONG\Documents\study\go\src\portal_news\assets`)

	//set render
	router.HTMLRender = createRender()

	//set auth router
	router.GET("/", controller.HomeGet)
	router.GET("/login", controller.LoginGet)
	router.POST("/login-auth", controller.LoginAuthPost)
	router.GET("/google-oauth", controller.GoogleOauthGet)
	router.GET("/google-oauth/callback", controller.GoogleOauthCallbackGet)
	router.GET("/logout", controller.LogoutGet)
	router.POST("/signup", controller.SignUpPost)

	//set news group router
	newsRouter := router.Group("/news")
	{
		newsRouter.GET("/naver", controller.NaverGet)
		newsRouter.GET("/naver/:language", controller.NaverLanguageGet)
		newsRouter.GET("/nate", controller.NateGet)
		newsRouter.GET("/daum", controller.DaumGet)
	}

	//set  mypage group router
	myPageRouter := router.Group("/mypage").Use(service.LoginCheck())
	{
		myPageRouter.GET("/", controller.MyPageGet)
	}

	//set review group router
	reviewRouter := router.Group("/review").Use(service.LoginCheck())
	{
		reviewRouter.GET("/*queryUrl", controller.WriteReviewGET)
		reviewRouter.POST("/*queryUrl", controller.WriteReviewPOST)
		reviewRouter.PUT("/*queryUrl", controller.WriteReviewPUT)
		reviewRouter.DELETE("/*queryUrl", controller.WriteReviewDELETE)
	}

	//set review group router
	searchRouter := router.Group("/search").Use(service.LoginCheck())
	{
		searchRouter.GET("/", controller.SearchGet)
		searchRouter.GET("/news", controller.SearchNewsGet)
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
	myPagePath := rootPath + `mypage.tmpl`
	searchPath := rootPath + `search.tmpl`

	defineRootPath := `C:\Users\SONG\Documents\study\go\src\portal_news\templates\define\`
	defineHeaderPath := defineRootPath + `define_header.tmpl`
	defineNavigationPath := defineRootPath + `define_navigation.tmpl`
	defineLoginPath := defineRootPath + `define_login.tmpl`

	r := multitemplate.NewRenderer()

	r.AddFromFilesFuncs(const_val.TmplFileNews, template.FuncMap{
		"AddHttpsString": service.AddHttpsString,
	}, newsPath, defineHeaderPath, defineNavigationPath)

	r.AddFromFilesFuncs(const_val.TmplFileMypage, template.FuncMap{
		"AddHttpsString": service.AddHttpsString,
	}, myPagePath, defineHeaderPath, defineNavigationPath)

	r.AddFromFilesFuncs(const_val.TmplFileSerch, template.FuncMap{
		"AddHttpsString": service.AddHttpsString,
	}, searchPath, defineHeaderPath, defineNavigationPath)

	r.AddFromFiles(const_val.TmplFileHome, homePath, defineHeaderPath, defineNavigationPath)
	r.AddFromFiles(const_val.TmplFileLogin, loginPath, defineHeaderPath, defineNavigationPath, defineLoginPath)
	r.AddFromFiles(const_val.TmplFileNotLogin, notLoginPath, defineHeaderPath, defineNavigationPath, defineLoginPath)
	r.AddFromFiles(const_val.TmplFileWriteReview, writeReviewPath, defineHeaderPath, defineNavigationPath)

	return r
}
