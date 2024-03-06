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

type CertificateInitParameters struct {

	// A domain name for which the certificate should be issued.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Set of domains that should be SANs in the issued certificate. domain_name attribute is automatically added as a Subject Alternative Name.
	// +listType=set
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CertificateObservation struct {

	// The ARN of the lightsail certificate.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The timestamp when the instance was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// A domain name for which the certificate should be issued.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Set of domain validation objects which can be used to complete certificate validation. Can have more than one element, e.g., if SANs are defined.
	DomainValidationOptions []DomainValidationOptionsObservation `json:"domainValidationOptions,omitempty" tf:"domain_validation_options,omitempty"`

	// The name of the lightsail certificate (matches name).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Set of domains that should be SANs in the issued certificate. domain_name attribute is automatically added as a Subject Alternative Name.
	// +listType=set
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CertificateParameters struct {

	// A domain name for which the certificate should be issued.
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Set of domains that should be SANs in the issued certificate. domain_name attribute is automatically added as a Subject Alternative Name.
	// +kubebuilder:validation:Optional
	// +listType=set
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DomainValidationOptionsInitParameters struct {
}

type DomainValidationOptionsObservation struct {

	// A domain name for which the certificate should be issued.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// The name of the Lightsail load balancer.
	ResourceRecordName *string `json:"resourceRecordName,omitempty" tf:"resource_record_name,omitempty"`

	ResourceRecordType *string `json:"resourceRecordType,omitempty" tf:"resource_record_type,omitempty"`

	ResourceRecordValue *string `json:"resourceRecordValue,omitempty" tf:"resource_record_value,omitempty"`
}

type DomainValidationOptionsParameters struct {
}

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateParameters `json:"forProvider"`
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
	InitProvider CertificateInitParameters `json:"initProvider,omitempty"`
}

// CertificateStatus defines the observed state of Certificate.
type CertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Certificate is the Schema for the Certificates API. Provides a lightsail certificate
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateSpec   `json:"spec"`
	Status            CertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateList contains a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

// Repository type metadata.
var (
	Certificate_Kind             = "Certificate"
	Certificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Certificate_Kind}.String()
	Certificate_KindAPIVersion   = Certificate_Kind + "." + CRDGroupVersion.String()
	Certificate_GroupVersionKind = CRDGroupVersion.WithKind(Certificate_Kind)
)

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}
