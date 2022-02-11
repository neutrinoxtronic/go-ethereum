package main

// GOAL: full sync of the ethereum chain in under an hour (on M1 Max)

import (
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/protocols/eth"
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
	//glogger.Verbosity(log.LvlInfo)
	glogger.Verbosity(log.LvlDebug)
	//glogger.Verbosity(log.LvlTrace)
	log.Root().SetHandler(glogger)

	// build up connections to peers
	config := p2p.Config{
		ListenAddr: ":30303",
		MaxPeers:   50,
		NAT:        nat.Any(),
	}
	urls := params.MainnetBootnodes
	//fmt.Println(urls)
	config.BootstrapNodes = make([]*enode.Node, 0, len(urls))
	config.Name = common.MakeName("Geth", "v1.10.15")

	/*protos := eth.MakeProtocols(nil, 0, nil)
	fmt.Println(protos)*/
	protocols := make([]p2p.Protocol, 1)
	protocols[0] = p2p.Protocol{
		Name:    "eth",
		Version: eth.ETH66,
		Length:  17,
		Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
			fmt.Println("protocol run", p, "rw", rw)
			for {
				msg, err := rw.ReadMsg()
				if err != nil {
					break
				}
				if msg.Code == eth.StatusMsg {
					status := eth.StatusPacket{}
					if err := msg.Decode(&status); err != nil {
						return fmt.Errorf("status didn't decode")
					}
					fmt.Println(status)
					// TODO: send back not a mirror
					p2p.Send(rw, eth.StatusMsg, &eth.StatusPacket{
						ProtocolVersion: status.ProtocolVersion,
						NetworkID:       status.NetworkID,
						TD:              status.TD,
						Head:            status.Head,
						Genesis:         status.Genesis,
						ForkID:          status.ForkID,
					})
				} else {
					//fmt.Println("other message", msg.Code, msg)
				}
			}
			return nil
		},
	}

	//(*ethHandler)(s.handler), s.networkID, s.ethDialCandidates)
	config.Protocols = protocols

	server := p2p.Server{Config: config}

	/*fmt.Println(server.Protocols)
	os.Exit(0)*/

	testNodeKey, _ := crypto.GenerateKey()
	server.PrivateKey = testNodeKey

	check(server.Start())

	for i := 0; i < 100; i++ {
		time.Sleep(1000 * time.Millisecond)
		/*for _, p := range server.Peers() {
			fmt.Println("Node", p.Node())
		}*/
	}

	/*

		memdb := memorydb.New()
		database := rawdb.NewDatabase(memdb)

		// download the block headers
		// 14186270 blocks, block header is ~0.5k? = 7 GB

		d := downloader.New(0, database, nil, nil, nil, nil)
		d.Synchronise("test 1")*/

}
