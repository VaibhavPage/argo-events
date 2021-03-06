// +build !ignore_autogenerated

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
// Code generated by main. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.ArtifactLocation":  schema_pkg_apis_sensor_v1alpha1_ArtifactLocation(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.ArtifactSignal":    schema_pkg_apis_sensor_v1alpha1_ArtifactSignal(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.CalendarSignal":    schema_pkg_apis_sensor_v1alpha1_CalendarSignal(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.EscalationPolicy":  schema_pkg_apis_sensor_v1alpha1_EscalationPolicy(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.GroupVersionKind":  schema_pkg_apis_sensor_v1alpha1_GroupVersionKind(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Message":           schema_pkg_apis_sensor_v1alpha1_Message(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.NodeStatus":        schema_pkg_apis_sensor_v1alpha1_NodeStatus(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.ResourceFilter":    schema_pkg_apis_sensor_v1alpha1_ResourceFilter(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.ResourceObject":    schema_pkg_apis_sensor_v1alpha1_ResourceObject(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.ResourceSignal":    schema_pkg_apis_sensor_v1alpha1_ResourceSignal(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.RetryStrategy":     schema_pkg_apis_sensor_v1alpha1_RetryStrategy(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.S3Artifact":        schema_pkg_apis_sensor_v1alpha1_S3Artifact(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.S3Bucket":          schema_pkg_apis_sensor_v1alpha1_S3Bucket(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.S3Filter":          schema_pkg_apis_sensor_v1alpha1_S3Filter(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Sensor":            schema_pkg_apis_sensor_v1alpha1_Sensor(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.SensorList":        schema_pkg_apis_sensor_v1alpha1_SensorList(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.SensorSpec":        schema_pkg_apis_sensor_v1alpha1_SensorSpec(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.SensorStatus":      schema_pkg_apis_sensor_v1alpha1_SensorStatus(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Signal":            schema_pkg_apis_sensor_v1alpha1_Signal(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.SignalConstraints": schema_pkg_apis_sensor_v1alpha1_SignalConstraints(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Stream":            schema_pkg_apis_sensor_v1alpha1_Stream(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.TimeConstraints":   schema_pkg_apis_sensor_v1alpha1_TimeConstraints(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Trigger":           schema_pkg_apis_sensor_v1alpha1_Trigger(ref),
		"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.WebhookSignal":     schema_pkg_apis_sensor_v1alpha1_WebhookSignal(ref),
	}
}

func schema_pkg_apis_sensor_v1alpha1_ArtifactLocation(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ArtifactLocation describes the location for an external artifact",
				Properties: map[string]spec.Schema{
					"s3": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.S3Artifact"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.S3Artifact"},
	}
}

func schema_pkg_apis_sensor_v1alpha1_ArtifactSignal(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ArtifactSignal describes an external object dependency",
				Properties: map[string]spec.Schema{
					"s3": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.S3Artifact"),
						},
					},
					"stream": {
						SchemaProps: spec.SchemaProps{
							Description: "NotificationStream is the stream to listen for artifact notifications",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Stream"),
						},
					},
				},
				Required: []string{"stream"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.S3Artifact", "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Stream"},
	}
}

func schema_pkg_apis_sensor_v1alpha1_CalendarSignal(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CalendarSignal describes a time based dependency. One of the fields (schedule, interval, or recurrence) must be passed. Schedule takes precedence over interval; interval takes precedence over recurrence",
				Properties: map[string]spec.Schema{
					"schedule": {
						SchemaProps: spec.SchemaProps{
							Description: "Schedule is a cron-like expression. For reference, see: https://en.wikipedia.org/wiki/Cron",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"interval": {
						SchemaProps: spec.SchemaProps{
							Description: "Interval is a string that describes an interval duration, e.g. 1s, 30m, 2h...",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"recurrence": {
						SchemaProps: spec.SchemaProps{
							Description: "List of RRULE, RDATE and EXDATE lines for a recurring event, as specified in RFC5545. RRULE is a recurrence rule which defines a repeating pattern for recurring events. RDATE defines the list of DATE-TIME values for recurring events. EXDATE defines the list of DATE-TIME exceptions for recurring events. the combination of these rules and dates combine to form a set of date times. NOTE: functionality currently only supports EXDATEs, but in the future could be expanded.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"schedule", "interval", "recurrence"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_sensor_v1alpha1_EscalationPolicy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EscalationPolicy describes the policy for escalating sensors in an Error state. NOTE: this functionality is currently experimental, but we believe serves as an important future enhancement around handling lifecycle error conditions of a sensor.",
				Properties: map[string]spec.Schema{
					"level": {
						SchemaProps: spec.SchemaProps{
							Description: "Level is the degree of importance",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "need someway to progressively get more serious notifications",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Message"),
						},
					},
				},
				Required: []string{"level", "message"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Message"},
	}
}

func schema_pkg_apis_sensor_v1alpha1_GroupVersionKind(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GroupVersionKind unambiguously identifies a kind.  It doesn't anonymously include GroupVersion to avoid automatic coercion.  It doesn't use a GroupVersion to avoid custom marshalling.",
				Properties: map[string]spec.Schema{
					"group": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"group", "version", "kind"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_sensor_v1alpha1_Message(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Message represents a message on a queue",
				Properties: map[string]spec.Schema{
					"body": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"stream": {
						SchemaProps: spec.SchemaProps{
							Description: "Stream descibes queue resources to send the message on",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Stream"),
						},
					},
				},
				Required: []string{"body"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Stream"},
	}
}

func schema_pkg_apis_sensor_v1alpha1_NodeStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NodeStatus describes the status for an individual node in the sensor's FSM. A single node can represent the status for signal or a trigger.",
				Properties: map[string]spec.Schema{
					"id": {
						SchemaProps: spec.SchemaProps{
							Description: "ID is a unique identifier of a node within a sensor It is a hash of the node name",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is a unique name in the node tree used to generate the node ID",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"displayName": {
						SchemaProps: spec.SchemaProps{
							Description: "DisplayName is the human readable representation of the node",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type is the type of the node",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase of the node",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"startedAt": {
						SchemaProps: spec.SchemaProps{
							Description: "StartedAt is the time at which this node started",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"resolvedAt": {
						SchemaProps: spec.SchemaProps{
							Description: "ResolvedAt is the time at which this node resolved",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "store data or something to save for signal notifications or trigger events",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"id", "name", "displayName", "type", "phase"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_sensor_v1alpha1_ResourceFilter(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ResourceFilter contains K8 ObjectMeta information to further filter resource signal objects",
				Properties: map[string]spec.Schema{
					"prefix": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"annotations": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"createdBy": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_sensor_v1alpha1_ResourceObject(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ResourceObject is the resource object to create on kubernetes",
				Properties: map[string]spec.Schema{
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Description: "Namespace in which to create this object optional",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"group": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"artifactLocation": {
						SchemaProps: spec.SchemaProps{
							Description: "Location in which the K8 resource file(s) are stored. If omitted, will attempt to use the default artifact location configured in the controller.",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.ArtifactLocation"),
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Description: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. This overrides any labels in the unstructured object with the same key.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"group", "version", "kind"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.ArtifactLocation"},
	}
}

func schema_pkg_apis_sensor_v1alpha1_ResourceSignal(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ResourceSignal refers to a dependency on a k8s resource.",
				Properties: map[string]spec.Schema{
					"group": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"filter": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.ResourceFilter"),
						},
					},
				},
				Required: []string{"group", "version", "kind", "namespace"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.ResourceFilter"},
	}
}

func schema_pkg_apis_sensor_v1alpha1_RetryStrategy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RetryStrategy represents a strategy for retrying operations",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_sensor_v1alpha1_S3Artifact(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "S3Artifact contains information about an artifact in S3",
				Properties: map[string]spec.Schema{
					"endpoint": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"bucket": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"region": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"insecure": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"accessKey": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.SecretKeySelector"),
						},
					},
					"secretKey": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.SecretKeySelector"),
						},
					},
					"key": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"event": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"filter": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.S3Filter"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.S3Filter", "k8s.io/api/core/v1.SecretKeySelector"},
	}
}

func schema_pkg_apis_sensor_v1alpha1_S3Bucket(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "S3Bucket contains information for an S3 Bucket",
				Properties: map[string]spec.Schema{
					"endpoint": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"bucket": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"region": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"insecure": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"accessKey": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.SecretKeySelector"),
						},
					},
					"secretKey": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.SecretKeySelector"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.SecretKeySelector"},
	}
}

func schema_pkg_apis_sensor_v1alpha1_S3Filter(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "S3Filter represents filters to apply to bucket nofifications for specifying constraints on objects",
				Properties: map[string]spec.Schema{
					"prefix": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"suffix": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"prefix", "suffix"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_sensor_v1alpha1_Sensor(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Sensor is the definition of a sensor resource",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.SensorSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.SensorStatus"),
						},
					},
				},
				Required: []string{"metadata", "spec", "status"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.SensorSpec", "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.SensorStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_sensor_v1alpha1_SensorList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SensorList is the list of Sensor resources",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Sensor"),
									},
								},
							},
						},
					},
				},
				Required: []string{"metadata", "items"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Sensor", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_sensor_v1alpha1_SensorSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SensorSpec represents desired sensor state",
				Properties: map[string]spec.Schema{
					"signals": {
						SchemaProps: spec.SchemaProps{
							Description: "Signals is a list of the things that this sensor is dependent on. These are the inputs to this sensor.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Signal"),
									},
								},
							},
						},
					},
					"triggers": {
						SchemaProps: spec.SchemaProps{
							Description: "Triggers is a list of the things that this sensor evokes. These are the outputs from this sensor.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Trigger"),
									},
								},
							},
						},
					},
					"escalation": {
						SchemaProps: spec.SchemaProps{
							Description: "Escalation describes the policy for signal failures and violations of the dependency's constraints.",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.EscalationPolicy"),
						},
					},
					"repeat": {
						SchemaProps: spec.SchemaProps{
							Description: "Repeat is a flag that determines if the sensor status should be reset after completion. NOTE: functionality is currently expiremental and part of an initiative to define a more concrete pattern or cycle for sensor reptition.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"signals", "triggers"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.EscalationPolicy", "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Signal", "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Trigger"},
	}
}

func schema_pkg_apis_sensor_v1alpha1_SensorStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SensorStatus contains information about the status of a sensor.",
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase is the high-level summary of the sensor",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"startedAt": {
						SchemaProps: spec.SchemaProps{
							Description: "StartedAt is the time at which this sensor was initiated",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"resolvedAt": {
						SchemaProps: spec.SchemaProps{
							Description: "ResolvedAt is the time at which this sensor was resolved",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Message is a human readable string indicating details about a sensor in its phase",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"nodes": {
						SchemaProps: spec.SchemaProps{
							Description: "Nodes is a mapping between a node ID and the node's status it records the states for the FSM of this sensor.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.NodeStatus"),
									},
								},
							},
						},
					},
					"escalated": {
						SchemaProps: spec.SchemaProps{
							Description: "Escalated is a flag for whether this sensor was escalated",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"phase"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.NodeStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_sensor_v1alpha1_Signal(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Signal describes a dependency",
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is a unique name of this dependency",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"deadline": {
						SchemaProps: spec.SchemaProps{
							Description: "Deadline is the duration in seconds after the StartedAt time of the sensor after which this signal is terminated. Note: this functionality is not yet respected, but it's theoretical behavior is as follows: This trumps the recurrence patterns of calendar signals and allows any signal to have a strict defined life. After the deadline is reached and this signal has not in a Resolved state, this signal is marked as Failed and proper escalations should proceed.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"stream": {
						SchemaProps: spec.SchemaProps{
							Description: "Stream defines a message stream dependency",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Stream"),
						},
					},
					"artifact": {
						SchemaProps: spec.SchemaProps{
							Description: "artifact defines an external file dependency",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.ArtifactSignal"),
						},
					},
					"calendar": {
						SchemaProps: spec.SchemaProps{
							Description: "Calendar defines a time based dependency",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.CalendarSignal"),
						},
					},
					"resource": {
						SchemaProps: spec.SchemaProps{
							Description: "Resource defines a dependency on a kubernetes resource -- this can be a pod, deployment or custom resource",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.ResourceSignal"),
						},
					},
					"webhook": {
						SchemaProps: spec.SchemaProps{
							Description: "Webhook defines a HTTP notification dependency",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.WebhookSignal"),
						},
					},
					"constraints": {
						SchemaProps: spec.SchemaProps{
							Description: "Constraints and rules governing tolerations of success and overrides",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.SignalConstraints"),
						},
					},
				},
				Required: []string{"name"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.ArtifactSignal", "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.CalendarSignal", "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.ResourceSignal", "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.SignalConstraints", "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Stream", "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.WebhookSignal"},
	}
}

func schema_pkg_apis_sensor_v1alpha1_SignalConstraints(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SignalConstraints defines constraints for a dependent signal.",
				Properties: map[string]spec.Schema{
					"time": {
						SchemaProps: spec.SchemaProps{
							Description: "Time constraints on the signal",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.TimeConstraints"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.TimeConstraints"},
	}
}

func schema_pkg_apis_sensor_v1alpha1_Stream(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Stream describes a queue stream resource",
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type of the stream resource",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"url": {
						SchemaProps: spec.SchemaProps{
							Description: "URL is the exposed endpoint for client connections to this service",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"attributes": {
						SchemaProps: spec.SchemaProps{
							Description: "Attributes contains additional fields specific to each service implementation",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"type", "url"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_sensor_v1alpha1_TimeConstraints(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TimeConstraints describes constraints in time",
				Properties: map[string]spec.Schema{
					"start": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"stop": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
				},
				Required: []string{"start", "stop"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_sensor_v1alpha1_Trigger(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Trigger is an action taken, output produced, an event created, a message sent",
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is a unique name of the action to take",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"resource": {
						SchemaProps: spec.SchemaProps{
							Description: "Resource describes the resource that will be created by this action",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.ResourceObject"),
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Message describes a message that will be sent on a queue",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Message"),
						},
					},
					"retryStrategy": {
						SchemaProps: spec.SchemaProps{
							Description: "RetryStrategy is the strategy to retry a trigger if it fails",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.RetryStrategy"),
						},
					},
				},
				Required: []string{"name", "retryStrategy"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.Message", "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.ResourceObject", "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1.RetryStrategy"},
	}
}

func schema_pkg_apis_sensor_v1alpha1_WebhookSignal(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WebhookSignal is a general purpose REST API",
				Properties: map[string]spec.Schema{
					"endpoint": {
						SchemaProps: spec.SchemaProps{
							Description: "REST API endpoint",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"port": {
						SchemaProps: spec.SchemaProps{
							Description: "Port to listen on",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"method": {
						SchemaProps: spec.SchemaProps{
							Description: "Method is HTTP request method that indicates the desired action to be performed for a given resource. See RFC7231 Hypertext Transfer Protocol (HTTP/1.1): Semantics and Content",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"endpoint", "port", "method"},
			},
		},
		Dependencies: []string{},
	}
}
