package service

import (
	"github.com/gocolly/colly"
	"portal_news/lambda_crawler/model"
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	naverNewsPrefixUrl = "news.naver.com/"
	naverNewsRootUrl             = httpsUrl + naverNewsPrefixUrl + `main/ranking/popularDay.nhn?rankingType=age&subType=30`
	naverCssSelectorUrl = ".ranking_list li .ranking_headline a"
	naverCssSelectorTitle = "#articleTitle"
	naverCssSelectorContent      = "#articleBodyContents"
	naverCssSelectorPress        = ".press_logo img"
	naverCssSelectorDate         = ".t11"
	deleteString = "\n\t\n\t\n\n\n\n// flash 오류를 우회하기 위한 함수 추가\nfunction _flash_removeCallback() {}\n\n\t\n\t"
)

type NaverNewsCrawler struct {
}


func (crawler NaverNewsCrawler)CrawlNews() []*model.RankingNews{
	naverNewsUrls := crawler.GetNewsUrls(naverNewsRootUrl)
	naverNews := crawler.GetNews(naverNewsUrls)
	setJpTitle(naverNews)

	return naverNews
}

//get Naver News url From nate root url
func (crawler NaverNewsCrawler) GetNewsUrls(rootUrl string) []string{
	urls := make([]string,0, NewsCount)
	c := colly.NewCollector()
	var wg sync.WaitGroup
	wg.Add(NewsCount)

	// Find and visit all links
	c.OnHTML(naverCssSelectorUrl, func(e *colly.HTMLElement) {
		if len(urls) < NewsCount {
			url := e.Attr("href")
			urls = append(urls, naverNewsPrefixUrl + url)
			wg.Done()
		}
	})

	c.Visit(rootUrl)
	wg.Wait()

	return urls
}

//get NaverNews Object from nate urls
func (crawler NaverNewsCrawler)GetNews(newsUrls []string) []*model.RankingNews {
	naverNews := make([]*model.RankingNews, NewsCount, NewsCount)
	for i:=0; i<len(naverNews); i++{
		naverNews[i] = &model.RankingNews{}
	}


	cSlice := make([]*colly.Collector, NewsCount, NewsCount)
	dateDuplicateCheck := make([]bool,10)
	var wg sync.WaitGroup

	//Set callback
	for i := 0; i < NewsCount; i++ {
		inIndex := i
		cSlice[i] = colly.NewCollector()
		cSlice[i].OnHTML(naverCssSelectorTitle, func(e *colly.HTMLElement) {
			naverNews[inIndex].Title = e.Text

			wg.Done()
		})

		cSlice[i].OnHTML(naverCssSelectorContent, func(e *colly.HTMLElement) {
			content := e.Text
			//str := strings.ReplaceAll(content, deleteString, "")

			index := strings.Index(content, deleteString)
			startIndex := index + len(deleteString)
			str := content[startIndex:]

			space := regexp.MustCompile(`\s+`)
			str = space.ReplaceAllString(str, " ")
			naverNews[inIndex].Content = trimmingContent(str,200)

			wg.Done()
		})

		cSlice[i].OnHTML(naverCssSelectorPress, func(e *colly.HTMLElement) {
			naverNews[inIndex].Press = e.Attr("title")

			wg.Done()
		})

		cSlice[i].OnHTML(naverCssSelectorDate, func(e *colly.HTMLElement) {
			if dateDuplicateCheck[inIndex]{
				return
			}
			dateDuplicateCheck[inIndex] = true

			dateStrng := e.Text[:10]
			replacedDateString := strings.ReplaceAll(dateStrng, ".", "-")
			naverNews[inIndex].Date, _ = time.Parse(layoutYYYYMMDD, replacedDateString)

			wg.Done()
		})
	}

	for i, url := range newsUrls {
		naverNews[i].Url = url
		naverNews[i].Portal = "naver"
		inUrl := url
		inIndex := i

		wg.Add(1 * setFieldCount)
		go func(c *colly.Collector) {
			c.Visit(httpsUrl + inUrl)
		}(cSlice[inIndex]) // i+1 is ranking
	}

	wg.Wait()

	return naverNews
}