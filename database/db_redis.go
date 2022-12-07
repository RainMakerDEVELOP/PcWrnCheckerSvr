package database

import (
	_ "github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

func Init(addr string, port string) error {
	DbInstance.Addr = addr
	DbInstance.Port = port

	err := errors.Errorf("db init error")
	return err
}
