package gin_route

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"portal_news/common"
	"portal_news/controller"
	"portal_news/infra/db"
	"portal_news/infra/db/repository_impl"
	"portal_news/infra/gin_route/middle_ware"
	"portal_news/service/impl"
)

func SetRouter() *gin.Engine{


	psDb := db.NewDbHandler()
	rankingNewsR := repository_impl.NewRankingNewsRepository(psDb)
	reviewR := repository_impl.NewReviewRepository(psDb)
	userR := repository_impl.NewUserRepository(psDb)
	newsR := repository_impl.NewNewsRepository(psDb)

	rankingNewsS := impl.NewRankingNewsService(rankingNewsR)
	reviewS := impl.NewReviewService(reviewR, newsR)
	searchS := impl.NewSearchService(newsR)
	userS := impl.NewUserService(userR)
	logoutS := impl.NewLogoutService()
	loginS := impl.NewLoginService(userR)
	myPageS := impl.NewMyPageService(reviewR, newsR)

	rankingNewsC := controller.NewRankingNewsController(rankingNewsS)
	reviewC := controller.NewReviewController(reviewS)
	myPageC := controller.NewMyPageController(myPageS)
	searchC := controller.NewSearchController(searchS)
	userC := controller.NewUserController(userS)
	loginC := controller.NewLoginController(loginS)
	logoutC := controller.NewLogoutController(logoutS)
	mainC := controller.NewMainController()

	r := GetRouter(
		mainC,
		rankingNewsC,
		reviewC,
		myPageC,
		searchC,
		loginC,
		logoutC,
		userC)

	return r
}


func GetRouter(
	mainController controller.MainController,
	rankingNewsController controller.RankingNewsController,
	reviewController controller.ReviewController,
	myPageController controller.MyPageController,
	searchController controller.SearchController,
	loginController controller.LoginController,
	logoutController controller.LogoutController,
	userController controller.UserController) *gin.Engine {

	router := gin.Default()

	store := cookie.NewStore([]byte(common.SessionKey))
	router.Use(sessions.Sessions("mySession", store))

	//temporary path for debug mode
	router.Static("/assets", `..\..\assets`)

	//set render
	router.HTMLRender = getRender()

	//set home router
	router.GET("/", mainController.HomeGet)

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
	rootPath := `./templates/`
	homePath := rootPath + `home.tmpl`
	newsPath := rootPath + `news.tmpl`
	loginPath := rootPath + `login.tmpl`
	notLoginPath := rootPath + `not_login.tmpl`
	writeReviewPath := rootPath + `write_review.tmpl`
	myPagePath := rootPath + `mypage.tmpl`
	searchPath := rootPath + `search.tmpl`

	defineRootPath := `./templates/define/`
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