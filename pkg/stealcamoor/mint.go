package stealcamoor

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func (sc *Stealcamoor) tryMint(creator common.Address, ids []uint64) {
	idsForSigs := make([]uint64, 0)
	sigs := make([][]byte, 0)

	// Fetch all signatures for the ids to mint
	for _, id := range ids {
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
	}

	if len(idsForSigs) != 0 {
		sc.mint(creator, idsForSigs, sigs)
	}
}

func (sc *Stealcamoor) mint(creator common.Address, ids []uint64, sigs [][]byte) {
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
			continue
		}
		log.Printf("Transaction broadcasted: %s", tx.Hash().Hex())

		go func(id uint64) {
			sc.reveal(creator, id)
		}(id)
	}
}
