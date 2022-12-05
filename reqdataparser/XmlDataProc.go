package reqdataparser

import (
	com_code "PcWrnChecker/PcWrnCheckerSvr/common"
	"net/http"
)

type XmlData struct {
	Writer http.ResponseWriter
	Reader *http.Request
}

func (xd XmlData) ReqDataParse() (bool, com_code.RestData) {
	retRestData := com_code.RestData{}

	return true, retRestData
}
