package p2p

// represents the remote node
type Peer interface {
}

/*
handlers the communication between the nodes in the network
this can be TCP, UDP, websockets
*/
type Transport interface {
	ListenAndAccept() error
}
