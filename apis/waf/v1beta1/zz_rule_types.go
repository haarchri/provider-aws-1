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

type RuleInitParameters struct {

	// The name or description for the Amazon CloudWatch metric of this rule. The name can contain only alphanumeric characters (A-Z, a-z, 0-9); the name can't contain whitespace.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// The name or description of the rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The objects to include in a rule (documented below).
	Predicates []RulePredicatesInitParameters `json:"predicates,omitempty" tf:"predicates,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RuleObservation struct {

	// The ARN of the WAF rule.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the WAF rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name or description for the Amazon CloudWatch metric of this rule. The name can contain only alphanumeric characters (A-Z, a-z, 0-9); the name can't contain whitespace.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// The name or description of the rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The objects to include in a rule (documented below).
	Predicates []RulePredicatesObservation `json:"predicates,omitempty" tf:"predicates,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RuleParameters struct {

	// The name or description for the Amazon CloudWatch metric of this rule. The name can contain only alphanumeric characters (A-Z, a-z, 0-9); the name can't contain whitespace.
	// +kubebuilder:validation:Optional
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// The name or description of the rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The objects to include in a rule (documented below).
	// +kubebuilder:validation:Optional
	Predicates []RulePredicatesParameters `json:"predicates,omitempty" tf:"predicates,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RulePredicatesInitParameters struct {

	// A unique identifier for a predicate in the rule, such as Byte Match Set ID or IPSet ID.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/waf/v1beta1.IPSet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	DataID *string `json:"dataId,omitempty" tf:"data_id,omitempty"`

	// Reference to a IPSet in waf to populate dataId.
	// +kubebuilder:validation:Optional
	DataIDRef *v1.Reference `json:"dataIdRef,omitempty" tf:"-"`

	// Selector for a IPSet in waf to populate dataId.
	// +kubebuilder:validation:Optional
	DataIDSelector *v1.Selector `json:"dataIdSelector,omitempty" tf:"-"`

	// Set this to false if you want to allow, block, or count requests
	// based on the settings in the specified waf_byte_match_set, waf_ipset, aws_waf_size_constraint_set, aws_waf_sql_injection_match_set or aws_waf_xss_match_set.
	// For example, if an IPSet includes the IP address 192.0.2.44, AWS WAF will allow or block requests based on that IP address.
	// If set to true, AWS WAF will allow, block, or count requests based on all IP addresses except 192.0.2.44.
	Negated *bool `json:"negated,omitempty" tf:"negated,omitempty"`

	// The type of predicate in a rule. Valid values: ByteMatch, GeoMatch, IPMatch, RegexMatch, SizeConstraint, SqlInjectionMatch, or XssMatch.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RulePredicatesObservation struct {

	// A unique identifier for a predicate in the rule, such as Byte Match Set ID or IPSet ID.
	DataID *string `json:"dataId,omitempty" tf:"data_id,omitempty"`

	// Set this to false if you want to allow, block, or count requests
	// based on the settings in the specified waf_byte_match_set, waf_ipset, aws_waf_size_constraint_set, aws_waf_sql_injection_match_set or aws_waf_xss_match_set.
	// For example, if an IPSet includes the IP address 192.0.2.44, AWS WAF will allow or block requests based on that IP address.
	// If set to true, AWS WAF will allow, block, or count requests based on all IP addresses except 192.0.2.44.
	Negated *bool `json:"negated,omitempty" tf:"negated,omitempty"`

	// The type of predicate in a rule. Valid values: ByteMatch, GeoMatch, IPMatch, RegexMatch, SizeConstraint, SqlInjectionMatch, or XssMatch.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RulePredicatesParameters struct {

	// A unique identifier for a predicate in the rule, such as Byte Match Set ID or IPSet ID.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-aws/apis/waf/v1beta1.IPSet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataID *string `json:"dataId,omitempty" tf:"data_id,omitempty"`

	// Reference to a IPSet in waf to populate dataId.
	// +kubebuilder:validation:Optional
	DataIDRef *v1.Reference `json:"dataIdRef,omitempty" tf:"-"`

	// Selector for a IPSet in waf to populate dataId.
	// +kubebuilder:validation:Optional
	DataIDSelector *v1.Selector `json:"dataIdSelector,omitempty" tf:"-"`

	// Set this to false if you want to allow, block, or count requests
	// based on the settings in the specified waf_byte_match_set, waf_ipset, aws_waf_size_constraint_set, aws_waf_sql_injection_match_set or aws_waf_xss_match_set.
	// For example, if an IPSet includes the IP address 192.0.2.44, AWS WAF will allow or block requests based on that IP address.
	// If set to true, AWS WAF will allow, block, or count requests based on all IP addresses except 192.0.2.44.
	// +kubebuilder:validation:Optional
	Negated *bool `json:"negated" tf:"negated,omitempty"`

	// The type of predicate in a rule. Valid values: ByteMatch, GeoMatch, IPMatch, RegexMatch, SizeConstraint, SqlInjectionMatch, or XssMatch.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// RuleSpec defines the desired state of Rule
type RuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleParameters `json:"forProvider"`
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
	InitProvider RuleInitParameters `json:"initProvider,omitempty"`
}

// RuleStatus defines the observed state of Rule.
type RuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Rule is the Schema for the Rules API. Provides a AWS WAF rule resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.metricName) || (has(self.initProvider) && has(self.initProvider.metricName))",message="spec.forProvider.metricName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   RuleSpec   `json:"spec"`
	Status RuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleList contains a list of Rules
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rule `json:"items"`
}

// Repository type metadata.
var (
	Rule_Kind             = "Rule"
	Rule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Rule_Kind}.String()
	Rule_KindAPIVersion   = Rule_Kind + "." + CRDGroupVersion.String()
	Rule_GroupVersionKind = CRDGroupVersion.WithKind(Rule_Kind)
)

func init() {
	SchemeBuilder.Register(&Rule{}, &RuleList{})
}
