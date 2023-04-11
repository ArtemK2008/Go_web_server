package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type API struct {
	//UNEXPORTED!!!
	config *Config
	logger *logrus.Logger
	router *mux.Router
	data   Data
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	api.logger.Info("starting api server at port ", api.config.BindAddr)

	api.configureRouterField()
	return http.ListenAndServe(":"+api.config.BindAddr, api.router)
}
