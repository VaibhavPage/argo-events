package main

import (
	"os"

	"github.com/argoproj/argo-events/common"
	"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1"
	"github.com/argoproj/argo-events/shared"
	plugin "github.com/hashicorp/go-plugin"
	"k8s.io/client-go/rest"
)

type resource struct {
	kubeConfig *rest.Config
}

func (r *resource) Fire(trigger *v1alpha1.Trigger, events []*v1alpha1.Event) error {
	//todo: needs access to the kubectl
	return nil
}

var config *rest.Config

func init() {
	var err error
	kubeConfig, _ := os.LookupEnv(common.EnvVarKubeConfig)
	config, err = common.GetClientConfig(kubeConfig)
	if err != nil {
		panic(err)
	}
}

func main() {
	resource := &resource{kubeConfig: config}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"Resource": shared.NewTriggerPlugin(resource),
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
