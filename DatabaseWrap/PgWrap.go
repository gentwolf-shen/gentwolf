package DatabaseWrap

import (
	_ "github.com/lib/pq"
)

type PqWrap struct {
	DbBase
}

func (this *PqWrap) Open(dsn string, maxOpenConnections int, maxIdleConnections int) error {
	return this.OpenDb("postgres", dsn, maxOpenConnections, maxIdleConnections)
}
