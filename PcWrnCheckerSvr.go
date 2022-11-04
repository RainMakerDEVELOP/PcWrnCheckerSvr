package main

import (
	"fmt"

	"pwcpkg/svr/pwc_svr_arg"
)

func main() {
	fmt.Println("-------------------------------------------------")
	fmt.Println("---------- Pc Wrn Checker Server Start ----------")
	fmt.Println("-------------------------------------------------")

	bRet := pwc_svr_arg.AddClient("a")
	fmt.Println(bRet)

	fmt.Println("-------------------------------------------------")
	fmt.Println("---------- Pc Wrn Checker Server End   ----------")
	fmt.Println("-------------------------------------------------")
}
