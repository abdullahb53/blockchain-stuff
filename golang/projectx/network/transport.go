package network

type NetAddr string

type RPC struct {
	From    string
	Payload []byte
}

type Transport interface {
	Consume() <-chan RPC
	Connect(*LocalTransport) error
	SendMessage(NetAddr, []byte) error
	Addr() NetAddr
}
