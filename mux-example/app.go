package main

import (
	"example.com/controller"
	"example.com/orch"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fmt"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
	"log"

	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize() {

	var err error

	a.DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	orch := orch.NewOrchestratorService()
	controller := controller.NewController(orch, a.Router)
	controller.InitializeRoutes()

}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":2999", a.Router))
}