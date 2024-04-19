package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	_ "github.com/go-sql-driver/mysql"
	project "github.com/tMinamiii/toybox/projectroot"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const dbname = "testdb"

func launchMySQL(ctx context.Context) testcontainers.Container {
	mysqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "mysql:8.0.32",
			Env: map[string]string{
				"MYSQL_ROOT_HOST":            "%",
				"MYSQL_DATABASE":             dbname,
				"MYSQL_ALLOW_EMPTY_PASSWORD": "yes",
			},
			ExposedPorts: []string{"3306/tcp"},
			WaitingFor:   wait.ForListeningPort("3306/tcp"),
			HostConfigModifier: func(cfg *container.HostConfig) {
				cfg.AutoRemove = true
				cfg.NetworkMode = network.NetworkBridge
				cfg.Mounts = []mount.Mount{
					{
						Type:     mount.TypeBind,
						Source:   project.Root() + "/ent/migrate/migrations",
						Target:   "/docker-entrypoint-initdb.d",
						ReadOnly: true,
					},
				}
			},
		},
		Started: true,
	})
	if err != nil {
		log.Fatalf("failed to create mysql container: %s", err)
	}

	return mysqlC
}

func showtable(ctx context.Context, mysqlC testcontainers.Container) {
	host, err := mysqlC.Host(ctx)
	if err != nil {
		log.Fatalf("failed to get host: %s", err)
	}

	port, err := mysqlC.MappedPort(ctx, "3306/tcp")
	if err != nil {
		log.Fatalf("failed to get externally mapped port: %s", err)
	}

	ds := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true",
		"root",
		"",
		host,
		port.Int(),
		dbname,
	)

	tdb, err := sql.Open("mysql", ds)
	if err != nil {
		log.Fatalf("failed to open database: %s", err)
	}

	if err := tdb.Ping(); err != nil {
		log.Fatalf("failed to verify a connection to the database: %s", err)
	}

	res, _ := tdb.Query("show tables;")
	var table string
	for res.Next() {
		res.Scan(&table)
		fmt.Println(table)
	}
}

func main() {
	ctx := context.Background()

	mysqlC := launchMySQL(ctx)
	defer func() {
		if err := mysqlC.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate mysql container: %s", err)
		}
	}()

	showtable(ctx, mysqlC)
}
