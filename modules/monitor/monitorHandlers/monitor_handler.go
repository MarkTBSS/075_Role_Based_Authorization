package monitorHandlers

import (
	_pkgConfig "github.com/MarkTBSS/075_Role_Based_Authorization/config"
	_pkgModulesEntities "github.com/MarkTBSS/075_Role_Based_Authorization/modules/entities"
	_pkgModulesMonitor "github.com/MarkTBSS/075_Role_Based_Authorization/modules/monitor"
	"github.com/gofiber/fiber/v2"
)

type IMontitorHandler interface {
	HealthCheck(c *fiber.Ctx) error
}

type monitorHandler struct {
	cfg _pkgConfig.IConfig
}

func MonitorHandler(cfg _pkgConfig.IConfig) IMontitorHandler {
	return &monitorHandler{
		cfg: cfg,
	}
}

func (h *monitorHandler) HealthCheck(c *fiber.Ctx) error {
	res := &_pkgModulesMonitor.Monitor{
		Name:    h.cfg.App().Name(),
		Version: h.cfg.App().Version(),
	}
	//return c.Status(fiber.StatusOK).JSON(res)
	return _pkgModulesEntities.NewResponse(c).Success(fiber.StatusOK, res).Res()
}
