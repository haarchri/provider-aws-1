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

type SAMLOptionsInitParameters struct {

	// Group attribute for this SAML integration.
	GroupAttribute *string `json:"groupAttribute,omitempty" tf:"group_attribute,omitempty"`

	// The XML IdP metadata file generated from your identity provider.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Session timeout, in minutes. Minimum is 5 minutes and maximum is 720 minutes (12 hours). Default is 60 minutes.
	SessionTimeout *float64 `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`

	// User attribute for this SAML integration.
	UserAttribute *string `json:"userAttribute,omitempty" tf:"user_attribute,omitempty"`
}

type SAMLOptionsObservation struct {

	// Group attribute for this SAML integration.
	GroupAttribute *string `json:"groupAttribute,omitempty" tf:"group_attribute,omitempty"`

	// The XML IdP metadata file generated from your identity provider.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Session timeout, in minutes. Minimum is 5 minutes and maximum is 720 minutes (12 hours). Default is 60 minutes.
	SessionTimeout *float64 `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`

	// User attribute for this SAML integration.
	UserAttribute *string `json:"userAttribute,omitempty" tf:"user_attribute,omitempty"`
}

type SAMLOptionsParameters struct {

	// Group attribute for this SAML integration.
	// +kubebuilder:validation:Optional
	GroupAttribute *string `json:"groupAttribute,omitempty" tf:"group_attribute,omitempty"`

	// The XML IdP metadata file generated from your identity provider.
	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata" tf:"metadata,omitempty"`

	// Session timeout, in minutes. Minimum is 5 minutes and maximum is 720 minutes (12 hours). Default is 60 minutes.
	// +kubebuilder:validation:Optional
	SessionTimeout *float64 `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`

	// User attribute for this SAML integration.
	// +kubebuilder:validation:Optional
	UserAttribute *string `json:"userAttribute,omitempty" tf:"user_attribute,omitempty"`
}

type SecurityConfigInitParameters struct {

	// Description of the security configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Configuration block for SAML options.
	SAMLOptions *SAMLOptionsInitParameters `json:"samlOptions,omitempty" tf:"saml_options,omitempty"`
}

type SecurityConfigObservation struct {

	// Version of the configuration.
	ConfigVersion *string `json:"configVersion,omitempty" tf:"config_version,omitempty"`

	// Description of the security configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration block for SAML options.
	SAMLOptions *SAMLOptionsObservation `json:"samlOptions,omitempty" tf:"saml_options,omitempty"`

	// Type of configuration. Must be saml.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SecurityConfigParameters struct {

	// Description of the security configuration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configuration block for SAML options.
	// +kubebuilder:validation:Optional
	SAMLOptions *SAMLOptionsParameters `json:"samlOptions,omitempty" tf:"saml_options,omitempty"`

	// Type of configuration. Must be saml.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// SecurityConfigSpec defines the desired state of SecurityConfig
type SecurityConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityConfigParameters `json:"forProvider"`
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
	InitProvider SecurityConfigInitParameters `json:"initProvider,omitempty"`
}

// SecurityConfigStatus defines the observed state of SecurityConfig.
type SecurityConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecurityConfig is the Schema for the SecurityConfigs API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SecurityConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.samlOptions) || (has(self.initProvider) && has(self.initProvider.samlOptions))",message="spec.forProvider.samlOptions is a required parameter"
	Spec   SecurityConfigSpec   `json:"spec"`
	Status SecurityConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityConfigList contains a list of SecurityConfigs
type SecurityConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityConfig `json:"items"`
}

// Repository type metadata.
var (
	SecurityConfig_Kind             = "SecurityConfig"
	SecurityConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityConfig_Kind}.String()
	SecurityConfig_KindAPIVersion   = SecurityConfig_Kind + "." + CRDGroupVersion.String()
	SecurityConfig_GroupVersionKind = CRDGroupVersion.WithKind(SecurityConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityConfig{}, &SecurityConfigList{})
}
