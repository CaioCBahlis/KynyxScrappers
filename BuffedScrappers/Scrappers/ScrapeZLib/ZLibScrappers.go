package ScrapeZLib

import (
	"BuffedScrappers/ProductStruct"
	"encoding/base64"
	"errors"
	"github.com/RomainMichau/cloudscraper_go/cloudscraper"
	"github.com/gocolly/colly/v2"
	"strings"
)

type ZlibProductScrapper struct {
	input string
}

func (Z *ZlibProductScrapper) Scrape(c *colly.Collector, Url string) (ProductStruct.Book, error) {
	MyBook := ProductStruct.Book{}

	client, _ := cloudscraper.Init(false, false)
	res2, _ := client.Get(Url, make(map[string]string), "")
	Z.input = res2.Body

	Title, err := Z.GetTitle()
	if err != nil {
		return ProductStruct.Book{}, err
	}
	MyBook.Name = Title

	Cover, err := Z.GetCover()
	if err != nil {
		return ProductStruct.Book{}, err
	}

	MyBook.CoverUrl = Cover

	DownloadLinks := Z.GetAll("openLinkNewTab('", "'", []string{})
	if len(DownloadLinks) > 0 {
		MyBook.DownloadLinks = DownloadLinks
	}

	return MyBook, nil
}

func (Z *ZlibProductScrapper) GetTitle() (string, error) {
	start := strings.Index(Z.input, "<h1>") + len("<h1")
	if start == -1 {
		return "", errors.New("cant find substring")
	}
	end := strings.Index(Z.input, "</h1>")
	return Z.input[start:end], nil
}

func (Z *ZlibProductScrapper) GetCover() (string, error) {
	start := strings.Index(Z.input, ".src=") + len(".src='")
	if start == -1 {
		return "", errors.New("couldn't find the subtring")
	}
	end := strings.Index(Z.input[start:], ";") + start

	return Z.input[start+1 : end-1], nil
}

func (Z *ZlibProductScrapper) GetAll(substr string, ending string, found []string) []string {
	Index := strings.Index(Z.input, substr)
	if Index == -1 {
		return found
	}

	start := Index + len(substr)
	end := strings.Index(Z.input[start:], ending) + start
	url, _ := base64.URLEncoding.DecodeString(Z.input[start:end])

	Z.input = Z.input[:Index] + Z.input[end:]
	found = append(found, string(url))
	return Z.GetAll(substr, ending, found)

}

type ZLibHubScrapper struct {
}

func (Z *ZLibHubScrapper) Scrape(c *colly.Collector, Url string) (ProductStruct.Book, error) {
	return ProductStruct.Book{}, nil
}
