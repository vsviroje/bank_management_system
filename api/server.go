package api

import (
	db "github.com/Golang/bank_management_system/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := Server{store: store}
	router := gin.Default()

	v, isOk := binding.Validator.Engine().(*validator.Validate)
	if isOk {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/account", server.CreateAccount)
	router.GET("/account/:id", server.GetAccount)
	router.GET("/accounts", server.GetAccountList)

	router.POST("/transfer", server.CreateTransfer)

	server.router = router
	return &server
}

func (Server *Server) Start(address string) error {
	return Server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
