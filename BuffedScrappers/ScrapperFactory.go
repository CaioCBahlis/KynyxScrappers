package main

import (
	"BuffedScrappers/ScrapperStruct"
	"BuffedScrappers/Scrappers/ScrapeGutenberg"
	"BuffedScrappers/Scrappers/ScrapeZLib"
	"fmt"
	"strings"
)

const (
	ZLibrary        = "zlibrary"
	HubScrapper     = "hub"
	ProductScrapper = "product"
	Gutenberg       = "gutenberg"
)

type Factory interface {
	Build(string) (ScrapperStruct.Scrapper, error)
}

func BuildScrapper(family string) (Factory, error) {
	family = strings.ToLower(family)
	switch family {
	case ZLibrary:
		return &ZLibFactory{}, nil
	case Gutenberg:
		return &GutenbergFactory{}, nil
	default:
		return nil, fmt.Errorf("unknown family %s", family)
	}
}

type ZLibFactory struct {
}

func (ZF *ZLibFactory) Build(Section string) (ScrapperStruct.Scrapper, error) {
	Section = strings.ToLower(Section)
	switch Section {
	case "hub":
		return nil, nil //&ScrapeZLib.ZLibHubScrapper, nil
	case "product":
		return &ScrapeZLib.ZlibProductScrapper{}, nil
	default:
		return nil, fmt.Errorf("unknown Section %s", Section)
	}
}

type GutenbergFactory struct {
}

func (GF *GutenbergFactory) Build(Section string) (ScrapperStruct.Scrapper, error) {
	Section = strings.ToLower(Section)
	switch Section {
	case "product":
		return &ScrapeGutenberg.GuttenbergProductScrapper{}, nil
	default:
		return nil, nil
	}
}
