/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type TrafficFilter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrafficFilterSpec   `json:"spec,omitempty"`
	Status            TrafficFilterStatus `json:"status,omitempty"`
}

type TrafficFilterSpecRule struct {
	// Optional rule description
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Computed rule ID
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// Required traffic filter source: IP address, CIDR mask, or VPC endpoint ID
	Source *string `json:"source" tf:"source"`
}

type TrafficFilterSpec struct {
	State *TrafficFilterSpecResource `json:"state,omitempty" tf:"-"`

	Resource TrafficFilterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type TrafficFilterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Optional ruleset description
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Should the ruleset be automatically included in the new deployments (Defaults to false)
	// +optional
	IncludeByDefault *bool `json:"includeByDefault,omitempty" tf:"include_by_default"`
	// Required name of the ruleset
	Name *string `json:"name" tf:"name"`
	// Required filter region, the ruleset can only be attached to deployments in the specific region
	Region *string `json:"region" tf:"region"`
	// Required list of rules, which the ruleset is made of.
	// +kubebuilder:validation:MinItems=1
	Rule []TrafficFilterSpecRule `json:"rule" tf:"rule"`
	// Required type of the ruleset ("ip" or "vpce")
	Type *string `json:"type" tf:"type"`
}

type TrafficFilterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// TrafficFilterList is a list of TrafficFilters
type TrafficFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TrafficFilter CRD objects
	Items []TrafficFilter `json:"items,omitempty"`
}
