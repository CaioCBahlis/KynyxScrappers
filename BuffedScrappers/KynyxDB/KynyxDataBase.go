package KynyxDB

import (
	"BuffedScrappers/ProductStruct"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "kynyxdb"
)

var MyDatabase *sql.DB

func OpenDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if MyDatabase == nil {
		MyDatabase = db
	}
	return MyDatabase

}

func CRUD(db *sql.DB, MyBooks []ProductStruct.Book) {

	for _, Product := range MyBooks {
		sqlStatement := `
			INSERT INTO Books (Title, ImgUrl, DownloadLinks) VALUES ($1, $2, $3)
			`

		res, err := db.Exec(sqlStatement, Product.Name, Product.CoverUrl, pq.Array(Product.DownloadLinks))
		fmt.Println(res, err)
	}

}
