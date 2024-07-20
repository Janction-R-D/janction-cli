package types

type NonceRes struct {
	Data struct {
		Nonce string `json:"nonce"`
	}
}

type LoginReq struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
	IsNode    bool   `json:"is_node"`
}

type LoginResp struct {
	Data struct {
		Token string `json:"token"`
	}
}
