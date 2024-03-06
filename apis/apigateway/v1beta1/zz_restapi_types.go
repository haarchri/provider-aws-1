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

type RestAPIEndpointConfigurationInitParameters struct {

	// List of endpoint types. This resource currently only supports managing a single value. Valid values: EDGE, REGIONAL or PRIVATE. If unspecified, defaults to EDGE. If set to PRIVATE recommend to set put_rest_api_mode = merge to not cause the endpoints and associated Route53 records to be deleted. Refer to the documentation for more information on the difference between edge-optimized and regional APIs.
	Types []*string `json:"types,omitempty" tf:"types,omitempty"`

	// Set of VPC Endpoint identifiers. It is only supported for PRIVATE endpoint type. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-endpoint-configuration extension vpcEndpointIds property. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	// +listType=set
	VPCEndpointIds []*string `json:"vpcEndpointIds,omitempty" tf:"vpc_endpoint_ids,omitempty"`
}

type RestAPIEndpointConfigurationObservation struct {

	// List of endpoint types. This resource currently only supports managing a single value. Valid values: EDGE, REGIONAL or PRIVATE. If unspecified, defaults to EDGE. If set to PRIVATE recommend to set put_rest_api_mode = merge to not cause the endpoints and associated Route53 records to be deleted. Refer to the documentation for more information on the difference between edge-optimized and regional APIs.
	Types []*string `json:"types,omitempty" tf:"types,omitempty"`

	// Set of VPC Endpoint identifiers. It is only supported for PRIVATE endpoint type. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-endpoint-configuration extension vpcEndpointIds property. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	// +listType=set
	VPCEndpointIds []*string `json:"vpcEndpointIds,omitempty" tf:"vpc_endpoint_ids,omitempty"`
}

type RestAPIEndpointConfigurationParameters struct {

	// List of endpoint types. This resource currently only supports managing a single value. Valid values: EDGE, REGIONAL or PRIVATE. If unspecified, defaults to EDGE. If set to PRIVATE recommend to set put_rest_api_mode = merge to not cause the endpoints and associated Route53 records to be deleted. Refer to the documentation for more information on the difference between edge-optimized and regional APIs.
	// +kubebuilder:validation:Optional
	Types []*string `json:"types" tf:"types,omitempty"`

	// Set of VPC Endpoint identifiers. It is only supported for PRIVATE endpoint type. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-endpoint-configuration extension vpcEndpointIds property. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	// +kubebuilder:validation:Optional
	// +listType=set
	VPCEndpointIds []*string `json:"vpcEndpointIds,omitempty" tf:"vpc_endpoint_ids,omitempty"`
}

type RestAPIInitParameters struct {

	// Source of the API key for requests. Valid values are HEADER (default) and AUTHORIZER. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-api-key-source extension. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	APIKeySource *string `json:"apiKeySource,omitempty" tf:"api_key_source,omitempty"`

	// List of binary media types supported by the REST API. By default, the REST API supports only UTF-8-encoded text payloads. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-binary-media-types extension. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	BinaryMediaTypes []*string `json:"binaryMediaTypes,omitempty" tf:"binary_media_types,omitempty"`

	// OpenAPI specification that defines the set of routes and integrations to create as part of the REST API. This configuration, and any updates to it, will replace all REST API configuration except values overridden in this resource configuration and other resource updates applied after this resource but before any aws_api_gateway_deployment creation. More information about REST API OpenAPI support can be found in the API Gateway Developer Guide.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Description of the REST API. If importing an OpenAPI specification via the body argument, this corresponds to the info.description field. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether clients can invoke your API by using the default execute-api endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint. Defaults to false. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-endpoint-configuration extension disableExecuteApiEndpoint property. If the argument value is true and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	DisableExecuteAPIEndpoint *bool `json:"disableExecuteApiEndpoint,omitempty" tf:"disable_execute_api_endpoint,omitempty"`

	// Configuration block defining API endpoint configuration including endpoint type. Defined below.
	EndpointConfiguration []RestAPIEndpointConfigurationInitParameters `json:"endpointConfiguration,omitempty" tf:"endpoint_configuration,omitempty"`

	// Whether warnings while API Gateway is creating or updating the resource should return an error or not. Defaults to false
	FailOnWarnings *bool `json:"failOnWarnings,omitempty" tf:"fail_on_warnings,omitempty"`

	// Minimum response size to compress for the REST API. String containing an integer value between -1 and 10485760 (10MB). -1 will disable an existing compression configuration, and all other values will enable compression with the configured size. New resources can simply omit this argument to disable compression, rather than setting the value to -1. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-minimum-compression-size extension. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	MinimumCompressionSize *string `json:"minimumCompressionSize,omitempty" tf:"minimum_compression_size,omitempty"`

	// Name of the REST API. If importing an OpenAPI specification via the body argument, this corresponds to the info.title field. If the argument value is different than the OpenAPI value, the argument value will override the OpenAPI value.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Map of customizations for importing the specification in the body argument. For example, to exclude DocumentationParts from an imported API, set ignore equal to documentation. Additional documentation, including other parameters such as basepath, can be found in the API Gateway Developer Guide.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Mode of the PutRestApi operation when importing an OpenAPI specification via the body argument (create or update operation). Valid values are merge and overwrite. If unspecificed, defaults to overwrite (for backwards compatibility). This corresponds to the x-amazon-apigateway-put-integration-method extension. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	PutRestAPIMode *string `json:"putRestApiMode,omitempty" tf:"put_rest_api_mode,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RestAPIObservation struct {

	// Source of the API key for requests. Valid values are HEADER (default) and AUTHORIZER. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-api-key-source extension. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	APIKeySource *string `json:"apiKeySource,omitempty" tf:"api_key_source,omitempty"`

	// ARN
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// List of binary media types supported by the REST API. By default, the REST API supports only UTF-8-encoded text payloads. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-binary-media-types extension. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	BinaryMediaTypes []*string `json:"binaryMediaTypes,omitempty" tf:"binary_media_types,omitempty"`

	// OpenAPI specification that defines the set of routes and integrations to create as part of the REST API. This configuration, and any updates to it, will replace all REST API configuration except values overridden in this resource configuration and other resource updates applied after this resource but before any aws_api_gateway_deployment creation. More information about REST API OpenAPI support can be found in the API Gateway Developer Guide.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Creation date of the REST API
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	// Description of the REST API. If importing an OpenAPI specification via the body argument, this corresponds to the info.description field. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether clients can invoke your API by using the default execute-api endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint. Defaults to false. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-endpoint-configuration extension disableExecuteApiEndpoint property. If the argument value is true and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	DisableExecuteAPIEndpoint *bool `json:"disableExecuteApiEndpoint,omitempty" tf:"disable_execute_api_endpoint,omitempty"`

	// Configuration block defining API endpoint configuration including endpoint type. Defined below.
	EndpointConfiguration []RestAPIEndpointConfigurationObservation `json:"endpointConfiguration,omitempty" tf:"endpoint_configuration,omitempty"`

	// Execution ARN part to be used in lambda_permission's source_arn
	// when allowing API Gateway to invoke a Lambda function,
	// e.g., arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j, which can be concatenated with allowed stage, method and resource path.
	ExecutionArn *string `json:"executionArn,omitempty" tf:"execution_arn,omitempty"`

	// Whether warnings while API Gateway is creating or updating the resource should return an error or not. Defaults to false
	FailOnWarnings *bool `json:"failOnWarnings,omitempty" tf:"fail_on_warnings,omitempty"`

	// ID of the REST API
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Minimum response size to compress for the REST API. String containing an integer value between -1 and 10485760 (10MB). -1 will disable an existing compression configuration, and all other values will enable compression with the configured size. New resources can simply omit this argument to disable compression, rather than setting the value to -1. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-minimum-compression-size extension. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	MinimumCompressionSize *string `json:"minimumCompressionSize,omitempty" tf:"minimum_compression_size,omitempty"`

	// Name of the REST API. If importing an OpenAPI specification via the body argument, this corresponds to the info.title field. If the argument value is different than the OpenAPI value, the argument value will override the OpenAPI value.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Map of customizations for importing the specification in the body argument. For example, to exclude DocumentationParts from an imported API, set ignore equal to documentation. Additional documentation, including other parameters such as basepath, can be found in the API Gateway Developer Guide.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// JSON formatted policy document that controls access to the API Gateway. We recommend using the aws_api_gateway_rest_api_policy resource instead. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-policy extension. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Mode of the PutRestApi operation when importing an OpenAPI specification via the body argument (create or update operation). Valid values are merge and overwrite. If unspecificed, defaults to overwrite (for backwards compatibility). This corresponds to the x-amazon-apigateway-put-integration-method extension. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	PutRestAPIMode *string `json:"putRestApiMode,omitempty" tf:"put_rest_api_mode,omitempty"`

	// Resource ID of the REST API's root
	RootResourceID *string `json:"rootResourceId,omitempty" tf:"root_resource_id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RestAPIParameters struct {

	// Source of the API key for requests. Valid values are HEADER (default) and AUTHORIZER. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-api-key-source extension. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	// +kubebuilder:validation:Optional
	APIKeySource *string `json:"apiKeySource,omitempty" tf:"api_key_source,omitempty"`

	// List of binary media types supported by the REST API. By default, the REST API supports only UTF-8-encoded text payloads. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-binary-media-types extension. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	// +kubebuilder:validation:Optional
	BinaryMediaTypes []*string `json:"binaryMediaTypes,omitempty" tf:"binary_media_types,omitempty"`

	// OpenAPI specification that defines the set of routes and integrations to create as part of the REST API. This configuration, and any updates to it, will replace all REST API configuration except values overridden in this resource configuration and other resource updates applied after this resource but before any aws_api_gateway_deployment creation. More information about REST API OpenAPI support can be found in the API Gateway Developer Guide.
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Description of the REST API. If importing an OpenAPI specification via the body argument, this corresponds to the info.description field. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether clients can invoke your API by using the default execute-api endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint. Defaults to false. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-endpoint-configuration extension disableExecuteApiEndpoint property. If the argument value is true and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	// +kubebuilder:validation:Optional
	DisableExecuteAPIEndpoint *bool `json:"disableExecuteApiEndpoint,omitempty" tf:"disable_execute_api_endpoint,omitempty"`

	// Configuration block defining API endpoint configuration including endpoint type. Defined below.
	// +kubebuilder:validation:Optional
	EndpointConfiguration []RestAPIEndpointConfigurationParameters `json:"endpointConfiguration,omitempty" tf:"endpoint_configuration,omitempty"`

	// Whether warnings while API Gateway is creating or updating the resource should return an error or not. Defaults to false
	// +kubebuilder:validation:Optional
	FailOnWarnings *bool `json:"failOnWarnings,omitempty" tf:"fail_on_warnings,omitempty"`

	// Minimum response size to compress for the REST API. String containing an integer value between -1 and 10485760 (10MB). -1 will disable an existing compression configuration, and all other values will enable compression with the configured size. New resources can simply omit this argument to disable compression, rather than setting the value to -1. If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-minimum-compression-size extension. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	// +kubebuilder:validation:Optional
	MinimumCompressionSize *string `json:"minimumCompressionSize,omitempty" tf:"minimum_compression_size,omitempty"`

	// Name of the REST API. If importing an OpenAPI specification via the body argument, this corresponds to the info.title field. If the argument value is different than the OpenAPI value, the argument value will override the OpenAPI value.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Map of customizations for importing the specification in the body argument. For example, to exclude DocumentationParts from an imported API, set ignore equal to documentation. Additional documentation, including other parameters such as basepath, can be found in the API Gateway Developer Guide.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Mode of the PutRestApi operation when importing an OpenAPI specification via the body argument (create or update operation). Valid values are merge and overwrite. If unspecificed, defaults to overwrite (for backwards compatibility). This corresponds to the x-amazon-apigateway-put-integration-method extension. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	// +kubebuilder:validation:Optional
	PutRestAPIMode *string `json:"putRestApiMode,omitempty" tf:"put_rest_api_mode,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RestAPISpec defines the desired state of RestAPI
type RestAPISpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RestAPIParameters `json:"forProvider"`
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
	InitProvider RestAPIInitParameters `json:"initProvider,omitempty"`
}

// RestAPIStatus defines the observed state of RestAPI.
type RestAPIStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RestAPIObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RestAPI is the Schema for the RestAPIs API. Manages an API Gateway REST API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RestAPI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   RestAPISpec   `json:"spec"`
	Status RestAPIStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RestAPIList contains a list of RestAPIs
type RestAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RestAPI `json:"items"`
}

// Repository type metadata.
var (
	RestAPI_Kind             = "RestAPI"
	RestAPI_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RestAPI_Kind}.String()
	RestAPI_KindAPIVersion   = RestAPI_Kind + "." + CRDGroupVersion.String()
	RestAPI_GroupVersionKind = CRDGroupVersion.WithKind(RestAPI_Kind)
)

func init() {
	SchemeBuilder.Register(&RestAPI{}, &RestAPIList{})
}
