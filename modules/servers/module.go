package servers

import (
	"github.com/MarkTBSS/075_Role_Based_Authorization/modules/middlewares/middlewaresHandlers"
	"github.com/MarkTBSS/075_Role_Based_Authorization/modules/middlewares/middlewaresRepositories"
	"github.com/MarkTBSS/075_Role_Based_Authorization/modules/middlewares/middlewaresUsecases"
	_pkgModulesMonitorMonitorHandlers "github.com/MarkTBSS/075_Role_Based_Authorization/modules/monitor/monitorHandlers"
	"github.com/MarkTBSS/075_Role_Based_Authorization/modules/users/usersHandlers"
	"github.com/MarkTBSS/075_Role_Based_Authorization/modules/users/usersRepositories"
	"github.com/MarkTBSS/075_Role_Based_Authorization/modules/users/usersUsecases"
	"github.com/gofiber/fiber/v2"
)

type IModuleFactory interface {
	MonitorModule()
	UsersModule()
}

type moduleFactory struct {
	r   fiber.Router
	s   *server
	mid middlewaresHandlers.IMiddlewaresHandler
}

func InitModule(r fiber.Router, s *server, mid middlewaresHandlers.IMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		r:   r,
		s:   s,
		mid: mid,
	}
}

func InitMiddlewares(s *server) middlewaresHandlers.IMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecases.MiddlewaresUsecase(repository)
	handler := middlewaresHandlers.MiddlewaresHandler(usecase, s.cfg)
	return handler
	//return middlewaresHandlers.MiddlewaresHandler(usecase, s.cfg)
}

func (m *moduleFactory) MonitorModule() {
	handler := _pkgModulesMonitorMonitorHandlers.MonitorHandler(m.s.cfg)
	m.r.Get("/", handler.HealthCheck)
}

func (m *moduleFactory) UsersModule() {
	repository := usersRepositories.UsersRepository(m.s.db)
	usecase := usersUsecases.UsersUsecase(m.s.cfg, repository)
	handler := usersHandlers.UsersHandler(m.s.cfg, usecase)

	router := m.r.Group("/users")

	router.Post("/signup", handler.SignUpCustomer)
	router.Post("/signin", handler.SignIn)
	router.Post("/refresh", handler.RefreshPassport)
	router.Post("/signout", handler.SignOut)
	router.Post("/signup-admin", handler.SignUpAdmin)
	router.Get("/:user_id", m.mid.JwtAuth(), m.mid.ParamsCheck(), handler.GetUserProfile)
	// Admin only
	router.Get("/admin/secret", m.mid.JwtAuth(), m.mid.Authorize(2), handler.GenerateAdminToken)
	// Admin and Customer Role
	// Insert other roles by roleId
	// router.Get("/admin/secret", m.mid.JwtAuth(), m.mid.Authorize(2, 1), handler.GenerateAdminToken)
}
