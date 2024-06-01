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

type ApplicationVersionInitParameters struct {

	// Name of the Beanstalk Application the version is associated with.
	Application *string `json:"application,omitempty" tf:"application,omitempty"`

	// S3 bucket that contains the Application Version source bundle.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Short description of the Application Version.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// On delete, force an Application Version to be deleted when it may be in use by multiple Elastic Beanstalk Environments.
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// S3 object that is the Application Version source bundle.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/s3/v1beta1.Object
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Reference to a Object in s3 to populate key.
	// +kubebuilder:validation:Optional
	KeyRef *v1.Reference `json:"keyRef,omitempty" tf:"-"`

	// Selector for a Object in s3 to populate key.
	// +kubebuilder:validation:Optional
	KeySelector *v1.Selector `json:"keySelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ApplicationVersionObservation struct {

	// Name of the Beanstalk Application the version is associated with.
	Application *string `json:"application,omitempty" tf:"application,omitempty"`

	// ARN assigned by AWS for this Elastic Beanstalk Application.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// S3 bucket that contains the Application Version source bundle.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Short description of the Application Version.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// On delete, force an Application Version to be deleted when it may be in use by multiple Elastic Beanstalk Environments.
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// S3 object that is the Application Version source bundle.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ApplicationVersionParameters struct {

	// Name of the Beanstalk Application the version is associated with.
	// +kubebuilder:validation:Optional
	Application *string `json:"application,omitempty" tf:"application,omitempty"`

	// S3 bucket that contains the Application Version source bundle.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Short description of the Application Version.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// On delete, force an Application Version to be deleted when it may be in use by multiple Elastic Beanstalk Environments.
	// +kubebuilder:validation:Optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// S3 object that is the Application Version source bundle.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/s3/v1beta1.Object
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Reference to a Object in s3 to populate key.
	// +kubebuilder:validation:Optional
	KeyRef *v1.Reference `json:"keyRef,omitempty" tf:"-"`

	// Selector for a Object in s3 to populate key.
	// +kubebuilder:validation:Optional
	KeySelector *v1.Selector `json:"keySelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ApplicationVersionSpec defines the desired state of ApplicationVersion
type ApplicationVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationVersionParameters `json:"forProvider"`
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
	InitProvider ApplicationVersionInitParameters `json:"initProvider,omitempty"`
}

// ApplicationVersionStatus defines the observed state of ApplicationVersion.
type ApplicationVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ApplicationVersion is the Schema for the ApplicationVersions API. Provides an Elastic Beanstalk Application Version Resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ApplicationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.application) || (has(self.initProvider) && has(self.initProvider.application))",message="spec.forProvider.application is a required parameter"
	Spec   ApplicationVersionSpec   `json:"spec"`
	Status ApplicationVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationVersionList contains a list of ApplicationVersions
type ApplicationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationVersion `json:"items"`
}

// Repository type metadata.
var (
	ApplicationVersion_Kind             = "ApplicationVersion"
	ApplicationVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationVersion_Kind}.String()
	ApplicationVersion_KindAPIVersion   = ApplicationVersion_Kind + "." + CRDGroupVersion.String()
	ApplicationVersion_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationVersion{}, &ApplicationVersionList{})
}
