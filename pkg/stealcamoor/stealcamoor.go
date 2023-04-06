package stealcamoor

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/0xmichalis/stealcamoor/pkg/abis"
	email "github.com/0xmichalis/stealcamoor/pkg/client/email"
	"github.com/0xmichalis/stealcamoor/pkg/client/stealcam"
)

type Stealcamoor struct {
	creators []common.Address

	/* Email-related config */
	emailClient *email.EmailClient
	// This cache is used to keep track of ids for which
	// an email has already been sent. Since it's just an
	// in-memory cache, this means that service restarts
	// may result in resending already sent emails.
	emailCacheLock *sync.Mutex
	emailCache     map[uint64]bool
	// where to send emails to
	to string

	/* Blockchain-related config */
	mintCacheLock    *sync.Mutex
	mintCache        map[uint64]bool
	stealcamContract *abis.Stealcam
	client           *ethclient.Client
	explorerURL      string
	stealcamAddress  common.Address
	ourAddress       common.Address
	txOpts           *bind.TransactOpts

	/* Backend-related config */
	apiClient          *stealcam.ApiClient
	apiRequestInterval time.Duration
}

func (sc *Stealcamoor) Start() {
	go sc.startChainListener()
	sc.startApiListener()
}
