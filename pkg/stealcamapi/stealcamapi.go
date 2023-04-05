package stealcamapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

type Memory struct {
	ID    uint64
	Owner *string
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

func (a ApiClient) GetMemory(id int) (*Memory, error) {
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

	return m, nil
}
