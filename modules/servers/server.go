package servers

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"

	_pkgConfig "github.com/MarkTBSS/075_Role_Based_Authorization/config"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type IServer interface {
	Start()
}

type server struct {
	cfg _pkgConfig.IConfig
	app *fiber.App
	db  *sqlx.DB
}

func NewServer(cfg _pkgConfig.IConfig, db *sqlx.DB) IServer {
	return &server{
		cfg: cfg,
		db:  db,
		app: fiber.New(fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
	}
}

func (s *server) Start() {
	// Middlewares
	middlewares := InitMiddlewares(s)
	s.app.Use(middlewares.Logger())
	s.app.Use(middlewares.Cors())

	// Modules
	v1 := s.app.Group("v1")
	modules := InitModule(v1, s, middlewares)
	modules.MonitorModule()
	modules.UsersModule()

	s.app.Use(middlewares.RouterCheck())

	// Graceful Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Println("server is shutting down...")
		_ = s.app.Shutdown()
	}()

	// Listen to host:port
	log.Printf("server is starting on %v", s.cfg.App().Url())
	s.app.Listen(s.cfg.App().Url())
}
