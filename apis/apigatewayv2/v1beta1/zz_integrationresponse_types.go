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

type IntegrationResponseInitParameters struct {

	// API identifier.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/apigatewayv2/v1beta1.API
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a API in apigatewayv2 to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a API in apigatewayv2 to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// How to handle response payload content type conversions. Valid values: CONVERT_TO_BINARY, CONVERT_TO_TEXT.
	ContentHandlingStrategy *string `json:"contentHandlingStrategy,omitempty" tf:"content_handling_strategy,omitempty"`

	// Identifier of the aws_apigatewayv2_integration.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/apigatewayv2/v1beta1.Integration
	IntegrationID *string `json:"integrationId,omitempty" tf:"integration_id,omitempty"`

	// Reference to a Integration in apigatewayv2 to populate integrationId.
	// +kubebuilder:validation:Optional
	IntegrationIDRef *v1.Reference `json:"integrationIdRef,omitempty" tf:"-"`

	// Selector for a Integration in apigatewayv2 to populate integrationId.
	// +kubebuilder:validation:Optional
	IntegrationIDSelector *v1.Selector `json:"integrationIdSelector,omitempty" tf:"-"`

	// Integration response key.
	IntegrationResponseKey *string `json:"integrationResponseKey,omitempty" tf:"integration_response_key,omitempty"`

	// Map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
	// +mapType=granular
	ResponseTemplates map[string]*string `json:"responseTemplates,omitempty" tf:"response_templates,omitempty"`

	// The template selection expression for the integration response.
	TemplateSelectionExpression *string `json:"templateSelectionExpression,omitempty" tf:"template_selection_expression,omitempty"`
}

type IntegrationResponseObservation struct {

	// API identifier.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// How to handle response payload content type conversions. Valid values: CONVERT_TO_BINARY, CONVERT_TO_TEXT.
	ContentHandlingStrategy *string `json:"contentHandlingStrategy,omitempty" tf:"content_handling_strategy,omitempty"`

	// Integration response identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Identifier of the aws_apigatewayv2_integration.
	IntegrationID *string `json:"integrationId,omitempty" tf:"integration_id,omitempty"`

	// Integration response key.
	IntegrationResponseKey *string `json:"integrationResponseKey,omitempty" tf:"integration_response_key,omitempty"`

	// Map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
	// +mapType=granular
	ResponseTemplates map[string]*string `json:"responseTemplates,omitempty" tf:"response_templates,omitempty"`

	// The template selection expression for the integration response.
	TemplateSelectionExpression *string `json:"templateSelectionExpression,omitempty" tf:"template_selection_expression,omitempty"`
}

type IntegrationResponseParameters struct {

	// API identifier.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/apigatewayv2/v1beta1.API
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a API in apigatewayv2 to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a API in apigatewayv2 to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// How to handle response payload content type conversions. Valid values: CONVERT_TO_BINARY, CONVERT_TO_TEXT.
	// +kubebuilder:validation:Optional
	ContentHandlingStrategy *string `json:"contentHandlingStrategy,omitempty" tf:"content_handling_strategy,omitempty"`

	// Identifier of the aws_apigatewayv2_integration.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/apigatewayv2/v1beta1.Integration
	// +kubebuilder:validation:Optional
	IntegrationID *string `json:"integrationId,omitempty" tf:"integration_id,omitempty"`

	// Reference to a Integration in apigatewayv2 to populate integrationId.
	// +kubebuilder:validation:Optional
	IntegrationIDRef *v1.Reference `json:"integrationIdRef,omitempty" tf:"-"`

	// Selector for a Integration in apigatewayv2 to populate integrationId.
	// +kubebuilder:validation:Optional
	IntegrationIDSelector *v1.Selector `json:"integrationIdSelector,omitempty" tf:"-"`

	// Integration response key.
	// +kubebuilder:validation:Optional
	IntegrationResponseKey *string `json:"integrationResponseKey,omitempty" tf:"integration_response_key,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ResponseTemplates map[string]*string `json:"responseTemplates,omitempty" tf:"response_templates,omitempty"`

	// The template selection expression for the integration response.
	// +kubebuilder:validation:Optional
	TemplateSelectionExpression *string `json:"templateSelectionExpression,omitempty" tf:"template_selection_expression,omitempty"`
}

// IntegrationResponseSpec defines the desired state of IntegrationResponse
type IntegrationResponseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntegrationResponseParameters `json:"forProvider"`
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
	InitProvider IntegrationResponseInitParameters `json:"initProvider,omitempty"`
}

// IntegrationResponseStatus defines the observed state of IntegrationResponse.
type IntegrationResponseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntegrationResponseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IntegrationResponse is the Schema for the IntegrationResponses API. Manages an Amazon API Gateway Version 2 integration response.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IntegrationResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.integrationResponseKey) || (has(self.initProvider) && has(self.initProvider.integrationResponseKey))",message="spec.forProvider.integrationResponseKey is a required parameter"
	Spec   IntegrationResponseSpec   `json:"spec"`
	Status IntegrationResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationResponseList contains a list of IntegrationResponses
type IntegrationResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IntegrationResponse `json:"items"`
}

// Repository type metadata.
var (
	IntegrationResponse_Kind             = "IntegrationResponse"
	IntegrationResponse_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IntegrationResponse_Kind}.String()
	IntegrationResponse_KindAPIVersion   = IntegrationResponse_Kind + "." + CRDGroupVersion.String()
	IntegrationResponse_GroupVersionKind = CRDGroupVersion.WithKind(IntegrationResponse_Kind)
)

func init() {
	SchemeBuilder.Register(&IntegrationResponse{}, &IntegrationResponseList{})
}
