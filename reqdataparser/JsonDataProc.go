package reqdataparser

import (
	com_code "PcWrnChecker/PcWrnCheckerSvr/common"
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonData struct {
	Writer http.ResponseWriter
	Reader http.Request
}

func (data JsonData) ReqDataParse() (bool, com_code.RestData_Common) {
	jsonRestData := com_code.RestData_Json{}
	retRestData := com_code.RestData_Common{}

	switch data.Reader.Method {
	case http.MethodGet: // 조회
		err := json.NewDecoder(data.Reader.Body).Decode(&jsonRestData)

		if err != nil {
			fmt.Printf("GET json.NewDecoder Error : '%v'\n", err.Error())
			http.Error(data.Writer, err.Error(), http.StatusBadRequest)
			return false, retRestData
		}

		retRestData.ItemName = jsonRestData.ItemName
		retRestData.Value = jsonRestData.Value

		// 응답 테스트용
		// json.NewEncoder(w).Encode("GET")
		json.NewEncoder(data.Writer).Encode(jsonRestData) // 테스트용 Echo 데이터 설정

		// Client IP에 해당하는 restData.itemname 의 데이터를 찾는다.

		// 없으면 없음 응답 데이터 리턴

		// 조회된 데이터를 http 응답 데이터에 설정한다.

	case http.MethodPost: // 등록
		err := json.NewDecoder(data.Reader.Body).Decode(&jsonRestData)

		if err != nil {
			fmt.Printf("POST json.NewDecoder Error : '%v'\n", err.Error())
			http.Error(data.Writer, err.Error(), http.StatusBadRequest)
			return false, retRestData
		}

		retRestData.ItemName = jsonRestData.ItemName
		retRestData.Value = jsonRestData.Value

		// 응답 테스트용
		// json.NewEncoder(w).Encode("POST")
		json.NewEncoder(data.Writer).Encode(jsonRestData) // 테스트용 Echo 데이터 설정

		// Client IP에 해당하는 restData.itemname 의 데이터를 찾는다.

		// 없으면 신규 데이터 등록

		// 있으면 기존 데이터에 추가

	}

	(data.Writer).WriteHeader(http.StatusOK)
	(data.Writer).Header().Set("Content-Type", "application/json")

	return true, retRestData
}

// func (data JsonData) SetMapClientInfo(mapClientInfo *map[string]*pwc_svr_arg.PwcArg) {
// 	m_mapClientInfo = *mapClientInfo
// }

// func (data JsonData) GetMapClientInfo() *map[string]*pwc_svr_arg.PwcArg {
// 	return &m_mapClientInfo
// }
