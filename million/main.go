package main

// GOAL: full sync of the ethereum chain in under an hour (on M1 Max)

import (
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/nat"
	"github.com/ethereum/go-ethereum/params"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("hello geth")

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(true)))
	glogger.Verbosity(log.LvlTrace)
	log.Root().SetHandler(glogger)

	// build up connections to peers
	config := p2p.Config{
		ListenAddr: ":30303",
		MaxPeers:   50,
		NAT:        nat.Any(),
	}
	urls := params.MainnetBootnodes
	config.BootstrapNodes = make([]*enode.Node, 0, len(urls))
	server := p2p.Server{Config: config}

	testNodeKey, _ := crypto.GenerateKey()
	server.PrivateKey = testNodeKey

	check(server.Start())

	for i := 0; i < 100; i++ {
		time.Sleep(1000 * time.Millisecond)
	}

	/*

		memdb := memorydb.New()
		database := rawdb.NewDatabase(memdb)

		// download the block headers
		// 14186270 blocks, block header is ~0.5k? = 7 GB

		d := downloader.New(0, database, nil, nil, nil, nil)
		d.Synchronise("test 1")*/

}
