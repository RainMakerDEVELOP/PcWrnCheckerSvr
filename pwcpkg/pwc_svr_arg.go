package pwc_svr_arg

import (
	"time"
)

// 모니터링할 PC 정보 상세( CPU 점유율, RAM 사용량 등)
type PwcArg_Item struct {
	StartTime time.Time
	EndTime   time.Time
}

type PwcArg struct {
	mapMon_Item map[string]PwcArg_Item
}

func (pa PwcArg) AddClient(itemName string) bool {
	var pi PwcArg_Item
	pi.StartTime = time.Now() // 신규 추가시에는 시작 시간만 기록한다
	pa.mapMon_Item[itemName] = pi

	return true
}

func (pa PwcArg) ExistClient(itemName string) (PwcArg_Item, bool) {
	vItemMap, ok := pa.mapMon_Item[itemName]

	return vItemMap, ok
}
