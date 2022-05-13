package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://aurora-testnet.infura.io/v3/c56cd359f90e484aa886b337072ce6d2"

func main() {

	//create a wallet.
	/*ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)

	ks.NewAccount("p@ssw0rd")
	ks.NewAccount("p@ssw0rd")*/

	//address 1 = 4b3721e97b8a2e167cff69dd697a35ede8afabe0
	//address 2 = 0208d139537631af583b681d39fe6de2a4244d42

	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("An error occured while trying to create a client: %v", err)
	}
	defer client.Close()

	a1 := common.HexToAddress("4b3721e97b8a2e167cff69dd697a35ede8afabe0")
	a2 := common.HexToAddress("0208d139537631af583b681d39fe6de2a4244d42")
	fmt.Println("send eth to this address: ", a1)
	b1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal(err)
	}

	b2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether  = 1000000000000000000 wei
	result1 := new(big.Float).Quo(new(big.Float).SetInt(b1), big.NewFloat(1000000000000000000))
	result2 := new(big.Float).Quo(new(big.Float).SetInt(b2), big.NewFloat(1000000000000000000))
	fmt.Println("Balance 1: ", result1)
	fmt.Println("Balance 2: ", result2)

	nonce, err := client.PendingNonceAt(context.Background(), a1)
	fmt.Println("Nonce count for acct 1: ", nonce)
	if err != nil {
		log.Fatal(err)
	}

	totalEther := 0.000001
	calc_amount := int64(totalEther * 1000000000000000000)
	amount := big.NewInt(calc_amount)
	gasprice, err := client.SuggestGasPrice(context.Background())
	transaction := types.NewTransaction(nonce, a2, amount, 21000, gasprice, nil)

	b, err := ioutil.ReadFile("../wallet/UTC--2022-05-11T11-32-36.314336300Z--4b3721e97b8a2e167cff69dd697a35ede8afabe0")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "p@ssw0rd")
	fmt.Println("key: ", key)
	if err != nil {
		log.Fatal(err)
	}

	networkChainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	s_transaction, err := types.SignTx(transaction, types.NewEIP155Signer(networkChainID), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), s_transaction)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("transaction-hash: %s", s_transaction.Hash().Hex())
}
