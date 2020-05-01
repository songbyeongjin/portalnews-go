package service

import (
	"portal_news/lambda_crawler/db"
	"portal_news/lambda_crawler/model"
	"portal_news/service"
	"unicode/utf8"
)

const(
	NewsCount              = 10
	layoutYYYYMMDD         = "2006-01-02"
	httpsUrl               = `https://`
	cssSelectorOrCondition = `, `
	setFieldCount          = 4 //Title,Content,Press,Date
)

func SaveNews(news []model.RankingNews){
	for _, r := range news {
		//save news to ranking news
		db.Instance.Create(&r)


		//save news record only unique news
		var news = new(model.News)
		existFlag :=  db.Instance.Where("url = ?", r.Url).First(news).RecordNotFound()
		if existFlag == true{
			news = &model.News{
				Title: r.Title,
				Content: r.Content,
				Press: r.Press,
				Date: r.Date,
				Url: r.Url,
				Portal: r.Portal,
				CreatedAt: r.CreatedAt,
				UpdatedAt: r.UpdatedAt,
			}

			db.Instance.Create(news)
		}
	}
}


func AddHttpsString(url string) string {
	return service.HttpsUrl + url
}

func trimmingContent(original string, limit int) string{
	ret := ""
	//Minimize String By limit Rune
	for _, rune := range original{
		ret += string(rune)

		if utf8.RuneCountInString(ret) >= limit{
			break
		}
	}

	ret += "..."

	return ret
}