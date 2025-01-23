package agent

import (
	"crypto/tls"
	"fmt"
	"github.com/anemec/proglog/internal/log"
	"github.com/anemec/proglog/internal/server"
	"net"
	"sync"

	api "github.com/anemec/proglog/api/v1"
	"github.com/anemec/proglog/internal/discovery"
	"google.golang.org/grpc"
)

type Agent struct {
	Config

	log        *log.Log
	server     *grpc.Server
	membership *discovery.Membership
	replicator *log.Replicator

	shutdown     bool
	shutdowns    chan struct{}
	shutdownLock sync.Mutex
}

type Config struct {
	ServerTLSConfig *tls.Config
	PeerTLSConfig   *tls.Config
	DataDir         string
	BindAddr        string
	RPCPort         int
	NodeName        string
	StartJoinAddrs  []string
	ACLModelFile    string
	ACLPolicyFile   string
}

func (c Config) RPCAddr() (string, error) {
	host, _, err := net.SplitHostPort(c.BindAddr)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%d", host, c.RPCPort), nil
}
