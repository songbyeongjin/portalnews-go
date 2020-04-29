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
	NewsCount = 20

	HttpsUrl = `https://`
	NateNewsRootUrl         = HttpsUrl+`news.nate.com/rank/interest?sc=all&p=day&date=20999999`
	CssSelectorFirstToFifth = ".mlt01 a"
	CssSelectorSixthTo      = ".mduSubject a"
	CssSelectorTitle      = ".articleSubecjt"
	CssSelectorContent      = "#realArtcContents"
	CssSelectorPress      = ".articleInfo .medium"
	CssSelectorDate      = ".firstDate em"
	CssSelectorOrCondition  = `, `
	SetFieldCount = 4 //Title,Content,Press,Date
)

func CrawlNateNews(){
	nateNewsUrls := getNateNewsUrls(NateNewsRootUrl)
	nateNews := getNateNews(nateNewsUrls)

	for _,news := range nateNews{
		fmt.Println(news)
	}
}

//get Nate News url From nate root url
func getNateNewsUrls(nateRootUrl string) []string{
	urls := make([]string,0,NewsCount)
	c := colly.NewCollector()
	var wg2 sync.WaitGroup
	wg2.Add(20)

	// Find and visit all links
	c.OnHTML(CssSelectorFirstToFifth + CssSelectorOrCondition + CssSelectorSixthTo, func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		if len(urls) < NewsCount{
			title := e.Attr("href")[2:]
			urls = append(urls, title)
			wg2.Done()
		}
	})

	c.Visit(nateRootUrl)
	wg2.Wait()

	return urls
}

//get NateNews Object from nate urls
func getNateNews(nateNewsUrls []string) []model.News {
	nateNews := make([]model.News, NewsCount, NewsCount)

	cSlice := make([]*colly.Collector, NewsCount, NewsCount)
	var wg sync.WaitGroup

	for i := 0; i < NewsCount; i++ {
		inIndex := i
		cSlice[i] = colly.NewCollector()
		cSlice[i].OnHTML(CssSelectorTitle, func(e *colly.HTMLElement) {
			nateNews[inIndex].Title = e.Text
			wg.Done()
		})
		cSlice[i].OnHTML(CssSelectorContent, func(e *colly.HTMLElement) {
			space := regexp.MustCompile(`\s+`)
			str := space.ReplaceAllString(e.Text[:200]+"...", " ")
			nateNews[inIndex].Content = str

			wg.Done()
		})
		cSlice[i].OnHTML(CssSelectorPress, func(e *colly.HTMLElement) {
			nateNews[inIndex].Press = e.Text
			wg.Done()
		})
		cSlice[i].OnHTML(CssSelectorDate, func(e *colly.HTMLElement) {
			nateNews[inIndex].Date, _ = time.Parse("2006-01-02", e.Text[:10])
			wg.Done()
		})
	}

	for i, url := range nateNewsUrls {
		nateNews[i].Url = url
		nateNews[i].Portal = "nate"
		inUrl := url
		inIndex := i

		wg.Add(1 * SetFieldCount)
		go func(c *colly.Collector) {
			c.Visit(HttpsUrl + inUrl)
		}(cSlice[inIndex]) // i+1 is ranking
	}
	wg.Wait()

	return nateNews
}
