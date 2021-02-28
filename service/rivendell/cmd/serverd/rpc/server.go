package rpc

import (
	"github.com/pawzio/arda/service/rivendell/genproto"
)

// Server controls RPC service responses.
type Server struct{
	genproto.UnimplementedRivendellServer
}

func NewServer() Server {
	return Server{}
}
