package stealcamoor

import (
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	email "github.com/0xmichalis/stealcamoor/pkg/email"
	"github.com/0xmichalis/stealcamoor/pkg/stealcamapi"
)

type Stealcamoor struct {
	creators []common.Address

	/* Email-related config */
	emailClient *email.EmailClient

	/* Blockchain-related config */
	client          *ethclient.Client
	explorerURL     string
	stealcamAddress common.Address
	txOpts          *bind.TransactOpts

	/* Backend-related config */
	apiClient          *stealcamapi.ApiClient
	apiRequestInterval time.Duration
}

func (sc *Stealcamoor) Start() {
	sc.startApiListener()
}
