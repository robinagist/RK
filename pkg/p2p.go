package rk

import (
    "fmt"
    "os"

    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/p2p"
)

const messageId = 0



func RK1Protocol() p2p.Protocol {
	return p2p.Protocol{
		Name:    "RK1Protocol",
		Version: 1,
		Length:  1,
		Run:     rk1Runner,
	}
}

func Start() {
	nodekey, _ := crypto.GenerateKey()
	cfg := p2p.Config{
		MaxPeers:   10,
		PrivateKey: nodekey,
		Name:       "rk1-mainpeer",
		ListenAddr: ":30300",
		Protocols:  []p2p.Protocol{RK1Protocol()},
	}

	srv := p2p.Server{
		Config: cfg,
	}

	if err := srv.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	nodeinfo := srv.NodeInfo()
	fmt.Println("server started", "enode", nodeinfo.Enode, "name", nodeinfo.Name, "ID", nodeinfo.ID, "IP", nodeinfo.IP)

	select {}
}

func rk1Runner(peer *p2p.Peer, ws p2p.MsgReadWriter) error {
    envelope, err := ExtractEnvelope(ws)
    if err != nil {
    	panic(err)
	}
    handler := new(RKBasicHandler)
    handler.ws = &ws

    handler.Handle(envelope)

	return nil
}