package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type PostData struct {
	URL string
	Key string
}
type Response struct {
	Status string
	Header http.Header
	Body   *SchemaResponse
}

type SchemaQuery struct {
	Query string `json:"query"`
}

type SchemaResponse struct {
	Data   interface{} `json:"data"`
	Errors []struct {
		Message  string `json:"message"`
		Location []struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"location"`
	} `json:"errors"`
}

func (s *SchemaResponse) ConvertData(t interface{}) error {
	m, err := json.Marshal(s.Data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(m, &t)
	if err != nil {
		return err
	}
	return nil
}

func NewPost(url string) *PostData {
	return &PostData{
		URL: url,
	}
}

func (p *PostData) Send(body string) (Response, error) {
	respBody := Response{
		Body: &SchemaResponse{},
	}
	jsonBody, err := json.Marshal(SchemaQuery{
		Query: body,
	})
	if err != nil {
		return respBody, err
	}

	req, errReq := http.NewRequest("POST", p.URL, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	if errReq != nil {
		return respBody, errReq
	}

	client := &http.Client{}
	resp, errSend := client.Do(req)
	if errSend != nil {
		return respBody, errSend
	}
	defer resp.Body.Close()

	respBody.Status = resp.Status
	respBody.Header = resp.Header

	bodyByte, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return respBody, errRead
	}

	errUnMarshal := json.Unmarshal(bodyByte, &respBody.Body)
	if errUnMarshal != nil {
		return respBody, errUnMarshal
	}

	return respBody, nil
}
