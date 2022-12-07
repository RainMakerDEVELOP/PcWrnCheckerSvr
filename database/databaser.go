package database

type ConnInfo struct {
	Addr string
	Port string
}

// singleton 객체
var DbInstance *ConnInfo

type databaser interface {
	Init(string, string) error
	GetDbInstance() *ConnInfo
}
