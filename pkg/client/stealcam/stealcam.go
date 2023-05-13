package stealcam

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

type Memory struct {
	ID        uint64
	Owner     *string
	Signature string
}

type ApiClient struct {
	baseURL string
	c       *http.Client
}

func New(baseURL string) *ApiClient {
	return &ApiClient{
		baseURL: baseURL,
		c:       &http.Client{},
	}
}

func (a ApiClient) GetMemories(creator common.Address) ([]Memory, error) {
	resp, err := a.c.Get(a.baseURL + "/memories/created/" + creator.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	memories := new([]Memory)
	if err := json.Unmarshal(body, memories); err != nil {
		return nil, err
	}

	return *memories, nil
}

func (a ApiClient) GetMemory(id uint64) (*Memory, error) {
	resp, err := a.c.Get(fmt.Sprintf("%s/memories/%d", a.baseURL, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	m := &Memory{}
	if err := json.Unmarshal(body, m); err != nil {
		return nil, err
	}

	// drop the 0x prefix
	m.Signature = m.Signature[2:]

	return m, nil
}

func (a ApiClient) RevealMemory(id uint64, address common.Address, signature string) (string, error) {
	reqBody := struct {
		Address   string `json:"address"`
		Signature string `json:"signature"`
	}{
		Address:   address.String(),
		Signature: signature,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/memories/%d", a.baseURL, id), bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// return the URL to the stolen image
	r := struct {
		MediaUrl string
	}{}
	if err := json.Unmarshal(respBody, &r); err != nil {
		return "", err
	}

	return r.MediaUrl, nil
}
