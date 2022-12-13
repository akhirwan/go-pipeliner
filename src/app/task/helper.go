package task

import (
	"go-pipeliner/src/infrastructure/config"
	"go-pipeliner/src/infrastructure/repository/db"
	"go-pipeliner/src/infrastructure/tunnel"
	"strconv"
)

func createDWHDevConfig(config config.Config) *db.MySQLConfig {
	var tunnelConfig *tunnel.SSHConfig
	// Check if db connection uses tunnel
	if config.Get("DEV_MYSQL_TUNNEL") == "true" {
		sshPort, _ := strconv.Atoi(config.Get("TUNNEL_SSH_PORT"))
		tunnelConfig = &tunnel.SSHConfig{
			Name:               config.Get("TUNNEL_SSH_NAME"),
			Host:               config.Get("TUNNEL_SSH_HOST"),
			Port:               sshPort,
			User:               config.Get("TUNNEL_SSH_USERNAME"),
			Password:           config.Get("TUNNEL_SSH_PASSWORD"),
			PrivateKeyFile:     config.Get("TUNNEL_SSH_PRIVATEKEYFILE"),
			PrivateKeyPassword: config.Get("TUNNEL_SSH_PRIVATEKEYPASSWORD"),
		}
	}

	mysqlPort, _ := strconv.Atoi(config.Get("DEV_MYSQL_PORT"))
	return &db.MySQLConfig{
		Host:     config.Get("DEV_MYSQL_HOST"),
		Port:     mysqlPort,
		User:     config.Get("DEV_MYSQL_USERNAME"),
		Password: config.Get("DEV_MYSQL_PASSWORD"),
		Database: config.Get("DEV_MYSQL_DATABASE"),
		Tunnel:   tunnelConfig,
	}
}

func createDWHProdConfig(config config.Config) *db.MySQLConfig {
	var tunnelConfig *tunnel.SSHConfig
	// Check if db connection uses tunnel
	if config.Get("PRD_MYSQL_TUNNEL") == "true" {
		sshPort, _ := strconv.Atoi(config.Get("TUNNEL_SSH_PORT"))
		tunnelConfig = &tunnel.SSHConfig{
			Name:               config.Get("TUNNEL_SSH_NAME"),
			Host:               config.Get("TUNNEL_SSH_HOST"),
			Port:               sshPort,
			User:               config.Get("TUNNEL_SSH_USERNAME"),
			Password:           config.Get("TUNNEL_SSH_PASSWORD"),
			PrivateKeyFile:     config.Get("TUNNEL_SSH_PRIVATEKEYFILE"),
			PrivateKeyPassword: config.Get("TUNNEL_SSH_PRIVATEKEYPASSWORD"),
		}
	}

	mysqlPort, _ := strconv.Atoi(config.Get("PRD_MYSQL_PORT"))
	return &db.MySQLConfig{
		Host:     config.Get("PRD_MYSQL_HOST"),
		Port:     mysqlPort,
		User:     config.Get("PRD_MYSQL_USERNAME"),
		Password: config.Get("PRD_MYSQL_PASSWORD"),
		Database: config.Get("PRD_MYSQL_DATABASE"),
		Tunnel:   tunnelConfig,
	}
}
