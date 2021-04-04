package main

import (
	"flag"
	"fmt"
	"library-system/console"
	"library-system/db"
	"library-system/repositories"
	"library-system/services"
	"os"
)

var helpFlag = flag.Bool("help", false, "Display a helpful message")

func main() {
	dbConnection := db.GetConnection()
	defer dbConnection.Close()

	bookRepository := repositories.NewBookRepository(dbConnection)
	bookService := services.NewBookService(bookRepository)

	flag.Parse()
	c := console.NewBookConsole(bookService)

	if *helpFlag || len(os.Args) == 1 {
		c.Help()
		return
	}

	err := c.Console()
	if err != nil {
		fmt.Printf("console option error: %v\n", err)
		os.Exit(2)
	}
}
