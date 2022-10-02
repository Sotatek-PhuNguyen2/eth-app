package main

import (
	"fmt"
	"log"
	"os"

	todo "eth-app/gen"

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
	defer client.Close()

	cAdd := common.HexToAddress(os.Getenv("ADDRESS_CONTRACT"))
	t, err := todo.NewContract(cAdd, client)
	if err != nil {
		log.Fatal(err)
	}

	users, err := t.List(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)

}
