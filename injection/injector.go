package injection

import (
	"portal_news/controller"
	"portal_news/domain/repository_interface"
	"portal_news/infra/db"
	"portal_news/infra/db/repository_impl"
	"portal_news/service"
)

var dbInstance *db.DbHandler

func injectDB() db.DbHandler {
	if dbInstance == nil{
		dbInstance = db.NewDbHandler()
	}

	return *dbInstance
}

// *** ranking news injection

func injectRankingNewsRepository() repository_interface.RankingNewsRepository {
	dbHandler := injectDB()
	return repository_impl.NewRankingNewsRepository(dbHandler)
}

func injectRankingNewsService() service.RankingNewsService {
	RankingNewsRepository := injectRankingNewsRepository()
	return service.NewRankingNewsService(RankingNewsRepository)
}

func InjectRankingNewsController() controller.RankingNewsController {
	return controller.NewRankingNewsController(injectRankingNewsService())
}


// ranking news injection ***


// *** news injection

func injectNewsRepository() repository_interface.NewsRepository {
	dbHandler := injectDB()
	return repository_impl.NewNewsRepository(dbHandler)
}

//  news injection ***


// *** review injection

func injectReviewRepository() repository_interface.ReviewRepository {
	dbHandler := injectDB()
	return repository_impl.NewReviewRepository(dbHandler)
}

func injectReviewService() service.ReviewService {
	ReviewRepository := injectReviewRepository()
	newsRepository := injectNewsRepository()
	return service.NewReviewService(ReviewRepository, newsRepository)
}

func InjectReviewController() controller.ReviewController {
	return controller.NewReviewController(injectReviewService())
}

//  review injection ***


// *** my page injection

func injectMyPageService() service.MyPageService {
	ReviewRepository := injectReviewRepository()
	newsRepository := injectNewsRepository()
	return service.NewMyPageService(ReviewRepository, newsRepository)
}

func InjectMyPageController() controller.MyPageController {
	return controller.NewMyPageController(injectMyPageService())
}

//  my page injection ***


// *** search injection

func injectSearchService() service.SearchService {
	newsRepository := injectNewsRepository()
	return service.NewSearchService(newsRepository)
}

func InjectSearchController() controller.SearchController {
	return controller.NewSearchController(injectSearchService())
}

//  search injection ***

// *** user injection

func injectUserRepository() repository_interface.UserRepository {
	dbHandler := injectDB()
	return repository_impl.NewUserRepository(dbHandler)
}

// user injection  ***

// *** login injection

func injectLoginService() service.LoginService {
	userRepository := injectUserRepository()
	return service.NewLoginService(userRepository)
}

func InjectLoginController() controller.LoginController {
	return controller.NewLoginController(injectLoginService())
}

//  login injection ***

// *** logout injection

func injectLogoutService() service.LogoutService {
	return service.NewLogoutService()
}

func InjectLogoutController() controller.LogoutController {
	return controller.NewLogoutController(injectLogoutService())
}

//  logout injection ***


// *** user injection

func injectUserService() service.UserService {
	userRepository := injectUserRepository()
	return service.NewUserService(userRepository)
}

func InjectUserController() controller.UserController {
	return controller.NewUserController(injectUserService())
}

//  user injection ***


// *** main injection

func InjectMainController() controller.MainController {
	return controller.NewMainController()
}

//  main injection ***