package DatabaseWrap

import (
	_ "github.com/go-sql-driver/mysql"
)

type MySqlWrap struct {
	DbBase
}

func (this *MySqlWrap) Open(dsn string, maxOpenConnections int, maxIdleConnections int) error {
	return this.OpenDb("mysql", dsn, maxOpenConnections, maxIdleConnections)
}
