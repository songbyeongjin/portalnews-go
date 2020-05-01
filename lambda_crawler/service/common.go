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

		//To Do save news to news
		//To Do filtering exist news
		//save only new news
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