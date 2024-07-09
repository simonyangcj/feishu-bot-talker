package feishu_api

type FeiShuDataPost struct {
	// App ID
	AppID string `json:"app_id"`
	// App Secret
	AppSecret string `json:"app_secret"`
	// Receive ID Type
	ReceiveIdType string `json:"receive_id_type"`
	// content
	Context string `json:"context"`
}

type Token struct {
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int    `json:"expire"`
}
