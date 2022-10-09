package gentwolf

import (
	"github.com/gentwolf-shen/gentwolf/DatabaseWrap"
)

type dbFactory struct {
	drivers map[string]*DatabaseWrap.DbBase
}

var DbFactory *dbFactory

func init() {
	DbFactory = &dbFactory{}
}

func (this *dbFactory) Init(configs map[string]DbConfig) error {
	this.drivers = make(map[string]*DatabaseWrap.DbBase, len(configs))

	for k, cfg := range configs {
		tmp := DatabaseWrap.DbBase{}
		err := tmp.OpenDb(cfg.Type, cfg.Dsn, cfg.MaxOpenConnections, cfg.MaxIdleConnections)
		if err != nil {
			return err
		} else {
			this.drivers[k] = &tmp
		}
	}
	return nil
}

func (this *dbFactory) Driver(key string) *DatabaseWrap.DbBase {
	return this.drivers[key]
}

func (this *dbFactory) Close(key string) {
	this.drivers[key].Close()
}

func (this *dbFactory) CloseAll() {
	for key, _ := range this.drivers {
		this.drivers[key].Close()
	}
}
