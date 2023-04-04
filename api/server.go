package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/mattchw/smart-bank/db/sqlc"
	"github.com/mattchw/smart-bank/internal/token"
	token_interfaces "github.com/mattchw/smart-bank/internal/token/interfaces"
	"github.com/mattchw/smart-bank/util"
)

type Server struct {
	store      db.Store
	tokenMaker token_interfaces.TokenMaker
	router     *gin.Engine
	config     util.Config
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validateCurrency)
	}

	server.initRoutes()
	return server, nil
}

func (server *Server) initRoutes() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	router.GET("/accounts/", server.listAccounts)
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)

	router.POST("/transfers/", server.createTransfer)

	server.router = router
}
