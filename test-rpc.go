package main

import (
	"encoding/json"
	"github.com/btcsuite/btcd/rpcclient"
	"log"
	"os"
	"strconv"
)

type btcutil struct {
	Amount []float64 `json:"btcutil.Amount"`
}

func main() {
	// Connect to local dash RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:         "localhost:9998",
		User:         "dash",
		Pass:         os.Getenv("DASH_PASS"),
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Get the current wallet balance.
	walletBalance, err := client.GetBalance("1")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wallet balance: %d", walletBalance)
	log.Printf("%d", walletBalance)

	bytes := []byte(walletBalance)
	var Balance int64
	json.Unmarshal(bytes, &Balance)
	log.Printf(strconv.FormatInt(Balance, 64))
	//	return Balance
}
