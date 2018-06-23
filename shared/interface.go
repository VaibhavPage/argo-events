/*
Copyright 2018 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package shared

import (
	"net/rpc"

	"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1"
	plugin "github.com/hashicorp/go-plugin"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

const (
	CloudEventsVersion       = "v1.0"
	ContextExtensionErrorKey = "error"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "SIGNAL_PLUGIN",
	MagicCookieValue: "signal",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	// signals
	"NATS":     &signalPlugin{},
	"KAFKA":    &signalPlugin{},
	"AMQP":     &signalPlugin{},
	"MQTT":     &signalPlugin{},
	"Artifact": &signalPlugin{},
	"Calendar": &signalPlugin{},
	"Resource": &signalPlugin{},
	"Webhook":  &signalPlugin{},

	// todo: triggers
}

// Signaler is the interface for signaling
type Signaler interface {
	Start(*v1alpha1.Signal) (<-chan *v1alpha1.Event, error)
	Stop() error
}

// ArtifactSignaler is the interface for signaling with artifacts
// In addition to including the basic Signaler interface, this also
// enables access to read an artifact object to include in the event data payload
type ArtifactSignaler interface {
	Signaler
	// todo: change to use io.Reader and io.Closer interfaces?
	Read(*v1alpha1.ArtifactLocation, string) ([]byte, error)
}

// Triggerer is the interface for triggering
type Triggerer interface {
	Fire(*v1alpha1.Trigger, []*v1alpha1.Event) error
}

// NewSignalPlugin creates a base signal plugin
func NewSignalPlugin(impl Signaler) plugin.Plugin {
	return &signalPlugin{Impl: impl}
}

// NewArtifactSignalPlugin creates an artifact signal plugin
func NewArtifactSignalPlugin(impl ArtifactSignaler) plugin.Plugin {
	return &signalPlugin{Impl: impl}
}

// NewTriggerPlugin creates a base trigger plugin
func NewTriggerPlugin(impl Triggerer) plugin.Plugin {
	return &triggerPlugin{Impl: impl}
}

// signalPlugin is the implementation of plugin.Plugin so we can serve/consume this.
type signalPlugin struct {
	Impl Signaler
}

func (p *signalPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RPCSignalServer{Impl: p.Impl}, nil
}

func (p *signalPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RPCSignalClient{client: c}, nil
}

func (p *signalPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterSignalServer(s, &GRPCSignalServer{Impl: p.Impl})
	return nil
}

func (p *signalPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCSignalClient{client: NewSignalClient(c)}, nil
}

type triggerPlugin struct {
	Impl Triggerer
}

func (p *triggerPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RPCTriggerServer{Impl: p.Impl}, nil
}

func (p *triggerPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RPCTriggerClient{client: c}, nil
}

func (p *triggerPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterTriggerServer(s, &GRPCTriggerServer{Impl: p.Impl})
	return nil
}

func (p *triggerPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCTriggerClient{client: NewTriggerClient(c)}, nil
}
