package main

import (
	"os"

	_pkgConfig "github.com/MarkTBSS/075_Role_Based_Authorization/config"
	_pkgModulesServers "github.com/MarkTBSS/075_Role_Based_Authorization/modules/servers"
	_pkgDatabase "github.com/MarkTBSS/075_Role_Based_Authorization/pkg/databases"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := _pkgConfig.LoadConfig(envPath())
	db := _pkgDatabase.DbConnect(cfg.Db())
	defer db.Close()
	_pkgModulesServers.NewServer(cfg, db).Start()
}
