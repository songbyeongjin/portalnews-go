package service

import (
	"net/url"
	"portal_news/db"
	"portal_news/model"
)

func GetSearchNews(values url.Values) (*[]model.News, string){
	news := &[]model.News{}

	var portalStr []string
	var target string
	var language string
	var content string

	if val, ok := values["check-portal"]; ok{
		if val[0] != "all"{
			for _, r := range val{
				if r != "all"{
					portalStr = append(portalStr, r)
				}
			}
		}
	}

	if val, ok := values["select-target"]; ok{
		if val[0] == "title"{
			target = "title"
		}else{
			target = "content"
		}
	}
	var retLang string
	if val, ok := values["radio-language"]; ok{
		if val[0] == "korean"{
			language = ""
			retLang = "korean"
		}else{
			language = "_ja"
			retLang = "japanese"

		}
	}

	if val, ok := values["text-content"]; ok{
		content = val[0]
		if content == ""{
			return nil,""
		}
	}

	whereTarget := target + language
	if len(portalStr) == 0{
		db.Instance.Where(whereTarget + " LIKE ?" , `%`+ content + `%`).Find(news)
	}else{
		db.Instance.Where("portal IN (?) AND " + whereTarget + " LIKE  ?" , portalStr, `%` + content+ `%`).Find(news)
	}

	if len(*news) == 0{
		return nil,""
	}
	return news, retLang
}