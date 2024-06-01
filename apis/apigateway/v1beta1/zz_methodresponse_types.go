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

type MethodResponseInitParameters struct {

	// The HTTP verb of the method resource (GET, POST, PUT, DELETE, HEAD, OPTIONS, ANY).
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/apigateway/v1beta1.Method
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("http_method",false)
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// Reference to a Method in apigateway to populate httpMethod.
	// +kubebuilder:validation:Optional
	HTTPMethodRef *v1.Reference `json:"httpMethodRef,omitempty" tf:"-"`

	// Selector for a Method in apigateway to populate httpMethod.
	// +kubebuilder:validation:Optional
	HTTPMethodSelector *v1.Selector `json:"httpMethodSelector,omitempty" tf:"-"`

	// The Resource identifier for the method resource.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/apigateway/v1beta1.Resource
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a Resource in apigateway to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a Resource in apigateway to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`

	// A map specifying the model resources used for the response's content type. Response models are represented as a key/value map, with a content type as the key and a Model name as the value.
	// +mapType=granular
	ResponseModels map[string]*string `json:"responseModels,omitempty" tf:"response_models,omitempty"`

	// A map specifying required or optional response parameters that API Gateway can send back to the caller. A key defines a method response header name and the associated value is a boolean flag indicating whether the method response parameter is required. The method response header names must match the pattern of method.response.header.{name}, where name is a valid and unique header name.
	// +mapType=granular
	ResponseParameters map[string]*bool `json:"responseParameters,omitempty" tf:"response_parameters,omitempty"`

	// The string identifier of the associated REST API.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/apigateway/v1beta1.RestAPI
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDRef *v1.Reference `json:"restApiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDSelector *v1.Selector `json:"restApiIdSelector,omitempty" tf:"-"`

	// The method response's status code.
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type MethodResponseObservation struct {

	// The HTTP verb of the method resource (GET, POST, PUT, DELETE, HEAD, OPTIONS, ANY).
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Resource identifier for the method resource.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// A map specifying the model resources used for the response's content type. Response models are represented as a key/value map, with a content type as the key and a Model name as the value.
	// +mapType=granular
	ResponseModels map[string]*string `json:"responseModels,omitempty" tf:"response_models,omitempty"`

	// A map specifying required or optional response parameters that API Gateway can send back to the caller. A key defines a method response header name and the associated value is a boolean flag indicating whether the method response parameter is required. The method response header names must match the pattern of method.response.header.{name}, where name is a valid and unique header name.
	// +mapType=granular
	ResponseParameters map[string]*bool `json:"responseParameters,omitempty" tf:"response_parameters,omitempty"`

	// The string identifier of the associated REST API.
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// The method response's status code.
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type MethodResponseParameters struct {

	// The HTTP verb of the method resource (GET, POST, PUT, DELETE, HEAD, OPTIONS, ANY).
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/apigateway/v1beta1.Method
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("http_method",false)
	// +kubebuilder:validation:Optional
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// Reference to a Method in apigateway to populate httpMethod.
	// +kubebuilder:validation:Optional
	HTTPMethodRef *v1.Reference `json:"httpMethodRef,omitempty" tf:"-"`

	// Selector for a Method in apigateway to populate httpMethod.
	// +kubebuilder:validation:Optional
	HTTPMethodSelector *v1.Selector `json:"httpMethodSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Resource identifier for the method resource.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/apigateway/v1beta1.Resource
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a Resource in apigateway to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a Resource in apigateway to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`

	// A map specifying the model resources used for the response's content type. Response models are represented as a key/value map, with a content type as the key and a Model name as the value.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ResponseModels map[string]*string `json:"responseModels,omitempty" tf:"response_models,omitempty"`

	// A map specifying required or optional response parameters that API Gateway can send back to the caller. A key defines a method response header name and the associated value is a boolean flag indicating whether the method response parameter is required. The method response header names must match the pattern of method.response.header.{name}, where name is a valid and unique header name.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ResponseParameters map[string]*bool `json:"responseParameters,omitempty" tf:"response_parameters,omitempty"`

	// The string identifier of the associated REST API.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/apigateway/v1beta1.RestAPI
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDRef *v1.Reference `json:"restApiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDSelector *v1.Selector `json:"restApiIdSelector,omitempty" tf:"-"`

	// The method response's status code.
	// +kubebuilder:validation:Optional
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

// MethodResponseSpec defines the desired state of MethodResponse
type MethodResponseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MethodResponseParameters `json:"forProvider"`
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
	InitProvider MethodResponseInitParameters `json:"initProvider,omitempty"`
}

// MethodResponseStatus defines the observed state of MethodResponse.
type MethodResponseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MethodResponseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MethodResponse is the Schema for the MethodResponses API. Provides an HTTP Method Response for an API Gateway Resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MethodResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.statusCode) || (has(self.initProvider) && has(self.initProvider.statusCode))",message="spec.forProvider.statusCode is a required parameter"
	Spec   MethodResponseSpec   `json:"spec"`
	Status MethodResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MethodResponseList contains a list of MethodResponses
type MethodResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MethodResponse `json:"items"`
}

// Repository type metadata.
var (
	MethodResponse_Kind             = "MethodResponse"
	MethodResponse_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MethodResponse_Kind}.String()
	MethodResponse_KindAPIVersion   = MethodResponse_Kind + "." + CRDGroupVersion.String()
	MethodResponse_GroupVersionKind = CRDGroupVersion.WithKind(MethodResponse_Kind)
)

func init() {
	SchemeBuilder.Register(&MethodResponse{}, &MethodResponseList{})
}
