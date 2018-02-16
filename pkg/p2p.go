package rk

import (
"fmt"
"os"

"github.com/ethereum/go-ethereum/crypto"
"github.com/ethereum/go-ethereum/p2p"
	"encoding/json"
)

const messageId = 0

type Envelope struct {
	Origin string
	Relay string
	MessageType uint
	Timestamp string
	Ttl int
	Message []byte
}

func RK1Protocol() p2p.Protocol {
	return p2p.Protocol{
		Name:    "RK1Protocol",
		Version: 1,
		Length:  1,
		Run:     rk1Handler,
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

	select {}
}

func rk1Handler(peer *p2p.Peer, ws p2p.MsgReadWriter) error {
	for {
		msg, err := ws.ReadMsg()
		if err != nil {
			return err
		}

		var envelope Envelope
		err = msg.Decode(&envelope)
		if err != nil {
			// handle decode error
			return err
		}

		switch envelope.MessageType {
		// standard transaction
		case 1:
            var tx Transaction
			err := json.Unmarshal(envelope.Message, &tx)
			if err != nil {
				return err
			}

			err = p2p.SendItems(ws, messageId, "bar")
			if err != nil {
				return err
			}

		// standard block find broadcast
		case 2:
			var block Block
			err := json.Unmarshal(envelope.Message, &block)
			if err != nil {
				return err
			}

		default:
			fmt.Println("recv:", envelope)
		}
	}

	return nil
}