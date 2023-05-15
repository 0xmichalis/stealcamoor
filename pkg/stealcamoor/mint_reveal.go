package stealcamoor

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"mime"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"

	mail "github.com/0xmichalis/stealcamoor/pkg/client/email"
)

func (sc *Stealcamoor) tryMint(creator common.Address, ids []uint64) {
	idsForSigs := make([]uint64, 0)
	sigs := make([][]byte, 0)

	sc.mintCacheLock.Lock()
	defer sc.mintCacheLock.Unlock()

	// Fetch all signatures for the ids to mint
	for _, id := range ids {
		if sc.mintCache[id] {
			continue
		}

		m, err := sc.apiClient.GetMemory(id)
		if err != nil {
			log.Printf("Cannot get memory %d: %v", id, err)
			// avoid DDoSing the Stealcam API
			time.Sleep(10 * time.Millisecond)
			continue
		}

		signature, err := hex.DecodeString(m.Signature)
		if err != nil {
			log.Printf("Cannot decode signature %s: %v", m.Signature, err)
			// avoid DDoSing the Stealcam API
			time.Sleep(10 * time.Millisecond)
			continue
		}

		idsForSigs = append(idsForSigs, id)
		sigs = append(sigs, signature)
		sc.mintCache[id] = true
	}

	sc.mint(creator, idsForSigs, sigs)
	sc.reveal(creator, idsForSigs)
}

func (sc *Stealcamoor) mint(creator common.Address, ids []uint64, sigs [][]byte) {
	if len(ids) == 0 {
		return
	}
	if len(ids) != len(sigs) {
		log.Printf("Cannot mint %d memories: %d signatures provided", len(ids), len(sigs))
		return
	}

	nonce, err := sc.client.PendingNonceAt(context.Background(), sc.ourAddress)
	if err != nil {
		log.Printf("Cannot get pending nonce: %v", err)
		return
	}
	log.Printf("User nonce: %v", nonce)

	gasPrice, err := sc.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("Cannot get gas price: %v", err)
		return
	}
	log.Printf("Gas price: %v", gasPrice)

	// TODO: Deploy a contract to execute batch minting
	for i, id := range ids {
		sc.txOpts.Nonce = big.NewInt(int64(nonce + uint64(i)))
		sc.txOpts.GasPrice = gasPrice
		sc.txOpts.Value = big.NewInt(0)

		// TODO: Handle reverts
		tx, err := sc.stealcamContract.StealcamTransactor.Mint(
			sc.txOpts, big.NewInt(int64(id)), creator, sigs[i],
		)
		if err != nil {
			log.Printf("Failed to mint memory %d: %v", id, err)
			sc.mintCache[id] = false
			continue
		}
		log.Printf("Transaction broadcasted: %s", tx.Hash().Hex())
	}
}

func (sc *Stealcamoor) reveal(creator common.Address, ids []uint64) {
	if len(ids) == 0 {
		return
	}
	log.Printf("Revealing memories %v for creator %s...", ids, creator)

	for _, id := range ids {
		if !sc.mintCache[id] {
			// Skip if not successfully minted
			continue
		}
		sc.sendEmailForMemory(creator, id)
	}
}

func (sc *Stealcamoor) sendEmailForMemory(creator common.Address, id uint64) {
	url, err := sc.apiClient.RevealMemoryWithRetries(id, sc.ourAddress, sc.ourSignature, 30)
	if err != nil {
		log.Printf("Cannot reveal memory %d: %v", id, err)
		return
	}

	// Download image
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Cannot download image: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Cannot save image: %v", err)
		return
	}

	// Read the first 512 bytes of the file to try to detect the content type
	buffer := make([]byte, 512)
	if _, err = bytes.NewReader(body).Read(buffer); err != nil {
		log.Printf("Cannot read body: %v", err)
		return
	}
	contentType := http.DetectContentType(buffer)
	log.Println("Detected content type: " + contentType)
	fileExtension, err := mime.ExtensionsByType(contentType)
	if err != nil {
		log.Printf("Cannot get file extension: %v", err)
		return
	}

	// Pack as an email attachment
	msg := fmt.Sprintf("Check attachments for revealed memory %d for creator %s", id, creator.String())
	attachment := mail.Attachment{
		Name:        creator.String() + "_" + fmt.Sprintf("%d", id) + fileExtension[0],
		ContentType: contentType,
		Content:     body,
	}

	// Send email
	log.Printf("Sending email for memory %d...", id)
	if err := sc.emailClient.Send(msg, attachment); err != nil {
		log.Printf("Cannot send email: %v", err)
	}
}
