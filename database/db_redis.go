package database

import (
	"fmt"

	_ "github.com/go-redis/redis/v8"
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
