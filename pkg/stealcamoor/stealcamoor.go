package stealcamoor

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/0xmichalis/stealcamoor/pkg/stealcamapi"
	stealcamutils "github.com/0xmichalis/stealcamoor/pkg/stealcamapi/utils"
)

type Stealcamoor struct {
	creators []common.Address

	/* Blockchain-related config */

	// Node connection
	client          *ethclient.Client
	explorerURL     string
	stealcamAddress common.Address

	// TODO: Figure out whether it is faster to always
	// instantiate this vs deep-copying to avoid mutations
	// or whether we don't care about mutations as these
	// will always be in specific fields, ie., gas stuff
	TxOpts *bind.TransactOpts

	/* Backend-related config */

	apiClient          *stealcamapi.ApiClient
	apiRequestInterval time.Duration
	apiURL             string
}

var (
	noOpts = new(bind.CallOpts)
	zero   = big.NewInt(0)
)

func New() (*Stealcamoor, error) {
	sc := &Stealcamoor{}

	// Run validations
	if err := sc.validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	// Connect to node
	// TODO: Make timeout configurable
	client, err := ethclient.Dial(os.Getenv("NODE_API_URL"))
	if err != nil {
		return nil, fmt.Errorf("cannot connect to node: %w", err)
	}
	sc.client = client

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("cannot get chain id: %w", err)
	}
	fmt.Println("Chain ID:", chainID)

	// Load private key
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return nil, fmt.Errorf("cannot load private key: %w", err)
	}

	// Extract address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot cast public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("Our address: %s/address/%s\n", sc.explorerURL, address)

	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("cannot create authorized transactor: %w", err)
	}
	sc.TxOpts = txOpts

	sc.apiClient = stealcamapi.New(sc.apiURL)

	return sc, nil
}

func (sc *Stealcamoor) validate() error {
	creatorStrings := strings.Split(os.Getenv("CREATORS"), ",")
	creators := make([]common.Address, 0)
	fmt.Printf("Tracking %d creators:\n", len(creatorStrings))
	for _, creator := range creatorStrings {
		fmt.Println(creator)
		creators = append(creators, common.HexToAddress(creator))
	}
	if len(creators) == 0 {
		return errors.New("Need at least one creator provided in CREATORS (comma-separated list)")
	}
	sc.creators = creators

	explorerURL := os.Getenv("BLOCKCHAIN_EXPLORER_URL")
	if explorerURL == "" {
		return errors.New("BLOCKCHAIN_EXPLORER_URL cannot be empty")
	}
	sc.explorerURL = explorerURL

	apiURL := os.Getenv("STEALCAM_API_URL")
	if apiURL == "" {
		return errors.New("STEALCAM_API_URL cannot be empty")
	}
	sc.apiURL = apiURL

	if os.Getenv("STEALCAM_API_REQUEST_INTERVAL") == "" {
		return errors.New("STEALCAM_API_REQUEST_INTERVAL cannot be empty")
	}
	apiRequestInterval, err := time.ParseDuration(os.Getenv("STEALCAM_API_REQUEST_INTERVAL"))
	if err != nil {
		return err
	}
	sc.apiRequestInterval = apiRequestInterval

	stealcamAddress := os.Getenv("STEALCAM_ADDRESS")
	if stealcamAddress == "" {
		return errors.New("STEALCAM_ADDRESS cannot be empty")
	}
	sc.stealcamAddress = common.HexToAddress(stealcamAddress)
	fmt.Printf("Stealcam address: %s/address/%s\n", sc.explorerURL, sc.stealcamAddress.String())

	if os.Getenv("PRIVATE_KEY") == "" {
		return errors.New("PRIVATE_KEY cannot be empty")
	}

	if os.Getenv("NODE_API_URL") == "" {
		return errors.New("NODE_API_URL cannot be empty")
	}

	return nil
}

func (sc *Stealcamoor) startChainListener() {
	headers := make(chan *types.Header)
	sub, err := sc.client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatalf("Failed to subscribe to headers: %v", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Got subscription error: %v", err)

		case header := <-headers:
			log.Printf("Processing block %d", header.Number.Uint64())

		}
	}
}

func (sc *Stealcamoor) startApiListener() {
	tick := time.Tick(sc.apiRequestInterval)
	for {
		select {
		case <-tick:
			fmt.Printf("Producer: TICK %v\n", time.Now())
			for _, creator := range sc.creators {
				memories, err := sc.apiClient.GetMemories(creator)
				if err != nil {
					log.Println(err)
					continue
				}
				unminted := stealcamutils.FilterUnmintedMemories(memories)
				if len(unminted) == 0 {
					fmt.Printf("No unminted memories for %s\n", creator)
					continue
				}
				for _, id := range unminted {
					fmt.Println("Unminted memory id %s for %s\n", id, creator)
				}
			}
		}
	}
}

func (sc *Stealcamoor) Start() {
	sc.startApiListener()
}
