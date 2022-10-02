package main

import (
	contract "eth-app/gen"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("././.env")
	if err != nil {
		fmt.Printf("Error loading .env file")
	}

	client, err := ethclient.Dial(os.Getenv("LINK"))
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(os.Getenv("ADDRESS_CONTRACT"))

	t, err := contract.NewContract(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	var opts *bind.WatchOpts
	logs := make(chan *contract.ContractUnsubscribe, 10000)
	sub, err := t.WatchUnsubscribe(opts, logs)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
