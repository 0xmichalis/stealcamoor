package stealcamoor

import (
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/0xmichalis/stealcamoor/pkg/abis"
	email "github.com/0xmichalis/stealcamoor/pkg/client/email"
	"github.com/0xmichalis/stealcamoor/pkg/client/stealcam"
)

type Stealcamoor struct {
	/* Misc config */
	backupDir string
	creators  []common.Address

	/* Email-related config */
	emailClient *email.EmailClient

	/* Blockchain-related config */
	stealcamContract *abis.Stealcam
	client           *ethclient.Client
	explorerURL      string
	stealcamAddress  common.Address
	ourAddress       common.Address
	ourSignature     string
	txOpts           *bind.TransactOpts

	/* API-related config */
	apiClient          *stealcam.ApiClient
	apiRequestInterval time.Duration
	addressToTwitter   map[string]string
}

func (sc *Stealcamoor) Start() {
	sc.startApiListener()
}
