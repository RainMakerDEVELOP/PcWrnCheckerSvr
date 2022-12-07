package database

type ConnInfo struct {
	Addr string
	Port string
}

// // singleton 객체
// var DbInstance *databaser

type databaser interface {
	Init(string, string) error
	// GetDbInstance() *databaser
}
