package main

import (
	"fmt"
	"github.com/OpenBazaar/spvwallet"
	"github.com/OpenBazaar/spvwallet/db"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/op/go-logging"
	"sync"
)

var stdoutLogFormat = logging.MustStringFormatter(
	`%{color:reset}%{color}%{time:15:04:05.000} [%{shortfunc}] [%{level}] %{message}`,
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	// Create a new config
	config := spvwallet.NewDefaultConfig()

	// Use testnet
	config.Params = &chaincfg.TestNet3Params

	// Select wallet datastore
	sqliteDatastore, _ := db.Create(config.RepoPath)
	config.DB = sqliteDatastore

	// Create the wallet
	wallet, err := spvwallet.NewSPVWallet(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start it!
	go wallet.Start()
	wg.Wait()
}