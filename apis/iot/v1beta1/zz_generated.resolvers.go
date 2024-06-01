// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/crossplane-contrib/provider-upjet-aws/config/common"
	apisresolver "github.com/crossplane-contrib/provider-upjet-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *LoggingOptions) ResolveReferences( // ResolveReferences of this LoggingOptions.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.RoleArnRef,
			Selector:     mg.Spec.ForProvider.RoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.RoleArnRef,
			Selector:     mg.Spec.InitProvider.RoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleArn")
	}
	mg.Spec.InitProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PolicyAttachment.
func (mg *PolicyAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iot.aws.upbound.io", "v1beta1", "Policy", "PolicyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Policy),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.PolicyRef,
			Selector:     mg.Spec.ForProvider.PolicySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Policy")
	}
	mg.Spec.ForProvider.Policy = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iot.aws.upbound.io", "v1beta1", "Certificate", "CertificateList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Target),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.TargetRef,
			Selector:     mg.Spec.ForProvider.TargetSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Target")
	}
	mg.Spec.ForProvider.Target = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TargetRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iot.aws.upbound.io", "v1beta1", "Policy", "PolicyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Policy),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.PolicyRef,
			Selector:     mg.Spec.InitProvider.PolicySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Policy")
	}
	mg.Spec.InitProvider.Policy = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PolicyRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iot.aws.upbound.io", "v1beta1", "Certificate", "CertificateList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Target),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.TargetRef,
			Selector:     mg.Spec.InitProvider.TargetSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Target")
	}
	mg.Spec.InitProvider.Target = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TargetRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ProvisioningTemplate.
func (mg *ProvisioningTemplate) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProvisioningRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.ProvisioningRoleArnRef,
			Selector:     mg.Spec.ForProvider.ProvisioningRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProvisioningRoleArn")
	}
	mg.Spec.ForProvider.ProvisioningRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProvisioningRoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ProvisioningRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.ProvisioningRoleArnRef,
			Selector:     mg.Spec.InitProvider.ProvisioningRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ProvisioningRoleArn")
	}
	mg.Spec.InitProvider.ProvisioningRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProvisioningRoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RoleAlias.
func (mg *RoleAlias) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.RoleArnRef,
			Selector:     mg.Spec.ForProvider.RoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.RoleArnRef,
			Selector:     mg.Spec.InitProvider.RoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleArn")
	}
	mg.Spec.InitProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ThingGroup.
func (mg *ThingGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iot.aws.upbound.io", "v1beta1", "ThingGroup", "ThingGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ParentGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ParentGroupNameRef,
			Selector:     mg.Spec.ForProvider.ParentGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ParentGroupName")
	}
	mg.Spec.ForProvider.ParentGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParentGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iot.aws.upbound.io", "v1beta1", "ThingGroup", "ThingGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ParentGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ParentGroupNameRef,
			Selector:     mg.Spec.InitProvider.ParentGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ParentGroupName")
	}
	mg.Spec.InitProvider.ParentGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ParentGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ThingPrincipalAttachment.
func (mg *ThingPrincipalAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iot.aws.upbound.io", "v1beta1", "Certificate", "CertificateList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Principal),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.PrincipalRef,
			Selector:     mg.Spec.ForProvider.PrincipalSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Principal")
	}
	mg.Spec.ForProvider.Principal = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrincipalRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iot.aws.upbound.io", "v1beta1", "Thing", "ThingList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Thing),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ThingRef,
			Selector:     mg.Spec.ForProvider.ThingSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Thing")
	}
	mg.Spec.ForProvider.Thing = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ThingRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iot.aws.upbound.io", "v1beta1", "Certificate", "CertificateList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Principal),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.PrincipalRef,
			Selector:     mg.Spec.InitProvider.PrincipalSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Principal")
	}
	mg.Spec.InitProvider.Principal = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PrincipalRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iot.aws.upbound.io", "v1beta1", "Thing", "ThingList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Thing),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ThingRef,
			Selector:     mg.Spec.InitProvider.ThingSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Thing")
	}
	mg.Spec.InitProvider.Thing = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ThingRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TopicRule.
func (mg *TopicRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ErrorAction); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ErrorAction[i3].Sns); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].RoleArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].RoleArnRef,
					Selector:     mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].RoleArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].RoleArn")
			}
			mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].RoleArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ErrorAction); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ErrorAction[i3].Sns); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].TargetArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].TargetArnRef,
					Selector:     mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].TargetArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].TargetArn")
			}
			mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].TargetArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ErrorAction[i3].Sns[i4].TargetArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Sns); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Sns[i3].RoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.Sns[i3].RoleArnRef,
				Selector:     mg.Spec.ForProvider.Sns[i3].RoleArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Sns[i3].RoleArn")
		}
		mg.Spec.ForProvider.Sns[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Sns[i3].RoleArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Sns); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Sns[i3].TargetArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.Sns[i3].TargetArnRef,
				Selector:     mg.Spec.ForProvider.Sns[i3].TargetArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Sns[i3].TargetArn")
		}
		mg.Spec.ForProvider.Sns[i3].TargetArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Sns[i3].TargetArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ErrorAction); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.ErrorAction[i3].Sns); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ErrorAction[i3].Sns[i4].RoleArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.ErrorAction[i3].Sns[i4].RoleArnRef,
					Selector:     mg.Spec.InitProvider.ErrorAction[i3].Sns[i4].RoleArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.ErrorAction[i3].Sns[i4].RoleArn")
			}
			mg.Spec.InitProvider.ErrorAction[i3].Sns[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.ErrorAction[i3].Sns[i4].RoleArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ErrorAction); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.ErrorAction[i3].Sns); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ErrorAction[i3].Sns[i4].TargetArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.ErrorAction[i3].Sns[i4].TargetArnRef,
					Selector:     mg.Spec.InitProvider.ErrorAction[i3].Sns[i4].TargetArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.ErrorAction[i3].Sns[i4].TargetArn")
			}
			mg.Spec.InitProvider.ErrorAction[i3].Sns[i4].TargetArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.ErrorAction[i3].Sns[i4].TargetArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Sns); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Sns[i3].RoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.Sns[i3].RoleArnRef,
				Selector:     mg.Spec.InitProvider.Sns[i3].RoleArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Sns[i3].RoleArn")
		}
		mg.Spec.InitProvider.Sns[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Sns[i3].RoleArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Sns); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Sns[i3].TargetArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.Sns[i3].TargetArnRef,
				Selector:     mg.Spec.InitProvider.Sns[i3].TargetArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Sns[i3].TargetArn")
		}
		mg.Spec.InitProvider.Sns[i3].TargetArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Sns[i3].TargetArnRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this TopicRuleDestination.
func (mg *TopicRuleDestination) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPCConfiguration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCConfiguration[i3].RoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.VPCConfiguration[i3].RoleArnRef,
				Selector:     mg.Spec.ForProvider.VPCConfiguration[i3].RoleArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCConfiguration[i3].RoleArn")
		}
		mg.Spec.ForProvider.VPCConfiguration[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.VPCConfiguration[i3].RoleArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPCConfiguration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCConfiguration[i3].SecurityGroups),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.VPCConfiguration[i3].SecurityGroupRefs,
				Selector:      mg.Spec.ForProvider.VPCConfiguration[i3].SecurityGroupSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCConfiguration[i3].SecurityGroups")
		}
		mg.Spec.ForProvider.VPCConfiguration[i3].SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.VPCConfiguration[i3].SecurityGroupRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPCConfiguration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCConfiguration[i3].SubnetIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.VPCConfiguration[i3].SubnetIDRefs,
				Selector:      mg.Spec.ForProvider.VPCConfiguration[i3].SubnetIDSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCConfiguration[i3].SubnetIds")
		}
		mg.Spec.ForProvider.VPCConfiguration[i3].SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.VPCConfiguration[i3].SubnetIDRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPCConfiguration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCConfiguration[i3].VPCID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.VPCConfiguration[i3].VPCIDRef,
				Selector:     mg.Spec.ForProvider.VPCConfiguration[i3].VPCIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCConfiguration[i3].VPCID")
		}
		mg.Spec.ForProvider.VPCConfiguration[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.VPCConfiguration[i3].VPCIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.VPCConfiguration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCConfiguration[i3].RoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.VPCConfiguration[i3].RoleArnRef,
				Selector:     mg.Spec.InitProvider.VPCConfiguration[i3].RoleArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.VPCConfiguration[i3].RoleArn")
		}
		mg.Spec.InitProvider.VPCConfiguration[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.VPCConfiguration[i3].RoleArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.VPCConfiguration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.VPCConfiguration[i3].SecurityGroups),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.InitProvider.VPCConfiguration[i3].SecurityGroupRefs,
				Selector:      mg.Spec.InitProvider.VPCConfiguration[i3].SecurityGroupSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.VPCConfiguration[i3].SecurityGroups")
		}
		mg.Spec.InitProvider.VPCConfiguration[i3].SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.VPCConfiguration[i3].SecurityGroupRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.VPCConfiguration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.VPCConfiguration[i3].SubnetIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.InitProvider.VPCConfiguration[i3].SubnetIDRefs,
				Selector:      mg.Spec.InitProvider.VPCConfiguration[i3].SubnetIDSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.VPCConfiguration[i3].SubnetIds")
		}
		mg.Spec.InitProvider.VPCConfiguration[i3].SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.VPCConfiguration[i3].SubnetIDRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.VPCConfiguration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCConfiguration[i3].VPCID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.VPCConfiguration[i3].VPCIDRef,
				Selector:     mg.Spec.InitProvider.VPCConfiguration[i3].VPCIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.VPCConfiguration[i3].VPCID")
		}
		mg.Spec.InitProvider.VPCConfiguration[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.VPCConfiguration[i3].VPCIDRef = rsp.ResolvedReference

	}

	return nil
}
