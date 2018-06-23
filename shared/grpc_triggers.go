package shared

import (
	v1alpha1 "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1"
	empty "github.com/golang/protobuf/ptypes/empty"
	context "golang.org/x/net/context"
)

// GRPCTriggerClient is an implementation of TriggerClient that talks over gRPC.
type GRPCTriggerClient struct{ client TriggerClient }

func (c *GRPCTriggerClient) Fire(trigger *v1alpha1.Trigger, events []*v1alpha1.Event) error {
	req := TriggerRequest{
		Trigger: trigger,
		Events:  events,
	}
	_, err := c.client.Fire(context.Background(), &req)
	return err
}

// GRPCTriggerServer is the gRPC server that GRPCClient talks to.
type GRPCTriggerServer struct {
	// This is the real implementation
	Impl Triggerer
}

func (s *GRPCTriggerServer) Fire(ctx context.Context, req *TriggerRequest) (*empty.Empty, error) {
	return &empty.Empty{}, s.Impl.Fire(req.Trigger, req.Events)
}
