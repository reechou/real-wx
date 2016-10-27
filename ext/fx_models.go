package ext

type CreateFxAccountReq struct {
	UnionId   string `json:"unionId"`
	WXAccount string `json:"wxAccount"`
	OpenId    string `json:"openId"`
	Name      string `json:"name"`
	Superior  string `json:"superiorId"`
}

type FxResponse struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
