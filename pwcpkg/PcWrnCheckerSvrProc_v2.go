// 해당 패키지에 큰 변동이 있을 경우,
// 1. 신규 go 파일을 버전명으로 하나 더 생성
// 2. 기존 go 파일의 확장자에 .bak 를 추가하여 미사용 처리

package svrproc

import (
	pwc_svr_arg "PcWrnChecker/PcWrnCheckerSvr/pwcpkg/arg"
	"bufio"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
)

// const addr = "localhost:1234"
const port = ":1234"

var m_mapClientInfo map[string]*pwc_svr_arg.PwcArg

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

	if ok == true { // 해당 주소가 목록에 있으면
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

		if ok == true { // 해당 항목이 있으면, 항목의 모니터링 값을 추가
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

// getIP returns the ip address from the http request
func getIP(r *http.Request) (string, error) {
	ips := r.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")

	if len(splitIps) > 0 {
		// get last IP in list since ELB prepends other user defined IPs, meaning the last one is the actual client IP.
		netIP := net.ParseIP(splitIps[len(splitIps)-1])
		if netIP != nil {
			return netIP.String(), nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	netIP := net.ParseIP(ip)
	if netIP != nil {
		ip := netIP.String()
		if ip == "::1" {
			return "127.0.0.1", nil
		}
		return ip, nil
	}

	return "", errors.New("IP not found")
}

// v1 : 소켓 서버 방식
// v2 : HTTP 서버 방식
func Run() {
	m_mapClientInfo = make(map[string]*pwc_svr_arg.PwcArg)

	// CPU 사용량
	http.HandleFunc("/USEDCPU", UsedCpuHandler)

	http.ListenAndServe(port, nil)
}

func UsedCpuHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Enter UsedCpuHandler Function")

	ip, err := getIP(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch req.Method {
	case http.MethodGet: // 조회
		break
	case http.MethodPost: // 등록
		break
	}

	fmt.Printf("ClientIP : %v\n", ip)

	fmt.Println("Proc End UsedCpuHandler Function")
}
