package canned

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/docker/docker/client"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type Mysql struct {
	Container    testcontainers.Container
	DockerClient *client.Client

	Host     string
	Port     string
	Database string
	Username string
	Password string
	Client   *sql.DB
}

func NewMysql8(ctx context.Context) (*Mysql, error) {
	return newMysql(ctx, "8.0")
}

func NewMysql(ctx context.Context) (*Mysql, error) {
	return newMysql(ctx, "5.6")
}
func newMysql(ctx context.Context, version string) (*Mysql, error) {
	os.Setenv("TC_HOST", "localhost")
	req := testcontainers.ContainerRequest{
		Image:        "mysql:" + version,
		ExposedPorts: []string{"3306/tcp"},
		WaitingFor:   wait.ForListeningPort("3306"),
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "password",
			"MYSQL_DATABASE":      "test",
		},
		AutoRemove: true,
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	host, _ := container.Host(ctx)
	port, _ := container.MappedPort(ctx, "3306")

	dbClient, err := sql.Open("mysql", fmt.Sprintf("root:password@tcp(%s:%s)/test", host, port.Port()))
	if err != nil {
		return nil, err
	}

	if _, err := dbClient.Exec("USE test"); err != nil {
		return nil, fmt.Errorf("error setting database test, error: %v", err)
	}

	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, fmt.Errorf("error creating docker client, error: %v", err)
	}

	return &Mysql{
		Container:    container,
		DockerClient: dockerClient,

		Host:     host,
		Port:     port.Port(),
		Database: "test",
		Username: "root",
		Password: "password",
		Client:   dbClient,
	}, nil
}
