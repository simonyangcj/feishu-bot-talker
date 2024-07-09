package feishu_api

import "encoding/json"

type ContentRender interface {
	Render([]byte) ([]byte, error)
}

type ContentRenderText struct {
	Text string `json:"text"`
}

func (content *ContentRenderText) Render(value []byte) ([]byte, error) {
	err := json.Unmarshal(value, content)
	if err != nil {
		return nil, err
	}
	return json.Marshal(content)
}
