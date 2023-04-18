package controller

import (
	"encoding/json"
	"example.com/dto"
	"example.com/orch"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type IController interface {
	InitializeRoutes()
	TestKits(w http.ResponseWriter, r *http.Request)
}

type Controller struct {
	orchestratorService orch.IOrchestratorService
	Router              *mux.Router
}

func NewController(orchestratorService orch.IOrchestratorService, router *mux.Router) IController {
	return &Controller{
		orchestratorService: orchestratorService,
		Router:              router,
	}
}
func (c *Controller) InitializeRoutes() {
	c.Router.HandleFunc("/testkits", c.TestKits).Methods("GET")
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

type errorResponse struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	errResp := errorResponse{Error: message}
	respondWithJSON(w, status, errResp)
}

func (c *Controller) TestKits(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	p, err := c.orchestratorService.TestKits(start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())

	}

	respondWithJSON(w, http.StatusOK, p)

}
