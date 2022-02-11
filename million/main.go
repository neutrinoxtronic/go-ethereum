package main

// GOAL: full sync of the ethereum chain in under an hour (on M1 Pro)

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
)

func main() {
	fmt.Println("hello geth")
	memdb := memorydb.New()
	database := rawdb.NewDatabase(memdb)

	downloader = downloader.New(0, database, nil, nil, nil, nil)

}
