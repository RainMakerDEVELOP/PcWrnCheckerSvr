package reqdataparser

import (
	com_code "PcWrnChecker/PcWrnCheckerSvr/common"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type XmlData struct {
	Writer *http.ResponseWriter
	Reader *http.Request
}

func (data XmlData) ReqDataParse() bool {
	rRestData := com_code.RestData_Xml{}
	rData, err := io.ReadAll(data.Reader.Body)

	if err != nil {
		fmt.Println("xml Body Read Failed")
		http.Error(*data.Writer, err.Error(), http.StatusBadRequest)
		return false
	}

	switch data.Reader.Method {
	case http.MethodGet: // 조회
		err := xml.Unmarshal(rData, &rRestData)

		if err != nil {
			fmt.Printf("GET xml.Unmarshal Error : '%v'\n", err.Error())
			http.Error(*data.Writer, err.Error(), http.StatusBadRequest)
			return false
		}

		// 응답 테스트용
		xml.NewEncoder(*data.Writer).Encode(rRestData) // 테스트용 Echo 데이터 설정

		// Client IP에 해당하는 restData.itemname 의 데이터를 찾는다.

		// 없으면 없음 응답 데이터 리턴

		// 조회된 데이터를 http 응답 데이터에 설정한다.

	case http.MethodPost: // 등록
		err := xml.Unmarshal(rData, &rRestData)

		if err != nil {
			fmt.Printf("POST xml.Unmarshal Error : '%v'\n", err.Error())
			http.Error(*data.Writer, err.Error(), http.StatusBadRequest)
			return false
		}

		// 응답 테스트용
		xml.NewEncoder(*data.Writer).Encode(rRestData) // 테스트용 Echo 데이터 설정

		// Client IP에 해당하는 restData.itemname 의 데이터를 찾는다.

		// 없으면 신규 데이터 등록

		// 있으면 기존 데이터에 추가

	}

	(*data.Writer).WriteHeader(http.StatusOK)
	(*data.Writer).Header().Set("Content-Type", "application/xml")

	return true
}
