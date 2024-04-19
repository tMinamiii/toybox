package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	_ "github.com/go-sql-driver/mysql"
	project "github.com/tMinamiii/toybox/projectroot"
	"github.com/testcontainers/testcontainers-go"
	tcnetwork "github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
)

const dbname = "testdb"
const dbport = "3306/tcp"

func launchMySQL(ctx context.Context) testcontainers.Container {
	net, err := tcnetwork.New(ctx, tcnetwork.WithCheckDuplicate())
	if err != nil {
		log.Fatalf("failed to create network: %s", err)
	}

	req := testcontainers.ContainerRequest{
		Image: "mysql:8.0.32",
		Env: map[string]string{
			"MYSQL_ROOT_HOST":            "%",
			"MYSQL_DATABASE":             dbname,
			"MYSQL_ALLOW_EMPTY_PASSWORD": "yes",
		},
		Networks:     []string{net.Name},
		ExposedPorts: []string{dbport},
		WaitingFor:   wait.ForListeningPort(dbport),
		HostConfigModifier: func(cfg *container.HostConfig) {
			cfg.AutoRemove = true
			cfg.NetworkMode = network.NetworkBridge
		},
	}

	mysqlC, err := testcontainers.GenericContainer(ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})
	if err != nil {
		log.Fatalf("failed to create mysql container: %s", err)
	}

	return mysqlC
}

func migrate(ctx context.Context, mysqlC testcontainers.Container) {
	ip, err := mysqlC.ContainerIP(ctx)
	if err != nil {
		log.Fatalf("failed to get ip address: %s", err)
	}

	networks, err := mysqlC.Networks(ctx)
	if err != nil {
		log.Fatalf("failed to get ip networks: %s", err)
	}

	port := nat.Port(dbport).Int()
	url := fmt.Sprintf("mysql://root:@%s:%d/%s", ip, port, dbname)

	req := testcontainers.ContainerRequest{
		Image: "arigaio/atlas",
		Cmd:   []string{"migrate", "apply", "-u", url},
		HostConfigModifier: func(cfg *container.HostConfig) {
			cfg.AutoRemove = true
			cfg.NetworkMode = network.NetworkBridge
			cfg.Mounts = []mount.Mount{
				{
					Type:     mount.TypeBind,
					Source:   project.Root() + "/testcontainers/migrations",
					Target:   "/migrations",
					ReadOnly: true,
				},
			}
		},
		Networks:   networks,
		WaitingFor: wait.ForExit(),
	}

	_, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("failed to create atlas container: %s", err)
	}
}

func showtable(ctx context.Context, mysqlC testcontainers.Container) {
	host, err := mysqlC.Host(ctx)
	if err != nil {
		log.Fatalf("failed to get host: %s", err)
	}

	port, err := mysqlC.MappedPort(ctx, dbport)
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

	migrate(ctx, mysqlC)

	// check result
	showtable(ctx, mysqlC)
}
