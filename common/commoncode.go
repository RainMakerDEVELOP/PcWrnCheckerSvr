package commoncode

type RestData_Json struct {
	ItemName string `json:"itemname"`
	Value    string `json:"value"`
}

type RestData_Xml struct {
	ItemName string
	Value    string
}

type RestData_Common struct {
	ItemName string
	Value    string
}
