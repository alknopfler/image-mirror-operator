/*
Copyright 2021.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MirroredSpec defines the desired state of Mirrored
type MirroredSpec struct {
	Ocp            Ocp         `json:"ocp"`
	Operators      []Operators `json:"operators"`
	TargetRegistry []string    `json:"targetRegistry"`
}

// Ocp struct defines the ocp version and channel to be mirrored
type Ocp struct {
	Channel []string `json:"versions"`
	Version []string `json:"tag"`
}

// Operator struct defines the packages and the catalag to be mirrored from
type Operators struct {
	Catalog  string     `json:"catalog"`
	Packages []Packages `json:"pacakges,omitempty"`
}
type Packages struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
	Channel string `json:"channel,omitempty"`
}

// MirroredStatus defines the observed state of Mirrored
type MirroredStatus struct {
	Ready       bool `json:"ready"`
	Initialized bool `json:"initialized"`
	Waiting     bool `json:"waiting"`
	Failed      bool `json:"failed"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Mirrored is the Schema for the mirroreds API
type Mirrored struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MirroredSpec   `json:"spec,omitempty"`
	Status MirroredStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MirroredList contains a list of Mirrored
type MirroredList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Mirrored `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Mirrored{}, &MirroredList{})
}
