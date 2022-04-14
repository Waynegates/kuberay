/*
Copyright 2022.

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

package v1alpha1

import (
	rayv1alpha1 "github.com/ray-project/kuberay/ray-operator/api/raycluster/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type ServeConfigSpec struct {
	Name                      string             `json:"name"`
	ImportPath                string             `json:"import_path"`
	InitArgs                  []string           `json:"init_args,omitempty"`
	InitKwargs                map[string]string  `json:"init_kwargs,omitempty"`
	NumReplicas               *int32             `json:"num_replicas,omitempty"`
	RoutePrefix               string             `json:"route_prefix,omitempty"`
	MaxConcurrentQueries      *int32             `json:"max_concurrent_queries,omitempty"`
	UserConfig                map[string]string  `json:"user_config,omitempty"`
	AutoscalingConfig         map[string]string  `json:"autoscaling_config,omitempty"`
	GracefulShutdownWaitLoopS *int32             `json:"graceful_shutdown_wait_loop_s,omitempty"`
	GracefulShutdownTimeoutS  *int32             `json:"graceful_shutdown_timeout_s,omitempty"`
	HealthCheckPeriodS        *int32             `json:"health_check_period_s,omitempty"`
	HealthCheckTimeoutS       *int32             `json:"health_check_timeout_s,omitempty"`
	RayActorOptions           RayActorOptionSpec `json:"ray_actor_options,omitempty"`
}

type RayActorOptionSpec struct {
	RuntimeEnv        map[string][]string `json:"runtime_env,omitempty"`
	NumCpus           *int32              `json:"num_cpus,omitempty"`
	NumGpus           *int32              `json:"num_gpus,omitempty"`
	Memory            *int32              `json:"memory,omitempty"`
	ObjectStoreMemory *int32              `json:"object_store_memory,omitempty"`
	Resources         map[string]string   `json:"resources,omitempty"`
	AcceleratorType   string              `json:"accelerator_type,omitempty"`
}

// ServingClusterSpec defines the desired state of ServingCluster
type ServingClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ServingCluster. Edit servingcluster_types.go to remove/update

	HealthCheckProbe *v1.Probe                  `json:"healthCheckConfig,omitempty"`
	ServeConfigSpecs []ServeConfigSpec          `json:"serveConfigs,omitempty"`
	RayClusterSpec   rayv1alpha1.RayClusterSpec `json:"rayClusterConfig,omitempty"`
}

type ServeStatus struct {
	NumReplicas *int32 `json:"num_replicas,omitempty"`
}

// ServingClusterStatus defines the observed state of ServingCluster
type ServingClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ServeStatuses    []ServeStatus                `json:"serveStatuses,omitempty"`
	RayClusterStatus rayv1alpha1.RayClusterStatus `json:"rayClusterStatus,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ServingCluster is the Schema for the servingclusters API
type ServingCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServingClusterSpec   `json:"spec,omitempty"`
	Status ServingClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ServingClusterList contains a list of ServingCluster
type ServingClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServingCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServingCluster{}, &ServingClusterList{}, &rayv1alpha1.RayCluster{}, &rayv1alpha1.RayClusterList{})
}
