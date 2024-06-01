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

func (mg *APIDestination) ResolveReferences( // ResolveReferences of this APIDestination.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cloudwatchevents.aws.upbound.io", "v1beta1", "Connection", "ConnectionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConnectionArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.ConnectionArnRef,
			Selector:     mg.Spec.ForProvider.ConnectionArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConnectionArn")
	}
	mg.Spec.ForProvider.ConnectionArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConnectionArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("cloudwatchevents.aws.upbound.io", "v1beta1", "Connection", "ConnectionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ConnectionArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.ConnectionArnRef,
			Selector:     mg.Spec.InitProvider.ConnectionArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ConnectionArn")
	}
	mg.Spec.InitProvider.ConnectionArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ConnectionArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Archive.
func (mg *Archive) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cloudwatchevents.aws.upbound.io", "v1beta1", "Bus", "BusList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventSourceArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.EventSourceArnRef,
			Selector:     mg.Spec.ForProvider.EventSourceArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventSourceArn")
	}
	mg.Spec.ForProvider.EventSourceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventSourceArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("cloudwatchevents.aws.upbound.io", "v1beta1", "Bus", "BusList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventSourceArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.EventSourceArnRef,
			Selector:     mg.Spec.InitProvider.EventSourceArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EventSourceArn")
	}
	mg.Spec.InitProvider.EventSourceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EventSourceArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BusPolicy.
func (mg *BusPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cloudwatchevents.aws.upbound.io", "v1beta1", "Bus", "BusList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventBusName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.EventBusNameRef,
			Selector:     mg.Spec.ForProvider.EventBusNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventBusName")
	}
	mg.Spec.ForProvider.EventBusName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventBusNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("cloudwatchevents.aws.upbound.io", "v1beta1", "Bus", "BusList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventBusName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.EventBusNameRef,
			Selector:     mg.Spec.InitProvider.EventBusNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EventBusName")
	}
	mg.Spec.InitProvider.EventBusName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EventBusNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Permission.
func (mg *Permission) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Condition); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("organizations.aws.upbound.io", "v1beta1", "Organization", "OrganizationList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Condition[i3].Value),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Condition[i3].ValueRef,
				Selector:     mg.Spec.ForProvider.Condition[i3].ValueSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Condition[i3].Value")
		}
		mg.Spec.ForProvider.Condition[i3].Value = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Condition[i3].ValueRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("cloudwatchevents.aws.upbound.io", "v1beta1", "Bus", "BusList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventBusName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.EventBusNameRef,
			Selector:     mg.Spec.ForProvider.EventBusNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventBusName")
	}
	mg.Spec.ForProvider.EventBusName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventBusNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Condition); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("organizations.aws.upbound.io", "v1beta1", "Organization", "OrganizationList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Condition[i3].Value),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Condition[i3].ValueRef,
				Selector:     mg.Spec.InitProvider.Condition[i3].ValueSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Condition[i3].Value")
		}
		mg.Spec.InitProvider.Condition[i3].Value = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Condition[i3].ValueRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("cloudwatchevents.aws.upbound.io", "v1beta1", "Bus", "BusList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventBusName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.EventBusNameRef,
			Selector:     mg.Spec.InitProvider.EventBusNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EventBusName")
	}
	mg.Spec.InitProvider.EventBusName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EventBusNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Rule.
func (mg *Rule) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cloudwatchevents.aws.upbound.io", "v1beta1", "Bus", "BusList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventBusName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.EventBusNameRef,
			Selector:     mg.Spec.ForProvider.EventBusNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventBusName")
	}
	mg.Spec.ForProvider.EventBusName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventBusNameRef = rsp.ResolvedReference
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
		m, l, err = apisresolver.GetManagedResource("cloudwatchevents.aws.upbound.io", "v1beta1", "Bus", "BusList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventBusName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.EventBusNameRef,
			Selector:     mg.Spec.InitProvider.EventBusNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EventBusName")
	}
	mg.Spec.InitProvider.EventBusName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EventBusNameRef = rsp.ResolvedReference
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

// ResolveReferences of this Target.
func (mg *Target) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.EcsTarget); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ecs.aws.upbound.io", "v1beta1", "TaskDefinition", "TaskDefinitionList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EcsTarget[i3].TaskDefinitionArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.EcsTarget[i3].TaskDefinitionArnRef,
				Selector:     mg.Spec.ForProvider.EcsTarget[i3].TaskDefinitionArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EcsTarget[i3].TaskDefinitionArn")
		}
		mg.Spec.ForProvider.EcsTarget[i3].TaskDefinitionArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EcsTarget[i3].TaskDefinitionArnRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("cloudwatchevents.aws.upbound.io", "v1beta1", "Bus", "BusList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventBusName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.EventBusNameRef,
			Selector:     mg.Spec.ForProvider.EventBusNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventBusName")
	}
	mg.Spec.ForProvider.EventBusName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventBusNameRef = rsp.ResolvedReference
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
		m, l, err = apisresolver.GetManagedResource("cloudwatchevents.aws.upbound.io", "v1beta1", "Rule", "RuleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Rule),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RuleRef,
			Selector:     mg.Spec.ForProvider.RuleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Rule")
	}
	mg.Spec.ForProvider.Rule = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RuleRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.EcsTarget); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ecs.aws.upbound.io", "v1beta1", "TaskDefinition", "TaskDefinitionList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EcsTarget[i3].TaskDefinitionArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.EcsTarget[i3].TaskDefinitionArnRef,
				Selector:     mg.Spec.InitProvider.EcsTarget[i3].TaskDefinitionArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.EcsTarget[i3].TaskDefinitionArn")
		}
		mg.Spec.InitProvider.EcsTarget[i3].TaskDefinitionArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.EcsTarget[i3].TaskDefinitionArnRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("cloudwatchevents.aws.upbound.io", "v1beta1", "Bus", "BusList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventBusName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.EventBusNameRef,
			Selector:     mg.Spec.InitProvider.EventBusNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EventBusName")
	}
	mg.Spec.InitProvider.EventBusName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EventBusNameRef = rsp.ResolvedReference
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
	{
		m, l, err = apisresolver.GetManagedResource("cloudwatchevents.aws.upbound.io", "v1beta1", "Rule", "RuleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Rule),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RuleRef,
			Selector:     mg.Spec.InitProvider.RuleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Rule")
	}
	mg.Spec.InitProvider.Rule = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RuleRef = rsp.ResolvedReference

	return nil
}
