// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NetworkPerformanceMetricSubscriptionInitParameters struct {

	// The target Region or Availability Zone that the metric subscription is enabled for. For example, eu-west-1.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// The metric used for the enabled subscription. Valid values: aggregate-latency. Default: aggregate-latency.
	Metric *string `json:"metric,omitempty" tf:"metric,omitempty"`

	// The source Region or Availability Zone that the metric subscription is enabled for. For example, us-east-1.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// The statistic used for the enabled subscription. Valid values: p50. Default: p50.
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`
}

type NetworkPerformanceMetricSubscriptionObservation struct {

	// The target Region or Availability Zone that the metric subscription is enabled for. For example, eu-west-1.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The metric used for the enabled subscription. Valid values: aggregate-latency. Default: aggregate-latency.
	Metric *string `json:"metric,omitempty" tf:"metric,omitempty"`

	// The data aggregation time for the subscription.
	Period *string `json:"period,omitempty" tf:"period,omitempty"`

	// The source Region or Availability Zone that the metric subscription is enabled for. For example, us-east-1.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// The statistic used for the enabled subscription. Valid values: p50. Default: p50.
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`
}

type NetworkPerformanceMetricSubscriptionParameters struct {

	// The target Region or Availability Zone that the metric subscription is enabled for. For example, eu-west-1.
	// +kubebuilder:validation:Optional
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// The metric used for the enabled subscription. Valid values: aggregate-latency. Default: aggregate-latency.
	// +kubebuilder:validation:Optional
	Metric *string `json:"metric,omitempty" tf:"metric,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The source Region or Availability Zone that the metric subscription is enabled for. For example, us-east-1.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// The statistic used for the enabled subscription. Valid values: p50. Default: p50.
	// +kubebuilder:validation:Optional
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`
}

// NetworkPerformanceMetricSubscriptionSpec defines the desired state of NetworkPerformanceMetricSubscription
type NetworkPerformanceMetricSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkPerformanceMetricSubscriptionParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider NetworkPerformanceMetricSubscriptionInitParameters `json:"initProvider,omitempty"`
}

// NetworkPerformanceMetricSubscriptionStatus defines the observed state of NetworkPerformanceMetricSubscription.
type NetworkPerformanceMetricSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkPerformanceMetricSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NetworkPerformanceMetricSubscription is the Schema for the NetworkPerformanceMetricSubscriptions API. Provides a resource to manage an Infrastructure Performance subscription.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NetworkPerformanceMetricSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destination) || (has(self.initProvider) && has(self.initProvider.destination))",message="spec.forProvider.destination is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.source) || (has(self.initProvider) && has(self.initProvider.source))",message="spec.forProvider.source is a required parameter"
	Spec   NetworkPerformanceMetricSubscriptionSpec   `json:"spec"`
	Status NetworkPerformanceMetricSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkPerformanceMetricSubscriptionList contains a list of NetworkPerformanceMetricSubscriptions
type NetworkPerformanceMetricSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkPerformanceMetricSubscription `json:"items"`
}

// Repository type metadata.
var (
	NetworkPerformanceMetricSubscription_Kind             = "NetworkPerformanceMetricSubscription"
	NetworkPerformanceMetricSubscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkPerformanceMetricSubscription_Kind}.String()
	NetworkPerformanceMetricSubscription_KindAPIVersion   = NetworkPerformanceMetricSubscription_Kind + "." + CRDGroupVersion.String()
	NetworkPerformanceMetricSubscription_GroupVersionKind = CRDGroupVersion.WithKind(NetworkPerformanceMetricSubscription_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkPerformanceMetricSubscription{}, &NetworkPerformanceMetricSubscriptionList{})
}
