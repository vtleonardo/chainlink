package synchronization

import (
	"context"
	"net/url"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/service"
	pb "github.com/smartcontractkit/chainlink/core/services/synchronization/telem"
	"github.com/smartcontractkit/chainlink/core/utils"

	"github.com/smartcontractkit/wsrpc"
	"github.com/smartcontractkit/wsrpc/examples/simple/keys"
)

// TelemetryIngressClient encapsulates all the functionality needed to
// send telemetry to the ingress server using wsrpc
type TelemetryIngressClient interface {
	service.Service
	Start() error
	Close() error
	Send(context.Context, []byte, common.Address)
}

type NoopTelemetryIngressClient struct{}

func (NoopTelemetryIngressClient) Start() error                                 { return nil }
func (NoopTelemetryIngressClient) Close() error                                 { return nil }
func (NoopTelemetryIngressClient) Send(context.Context, []byte, common.Address) {}
func (NoopTelemetryIngressClient) Healthy() error                               { return nil }
func (NoopTelemetryIngressClient) Ready() error                                 { return nil }

type telemetryIngressClient struct {
	utils.StartStopOnce
	url              *url.URL
	clientPrivKeyHex string
	serverPubKeyHex  string
	wsrpcClient      pb.TelemClient
	logging          bool

	mu      *sync.RWMutex
	isReady bool
	wgDone  sync.WaitGroup
	chDone  chan struct{}
}

// NewTelemetryIngressClient returns a client backed by wsrpc that
// can send telemetry to the telemetry ingress server
func NewTelemetryIngressClient(url *url.URL, serverPubKeyHex string, clientPrivKeyHex string, logging bool) TelemetryIngressClient {
	return &telemetryIngressClient{
		url:              url,
		clientPrivKeyHex: clientPrivKeyHex,
		serverPubKeyHex:  serverPubKeyHex,
		logging:          logging,
		mu:               new(sync.RWMutex),
	}
}

// Start connects the wsrpc client to the telemetry ingress server
func (tc *telemetryIngressClient) Start() error {
	return tc.StartOnce("TelemetryIngressClient", func() error {
		// TODO: get priv key here from keystore
		// privkey, err := s.getCSAPrivateKey()
		// if err != nil {
		//   return err
		// }

		tc.connect()

		return nil
	})
}

// Close disconnects the wsrpc client from the ingress server
func (tc *telemetryIngressClient) Close() error {
	return tc.StopOnce("TelemetryIngressClient", func() error {
		close(tc.chDone)
		tc.wgDone.Wait()
		return nil
	})
}

func (tc *telemetryIngressClient) connect() {
	tc.wgDone.Add(1)

	go func() {
		defer tc.wgDone.Done()

		clientPrivKey := keys.FromHex(tc.clientPrivKeyHex)
		serverPubKey := keys.FromHex(tc.serverPubKeyHex)

		conn, err := wsrpc.Dial(tc.url.String(), wsrpc.WithTransportCreds(clientPrivKey, serverPubKey))
		if err != nil {
			logger.Errorf("Error connecting to telemetry ingress server: %v", err)
			return
		}
		defer conn.Close()

		// Initialize a new wsrpc client caller
		// This is used to call RPC methods on the server
		tc.mu.Lock()
		tc.wsrpcClient = pb.NewTelemClient(conn)
		tc.isReady = true
		tc.mu.Unlock()

		// Wait for close
		<-tc.chDone

	}()
}

// Send sends telemetry to the ingress server using wsrpc if the client is ready
func (tc *telemetryIngressClient) Send(ctx context.Context, telemetry []byte, contractAddr common.Address) {
	telemReq := &pb.TelemRequest{Telemetry: telemetry, Address: contractAddr.String()}

	tc.mu.RLock()
	defer tc.mu.RUnlock()

	if !tc.isReady {
		logger.Error("Could not send telemetry, client is not ready")
		return
	}

	go func() {
		// Send telemetry to the ingress server, log any errors
		_, err := tc.wsrpcClient.Telem(ctx, telemReq)
		if err != nil {
			logger.Errorf("Some error ocurred sending telemetry: %v", err)
			return
		}
		if tc.logging {
			logger.Debugw("successfully sent telemetry to ingress server", "contractAddress", contractAddr.String(), "telemetry", telemetry)
		}
	}()
}
