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

type ProductInitParameters struct {

	// Language code. Valid values: en (English), jp (Japanese), zh (Chinese). Default value is en.
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language,omitempty"`

	// Description of the product.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Distributor (i.e., vendor) of the product.
	Distributor *string `json:"distributor,omitempty" tf:"distributor,omitempty"`

	// Name of the product.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Owner of the product.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Configuration block for provisioning artifact (i.e., version) parameters. Detailed below.
	ProvisioningArtifactParameters []ProvisioningArtifactParametersInitParameters `json:"provisioningArtifactParameters,omitempty" tf:"provisioning_artifact_parameters,omitempty"`

	// Support information about the product.
	SupportDescription *string `json:"supportDescription,omitempty" tf:"support_description,omitempty"`

	// Contact email for product support.
	SupportEmail *string `json:"supportEmail,omitempty" tf:"support_email,omitempty"`

	// Contact URL for product support.
	SupportURL *string `json:"supportUrl,omitempty" tf:"support_url,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Type of product. See AWS Docs for valid list of values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProductObservation struct {

	// Language code. Valid values: en (English), jp (Japanese), zh (Chinese). Default value is en.
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language,omitempty"`

	// ARN of the product.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Time when the product was created.
	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time,omitempty"`

	// Description of the product.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Distributor (i.e., vendor) of the product.
	Distributor *string `json:"distributor,omitempty" tf:"distributor,omitempty"`

	// Whether the product has a default path. If the product does not have a default path, call ListLaunchPaths to disambiguate between paths.  Otherwise, ListLaunchPaths is not required, and the output of ProductViewSummary can be used directly with DescribeProvisioningParameters.
	HasDefaultPath *bool `json:"hasDefaultPath,omitempty" tf:"has_default_path,omitempty"`

	// Product ID. For example, prod-dnigbtea24ste.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the product.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Owner of the product.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Configuration block for provisioning artifact (i.e., version) parameters. Detailed below.
	ProvisioningArtifactParameters []ProvisioningArtifactParametersObservation `json:"provisioningArtifactParameters,omitempty" tf:"provisioning_artifact_parameters,omitempty"`

	// Status of the product.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Support information about the product.
	SupportDescription *string `json:"supportDescription,omitempty" tf:"support_description,omitempty"`

	// Contact email for product support.
	SupportEmail *string `json:"supportEmail,omitempty" tf:"support_email,omitempty"`

	// Contact URL for product support.
	SupportURL *string `json:"supportUrl,omitempty" tf:"support_url,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Type of product. See AWS Docs for valid list of values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProductParameters struct {

	// Language code. Valid values: en (English), jp (Japanese), zh (Chinese). Default value is en.
	// +kubebuilder:validation:Optional
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language,omitempty"`

	// Description of the product.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Distributor (i.e., vendor) of the product.
	// +kubebuilder:validation:Optional
	Distributor *string `json:"distributor,omitempty" tf:"distributor,omitempty"`

	// Name of the product.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Owner of the product.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Configuration block for provisioning artifact (i.e., version) parameters. Detailed below.
	// +kubebuilder:validation:Optional
	ProvisioningArtifactParameters []ProvisioningArtifactParametersParameters `json:"provisioningArtifactParameters,omitempty" tf:"provisioning_artifact_parameters,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Support information about the product.
	// +kubebuilder:validation:Optional
	SupportDescription *string `json:"supportDescription,omitempty" tf:"support_description,omitempty"`

	// Contact email for product support.
	// +kubebuilder:validation:Optional
	SupportEmail *string `json:"supportEmail,omitempty" tf:"support_email,omitempty"`

	// Contact URL for product support.
	// +kubebuilder:validation:Optional
	SupportURL *string `json:"supportUrl,omitempty" tf:"support_url,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Type of product. See AWS Docs for valid list of values.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProvisioningArtifactParametersInitParameters struct {

	// Description of the provisioning artifact (i.e., version), including how it differs from the previous provisioning artifact.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether AWS Service Catalog stops validating the specified provisioning artifact template even if it is invalid.
	DisableTemplateValidation *bool `json:"disableTemplateValidation,omitempty" tf:"disable_template_validation,omitempty"`

	// Name of the provisioning artifact (for example, v1, v2beta). No spaces are allowed.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Template source as the physical ID of the resource that contains the template. Currently only supports CloudFormation stack ARN. Specify the physical ID as arn:[partition]:cloudformation:[region]:[account ID]:stack/[stack name]/[resource ID].
	TemplatePhysicalID *string `json:"templatePhysicalId,omitempty" tf:"template_physical_id,omitempty"`

	// Template source as URL of the CloudFormation template in Amazon S3.
	TemplateURL *string `json:"templateUrl,omitempty" tf:"template_url,omitempty"`

	// Type of provisioning artifact. See AWS Docs for valid list of values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProvisioningArtifactParametersObservation struct {

	// Description of the provisioning artifact (i.e., version), including how it differs from the previous provisioning artifact.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether AWS Service Catalog stops validating the specified provisioning artifact template even if it is invalid.
	DisableTemplateValidation *bool `json:"disableTemplateValidation,omitempty" tf:"disable_template_validation,omitempty"`

	// Name of the provisioning artifact (for example, v1, v2beta). No spaces are allowed.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Template source as the physical ID of the resource that contains the template. Currently only supports CloudFormation stack ARN. Specify the physical ID as arn:[partition]:cloudformation:[region]:[account ID]:stack/[stack name]/[resource ID].
	TemplatePhysicalID *string `json:"templatePhysicalId,omitempty" tf:"template_physical_id,omitempty"`

	// Template source as URL of the CloudFormation template in Amazon S3.
	TemplateURL *string `json:"templateUrl,omitempty" tf:"template_url,omitempty"`

	// Type of provisioning artifact. See AWS Docs for valid list of values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProvisioningArtifactParametersParameters struct {

	// Description of the provisioning artifact (i.e., version), including how it differs from the previous provisioning artifact.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether AWS Service Catalog stops validating the specified provisioning artifact template even if it is invalid.
	// +kubebuilder:validation:Optional
	DisableTemplateValidation *bool `json:"disableTemplateValidation,omitempty" tf:"disable_template_validation,omitempty"`

	// Name of the provisioning artifact (for example, v1, v2beta). No spaces are allowed.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Template source as the physical ID of the resource that contains the template. Currently only supports CloudFormation stack ARN. Specify the physical ID as arn:[partition]:cloudformation:[region]:[account ID]:stack/[stack name]/[resource ID].
	// +kubebuilder:validation:Optional
	TemplatePhysicalID *string `json:"templatePhysicalId,omitempty" tf:"template_physical_id,omitempty"`

	// Template source as URL of the CloudFormation template in Amazon S3.
	// +kubebuilder:validation:Optional
	TemplateURL *string `json:"templateUrl,omitempty" tf:"template_url,omitempty"`

	// Type of provisioning artifact. See AWS Docs for valid list of values.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// ProductSpec defines the desired state of Product
type ProductSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProductParameters `json:"forProvider"`
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
	InitProvider ProductInitParameters `json:"initProvider,omitempty"`
}

// ProductStatus defines the observed state of Product.
type ProductStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProductObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Product is the Schema for the Products API. Manages a Service Catalog Product
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Product struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.owner) || (has(self.initProvider) && has(self.initProvider.owner))",message="spec.forProvider.owner is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.provisioningArtifactParameters) || (has(self.initProvider) && has(self.initProvider.provisioningArtifactParameters))",message="spec.forProvider.provisioningArtifactParameters is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   ProductSpec   `json:"spec"`
	Status ProductStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProductList contains a list of Products
type ProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Product `json:"items"`
}

// Repository type metadata.
var (
	Product_Kind             = "Product"
	Product_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Product_Kind}.String()
	Product_KindAPIVersion   = Product_Kind + "." + CRDGroupVersion.String()
	Product_GroupVersionKind = CRDGroupVersion.WithKind(Product_Kind)
)

func init() {
	SchemeBuilder.Register(&Product{}, &ProductList{})
}
