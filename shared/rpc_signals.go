package shared

import (
	"net/rpc"

	v1alpha1 "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1"
)

type RPCSignalClient struct{ client *rpc.Client }

func (c *RPCSignalClient) Start(signal *v1alpha1.Signal) (<-chan *v1alpha1.Event, error) {
	var resp <-chan *v1alpha1.Event
	err := c.client.Call("Plugin.Start", map[string]interface{}{
		"signal": signal,
	}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *RPCSignalClient) Stop() error {
	var resp error
	err := c.client.Call("Plugin.Stop", map[string]interface{}{}, &resp)
	if err != nil {
		return err
	}
	return resp
}

// RPCSignalServer is the RPC server that RPCClient talks to, conforming to
// the requirements of net/rpc
type RPCSignalServer struct {
	// This is the real implementation
	Impl Signaler
}

func (s *RPCSignalServer) Start(args map[string]interface{}, resp *interface{}) error {
	events, err := s.Impl.Start(args["signal"].(*v1alpha1.Signal))
	*resp = events
	return err
}

func (s *RPCSignalServer) Stop(args map[string]interface{}, resp *interface{}) error {
	return s.Impl.Stop()
}
