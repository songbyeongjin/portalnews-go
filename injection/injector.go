package injection

import (
	"portal_news/controller_clean"
	"portal_news/domain_clean/repository_interface"
	"portal_news/infra_clean/db"
	"portal_news/infra_clean/db/repository_impl"
	"portal_news/service_clean"
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

func injectRankingNewsService() service_clean.RankingNewsService {
	RankingNewsRepository := injectRankingNewsRepository()
	return service_clean.NewRankingNewsService(RankingNewsRepository)
}

func InjectRankingNewsController() controller_clean.RankingNewsController {
	return controller_clean.NewRankingNewsController(injectRankingNewsService())
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

func injectReviewService() service_clean.ReviewService {
	ReviewRepository := injectReviewRepository()
	newsRepository := injectNewsRepository()
	return service_clean.NewReviewService(ReviewRepository, newsRepository)
}

func InjectReviewController() controller_clean.ReviewController {
	return controller_clean.NewReviewController(injectReviewService())
}

//  review injection ***


// *** my page injection

func injectMyPageService() service_clean.MyPageService {
	ReviewRepository := injectReviewRepository()
	newsRepository := injectNewsRepository()
	return service_clean.NewMyPageService(ReviewRepository, newsRepository)
}

func InjectMyPageController() controller_clean.MyPageController {
	return controller_clean.NewMyPageController(injectMyPageService())
}

//  my page injection ***


// *** search injection

func injectSearchService() service_clean.SearchService {
	newsRepository := injectNewsRepository()
	return service_clean.NewSearchService(newsRepository)
}

func InjectSearchController() controller_clean.SearchController {
	return controller_clean.NewSearchController(injectSearchService())
}

//  search injection ***

// *** user injection

func injectUserRepository() repository_interface.UserRepository {
	dbHandler := injectDB()
	return repository_impl.NewUserRepository(dbHandler)
}

// user injection  ***

// *** login injection

func injectLoginService() service_clean.LoginService {
	userRepository := injectUserRepository()
	return service_clean.NewLoginService(userRepository)
}

func InjectLoginController() controller_clean.LoginController {
	return controller_clean.NewLoginController(injectLoginService())
}

//  login injection ***

// *** logout injection

func injectLogoutService() service_clean.LogoutService {
	return service_clean.NewLogoutService()
}

func InjectLogoutController() controller_clean.LogoutController {
	return controller_clean.NewLogoutController(injectLogoutService())
}

//  logout injection ***


// *** user injection

func injectUserService() service_clean.UserService {
	userRepository := injectUserRepository()
	return service_clean.NewUserService(userRepository)
}

func InjectUserController() controller_clean.UserController {
	return controller_clean.NewUserController(injectUserService())
}

//  user injection ***


// *** main injection

func InjectMainController() controller_clean.MainController {
	return controller_clean.NewMainController()
}

//  main injection ***