package DatabaseWrap

import (
	_ "github.com/Go-SQL-Driver/MySQL"
)

type MySqlWrap struct {
	DbBase
}

func (this *MySqlWrap) Open(dsn string, maxOpenConnections int, maxIdleConnections int) error {
	return this.OpenDb("mysql", dsn, maxOpenConnections, maxIdleConnections)
}
