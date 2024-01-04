// description:
// @author renshiwei
// Date: 2023/11/3

package test

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestBase(t *testing.T) {
	client, err := ethclient.Dial("https://api.chainup.net/arbitrum/mainnet/dc59a7e90f05424a81d825d033d0f2de")
	require.NoError(t, err)

	blockNumber, err := client.BlockNumber(context.Background())
	require.NoError(t, err)

	fmt.Println("blockNumber:", blockNumber)

	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	require.NoError(t, err)

	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash())
	}
}
