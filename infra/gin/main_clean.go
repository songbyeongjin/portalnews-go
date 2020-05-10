package gin_infra

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"portal_news/common"
	c "portal_news/controller"
	"portal_news/infra/gin/middle_ware"
	"portal_news/injection"
)

func SetRouter() *gin.Engine{
	mainController := injection.InjectMainController()
	newsController := injection.InjectRankingNewsController()
	reviewController := injection.InjectReviewController()
	myPageController := injection.InjectMyPageController()
	searchController := injection.InjectSearchController()
	loginController := injection.InjectLoginController()
	logoutController := injection.InjectLogoutController()
	userController := injection.InjectUserController()

	r := GetRouter(
		mainController,
		newsController,
		reviewController,
		myPageController,
		searchController,
		loginController,
		logoutController,
		userController)

	return r
}


func GetRouter(
	maincontroller c.MainController,
	rankingNewsController c.RankingNewsController,
	reviewController c.ReviewController,
	myPageController c.MyPageController,
	searchController c.SearchController,
	loginController c.LoginController,
	logoutController c.LogoutController,
	userController c.UserController) *gin.Engine {

	router := gin.Default()

	store := cookie.NewStore([]byte(common.SessionKey))
	router.Use(sessions.Sessions("mySession", store))

	//temporary path for debug mode
	router.Static("/assets", `C:\Users\SONG\Documents\study\go\src\portal_news\assets`)

	//set render
	router.HTMLRender = getRender()

	//set home router
	router.GET("/", maincontroller.HomeGet)

	//set login group router
	loginRouter := router.Group("/login")
	{
		loginRouter.GET("/", loginController.LoginGet)
		loginRouter.POST("", loginController.LoginPost)
		loginRouter.GET("/google-oauth", loginController.GoogleOauthGet)
		loginRouter.GET("/google-oauth/callback", loginController.GoogleOauthCallbackGet)
	}

	logoutRouter := router.Group("/logout")
	{
		logoutRouter.GET("/", logoutController.LogoutGet)
	}

	userRouter := router.Group("/user")
	{
		userRouter.POST("/", userController.SignUpPost)
	}


	//set news group router
	newsRouter := router.Group("/ranking-news")
	{
		newsRouter.GET("/naver", rankingNewsController.NaverGet)
		newsRouter.GET("/naver/:language", rankingNewsController.NaverLanguageGet)
		newsRouter.GET("/nate", rankingNewsController.NateGet)
		newsRouter.GET("/daum", rankingNewsController.DaumGet)
	}

	//set  mypage group router
	myPageRouter := router.Group("/mypage").Use(middleWare.LoginCheck())
	{
		myPageRouter.GET("/", myPageController.MyPageGet)
	}

	//set review group router
	reviewRouter := router.Group("/review").Use(middleWare.LoginCheck())
	{
		reviewRouter.GET("/*queryUrl", reviewController.WriteReviewGET)
		reviewRouter.POST("/*queryUrl", reviewController.WriteReviewPOST)
		reviewRouter.PUT("/*queryUrl", reviewController.UpdateReviewPUT)
		reviewRouter.DELETE("/*queryUrl", reviewController.DeleteReviewDELETE)
	}

	//set review group router
	searchRouter := router.Group("/search")
	{
		searchRouter.GET("/", searchController.SearchGet)
		searchRouter.GET("/news", searchController.SearchNewsGet)
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

	r.AddFromFilesFuncs(common.TmplFileNews, template.FuncMap{
		"AddHttpsString": common.AddHttpsString,
	}, newsPath, defineHeaderPath, defineNavigationPath)

	r.AddFromFilesFuncs(common.TmplFileMypage, template.FuncMap{
		"AddHttpsString": common.AddHttpsString,
	}, myPagePath, defineHeaderPath, defineNavigationPath)

	r.AddFromFilesFuncs(common.TmplFileSerch, template.FuncMap{
		"AddHttpsString": common.AddHttpsString,
	}, searchPath, defineHeaderPath, defineNavigationPath)

	r.AddFromFiles(common.TmplFileHome, homePath, defineHeaderPath, defineNavigationPath)
	r.AddFromFiles(common.TmplFileLogin, loginPath, defineHeaderPath, defineNavigationPath, defineLoginPath)
	r.AddFromFiles(common.TmplFileNotLogin, notLoginPath, defineHeaderPath, defineNavigationPath, defineLoginPath)
	r.AddFromFiles(common.TmplFileWriteReview, writeReviewPath, defineHeaderPath, defineNavigationPath)

	return r
}