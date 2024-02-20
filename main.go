package main

import (
	"database/sql"
	"fmt"
	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"os"
	"unsafe"

	_ "modernc.org/sqlite"
)

func main() {
	dbFilePath := "./database.db"
	db, err := openDB(dbFilePath)
	errShouldNotHappen(err)

	err = applySchema(db)
	errShouldNotHappen(err)

	populateDatabaseWithStockData()
}

func populateDatabaseWithStockData() error {
	client := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    "aj",
		APISecret: "askdfj",
		BaseURL:   "https://paper-api.alpaca.markets",
	})
	_, err := client.GetAccount()
	if err != nil {
		return err
	}

	assets, err := client.GetAssets(alpaca.GetAssetsRequest{})
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", assets)
	return nil
}

func applySchema(db *sql.DB) error {
	schema, err := os.ReadFile("./schema.sql")
	errShouldNotHappen(err)

	schemaStrPtr := *(*string)(unsafe.Pointer(&schema))
	_, err = db.Exec(schemaStrPtr)
	return err
}

func openDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func errShouldNotHappen(err error) {
	if err != nil {
		panic(err)
	}
}
