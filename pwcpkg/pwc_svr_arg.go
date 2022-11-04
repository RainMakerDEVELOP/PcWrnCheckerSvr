package pwc_svr_arg

import (
	"time"
)

// 모니터링할 PC 정보 상세( CPU 점유율, RAM 사용량 등)
type PwcArg_Item struct {
	ItemName  string
	StartTime time.Time
	EndTime   time.Time
}

type PwcArg struct {
	ClientIPAddr    string
	Monitoring_Item []PwcArg_Item
}

func AddClient(string) bool {
	return false
}

func ExistClient(string) bool {
	return false
}
