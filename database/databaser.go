package database

type ConnInfo struct {
	Addr   string
	Port   string
	DbKind string
	DbVer  string
}

// // singleton 객체
// var DbInstance *databaser

type Databaser interface {
	Init(ConnInfo) error
	// GetDbInstance() *databaser
	ConnectDB() error
}
