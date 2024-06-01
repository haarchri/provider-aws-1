// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VPNConnectionRouteInitParameters struct {

	// The CIDR block associated with the local subnet of the customer network.
	DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty" tf:"destination_cidr_block,omitempty"`

	// The ID of the VPN connection.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/ec2/v1beta1.VPNConnection
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	VPNConnectionID *string `json:"vpnConnectionId,omitempty" tf:"vpn_connection_id,omitempty"`

	// Reference to a VPNConnection in ec2 to populate vpnConnectionId.
	// +kubebuilder:validation:Optional
	VPNConnectionIDRef *v1.Reference `json:"vpnConnectionIdRef,omitempty" tf:"-"`

	// Selector for a VPNConnection in ec2 to populate vpnConnectionId.
	// +kubebuilder:validation:Optional
	VPNConnectionIDSelector *v1.Selector `json:"vpnConnectionIdSelector,omitempty" tf:"-"`
}

type VPNConnectionRouteObservation struct {

	// The CIDR block associated with the local subnet of the customer network.
	DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty" tf:"destination_cidr_block,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the VPN connection.
	VPNConnectionID *string `json:"vpnConnectionId,omitempty" tf:"vpn_connection_id,omitempty"`
}

type VPNConnectionRouteParameters struct {

	// The CIDR block associated with the local subnet of the customer network.
	// +kubebuilder:validation:Optional
	DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty" tf:"destination_cidr_block,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the VPN connection.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/ec2/v1beta1.VPNConnection
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPNConnectionID *string `json:"vpnConnectionId,omitempty" tf:"vpn_connection_id,omitempty"`

	// Reference to a VPNConnection in ec2 to populate vpnConnectionId.
	// +kubebuilder:validation:Optional
	VPNConnectionIDRef *v1.Reference `json:"vpnConnectionIdRef,omitempty" tf:"-"`

	// Selector for a VPNConnection in ec2 to populate vpnConnectionId.
	// +kubebuilder:validation:Optional
	VPNConnectionIDSelector *v1.Selector `json:"vpnConnectionIdSelector,omitempty" tf:"-"`
}

// VPNConnectionRouteSpec defines the desired state of VPNConnectionRoute
type VPNConnectionRouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPNConnectionRouteParameters `json:"forProvider"`
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
	InitProvider VPNConnectionRouteInitParameters `json:"initProvider,omitempty"`
}

// VPNConnectionRouteStatus defines the observed state of VPNConnectionRoute.
type VPNConnectionRouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPNConnectionRouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VPNConnectionRoute is the Schema for the VPNConnectionRoutes API. Provides a static route between a VPN connection and a customer gateway.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPNConnectionRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destinationCidrBlock) || (has(self.initProvider) && has(self.initProvider.destinationCidrBlock))",message="spec.forProvider.destinationCidrBlock is a required parameter"
	Spec   VPNConnectionRouteSpec   `json:"spec"`
	Status VPNConnectionRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPNConnectionRouteList contains a list of VPNConnectionRoutes
type VPNConnectionRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNConnectionRoute `json:"items"`
}

// Repository type metadata.
var (
	VPNConnectionRoute_Kind             = "VPNConnectionRoute"
	VPNConnectionRoute_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPNConnectionRoute_Kind}.String()
	VPNConnectionRoute_KindAPIVersion   = VPNConnectionRoute_Kind + "." + CRDGroupVersion.String()
	VPNConnectionRoute_GroupVersionKind = CRDGroupVersion.WithKind(VPNConnectionRoute_Kind)
)

func init() {
	SchemeBuilder.Register(&VPNConnectionRoute{}, &VPNConnectionRouteList{})
}
