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
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/0xmichalis/stealcamoor/pkg/abis"
	email "github.com/0xmichalis/stealcamoor/pkg/client/email"
	"github.com/0xmichalis/stealcamoor/pkg/client/etherscan"
	"github.com/0xmichalis/stealcamoor/pkg/client/stealcam"
)

func New() (*Stealcamoor, error) {
	sc := &Stealcamoor{}
	if err := sc.initApiClient(); err != nil {
		return nil, fmt.Errorf("failed to initialize api client: %w", err)
	}
	if err := sc.initMisc(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	if err := sc.initEmailClient(); err != nil {
		return nil, fmt.Errorf("failed to initialize email client: %w", err)
	}
	if err := sc.initBlockchainClient(); err != nil {
		return nil, fmt.Errorf("failed to initialize blockchain client: %w", err)
	}

	return sc, nil
}

func (sc *Stealcamoor) initApiClient() error {
	apiURL := os.Getenv("STEALCAM_API_URL")
	if apiURL == "" {
		return errors.New("STEALCAM_API_URL cannot be empty")
	}
	sc.apiClient = stealcam.New(apiURL)

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
	host := os.Getenv("SMTP_HOST")
	if host == "" {
		return errors.New("SMTP_HOST cannot be empty")
	}
	port := os.Getenv("SMTP_PORT")
	if port == "" {
		return errors.New("SMTP_PORT cannot be empty")
	}
	username := os.Getenv("SMTP_USERNAME")
	if username == "" {
		return errors.New("SMTP_USERNAME cannot be empty")
	}
	password := os.Getenv("SMTP_PASSWORD")
	if password == "" {
		return errors.New("SMTP_PASSWORD cannot be empty")
	}
	from := os.Getenv("SMTP_FROM")
	if from == "" {
		return errors.New("SMTP_FROM cannot be empty")
	}
	to := os.Getenv("SMTP_TO")
	if to == "" {
		return errors.New("SMTP_TO cannot be empty")
	}
	sc.emailClient = email.New(host, port, username, password, from, strings.Split(to, ","))

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

	nodeURL := os.Getenv("NODE_API_URL")
	if nodeURL == "" {
		return errors.New("NODE_API_URL cannot be empty")
	}
	if !strings.HasPrefix(nodeURL, "wss://") {
		return errors.New("NODE_API_URL needs to be a Websocket RPC URL")
	}

	log.Println("Initializing node connection...")
	client, err := ethclient.Dial(nodeURL)
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
	sc.ourAddress = address
	log.Println("Our address:", etherscan.GetEtherscanAddress(sc.explorerURL, address))

	// Sign message and store signature once to be used during reveals
	message := "Confirming my ETH address to reveal a Stealcam photo"
	signature, err := signMessage(message, privateKey)
	if err != nil {
		return fmt.Errorf("cannot sign message: %w", err)
	}
	sc.ourSignature = signature

	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("cannot create authorized transactor: %w", err)
	}
	sc.txOpts = txOpts
	sc.txOpts.GasLimit = 10000000

	stealcam, err := abis.NewStealcam(sc.stealcamAddress, client)
	if err != nil {
		return fmt.Errorf("cannot instantiate stealcam contract client: %w", err)
	}
	sc.stealcamContract = stealcam

	return nil
}

func signMessage(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}
	signatureBytes[64] += 27
	return hexutil.Encode(signatureBytes), nil
}

func (sc *Stealcamoor) initMisc() error {
	explorerURL := os.Getenv("BLOCKCHAIN_EXPLORER_URL")
	if explorerURL == "" {
		return errors.New("BLOCKCHAIN_EXPLORER_URL cannot be empty")
	}
	sc.explorerURL = explorerURL

	creatorStrings := strings.Split(os.Getenv("CREATORS"), ",")
	creators := make([]common.Address, 0)
	isDuplicate := make(map[string]bool)
	addressToTwitter := make(map[string]string)
	for _, creator := range creatorStrings {
		if isDuplicate[creator] {
			return fmt.Errorf("duplicate creator %s", creator)
		}
		isDuplicate[creator] = true
		c := common.HexToAddress(creator)
		creators = append(creators, c)
		profile, err := sc.apiClient.GetProfile(c)
		if err != nil {
			return fmt.Errorf("cannot get profile for creator %s: %w", creator, err)
		}
		addressToTwitter[c.String()] = profile.Username
		logCreator := etherscan.GetEtherscanAddress(sc.explorerURL, c)
		if profile.Username != "" {
			logCreator = fmt.Sprintf("https://twitter.com/%s (%s)", profile.Username, creator)
		}
		log.Printf("Tracking creator %s", logCreator)
	}
	if len(creators) == 0 {
		return errors.New("Need at least one creator provided in CREATORS (comma-separated list)")
	}
	if len(creators) == 1 {
		log.Print("Tracking 1 creator.")
	} else {
		log.Printf("Tracking %d creators.", len(creatorStrings))
	}

	sc.creators = creators
	sc.addressToTwitter = addressToTwitter
	sc.backupDir = os.Getenv("BACKUP_DIR")

	return nil
}
