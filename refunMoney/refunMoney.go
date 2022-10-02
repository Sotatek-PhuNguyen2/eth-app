package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	todo "eth-app/gen"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
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

	privateKey, err := crypto.HexToECDSA(os.Getenv("PERSON2_PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To: &cAdd,
	})
	if err == nil {
		gasLimit += 20000 * params.Wei
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	t, err := todo.NewContract(cAdd, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	tx.GasLimit = uint64(gasLimit)
	tx.GasPrice = gasPrice
	tx.Value = big.NewInt(100000000)
	tra, err := t.RefundMoneyUser(tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tra.Hash())

}
