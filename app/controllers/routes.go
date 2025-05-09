package controllers

import "github.com/gorilla/mux"

func (server *Server) InitializeRoutes() {

	server.Router = mux.NewRouter()

}
