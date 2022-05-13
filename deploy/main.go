package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	todo "dapp.com/todo"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://aurora-testnet.infura.io/v3/c56cd359f90e484aa886b337072ce6d2"

func main() {

	b, err := ioutil.ReadFile("../wallet/UTC--2022-05-11T11-32-36.314336300Z--4b3721e97b8a2e167cff69dd697a35ede8afabe0")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "p@ssw0rd")
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("An error occured while trying to create a client: %v", err)
	}
	defer client.Close()

	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}

	gasprice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	networkChainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, networkChainID)
	if err != nil {
		log.Fatal(err)
	}

	auth.GasPrice = gasprice
	auth.GasLimit = uint64(3000000)
	auth.Nonce = big.NewInt(int64(nonce))
	contractAddress, transactionHash, _, err := todo.DeployTodo(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(contractAddress)
	fmt.Println(transactionHash.Hash().Hex())
}
