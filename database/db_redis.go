package database

import (
	"fmt"

	_ "github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

// singleton 객체
var dbInstance *stRedisDb

type stRedisDb struct {
	connInfo ConnInfo
}

func (db stRedisDb) Init(addr string, port string) error {
	dbInstance.connInfo = ConnInfo{Addr: addr, Port: port}

	fmt.Printf("REDIS connInfo = '%v'\n", db.connInfo)

	err := errors.Errorf("db init error")
	return err
}

func (db stRedisDb) GetDbInstance() *stRedisDb {
	return dbInstance
}
