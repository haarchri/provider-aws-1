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

type CustomerManagedPolicyAttachmentInitParameters struct {

	// Specifies the name and path of a customer managed policy. See below.
	CustomerManagedPolicyReference []CustomerManagedPolicyReferenceInitParameters `json:"customerManagedPolicyReference,omitempty" tf:"customer_managed_policy_reference,omitempty"`
}

type CustomerManagedPolicyAttachmentObservation struct {

	// Specifies the name and path of a customer managed policy. See below.
	CustomerManagedPolicyReference []CustomerManagedPolicyReferenceObservation `json:"customerManagedPolicyReference,omitempty" tf:"customer_managed_policy_reference,omitempty"`

	// Policy Name, Policy Path, Permission Set Amazon Resource Name (ARN), and SSO Instance ARN, each separated by a comma (,).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn *string `json:"instanceArn,omitempty" tf:"instance_arn,omitempty"`

	// The Amazon Resource Name (ARN) of the Permission Set.
	PermissionSetArn *string `json:"permissionSetArn,omitempty" tf:"permission_set_arn,omitempty"`
}

type CustomerManagedPolicyAttachmentParameters struct {

	// Specifies the name and path of a customer managed policy. See below.
	// +kubebuilder:validation:Optional
	CustomerManagedPolicyReference []CustomerManagedPolicyReferenceParameters `json:"customerManagedPolicyReference,omitempty" tf:"customer_managed_policy_reference,omitempty"`

	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	// +kubebuilder:validation:Required
	InstanceArn *string `json:"instanceArn" tf:"instance_arn,omitempty"`

	// The Amazon Resource Name (ARN) of the Permission Set.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/ssoadmin/v1beta1.PermissionSet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	PermissionSetArn *string `json:"permissionSetArn,omitempty" tf:"permission_set_arn,omitempty"`

	// Reference to a PermissionSet in ssoadmin to populate permissionSetArn.
	// +kubebuilder:validation:Optional
	PermissionSetArnRef *v1.Reference `json:"permissionSetArnRef,omitempty" tf:"-"`

	// Selector for a PermissionSet in ssoadmin to populate permissionSetArn.
	// +kubebuilder:validation:Optional
	PermissionSetArnSelector *v1.Selector `json:"permissionSetArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type CustomerManagedPolicyReferenceInitParameters struct {

	// Name of the customer managed IAM Policy to be attached.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/iam/v1beta1.Policy
	// +crossplane:generate:reference:refFieldName=PolicyNameRef
	// +crossplane:generate:reference:selectorFieldName=PolicyNameSelector
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The path to the IAM policy to be attached. The default is /. See IAM Identifiers for more information.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Reference to a Policy in iam to populate name.
	// +kubebuilder:validation:Optional
	PolicyNameRef *v1.Reference `json:"policyNameRef,omitempty" tf:"-"`

	// Selector for a Policy in iam to populate name.
	// +kubebuilder:validation:Optional
	PolicyNameSelector *v1.Selector `json:"policyNameSelector,omitempty" tf:"-"`
}

type CustomerManagedPolicyReferenceObservation struct {

	// Name of the customer managed IAM Policy to be attached.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The path to the IAM policy to be attached. The default is /. See IAM Identifiers for more information.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type CustomerManagedPolicyReferenceParameters struct {

	// Name of the customer managed IAM Policy to be attached.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/iam/v1beta1.Policy
	// +crossplane:generate:reference:refFieldName=PolicyNameRef
	// +crossplane:generate:reference:selectorFieldName=PolicyNameSelector
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The path to the IAM policy to be attached. The default is /. See IAM Identifiers for more information.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Reference to a Policy in iam to populate name.
	// +kubebuilder:validation:Optional
	PolicyNameRef *v1.Reference `json:"policyNameRef,omitempty" tf:"-"`

	// Selector for a Policy in iam to populate name.
	// +kubebuilder:validation:Optional
	PolicyNameSelector *v1.Selector `json:"policyNameSelector,omitempty" tf:"-"`
}

// CustomerManagedPolicyAttachmentSpec defines the desired state of CustomerManagedPolicyAttachment
type CustomerManagedPolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomerManagedPolicyAttachmentParameters `json:"forProvider"`
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
	InitProvider CustomerManagedPolicyAttachmentInitParameters `json:"initProvider,omitempty"`
}

// CustomerManagedPolicyAttachmentStatus defines the observed state of CustomerManagedPolicyAttachment.
type CustomerManagedPolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomerManagedPolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CustomerManagedPolicyAttachment is the Schema for the CustomerManagedPolicyAttachments API. Manages a customer managed policy for a Single Sign-On (SSO) Permission Set
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CustomerManagedPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.customerManagedPolicyReference) || (has(self.initProvider) && has(self.initProvider.customerManagedPolicyReference))",message="spec.forProvider.customerManagedPolicyReference is a required parameter"
	Spec   CustomerManagedPolicyAttachmentSpec   `json:"spec"`
	Status CustomerManagedPolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomerManagedPolicyAttachmentList contains a list of CustomerManagedPolicyAttachments
type CustomerManagedPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomerManagedPolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	CustomerManagedPolicyAttachment_Kind             = "CustomerManagedPolicyAttachment"
	CustomerManagedPolicyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomerManagedPolicyAttachment_Kind}.String()
	CustomerManagedPolicyAttachment_KindAPIVersion   = CustomerManagedPolicyAttachment_Kind + "." + CRDGroupVersion.String()
	CustomerManagedPolicyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(CustomerManagedPolicyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomerManagedPolicyAttachment{}, &CustomerManagedPolicyAttachmentList{})
}
