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

type AwsLambdaInitParameters struct {

	// The Amazon Resource Name (ARN) of the AWS Lambda function.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta1.Function
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	FunctionArn *string `json:"functionArn,omitempty" tf:"function_arn,omitempty"`

	// Reference to a Function in lambda to populate functionArn.
	// +kubebuilder:validation:Optional
	FunctionArnRef *v1.Reference `json:"functionArnRef,omitempty" tf:"-"`

	// Selector for a Function in lambda to populate functionArn.
	// +kubebuilder:validation:Optional
	FunctionArnSelector *v1.Selector `json:"functionArnSelector,omitempty" tf:"-"`

	// Additional JSON that provides supplemental data to the Lambda function used to transform objects.
	FunctionPayload *string `json:"functionPayload,omitempty" tf:"function_payload,omitempty"`
}

type AwsLambdaObservation struct {

	// The Amazon Resource Name (ARN) of the AWS Lambda function.
	FunctionArn *string `json:"functionArn,omitempty" tf:"function_arn,omitempty"`

	// Additional JSON that provides supplemental data to the Lambda function used to transform objects.
	FunctionPayload *string `json:"functionPayload,omitempty" tf:"function_payload,omitempty"`
}

type AwsLambdaParameters struct {

	// The Amazon Resource Name (ARN) of the AWS Lambda function.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta1.Function
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	FunctionArn *string `json:"functionArn,omitempty" tf:"function_arn,omitempty"`

	// Reference to a Function in lambda to populate functionArn.
	// +kubebuilder:validation:Optional
	FunctionArnRef *v1.Reference `json:"functionArnRef,omitempty" tf:"-"`

	// Selector for a Function in lambda to populate functionArn.
	// +kubebuilder:validation:Optional
	FunctionArnSelector *v1.Selector `json:"functionArnSelector,omitempty" tf:"-"`

	// Additional JSON that provides supplemental data to the Lambda function used to transform objects.
	// +kubebuilder:validation:Optional
	FunctionPayload *string `json:"functionPayload,omitempty" tf:"function_payload,omitempty"`
}

type ConfigurationInitParameters struct {

	// Allowed features. Valid values: GetObject-Range, GetObject-PartNumber.
	// +listType=set
	AllowedFeatures []*string `json:"allowedFeatures,omitempty" tf:"allowed_features,omitempty"`

	// Whether or not the CloudWatch metrics configuration is enabled.
	CloudWatchMetricsEnabled *bool `json:"cloudWatchMetricsEnabled,omitempty" tf:"cloud_watch_metrics_enabled,omitempty"`

	// Standard access point associated with the Object Lambda Access Point.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3control/v1beta1.AccessPoint
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	SupportingAccessPoint *string `json:"supportingAccessPoint,omitempty" tf:"supporting_access_point,omitempty"`

	// Reference to a AccessPoint in s3control to populate supportingAccessPoint.
	// +kubebuilder:validation:Optional
	SupportingAccessPointRef *v1.Reference `json:"supportingAccessPointRef,omitempty" tf:"-"`

	// Selector for a AccessPoint in s3control to populate supportingAccessPoint.
	// +kubebuilder:validation:Optional
	SupportingAccessPointSelector *v1.Selector `json:"supportingAccessPointSelector,omitempty" tf:"-"`

	// List of transformation configurations for the Object Lambda Access Point. See Transformation Configuration below for more details.
	TransformationConfiguration []TransformationConfigurationInitParameters `json:"transformationConfiguration,omitempty" tf:"transformation_configuration,omitempty"`
}

type ConfigurationObservation struct {

	// Allowed features. Valid values: GetObject-Range, GetObject-PartNumber.
	// +listType=set
	AllowedFeatures []*string `json:"allowedFeatures,omitempty" tf:"allowed_features,omitempty"`

	// Whether or not the CloudWatch metrics configuration is enabled.
	CloudWatchMetricsEnabled *bool `json:"cloudWatchMetricsEnabled,omitempty" tf:"cloud_watch_metrics_enabled,omitempty"`

	// Standard access point associated with the Object Lambda Access Point.
	SupportingAccessPoint *string `json:"supportingAccessPoint,omitempty" tf:"supporting_access_point,omitempty"`

	// List of transformation configurations for the Object Lambda Access Point. See Transformation Configuration below for more details.
	TransformationConfiguration []TransformationConfigurationObservation `json:"transformationConfiguration,omitempty" tf:"transformation_configuration,omitempty"`
}

type ConfigurationParameters struct {

	// Allowed features. Valid values: GetObject-Range, GetObject-PartNumber.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedFeatures []*string `json:"allowedFeatures,omitempty" tf:"allowed_features,omitempty"`

	// Whether or not the CloudWatch metrics configuration is enabled.
	// +kubebuilder:validation:Optional
	CloudWatchMetricsEnabled *bool `json:"cloudWatchMetricsEnabled,omitempty" tf:"cloud_watch_metrics_enabled,omitempty"`

	// Standard access point associated with the Object Lambda Access Point.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3control/v1beta1.AccessPoint
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	SupportingAccessPoint *string `json:"supportingAccessPoint,omitempty" tf:"supporting_access_point,omitempty"`

	// Reference to a AccessPoint in s3control to populate supportingAccessPoint.
	// +kubebuilder:validation:Optional
	SupportingAccessPointRef *v1.Reference `json:"supportingAccessPointRef,omitempty" tf:"-"`

	// Selector for a AccessPoint in s3control to populate supportingAccessPoint.
	// +kubebuilder:validation:Optional
	SupportingAccessPointSelector *v1.Selector `json:"supportingAccessPointSelector,omitempty" tf:"-"`

	// List of transformation configurations for the Object Lambda Access Point. See Transformation Configuration below for more details.
	// +kubebuilder:validation:Optional
	TransformationConfiguration []TransformationConfigurationParameters `json:"transformationConfiguration" tf:"transformation_configuration,omitempty"`
}

type ContentTransformationInitParameters struct {

	// Configuration for an AWS Lambda function. See AWS Lambda below for more details.
	AwsLambda []AwsLambdaInitParameters `json:"awsLambda,omitempty" tf:"aws_lambda,omitempty"`
}

type ContentTransformationObservation struct {

	// Configuration for an AWS Lambda function. See AWS Lambda below for more details.
	AwsLambda []AwsLambdaObservation `json:"awsLambda,omitempty" tf:"aws_lambda,omitempty"`
}

type ContentTransformationParameters struct {

	// Configuration for an AWS Lambda function. See AWS Lambda below for more details.
	// +kubebuilder:validation:Optional
	AwsLambda []AwsLambdaParameters `json:"awsLambda" tf:"aws_lambda,omitempty"`
}

type ObjectLambdaAccessPointInitParameters struct {

	// The AWS account ID for the owner of the bucket for which you want to create an Object Lambda Access Point.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// A configuration block containing details about the Object Lambda Access Point. See Configuration below for more details.
	Configuration []ConfigurationInitParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// The name for this Object Lambda Access Point.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ObjectLambdaAccessPointObservation struct {

	// The AWS account ID for the owner of the bucket for which you want to create an Object Lambda Access Point.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Alias for the S3 Object Lambda Access Point.
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Amazon Resource Name (ARN) of the Object Lambda Access Point.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A configuration block containing details about the Object Lambda Access Point. See Configuration below for more details.
	Configuration []ConfigurationObservation `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// The AWS account ID and access point name separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name for this Object Lambda Access Point.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ObjectLambdaAccessPointParameters struct {

	// The AWS account ID for the owner of the bucket for which you want to create an Object Lambda Access Point.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// A configuration block containing details about the Object Lambda Access Point. See Configuration below for more details.
	// +kubebuilder:validation:Optional
	Configuration []ConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// The name for this Object Lambda Access Point.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type TransformationConfigurationInitParameters struct {

	// The actions of an Object Lambda Access Point configuration. Valid values: GetObject.
	// +listType=set
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// The content transformation of an Object Lambda Access Point configuration. See Content Transformation below for more details.
	ContentTransformation []ContentTransformationInitParameters `json:"contentTransformation,omitempty" tf:"content_transformation,omitempty"`
}

type TransformationConfigurationObservation struct {

	// The actions of an Object Lambda Access Point configuration. Valid values: GetObject.
	// +listType=set
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// The content transformation of an Object Lambda Access Point configuration. See Content Transformation below for more details.
	ContentTransformation []ContentTransformationObservation `json:"contentTransformation,omitempty" tf:"content_transformation,omitempty"`
}

type TransformationConfigurationParameters struct {

	// The actions of an Object Lambda Access Point configuration. Valid values: GetObject.
	// +kubebuilder:validation:Optional
	// +listType=set
	Actions []*string `json:"actions" tf:"actions,omitempty"`

	// The content transformation of an Object Lambda Access Point configuration. See Content Transformation below for more details.
	// +kubebuilder:validation:Optional
	ContentTransformation []ContentTransformationParameters `json:"contentTransformation" tf:"content_transformation,omitempty"`
}

// ObjectLambdaAccessPointSpec defines the desired state of ObjectLambdaAccessPoint
type ObjectLambdaAccessPointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ObjectLambdaAccessPointParameters `json:"forProvider"`
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
	InitProvider ObjectLambdaAccessPointInitParameters `json:"initProvider,omitempty"`
}

// ObjectLambdaAccessPointStatus defines the observed state of ObjectLambdaAccessPoint.
type ObjectLambdaAccessPointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ObjectLambdaAccessPointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ObjectLambdaAccessPoint is the Schema for the ObjectLambdaAccessPoints API. Provides a resource to manage an S3 Object Lambda Access Point.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ObjectLambdaAccessPoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.configuration) || (has(self.initProvider) && has(self.initProvider.configuration))",message="spec.forProvider.configuration is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ObjectLambdaAccessPointSpec   `json:"spec"`
	Status ObjectLambdaAccessPointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectLambdaAccessPointList contains a list of ObjectLambdaAccessPoints
type ObjectLambdaAccessPointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectLambdaAccessPoint `json:"items"`
}

// Repository type metadata.
var (
	ObjectLambdaAccessPoint_Kind             = "ObjectLambdaAccessPoint"
	ObjectLambdaAccessPoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ObjectLambdaAccessPoint_Kind}.String()
	ObjectLambdaAccessPoint_KindAPIVersion   = ObjectLambdaAccessPoint_Kind + "." + CRDGroupVersion.String()
	ObjectLambdaAccessPoint_GroupVersionKind = CRDGroupVersion.WithKind(ObjectLambdaAccessPoint_Kind)
)

func init() {
	SchemeBuilder.Register(&ObjectLambdaAccessPoint{}, &ObjectLambdaAccessPointList{})
}
