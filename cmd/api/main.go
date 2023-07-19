package main

import (
	"net/http"
	"os"

	"github.com/ibanezv/fly-devs-practice_go/cmd/api/handlers"
	"github.com/ibanezv/fly-devs-practice_go/cmd/api/settings"
	"github.com/ibanezv/fly-devs-practice_go/internal/repository"
	"github.com/ibanezv/fly-devs-practice_go/internal/services"
	"github.com/rainycape/memcache"
)

func main() {
	configs := settings.LoadConfigurationDB()
	db := repository.NewDatabase(configs)
	cnn, err := db.GetConnection()
	if err != nil {
		os.Exit(1)
	}

	mc, err := memcache.New("127.0.0.1:11211")
	if err != nil {
		panic(err)
	}

	bookRepository := repository.NewRepository(cnn, mc)
	bookService := services.NewBookService(bookRepository)
	router := handlers.ApiRoutesMapper(*bookService)
	http.Handle("/", router)
	_ = http.ListenAndServe(":8090", nil)
}
