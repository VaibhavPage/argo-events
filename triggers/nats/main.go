package main

import (
	"fmt"

	"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1"
	"github.com/argoproj/argo-events/shared"
	plugin "github.com/hashicorp/go-plugin"
	natsio "github.com/nats-io/go-nats"
)

const (
	subjectKey = "subject"
)

type nats struct{}

func (n *nats) Fire(trigger *v1alpha1.Trigger, events []*v1alpha1.Event) error {
	// parse out the attributes
	subject, ok := trigger.Message.Stream.Attributes[subjectKey]
	if !ok {
		return shared.ErrMissingRequiredAttribute
	}
	conn, err := natsio.Connect(trigger.Message.Stream.URL)
	if err != nil {
		return fmt.Errorf("failed to connect to nats cluster url %s. Cause: %s", trigger.Message.Stream.URL, err)
	}
	defer conn.Close()

	// todo: things with the event payloads
	payload := []byte(trigger.Message.Body)
	return conn.Publish(subject, payload)
}

func main() {
	nats := &nats{}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"NATS": shared.NewTriggerPlugin(nats),
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
