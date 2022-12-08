package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

// singleton 객체
var dbInstance *St_Redis_Db

type St_Redis_Db struct {
	connInfo ConnInfo
}

func (db *St_Redis_Db) Init(connInfo ConnInfo) error {
	dbConn, errGetInst := db.GetDbInstance()

	if errGetInst != nil {
		fmt.Println(errGetInst.Error())
		return errGetInst
	}

	dbConn.connInfo = ConnInfo{Addr: connInfo.Addr, Port: connInfo.Port}

	fmt.Printf("REDIS connInfo = '%v'\n", dbConn.connInfo)

	return nil
	// err := errors.Errorf("REDIS DB Init Error")
	// return err
}

func (db *St_Redis_Db) ConnectDB() error {
	dbConn, errGetInst := db.GetDbInstance()

	if errGetInst != nil {
		fmt.Println(errGetInst.Error())
		return errGetInst
	}

	ctx := context.Background()

	connAddrInfo := []string{dbConn.connInfo.Addr, dbConn.connInfo.Port}
	fullconnAddr := strings.Join(connAddrInfo, ":")
	rdb := redis.NewClient(&redis.Options{Addr: fullconnAddr, Password: "", DB: 0})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		fmt.Printf("REDIS ConnectDB Ping Result Check Failed. err = '%v'\n", err.Error())
		return err
	}

	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Printf("REDIS ConnectDB rdb.Set Failed. err = '%v'\n", err.Error())
		return err
	}

	return nil
}

func (db *St_Redis_Db) GetDbInstance() (*St_Redis_Db, error) {
	if dbInstance == nil {
		dbInstance = new(St_Redis_Db)
	}

	if dbInstance == nil {
		err := errors.Errorf("REDIS DB Instance Initialize Failed")
		return nil, err
	}

	return dbInstance, nil
}
