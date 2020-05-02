package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"portal_news/lambda_crawler/db"
	"portal_news/lambda_crawler/model"
	"portal_news/service"
	"strconv"
	"strings"
	"unicode/utf8"
)

const(
	NewsCount              = 10
	layoutYYYYMMDD         = "2006-01-02"
	httpsUrl               = `https://`
	cssSelectorOrCondition = `, `
	setFieldCount          = 4 //Title,Content,Press,Date
)

func SaveNews(rankingNews []*model.RankingNews){
	for _, r := range rankingNews {
		//save news to ranking news
		db.Instance.Create(r)


		//save news record only unique news
		var news = new(model.News)
		existFlag :=  db.Instance.Where("url = ?", r.Url).First(news).RecordNotFound()
		if existFlag == true{
			news = &model.News{
				Title: r.Title,
				TitleJapanese: r.TitleJapanese,
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
func setJpTitle(newsArray []*model.RankingNews) {
	for _, news := range newsArray{
		title := news.Title

		//translate title ko to ja
		translatedTitle,err := translate("ko","ja",title)

		if err == nil{
			news.TitleJapanese = translatedTitle
		}else{
			news.TitleJapanese = title
		}
	}
}


func translate(sourceLang, targetLang, targetStr string) (string,error){
	papagoUrl := `https://openapi.naver.com/v1/papago/n2mt`

	data := url.Values{}
	data.Set("source", sourceLang)
	data.Set("target", targetLang)
	data.Set("text", targetStr)

	u, _ := url.ParseRequestURI(papagoUrl)
	urlStr := fmt.Sprintf("%v", u)
	u.RawQuery = data.Encode()

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	r.Header.Add("X-Naver-Client-Id", os.Getenv("naver_papago_id"))
	r.Header.Add("X-Naver-Client-Secret", os.Getenv("naver_papago_pass"))
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, _ := client.Do(r)

	if resp.StatusCode != 200{
		return"", fmt.Errorf("papago respone error code : " + strconv.Itoa(resp.StatusCode))
	}

	papagoRep := &model.PapagoRep{}
	body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(body, papagoRep)
	if err != nil{
		return "", err
	}
	return papagoRep.Message.Result.TranslatedText, nil
}

func AddHttpsString(url string) string {
	if strings.Index(url, "v.media.daum.net/") == -1{
		return service.HttpsUrl + url
	}else{
		return service.HttpUrl + url
	}
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