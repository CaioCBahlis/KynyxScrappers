package ScrapeGutenberg

import (
	"BuffedScrappers/ProductStruct"
	"github.com/gocolly/colly/v2"
)

type GuttenbergProductScrapper struct {
}

func (G *GuttenbergProductScrapper) Scrape(c *colly.Collector, Url string) (ProductStruct.Book, error) {

	MyBook := ProductStruct.Book{}

	c.OnHTML("div.page_content", func(e *colly.HTMLElement) {
		Title := e.DOM.Find("h1").Text()
		MyBook.Name = Title

		MainPage := e.DOM.Find("div.page-body")

		Cover, _ := MainPage.Find("img[class=cover-art]").Attr("src")
		MyBook.CoverUrl = Cover

		e.ForEach("tr[class=even]", func(i int, e *colly.HTMLElement) {
			DownloadLink, _ := e.DOM.Find("a").Attr("href")
			MyBook.DownloadLinks = append(MyBook.DownloadLinks, "https://www.gutenberg.org/"+DownloadLink)
		})

	})

	c.Visit(Url)
	return MyBook, nil

}
