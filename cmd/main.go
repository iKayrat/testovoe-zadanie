package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/iKayrat/test/internal/app/config"
	"github.com/iKayrat/test/internal/app/controller"
	sqlc "github.com/iKayrat/test/internal/app/db/sqlc"

	_ "github.com/lib/pq"
)

func main() {
	// Parse command-line arguments

	conf, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal(err)
	}

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please provide order numbers as arguments.")
		return
	}

	db, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	storage := sqlc.NewStore(db)
	if err != nil {
		return
	}

	ctx := context.Background()

	service := controller.New(ctx, storage)

	_, err = service.GetPages(args)
	if err != nil {
		log.Println("get pages err: ", err)
		return
	}

	service.BeautyPrint(args)
}
