package main

import (
	KynyxDataBase "BuffedScrappers/KynyxDB"
	"BuffedScrappers/ProductStruct"
	"fmt"
	"github.com/gocolly/colly/v2"
	"strconv"
	"time"
)

func main() {

	Collector := colly.NewCollector()
	MyFactory, _ := BuildScrapper(ZLibrary)
	Scrapper, _ := MyFactory.Build(ProductScrapper)
	start := 503179

	var MyBooks []ProductStruct.Book

	Db := KynyxDataBase.OpenDB()

	for i := start; i < (start + 20); i++ {

		Book, _ := Scrapper.Scrape(Collector, "https://breadl.org/d/"+strconv.Itoa(i))
		MyBooks = append(MyBooks, Book)
		fmt.Println(Book)
		time.Sleep(500 * time.Millisecond)

	}
	KynyxDataBase.CRUD(Db, MyBooks)
}
