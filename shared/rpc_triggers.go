package shared

import (
	"net/rpc"

	v1alpha1 "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1"
)

type RPCTriggerClient struct{ client *rpc.Client }

func (c *RPCTriggerClient) Fire(trigger *v1alpha1.Trigger, events []v1alpha1.Event) error {
	err := c.client.Call("Plugin.Fire", map[string]interface{}{
		"trigger": trigger,
		"events":  events,
	}, nil)
	return err
}

// RPCTriggerServer is the RPC server that RPCClient talks to, conforming to
// the requirements of net/rpc
type RPCTriggerServer struct {
	// This is the real implementation
	Impl Triggerer
}

func (s *RPCTriggerServer) Fire(args map[string]interface{}, resp *interface{}) error {
	return s.Impl.Fire(args["trigger"].(*v1alpha1.Trigger), args["events"].([]*v1alpha1.Event))
}
