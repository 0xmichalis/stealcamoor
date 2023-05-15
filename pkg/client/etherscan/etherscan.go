package etherscan

import (
	"github.com/ethereum/go-ethereum/common"
)

func GetEtherscanAddress(explorerURL string, address common.Address) string {
	return explorerURL + "/address/" + address.String()
}
