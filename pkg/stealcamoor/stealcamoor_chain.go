package stealcamoor

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/common"
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

		url, err := sc.apiClient.RevealMemory(id, sc.ourAddress, sc.ourSignature)
		if err != nil {
			log.Printf("Cannot reveal memory %d: %v", id, err)
			continue
		}

		// Download image
		filename := creator.String() + "_" + fmt.Sprintf("%d", id) + ".jpg"

		out, err := os.Create(filename)
		if err != nil {
			log.Printf("Cannot create file to store image: %v", err)
			continue
		}
		defer out.Close()

		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Cannot download image: %v", err)
			return
		}
		defer resp.Body.Close()

		// Write the image to disk
		// TODO: Update to send as attachment via email
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			log.Printf("Cannot save image: %v", err)
			return
		}

		log.Println("Image saved to", filepath.Join(".", filename))
	}
}
