package stealcamapi

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

type Memory struct {
	id string
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

func (a ApiClient) getMemories(creator common.Address) ([]Memory, error) {
	resp, err := a.c.Get(a.baseURL + "/memories/created/" + creator.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(body)

	return nil, nil
}

func (a ApiClient) getMemory(id string) (*Memory, error) {
	resp, err := a.c.Get(a.baseURL + "/memories/" + id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(body)

	return nil, nil
}
