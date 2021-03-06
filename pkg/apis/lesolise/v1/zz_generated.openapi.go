// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"rabbitmq-operator/pkg/apis/lesolise/v1.RabbitMQ":       schema_pkg_apis_lesolise_v1_RabbitMQ(ref),
		"rabbitmq-operator/pkg/apis/lesolise/v1.RabbitMQSpec":   schema_pkg_apis_lesolise_v1_RabbitMQSpec(ref),
		"rabbitmq-operator/pkg/apis/lesolise/v1.RabbitMQStatus": schema_pkg_apis_lesolise_v1_RabbitMQStatus(ref),
	}
}

func schema_pkg_apis_lesolise_v1_RabbitMQ(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RabbitMQ is the Schema for the rabbitmqs API",
				Type:        []string{"object"},
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
							Ref: ref("rabbitmq-operator/pkg/apis/lesolise/v1.RabbitMQSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("rabbitmq-operator/pkg/apis/lesolise/v1.RabbitMQStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "rabbitmq-operator/pkg/apis/lesolise/v1.RabbitMQSpec", "rabbitmq-operator/pkg/apis/lesolise/v1.RabbitMQStatus"},
	}
}

func schema_pkg_apis_lesolise_v1_RabbitMQSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RabbitMQSpec defines the desired state of RabbitMQ",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_lesolise_v1_RabbitMQStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RabbitMQStatus defines the observed state of RabbitMQ",
				Type:        []string{"object"},
			},
		},
	}
}
