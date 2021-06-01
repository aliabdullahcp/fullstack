package controllers

import (
	"net/http"

	"github.com/aliabdullahcp/fullstack/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to the API Home Page.")

}
