package model

import "time"

type Respon struct {
	Status string `json:"status"`
	No     string `json:"no"`
}
type RespGlobal struct {
	Status         string      `json:"status"`
	StatusDesc     string      `json:"status_desc"`
	StatusDateTime time.Time   `json:"status_date_time"`
	Result         interface{} `json:"result"`
}
type RespReqSend struct {
	No    string `json:"no"`
	Text  string `json:"text"`
	Index string `json:"index"`
}
