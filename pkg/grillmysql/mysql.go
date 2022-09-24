package grillmysql

import (
	"context"
	"database/sql"

	"github.com/swiggy-private/grill/canned"
)

type Mysql struct {
	mysql   *canned.Mysql
	version string
}

func (gm *Mysql) Start(ctx context.Context) error {
	var mysql = new(canned.Mysql)
	var err error
	if gm.version == "" {
		mysql, err = canned.NewMysql(ctx)
	} else {
		mysql, err = canned.NewMysql8(ctx)
	}
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
