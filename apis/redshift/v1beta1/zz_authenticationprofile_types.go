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

type AuthenticationProfileInitParameters struct {

	// The content of the authentication profile in JSON format. The maximum length of the JSON string is determined by a quota for your account.
	AuthenticationProfileContent *string `json:"authenticationProfileContent,omitempty" tf:"authentication_profile_content,omitempty"`
}

type AuthenticationProfileObservation struct {

	// The content of the authentication profile in JSON format. The maximum length of the JSON string is determined by a quota for your account.
	AuthenticationProfileContent *string `json:"authenticationProfileContent,omitempty" tf:"authentication_profile_content,omitempty"`

	// The name of the authentication profile.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AuthenticationProfileParameters struct {

	// The content of the authentication profile in JSON format. The maximum length of the JSON string is determined by a quota for your account.
	// +kubebuilder:validation:Optional
	AuthenticationProfileContent *string `json:"authenticationProfileContent,omitempty" tf:"authentication_profile_content,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// AuthenticationProfileSpec defines the desired state of AuthenticationProfile
type AuthenticationProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthenticationProfileParameters `json:"forProvider"`
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
	InitProvider AuthenticationProfileInitParameters `json:"initProvider,omitempty"`
}

// AuthenticationProfileStatus defines the observed state of AuthenticationProfile.
type AuthenticationProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthenticationProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AuthenticationProfile is the Schema for the AuthenticationProfiles API. Creates a Redshift authentication profile
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AuthenticationProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authenticationProfileContent) || (has(self.initProvider) && has(self.initProvider.authenticationProfileContent))",message="spec.forProvider.authenticationProfileContent is a required parameter"
	Spec   AuthenticationProfileSpec   `json:"spec"`
	Status AuthenticationProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthenticationProfileList contains a list of AuthenticationProfiles
type AuthenticationProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthenticationProfile `json:"items"`
}

// Repository type metadata.
var (
	AuthenticationProfile_Kind             = "AuthenticationProfile"
	AuthenticationProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthenticationProfile_Kind}.String()
	AuthenticationProfile_KindAPIVersion   = AuthenticationProfile_Kind + "." + CRDGroupVersion.String()
	AuthenticationProfile_GroupVersionKind = CRDGroupVersion.WithKind(AuthenticationProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthenticationProfile{}, &AuthenticationProfileList{})
}
