package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"portal_news/const_val"
	"portal_news/controller"
	c "portal_news/controller_clean"
	"portal_news/service"
)

func main() {
	newsController := InjectRankingNewsController()
	reviewController := InjectReviewController()
	r := GetRouterGetRouter(newsController, reviewController)
	r.Run(":8080")

}



func GetRouterGetRouter(
	rankingNewsController c.RankingNewsController,
	reviewController c.ReviewController,) *gin.Engine {
	router := gin.Default()

	store := cookie.NewStore([]byte(const_val.SessionKey))
	router.Use(sessions.Sessions("mySession", store))

	//temporary path for debug mode
	router.Static("/assets", `C:\Users\SONG\Documents\study\go\src\portal_news\assets`)

	//set render
	router.HTMLRender = getRender()

	//set auth router
	router.GET("/", controller.HomeGet)
	router.GET("/login", controller.LoginGet)
	router.POST("/login-auth", controller.LoginAuthPost)
	router.GET("/google-oauth", controller.GoogleOauthGet)
	router.GET("/google-oauth/callback", controller.GoogleOauthCallbackGet)
	router.GET("/logout", controller.LogoutGet)
	router.POST("/signup", controller.SignUpPost)

	//set news group router
	newsRouter := router.Group("/rankingNews")
	{
		newsRouter.GET("/naver", rankingNewsController.NaverGet)
		newsRouter.GET("/naver/:language", rankingNewsController.NaverLanguageGet)
		newsRouter.GET("/nate", rankingNewsController.NateGet)
		newsRouter.GET("/daum", rankingNewsController.DaumGet)
	}

	//set  mypage group router
	myPageRouter := router.Group("/mypage").Use(service.LoginCheck())
	{
		myPageRouter.GET("/", controller.MyPageGet)
	}

	//set review group router
	reviewRouter := router.Group("/review").Use(service.LoginCheck())
	{
		reviewRouter.GET("/*queryUrl", reviewController.WriteReviewGET)
		reviewRouter.POST("/*queryUrl", reviewController.WriteReviewPOST)
		reviewRouter.PUT("/*queryUrl", reviewController.UpdateReviewPUT)
		reviewRouter.DELETE("/*queryUrl", reviewController.DeleteReviewDELETE)
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
func getRender() multitemplate.Renderer {
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