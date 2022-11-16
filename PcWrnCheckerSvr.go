package main

import (
	svrproc "PcWrnChecker/PcWrnCheckerSvr/pwcpkg"
	"fmt"
)

func main() {
	fmt.Println("----------------------------------")
	fmt.Println("---------- Server Start ----------")
	fmt.Println("----------------------------------")

	// 실제 모든 데이터 처리를 PcWrnCheckerSvrProc 으로 명명된 svrproc 패키지에서 하도록 한다.
	svrproc.Run()

	fmt.Println("----------------------------------")
	fmt.Println("---------- Server End   ----------")
	fmt.Println("----------------------------------")
}
