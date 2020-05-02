package service

import (
	"github.com/gocolly/colly"
	"portal_news/lambda_crawler/model"
	"regexp"
	"sync"
	"time"
)

const (
	nateNewsRootUrl             = httpsUrl +`news.nate.com/rank/interest?sc=all&p=day&date=20999999`
	nateCssSelectorFirstToFifth = ".mlt01 a"
	nateCssSelectorSixthTo      = ".mduSubject a"
	nateCssSelectorTitle        = ".articleSubecjt"
	nateCssSelectorContent      = "#realArtcContents"
	nateCssSelectorPress        = ".articleInfo .medium"
	nateCssSelectorDate         = ".firstDate em"
)

type NateNewsCrawler struct {
}

func (crawler NateNewsCrawler)CrawlNews() []*model.RankingNews{
	nateNewsUrls := crawler.GetNewsUrls(nateNewsRootUrl)
	nateNews := crawler.GetNews(nateNewsUrls)
	setJpTitle(nateNews)

	return nateNews
}


//get Nate News url From nate root url
func (crawler NateNewsCrawler) GetNewsUrls(rootUrl string) []string{
	urls := make([]string,0, NewsCount)
	c := colly.NewCollector()
	var wg sync.WaitGroup
	wg.Add(NewsCount)

	// Find and visit all links
	c.OnHTML(nateCssSelectorFirstToFifth+cssSelectorOrCondition+nateCssSelectorSixthTo, func(e *colly.HTMLElement) {
		if len(urls) < NewsCount {
			url := e.Attr("href")[2:]//delete string("//") in title
			urls = append(urls, url)
			wg.Done()
		}
	})

	c.Visit(rootUrl)
	wg.Wait()

	return urls
}

//get NateNews Object from nate urls
func (crawler NateNewsCrawler) GetNews(newsUrls []string) []*model.RankingNews {
	nateNews := make([]*model.RankingNews, NewsCount, NewsCount)
	for i:=0; i<len(nateNews); i++{
		nateNews[i] = &model.RankingNews{}
	}

	cSlice := make([]*colly.Collector, NewsCount, NewsCount)
	var wg sync.WaitGroup

	//Set callback
	for i := 0; i < NewsCount; i++ {
		inIndex := i
		cSlice[i] = colly.NewCollector()
		cSlice[i].OnHTML(nateCssSelectorTitle, func(e *colly.HTMLElement) {
			nateNews[inIndex].Title = e.Text
			wg.Done()
		})

		cSlice[i].OnHTML(nateCssSelectorContent, func(e *colly.HTMLElement) {
			space := regexp.MustCompile(`\s+`)
			str := space.ReplaceAllString(e.Text, " ")
			nateNews[inIndex].Content = trimmingContent(str,200)

			wg.Done()
		})

		cSlice[i].OnHTML(nateCssSelectorPress, func(e *colly.HTMLElement) {
			nateNews[inIndex].Press = e.Text
			wg.Done()
		})

		cSlice[i].OnHTML(nateCssSelectorDate, func(e *colly.HTMLElement) {

			nateNews[inIndex].Date, _ = time.Parse(layoutYYYYMMDD, e.Text[:10])
			wg.Done()
		})
	}

	for i, url := range newsUrls {
		nateNews[i].Url = url
		nateNews[i].Portal = "nate"
		inUrl := url
		inIndex := i

		wg.Add(1 * setFieldCount)
		go func(c *colly.Collector) {
			c.Visit(httpsUrl + inUrl)
		}(cSlice[inIndex]) // i+1 is ranking
	}
	wg.Wait()

	return nateNews
}

