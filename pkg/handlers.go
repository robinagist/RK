package rk

import (
	"github.com/ethereum/go-ethereum/p2p"
	"fmt"
	"encoding/json"
)

type RKHandler interface {
	Handle(message *RKMessage)
	Send(rks *RKResponse)
}

type RKMessage struct {

}


const MSG_TRANSACTION_BROADCAST = 1
const MSG_BLOCK_FOUND_BROADCAST = 2



type RKResponse struct {

}

type RKEnvelope struct {
	Version float32      `json:"version"`
	Origin string        `json:"origin"`
	Relay string         `json:"relay"`
	MessageType uint     `json:"message"`
	Timestamp string     `json:"timestamp"`
	Ttl int              `json:"ttl"`
	Message []byte       `json:"message"`
}

func ExtractEnvelope(ws p2p.MsgReadWriter) (*RKEnvelope, error)  {
	var envelope RKEnvelope
	for {
		msg, err := ws.ReadMsg()
		if err != nil {
			return nil, err
		}

		err = msg.Decode(&envelope)
		if err != nil {
			// handle decode error
			return nil, err
		}
	}
	return &envelope,nil
}


type RKBasicHandler struct {
  ws *p2p.MsgReadWriter
}


func (h *RKBasicHandler) Handle(envelope *RKEnvelope) error {

	switch envelope.MessageType {
	// standard transaction
	case MSG_TRANSACTION_BROADCAST:
		var tx Transaction
		err := json.Unmarshal(envelope.Message, &tx)
		if err != nil {
			return err
		}

		err = p2p.SendItems(*h.ws, messageId, "bar")
		if err != nil {
			return err
		}
	// standard block find broadcast
	case MSG_BLOCK_FOUND_BROADCAST:
		var block Block
		err := json.Unmarshal(envelope.Message, &block)
		if err != nil {
			return err
		}

	default:
		fmt.Println("recv:", envelope)
	}
	return nil
}

func (h *RKBasicHandler) Send() error {
	err := p2p.SendItems(*h.ws, messageId, "bar")
	if err != nil {
		return err
	}
	return nil
}