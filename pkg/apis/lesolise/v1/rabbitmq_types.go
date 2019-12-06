package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RabbitMQSpec defines the desired state of RabbitMQ
// +k8s:openapi-gen=true
type RabbitMQSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	Image string `json:"image,omitempty"`
	// +kubebuilder:validation:Minimum=3
	Size int32 `json:"size,omitempty"`
	// resource requests and limits
	DiskLimit     string `json:"disk_limit,omitempty"`
	DiskRequest   string `json:"disk_request,omitempty"`
	MemoryRequest string `json:"memory_request,omitempty"`
	MemoryLimit   string `json:"memory_limit,omitempty"`
	CpuLimit      string `json:"cpu_limit,omitempty"`
	CpuRequest    string `json:"cpu_request,omitempty"`
	// we suggest to use local pv, so the storage class name must be set
	StorageClassName string `json:"storage_class_name"`
	// specify the hostname suffix for rabbitmq management UI,
	// for example, when ObjectMeta.Name = test and this field is .rmq.com,
	// we will generate a ingress whose rule host is test.rmq.com for kafka manager
	// then you can bind hosts test.rmq.com to access it
	// default value is .rmq.com
	ManagerHost string `json:"rabbitmq_manager_host,omitempty"`
}

// RabbitMQStatus defines the observed state of RabbitMQ
// +k8s:openapi-gen=true
type RabbitMQStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	RabbitmqUrl             string `json:rabbitmq_url`
	RabbitmqManagerUrl      string `json:rabbitmq_manager_url`
	RabbitmqManagerPassword string `json:rabbitmq_manager_password`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RabbitMQ is the Schema for the rabbitmqs API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=rabbitmqs,scope=Namespaced
type RabbitMQ struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RabbitMQSpec   `json:"spec,omitempty"`
	Status RabbitMQStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RabbitMQList contains a list of RabbitMQ
type RabbitMQList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RabbitMQ `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RabbitMQ{}, &RabbitMQList{})
}