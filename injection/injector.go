package injection

import (
	"portal_news/controller"
	"portal_news/domain/repository_interface"
	"portal_news/infra/db"
	"portal_news/infra/db/repository_impl"
	"portal_news/service"
)

var dbInstance *db.Handler

func injectDB() db.Handler {
	if dbInstance == nil{
		dbInstance = db.NewDbHandler()
	}

	return *dbInstance
}

// *** repository injection

func InjectRankingNewsRepository() repository_interface.RankingNewsRepository {
	dbHandler := injectDB()
	return repository_impl.NewRankingNewsRepository(dbHandler)
}

func InjectNewsRepository() repository_interface.NewsRepository {
	dbHandler := injectDB()
	return repository_impl.NewNewsRepository(dbHandler)
}

func InjectReviewRepository() repository_interface.ReviewRepository {
	dbHandler := injectDB()
	return repository_impl.NewReviewRepository(dbHandler)
}

func InjectUserRepository() repository_interface.UserRepository {
	dbHandler := injectDB()
	return repository_impl.NewUserRepository(dbHandler)
}

// repository injection ***



// *** service injection

func InjectRankingNewsService(rr repository_interface.RankingNewsRepository) service.RankingNewsService {
	RankingNewsRepository := rr
	return service.NewRankingNewsService(RankingNewsRepository)
}

func InjectReviewService(rr repository_interface.ReviewRepository, nr repository_interface.NewsRepository) service.ReviewService {
	ReviewRepository := rr
	newsRepository := nr
	return service.NewReviewService(ReviewRepository, newsRepository)
}

func InjectMyPageService(rr repository_interface.ReviewRepository, nr repository_interface.NewsRepository) service.MyPageService {
	ReviewRepository := rr
	newsRepository := nr
	return service.NewMyPageService(ReviewRepository, newsRepository)
}

func InjectLoginService(ur repository_interface.UserRepository) service.LoginService {
	userRepository := ur
	return service.NewLoginService(userRepository)
}

func InjectLogoutService() service.LogoutService {
	return service.NewLogoutService()
}

func InjectUserService(ur repository_interface.UserRepository) service.UserService {
	userRepository := ur
	return service.NewUserService(userRepository)
}

func InjectSearchService(nr repository_interface.NewsRepository) service.SearchService {
	newsRepository := nr
	return service.NewSearchService(newsRepository)
}

//  service injection ***



// *** controller injection

func InjectMainController() controller.MainController {
	return controller.NewMainController()
}

func InjectRankingNewsController(rs service.RankingNewsService) controller.RankingNewsController {
	return controller.NewRankingNewsController(rs)
}
func InjectReviewController(rs service.ReviewService) controller.ReviewController {
	return controller.NewReviewController(rs)
}

func InjectMyPageController(ms service.MyPageService) controller.MyPageController {
	return controller.NewMyPageController(ms)
}

func InjectSearchController(ss service.SearchService) controller.SearchController {
	return controller.NewSearchController(ss)
}

func InjectLoginController(ls service.LoginService) controller.LoginController {
	return controller.NewLoginController(ls)
}

func InjectLogoutController(ls service.LogoutService) controller.LogoutController {
	return controller.NewLogoutController(ls)
}

func InjectUserController(us service.UserService) controller.UserController {
	return controller.NewUserController(us)
}

//  controller injection ***