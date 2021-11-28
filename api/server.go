package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/sakhaei-wd/espresso/db/sqlc"
	"github.com/sakhaei-wd/espresso/token"
	"github.com/sakhaei-wd/espresso/util"
)

//This Server will serves all HTTP requests for our banking service
type Server struct {
	config     util.Config
	store      db.Store    //It will allow us to interact with the database when processing API requests from clients.
	router     *gin.Engine //This router will help us send each API request to the correct handler for processing.
	tokenMaker token.Maker
}

//This function will create a new Server instance, and setup all HTTP API routes for our service on that server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	//It requires a symmetric key string, so we will need to load this from environment variable
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
					 //token.NewJWTMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

//This function will take an error as input, and it will return a gin.H object
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

//Its role is to run the HTTP server on the input address to start listening for API requests
//Note that the server.router field is private, so it cannot be accessed from outside of this api package. That’s one of the reasons we have this public Start() function.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default() //we create a new router by  calling gin.Default()

	// router.POST("/accounts", server.createAccount)
	// router.GET("/accounts/:id", server.getAccount)
	// router.GET("/accounts", server.listAccount)

	// router.POST("/entries", server.createEntry)
	// router.GET("/entries/:id", server.getEntry)
	// router.GET("/entries", server.listEntry)

	// router.POST("/transfers", server.createTransfer)

	// router.GET("/user/:username", server.getUser)
	// router.POST("/user", server.createUser)
	// router.POST("/user/login", server.loginUser)

	server.router = router
}
