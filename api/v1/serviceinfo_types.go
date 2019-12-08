/*

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ServiceInfoSpec defines the desired state of ServiceInfo
type ServiceInfoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ServiceInfo. Edit ServiceInfo_types.go to remove/update
	Selector map[string]string `json:"selector,omitempty"`
}

// ServiceInfoStatus defines the observed state of ServiceInfo
type ServiceInfoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Nodes []string
}

// +kubebuilder:object:root=true

// ServiceInfo is the Schema for the serviceinfoes API
type ServiceInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceInfoSpec   `json:"spec,omitempty"`
	Status ServiceInfoStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceInfoList contains a list of ServiceInfo
type ServiceInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceInfo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceInfo{}, &ServiceInfoList{})
}
