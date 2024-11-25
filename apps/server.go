package apps

import (
	"github.com/gin-gonic/gin"
	"github.com/ij4l/go-htmx/apps/home"
	"github.com/ij4l/go-htmx/internal/utils"
)

type Server struct {
	router *gin.Engine
}

func NewServer() (*Server, error) {
	router := gin.Default()
	server := &Server{router: router}
	server.setupRouter()
	return server, nil
}

func (s Server) setupRouter() {
	s.router.HTMLRender = &utils.TemplRender{}
	s.router.GET("/", home.GetHome)
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
