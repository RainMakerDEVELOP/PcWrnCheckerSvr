package main

import (
	pwc_svr_arg "PcWrnChecker/PcWrnCheckerSvr/pwcpkg"
	"bufio"
	"fmt"
	"net"
)

const addr = "localhost:1234"

var m_mapClientInfo map[string]pwc_svr_arg.PwcArg

func proc_connection(conn net.Conn) {
	// conn에 리더(reader)를 설정한다(io.Reader)
	reader := bufio.NewReader(conn)

	remoteAddr := conn.RemoteAddr().String()

	fmt.Printf("remoteAddr : %s\n", remoteAddr)

	// 해당 주소가 목록에 있는지 조사
	vClientInfo, ok := m_mapClientInfo[remoteAddr]

	if ok == true { // 해당 주소가 목록에 있으면
	} else { // 해당 주소가 목록에 없으면
		// 읽어온 데이터의 첫 줄을 가져온다.
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("ReadString 에러 : %s\n", err.Error())
			return
		}

		// 해당 주소의 모니터링 정보에 모니터링하고자 하는 항목이 있는지 조사
		vItem, ok := vClientInfo.ExistClient(data)

		if ok == true { // 해당 항목이 있으면, 항목의 모니터링 값을 추가
			fmt.Printf("ExistClient vItem.StartTime = '%s'\n", vItem.StartTime.String())
		} else { // 해당 항목이 없으면 추가
			addRet := vClientInfo.AddClient(data)

			fmt.Printf("AddClient Ret = '%d'\n", addRet)
		}
		// 출력한 다음 데이터를 다시 보낸다.
		// fmt.Printf("Received : %s\n", data)
		// conn.Write([]byte(strings.ToUpper(data)))
		// m_mapClientInfo[remoteAddr] =
	}
}

func main() {
	fmt.Println("----------------------------------")
	fmt.Println("---------- Server Start ----------")
	fmt.Println("----------------------------------")

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	defer ln.Close()
	fmt.Printf("listening on : %s\n", addr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Accept 에러 : %s\n", err.Error())

			// 오류가 있으면 다시 시도한다
			continue
		}

		// 이 작업을 비동기로 처리하면 잠재적으로 워커 풀을 위해 좋은 사용 사례가 될 것이다
		go proc_connection(conn)
	}

	fmt.Println("----------------------------------")
	fmt.Println("---------- Server End   ----------")
	fmt.Println("----------------------------------")
}
