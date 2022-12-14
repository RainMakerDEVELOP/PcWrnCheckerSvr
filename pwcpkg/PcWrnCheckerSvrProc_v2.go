// 해당 패키지에 큰 변동이 있을 경우,
// 1. 신규 go 파일을 버전명으로 하나 더 생성
// 2. 기존 go 파일의 확장자에 .bak 를 추가하여 미사용 처리

package svrproc

import (
	commoncode "PcWrnChecker/PcWrnCheckerSvr/common"
	"PcWrnChecker/PcWrnCheckerSvr/config"
	"PcWrnChecker/PcWrnCheckerSvr/database"
	pwc_svr_arg "PcWrnChecker/PcWrnCheckerSvr/pwcpkg/arg"
	reqdataparser "PcWrnChecker/PcWrnCheckerSvr/reqdataparser"
	"PcWrnChecker/PcWrnCheckerSvr/util"
	"fmt"
	"net/http"
	"strings"
)

type HttpDataParser interface {
	ReqDataParse() (bool, commoncode.RestData_Common)
	// SetMapClientInfo(*map[string]*pwc_svr_arg.PwcArg)
	// GetMapClientInfo() *map[string]*pwc_svr_arg.PwcArg
}

// func httpDataParse(w *http.ResponseWriter, r *http.Request, hdp HttpDataParser) (bool, com_code.RestData) {
// 	return hdp.ReqDataParse()
// }

// const addr = "localhost:1234"
const port = ":1234"

var m_mapClientInfo map[string]*pwc_svr_arg.PwcArg
var m_config config.Config
var m_dbConn database.Databaser

/*
func proc_connection(conn net.Conn) {
	// conn에 리더(reader)를 설정한다(io.Reader)
	reader := bufio.NewReader(conn)

	remoteFullAddr := conn.RemoteAddr().String()

	fmt.Printf("remoteFullAddr : %s\n", remoteFullAddr)

	nColonIdx := strings.Index(remoteFullAddr, ":")

	var remoteIpAddr string

	if nColonIdx >= 0 {
		if beforeFullAddr, _, ok := strings.Cut(remoteFullAddr, ":"); ok {
			remoteIpAddr = beforeFullAddr
		} else {
			remoteIpAddr = beforeFullAddr
		}

	}

	// 해당 주소가 목록에 있는지 조사
	_, ok := m_mapClientInfo[remoteIpAddr]

	if ok { // 해당 주소가 목록에 있으면
		fmt.Printf("이미 맵에 존재하는 클라이언트 주소 (%v)\n", remoteIpAddr)
	} else { // 해당 주소가 목록에 없으면
		// 읽어온 데이터의 첫 줄을 가져온다.
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("ReadString 에러 : %s\n", err.Error())
			return
		}

		// 여기에 m_mapClientInfo[] 의 Value 에 해당하는 map 을 생성하는 루틴 필요할 듯 (2022.11.11)
		m_mapClientInfo[remoteIpAddr] = &pwc_svr_arg.PwcArg{}

		m_mapClientInfo[remoteIpAddr].ConstructMap()
		// vClientInfo.ConstructMap()
		// m_mapClientInfo[remoteIpAddr] = vClientInfo

		// 해당 주소의 모니터링 정보에 모니터링하고자 하는 항목이 있는지 조사
		vItem, ok := m_mapClientInfo[remoteIpAddr].ExistClient(data)

		if ok  { // 해당 항목이 있으면, 항목의 모니터링 값을 추가
			fmt.Printf("ExistClient vItem.StartTime = '%s'\n", vItem.StartTime.String())
		} else { // 해당 항목이 없으면 추가
			addRet := m_mapClientInfo[remoteIpAddr].AddClient(data)

			fmt.Printf("AddClient Ret = '%v'\n", addRet)
		}
		// 출력한 다음 데이터를 다시 보낸다.
		// fmt.Printf("Received : %s\n", data)
		// conn.Write([]byte(strings.ToUpper(data)))
		// m_mapClientInfo[remoteIpAddr] =
	}
}
*/

// v1 : 소켓 서버 방식
// v2 : HTTP 서버 방식
func Run() {
	m_mapClientInfo = make(map[string]*pwc_svr_arg.PwcArg)

	// 환경설정을 읽어온다.
	m_config = config.Config{DbAddr: "localhost", DbPort: "3389"}
	err := m_config.ReadConfig()
	if err != nil {
		fmt.Printf("ReadConfig Error. err = '%v'\n", err.Error())
		return
	}

	// DB연결
	connInfo := database.ConnInfo{Addr: m_config.DbAddr, Port: m_config.DbPort, DbKind: m_config.DbKind, DbVer: m_config.DbVer}

	if strings.Compare(connInfo.DbKind, "REDIS") == 0 {
		m_dbConn = &database.St_Redis_Db{}
	} else if strings.Compare(connInfo.DbKind, "MSSQL") == 0 {
		m_dbConn = &database.St_Redis_Db{}
	} else {
		fmt.Printf("Not Supported DBKind = '%v'\n", connInfo.DbKind)
		return
	}

	errDbInit := m_dbConn.Init(connInfo) // 초기화
	if errDbInit != nil {
		fmt.Printf("DB Initialize Error. err = '%v'\n", errDbInit.Error())
		return
	}

	errDbConn := m_dbConn.ConnectDB() // DB 연결
	if errDbConn != nil {
		fmt.Printf("DB Connect Error. err = '%v'\n", errDbConn.Error())
	}

	// CPU 사용량
	http.HandleFunc("/USEDCPU", UsedCpuHandler)

	http.ListenAndServe(port, nil)
}

func UsedCpuHandler(w http.ResponseWriter, r *http.Request) {
	// 요청 데이터의 ClientIP 확인
	ip, err := util.GetIP(r)
	if err != nil {
		fmt.Printf("getIP Error : '%v'\n", err.Error())
		http.Error(w, "Getting ClientIP Failed", http.StatusInternalServerError)
		return
	}
	fmt.Printf("ClientIP : %v\n", ip)

	writer := w
	reader := *r

	done := make(chan string) // http 연결이 유실되지 않도록 채널을 사용
	go proc_UsedCpuHandler(writer, reader, ip, done)
	fmt.Printf("return done = '%v'\n", <-done)
}

// USEDCPU 핸들러 데이터 처리 함수
func proc_UsedCpuHandler(w http.ResponseWriter, r http.Request, ip string, done chan<- string) {
	fmt.Println("----------- START proc_UsedCpuHandler")

	defer fmt.Println("----------- END proc_UsedCpuHandler")

	var parser HttpDataParser

	// Content-Type 별 처리
	if strings.Compare(r.Header.Get("Content-Type"), "application/json") == 0 {
		// JSON 인 경우 처리
		parser = reqdataparser.JsonData{Writer: w, Reader: r}
	} else if strings.Compare(r.Header.Get("Content-Type"), "application/xml") == 0 {
		// XML 인 경우 처리
		parser = reqdataparser.XmlData{Writer: w, Reader: r}
	} else {
		// 허용되지 않은 데이터 타입의 경우 처리
		strLog := "return! Request Data Content-Type is '" + r.Header.Get("Content-Type") + "'"
		fmt.Println(strLog)

		fmt.Println("Content-Type Error")
		http.Error(w, "Content-Type Error", http.StatusBadRequest)

		done <- ip
		return
	}

	// 2022.12.06 Content-Type 에 따라 파싱될 수 있도록 인터페이스로 처리
	bRet, restData := parser.ReqDataParse()

	fmt.Printf("ReqDataParse Ret = '%v', restData = '%v'\n", bRet, restData)

	nColonIdx := strings.Index(ip, ":")

	var remoteIpAddr string

	if nColonIdx >= 0 {
		if beforeFullAddr, _, ok := strings.Cut(ip, ":"); ok {
			remoteIpAddr = beforeFullAddr
		} else {
			remoteIpAddr = beforeFullAddr
		}
	} else {
		remoteIpAddr = ip
	}

	if len(remoteIpAddr) <= 0 {
		fmt.Println("IP 주소 조사 실패")
		done <- ip
		return
	}

	// 해당 주소가 목록에 있는지 조사
	_, ok := m_mapClientInfo[remoteIpAddr]

	if ok { // 해당 주소가 목록에 있으면
		fmt.Printf("이미 맵에 존재하는 클라이언트 주소 (%v)\n", remoteIpAddr)
	} else { // 해당 주소가 목록에 없으면
		// 여기에 m_mapClientInfo[] 의 Value 에 해당하는 map 을 생성하는 루틴 필요할 듯 (2022.11.11)
		m_mapClientInfo[remoteIpAddr] = &pwc_svr_arg.PwcArg{}

		m_mapClientInfo[remoteIpAddr].ConstructMap()

		// 해당 주소의 모니터링 정보에 모니터링하고자 하는 항목이 있는지 조사
		vItem, ok := m_mapClientInfo[remoteIpAddr].ExistClient(restData.ItemName)

		if ok { // 해당 항목이 있으면, 항목의 모니터링 값을 추가
			fmt.Printf("ExistClient vItem.StartTime = '%s'\n", vItem.StartTime.String())
		} else { // 해당 항목이 없으면 추가
			addRet := m_mapClientInfo[remoteIpAddr].AddClient(restData.ItemName)

			fmt.Printf("AddClient Ret = '%v'\n", addRet)
		}
	}

	// [PROC] DB 에 데이터 ISNERT, UPDATE 처리 필요

	done <- ip
}
