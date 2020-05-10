package repository_impl

import (
	"portal_news/domain/model"
	"portal_news/domain/repository_interface"
	"portal_news/infra/db"
)

type ReviewRepository struct {
	dbHandler db.DbHandler
}

func NewReviewRepository(dbHandler db.DbHandler) repository_interface.ReviewRepository {
	reviewRepository := ReviewRepository{dbHandler}
	return &reviewRepository
}


func (r *ReviewRepository) FindFirstByNewsUrlAndUserId(newsUrl, userID string)  *model.Review {
	review := &model.Review{}
	r.dbHandler.Conn.Where("news_url = ? AND user_id = ?", newsUrl, userID).First(review)

	if review.ID == ""{
		return nil
	}

	return review
}


func (r *ReviewRepository) FindByUserIdOrderByDateDESC(userID string)  *[]model.Review {
	reviews := &[]model.Review{}
	r.dbHandler.Conn.Order("date DESC", true).Find(reviews, "user_id=?", userID)

	return reviews
}

func (r *ReviewRepository) Create(review *model.Review){
	r.dbHandler.Conn.Create(review)
}

func (r *ReviewRepository) Update(review *model.Review, field map[string]interface{}){
	r.dbHandler.Conn.Model(review).Updates(field)
}

func (r *ReviewRepository) Delete(review *model.Review){
	r.dbHandler.Conn.Delete(review)
}