package stealcamoor

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	email "github.com/0xmichalis/stealcamoor/pkg/email"
	"github.com/0xmichalis/stealcamoor/pkg/etherscan"
	"github.com/0xmichalis/stealcamoor/pkg/stealcamapi"
)

func New() (*Stealcamoor, error) {
	sc := &Stealcamoor{}
	if err := sc.initApiClient(); err != nil {
		return nil, fmt.Errorf("failed to initialize api client: %w", err)
	}
	if err := sc.initEmailClient(); err != nil {
		return nil, fmt.Errorf("failed to initialize email client: %w", err)
	}
	if err := sc.initBlockchainClient(); err != nil {
		return nil, fmt.Errorf("failed to initialize blockchain client: %w", err)
	}
	if err := sc.initMisc(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return sc, nil
}

func (sc *Stealcamoor) initApiClient() error {
	apiURL := os.Getenv("STEALCAM_API_URL")
	if apiURL == "" {
		return errors.New("STEALCAM_API_URL cannot be empty")
	}
	sc.apiClient = stealcamapi.New(apiURL)

	if os.Getenv("STEALCAM_API_REQUEST_INTERVAL") == "" {
		return errors.New("STEALCAM_API_REQUEST_INTERVAL cannot be empty")
	}
	apiRequestInterval, err := time.ParseDuration(os.Getenv("STEALCAM_API_REQUEST_INTERVAL"))
	if err != nil {
		return err
	}
	sc.apiRequestInterval = apiRequestInterval

	return nil
}

func (sc *Stealcamoor) initEmailClient() error {
	username := os.Getenv("MAILTRAP_USERNAME")
	if username == "" {
		return errors.New("MAILTRAP_USERNAME cannot be empty")
	}
	password := os.Getenv("MAILTRAP_PASSWORD")
	if password == "" {
		return errors.New("MAILTRAP_PASSWORD cannot be empty")
	}
	from := os.Getenv("MAILTRAP_FROM")
	if from == "" {
		return errors.New("MAILTRAP_FROM cannot be empty")
	}
	sc.emailClient = email.New(username, password, from)
	to := os.Getenv("MAILTRAP_TO")
	if to == "" {
		return errors.New("MAILTRAP_TO cannot be empty")
	}
	sc.to = to

	return nil
}

func (sc *Stealcamoor) initBlockchainClient() error {
	stealcamAddress := os.Getenv("STEALCAM_ADDRESS")
	if stealcamAddress == "" {
		return errors.New("STEALCAM_ADDRESS cannot be empty")
	}
	sc.stealcamAddress = common.HexToAddress(stealcamAddress)
	log.Println("Stealcam address:", etherscan.GetEtherscanAddress(sc.explorerURL, sc.stealcamAddress))

	if os.Getenv("PRIVATE_KEY") == "" {
		return errors.New("PRIVATE_KEY cannot be empty")
	}

	if os.Getenv("NODE_API_URL") == "" {
		return errors.New("NODE_API_URL cannot be empty")
	}

	log.Println("Initializing node connection...")
	client, err := ethclient.Dial(os.Getenv("NODE_API_URL"))
	if err != nil {
		return fmt.Errorf("cannot connect to node: %w", err)
	}
	sc.client = client

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("cannot get chain id: %w", err)
	}
	log.Println("Chain ID:", chainID)

	// Load private key
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return fmt.Errorf("cannot load private key: %w", err)
	}

	// Extract address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("cannot cast public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Println("Our address:", etherscan.GetEtherscanAddress(sc.explorerURL, address))

	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("cannot create authorized transactor: %w", err)
	}
	sc.txOpts = txOpts

	return nil
}

func (sc *Stealcamoor) initMisc() error {
	explorerURL := os.Getenv("BLOCKCHAIN_EXPLORER_URL")
	if explorerURL == "" {
		return errors.New("BLOCKCHAIN_EXPLORER_URL cannot be empty")
	}
	sc.explorerURL = explorerURL

	creatorStrings := strings.Split(os.Getenv("CREATORS"), ",")
	creators := make([]common.Address, 0)
	log.Printf("Tracking %d creators:\n", len(creatorStrings))
	for _, creator := range creatorStrings {
		c := common.HexToAddress(creator)
		creators = append(creators, c)
		log.Println(etherscan.GetEtherscanAddress(sc.explorerURL, c))
	}
	if len(creators) == 0 {
		return errors.New("Need at least one creator provided in CREATORS (comma-separated list)")
	}
	sc.creators = creators

	return nil
}
