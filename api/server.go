package api

import (
	db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/nmpotential/simplebank/token"
	"github.com/nmpotential/simplebank/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(server *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.POST("/accounts/:id", server.createAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
