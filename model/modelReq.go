package model

type User struct {
	No    string `json:"no"`
	Text  string `json:"text"`
	Index string `json:"index"`
}
type ReqInquiry struct {
	User   string `json:"user"`
	IDSend string `json:"id_send"`
}
