package ScrapperStruct

import (
	"BuffedScrappers/ProductStruct"
	"github.com/gocolly/colly/v2"
)

type Scrapper interface {
	Scrape(c *colly.Collector, Url string) (ProductStruct.Book, error)
}
