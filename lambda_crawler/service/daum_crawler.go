package service

import (
	"fmt"
	"github.com/gocolly/colly"
	"portal_news/lambda_crawler/model"
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	daumNewsRootUrl             = httpsUrl +`media.daum.net/ranking/popular/`
	daumCssSelectorUrl 			= ".list_news2 .tit_thumb a"
	daumCssSelectorTitle        = ".tit_view"
	daumCssSelectorContent      = "#harmonyContainer"
	daumCssSelectorPress        = ".link_cp .thumb_g"
	daumCssSelectorDate         = ".info_view"

	dateBoundaryStr = "입력 "
)

type DaumNewsCrawler struct {
}
func (crawler DaumNewsCrawler) CrawlNews() []*model.RankingNews{
	daumNewsUrls := crawler.GetNewsUrls(daumNewsRootUrl)
	daumNews := crawler.GetNews(daumNewsUrls)
	setJpTitle(daumNews)

	return daumNews
}

//get Daum News url From nate root url
func (crawler DaumNewsCrawler) GetNewsUrls(rootUrl string) []string{
	urls := make([]string,0, NewsCount)
	c := colly.NewCollector()
	var wg sync.WaitGroup
	wg.Add(NewsCount)

	// Find and visit all links
	c.OnHTML(daumCssSelectorUrl, func(e *colly.HTMLElement) {
		if len(urls) < NewsCount {
			url := e.Attr("href")//delete string("//") in title

			// delete "http://" string
			url = strings.ReplaceAll(url ,"http://" ,"" )
			urls = append(urls, url)
			wg.Done()
		}
	})

	c.Visit(rootUrl)
	wg.Wait()

	return urls
}

//get DaumNews Object from nate urls
func (crawler DaumNewsCrawler) GetNews(newsUrls []string) []*model.RankingNews {
	daumNews := make([]*model.RankingNews, NewsCount, NewsCount)
	for i:=0; i<len(daumNews); i++{
		daumNews[i] = &model.RankingNews{}
	}

	cSlice := make([]*colly.Collector, NewsCount, NewsCount)
	var wg sync.WaitGroup

	//Set callback
	for i := 0; i < NewsCount; i++ {
		inIndex := i
		cSlice[i] = colly.NewCollector()
		cSlice[i].OnHTML(daumCssSelectorTitle, func(e *colly.HTMLElement) {
			daumNews[inIndex].Title = e.Text

			wg.Done()
		})

		cSlice[i].OnHTML(daumCssSelectorContent, func(e *colly.HTMLElement) {
			space := regexp.MustCompile(`\s+`)
			str := space.ReplaceAllString(e.Text, " ")
			daumNews[inIndex].Content = trimmingContent(str,200)

			wg.Done()
		})

		cSlice[i].OnHTML(daumCssSelectorPress, func(e *colly.HTMLElement) {
			daumNews[inIndex].Press = e.Attr("alt")

			wg.Done()
		})

		cSlice[i].OnHTML(daumCssSelectorDate, func(e *colly.HTMLElement) {
			date := e.Text
			fmt.Println(len(date))
			dateIndex := strings.Index(date, dateBoundaryStr)
			slicePoint := dateIndex + len(dateBoundaryStr)
			date = date[slicePoint:slicePoint+10]
			date = strings.ReplaceAll(date, ".", "-")
			daumNews[inIndex].Date, _ = time.Parse(layoutYYYYMMDD, date)

			wg.Done()
		})
	}

	for i, url := range newsUrls {
		daumNews[i].Url = url
		daumNews[i].Portal = "daum"
		inUrl := url
		inIndex := i

		wg.Add(1 * setFieldCount)
		go func(c *colly.Collector) {
			goUrl := "http://" + inUrl
			c.Visit(goUrl)
		}(cSlice[inIndex]) // i+1 is ranking
	}
	wg.Wait()

	return daumNews
}