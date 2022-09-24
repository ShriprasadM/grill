package grillmysql

import (
	"context"
	"database/sql"

	"github.com/swiggy-private/grill/canned"
)

type Mysql struct {
	mysql   *canned.Mysql
	Version string
}

func (gm *Mysql) Start(ctx context.Context) error {
	if gm.Version == "" {
		gm.Version = "5.6"
	}
	mysql, err := canned.NewMysql(ctx, gm.Version)
	if err != nil {
		return err
	}
	gm.mysql = mysql
	return nil
}

func (gm *Mysql) Client() *sql.DB {
	return gm.mysql.Client
}

func (gm *Mysql) Host() string {
	return gm.mysql.Host
}

func (gm *Mysql) Port() string {
	return gm.mysql.Port
}

func (gm *Mysql) Database() string {
	return gm.mysql.Database
}

func (gm *Mysql) Username() string {
	return gm.mysql.Username
}

func (gm *Mysql) Password() string {
	return gm.mysql.Password
}

func (gm *Mysql) Stop(ctx context.Context) error {
	return gm.mysql.Container.Terminate(ctx)
}
