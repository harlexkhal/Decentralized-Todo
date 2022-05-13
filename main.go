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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://ropsten.infura.io/v3/c56cd359f90e484aa886b337072ce6d2"

func main() {

	b, err := ioutil.ReadFile("wallet/UTC--2022-05-11T11-32-36.314336300Z--4b3721e97b8a2e167cff69dd697a35ede8afabe0")
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

	gasprice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	networkChainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xCf2aB7AAb4ad6F9545376fb8f309e9E938BEa99b")
	req, err := todo.NewTodo(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, networkChainID)
	if err != nil {
		log.Fatal(err)
	}

	tx.GasPrice = gasprice
	tx.GasLimit = uint64(3000000)

	//Add to todo list
	/*transHash, err := req.Add(tx, "First todo on the list")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(transHash.Hash())
	*/

	//List tasks
	/*
		publicAddress := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
		tasks, err := req.List(&bind.CallOpts{
			From: publicAddress,
		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tasks)
	*/

	//Update to todo list
	/*transHash, err := req.Update(tx, big.NewInt(0), "dang it porting to aurora")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(transHash.Hash())*/

	transHash, err := req.Remove(tx, big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(transHash.Hash())
}
