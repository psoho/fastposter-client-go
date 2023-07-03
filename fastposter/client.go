package fastposter

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

const (
	CLIENT_VERSION = "1.0.0"
	CLIENT_TYPE    = "go"
	ENDPOINT       = "https://api.fastposter.net"
)

type Poster struct {
	TraceID string
	Uuid    string
	Type    string
	Bytes   []byte
	Size    int
}

func (p *Poster) SaveTo(path string) error {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		fmt.Println("Failed to create directories:", err)
		return err
	}
	err = ioutil.WriteFile(path, p.Bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (p *Poster) Save() (string, error) {
	path := p.Uuid + "." + p.Type
	err := p.SaveTo(path)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (p *Poster) B64String() string {
	return base64.StdEncoding.EncodeToString(p.Bytes)
}

type _Client struct {
	Token    string
	Endpoint string
}

func Client(token string) *_Client {
	return &_Client{
		Token:    token,
		Endpoint: ENDPOINT,
	}
}

func ClientWithEndpoint(token string, endpoint *string) *_Client {
	return &_Client{
		Token:    token,
		Endpoint: *endpoint,
	}
}

func (c *_Client) BuildPoster(uuid string, params map[string]interface{}, posterType string) (*Poster, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	payloadB64 := base64.StdEncoding.EncodeToString(payload)

	body := map[string]interface{}{
		"uuid":    uuid,
		"payload": payloadB64,
		"type":    posterType,
	}

	//
	//if scale != 1.0 {
	//	body["scale"] = scale
	//}

	url := c.Endpoint + "/v1/build/poster"

	requestBody, err := json.Marshal(body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Client-Type", CLIENT_TYPE)
	req.Header.Set("Client-Version", CLIENT_VERSION)
	req.Header.Set("token", c.Token)
	//req.Header.Set("User-Agent", userAgent)
	req.Header.Set("cache-control", "no-cache")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}

	traceID := resp.Header.Get("fastposter-traceid")

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	poster := &Poster{
		TraceID: traceID,
		Uuid:    uuid,
		Type:    posterType,
		Bytes:   content,
		Size:    len(content),
	}

	return poster, nil
}
