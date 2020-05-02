package service

import (
	"fmt"
	"github.com/gocolly/colly"
	"portal_news/model"
	"regexp"
	"sync"
	"time"
)

const(
	newsCount = 20
	HttpsUrl                = `https://`
	HttpUrl                = `http://`
	nateNewsRootUrl         = HttpsUrl +`news.nate.com/rank/interest?sc=all&p=day&date=20999999`
	cssSelectorFirstToFifth = ".mlt01 a"
	cssSelectorSixthTo      = ".mduSubject a"
	cssSelectorTitle        = ".articleSubecjt"
	cssSelectorContent      = "#realArtcContents"
	cssSelectorPress        = ".articleInfo .medium"
	cssSelectorDate         = ".firstDate em"
	cssSelectorOrCondition  = `, `
	setFieldCount           = 4 //Title,Content,Press,Date
	layoutYYYYMMDD          = "2006-01-02"
)

func CrawlNaverNews(){
	fmt.Println("schedule call" + time.Now().String())
	nateNewsUrls := getNateNewsUrls(nateNewsRootUrl)
	nateNews := getNateNews(nateNewsUrls)

	for _,news := range nateNews{
		fmt.Println(news)
	}
}

//get Nate News url From nate root url
func getNateNewsUrls(nateRootUrl string) []string{
	urls := make([]string,0, newsCount)
	c := colly.NewCollector()
	var wg sync.WaitGroup
	wg.Add(newsCount)

	// Find and visit all links
	c.OnHTML(cssSelectorFirstToFifth + cssSelectorOrCondition + cssSelectorSixthTo, func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		if len(urls) < newsCount {
			title := e.Attr("href")[2:]//delete string("//") in title
			urls = append(urls, title)
			wg.Done()
		}
	})

	c.Visit(nateRootUrl)
	wg.Wait()

	return urls
}

//get NateNews Object from nate urls
func getNateNews(nateNewsUrls []string) []model.News {
	nateNews := make([]model.News, newsCount, newsCount)

	cSlice := make([]*colly.Collector, newsCount, newsCount)
	var wg sync.WaitGroup

	//Set callback
	for i := 0; i < newsCount; i++ {
		inIndex := i
		cSlice[i] = colly.NewCollector()
		cSlice[i].OnHTML(cssSelectorTitle, func(e *colly.HTMLElement) {
			nateNews[inIndex].Title = e.Text
			wg.Done()
		})
		cSlice[i].OnHTML(cssSelectorContent, func(e *colly.HTMLElement) {
			space := regexp.MustCompile(`\s+`)
			str := space.ReplaceAllString(e.Text[:200]+"...", " ")
			nateNews[inIndex].Content = str

			wg.Done()
		})
		cSlice[i].OnHTML(cssSelectorPress, func(e *colly.HTMLElement) {
			nateNews[inIndex].Press = e.Text
			wg.Done()
		})
		cSlice[i].OnHTML(cssSelectorDate, func(e *colly.HTMLElement) {

			nateNews[inIndex].Date, _ = time.Parse(layoutYYYYMMDD, e.Text[:10])
			wg.Done()
		})
	}

	for i, url := range nateNewsUrls {
		nateNews[i].Url = url
		nateNews[i].Portal = "nate"
		inUrl := url
		inIndex := i

		wg.Add(1 * setFieldCount)
		go func(c *colly.Collector) {
			c.Visit(HttpsUrl + inUrl)
		}(cSlice[inIndex]) // i+1 is ranking
	}
	wg.Wait()

	return nateNews
}
