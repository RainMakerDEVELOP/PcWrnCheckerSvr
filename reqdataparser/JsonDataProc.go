package reqdataparser

import (
	com_code "PcWrnChecker/PcWrnCheckerSvr/common"
	"net/http"
)

type JsonData struct {
	Writer http.ResponseWriter
	Reader *http.Request
}

func (jd JsonData) ReqDataParse() (bool, com_code.RestData) {
	retRestData := com_code.RestData{}

	return true, retRestData
}
