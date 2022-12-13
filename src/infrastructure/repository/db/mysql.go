package db

import (
	"context"
	"fmt"
	"go-pipeliner/src/infrastructure/tunnel"
	"log"
	"net"
	"time"

	// mysql driver
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Tunnel   *tunnel.SSHConfig
}

// NewMySQLDBConnection return instance of DB Connection
func NewMySQLDBConnection(c *MySQLConfig) *sqlx.DB {
	// network name default tcp optional as long as host and port wrapped in ()
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?timeout=3s&charset=utf8mb4&parseTime=true&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database)

	// Check if using Tunnel
	if c.Tunnel != nil {
		// 1. Create ssh client connection (sshCon)
		// 2. Create mysql dialer. When called it will dial to mysql server on top of sshCon
		// 3. Wrap MySQL host/address inside the sql dialer name
		sshConn := c.Tunnel.CreateSSHTunnel()

		mysql.RegisterDialContext(c.Tunnel.Name, func(ctx context.Context, addr string) (net.Conn, error) {
			// ctx.Deadline() // cancel if timeout dialing
			log.Println("Registering mysql custom dial function.")
			return sshConn.Dial("tcp", addr)
		})

		// Tunnel name must be set as dataSourceName
		dataSourceName = fmt.Sprintf("%s:%s@%s(%s:%v)/%s?timeout=3s&charset=utf8mb4&parseTime=true&loc=Local", c.User, c.Password, c.Tunnel.Name, c.Host, c.Port, c.Database)
	}

	db, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		log.Printf("Connect Database failed: %s", err.Error())
		return nil
	}

	// No need to ping manually to check if server is reachable
	// Sqlx call will it internally on connect
	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Server is unreachable. %s", err.Error())
		return nil
	}

	log.Printf("Successfully connected to the db.\n")
	// ctx, cancelFunc := context.WithCancel(context.Background())
	// defer cancelFunc()
	// if rows, err := db.QueryContext(ctx, "SELECT 1 AS id, 'aaa' AS name"); err == nil {
	// 	for rows.Next() {
	// 		var id int64
	// 		var name string
	// 		rows.Scan(&id, &name)
	// 		fmt.Printf("ID: %d  Name: %s\n", id, name)
	// 	}
	// 	rows.Close()
	// } else {
	// 	fmt.Printf("Failure: %s", err.Error())
	// }

	db.SetConnMaxLifetime(time.Minute * 15)
	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(30) // Increase for PP01/EV01

	log.Println("MySQL Database connected")

	return db
}
