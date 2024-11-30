package apps

import (
	"github.com/gin-gonic/gin"
	"github.com/ij4l/go-htmx/apps/contact"
	"github.com/ij4l/go-htmx/internal/config"
	"github.com/ij4l/go-htmx/internal/utils"
	"github.com/jackc/pgx/v5"
)

type Server struct {
	Router  *gin.Engine
	pgxConn *pgx.Conn
	config  config.Config
}

func NewServer(config config.Config, pgx *pgx.Conn) (*Server, error) {
	router := gin.Default()
	server := Server{Router: router, config: config, pgxConn: pgx}
	server.setupRouter()
	return &server, nil
}

func (s Server) setupRouter() {
	s.Router.HTMLRender = &utils.TemplRender{}

	s.Router.GET("/contact", contact.GetContact)
	s.Router.GET("/images/:imageName", serveImage)
	s.Router.GET("/icons/:iconName", serveIcon)
	s.Router.GET("/storages/:fileName", serveStorage)
}

func (s *Server) Start() error {
	return s.Router.Run(s.config.ServerAddress)
}
