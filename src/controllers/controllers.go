package controllers

import "encoding/json"

type Export struct{}

type ResponseStream struct {
	Status   string `json:"status"`
	Messages string `json:"messages"`
	Data     any    `json:"data"`
}

func (e Export) Response(d any, s string, m string) ([]byte, error) {

	data := &ResponseStream{
		Status:   s,
		Messages: m,
		Data:     d,
	}

	response, _ := json.Marshal(data)

	return response, nil
}
