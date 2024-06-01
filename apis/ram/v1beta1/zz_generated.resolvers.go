// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"

	apisresolver "github.com/crossplane-contrib/provider-upjet-aws/internal/apis"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *PrincipalAssociation) ResolveReferences( // ResolveReferences of this PrincipalAssociation.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("organizations.aws.upbound.io", "v1beta1", "Organization", "OrganizationList")
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
		m, l, err = apisresolver.GetManagedResource("ram.aws.upbound.io", "v1beta1", "ResourceShare", "ResourceShareList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceShareArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.ResourceShareArnRef,
			Selector:     mg.Spec.ForProvider.ResourceShareArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceShareArn")
	}
	mg.Spec.ForProvider.ResourceShareArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceShareArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("organizations.aws.upbound.io", "v1beta1", "Organization", "OrganizationList")
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
		m, l, err = apisresolver.GetManagedResource("ram.aws.upbound.io", "v1beta1", "ResourceShare", "ResourceShareList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceShareArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.ResourceShareArnRef,
			Selector:     mg.Spec.InitProvider.ResourceShareArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceShareArn")
	}
	mg.Spec.InitProvider.ResourceShareArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceShareArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ResourceAssociation.
func (mg *ResourceAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ram.aws.upbound.io", "v1beta1", "ResourceShare", "ResourceShareList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceShareArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.ResourceShareArnRef,
			Selector:     mg.Spec.ForProvider.ResourceShareArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceShareArn")
	}
	mg.Spec.ForProvider.ResourceShareArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceShareArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ram.aws.upbound.io", "v1beta1", "ResourceShare", "ResourceShareList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceShareArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.ResourceShareArnRef,
			Selector:     mg.Spec.InitProvider.ResourceShareArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceShareArn")
	}
	mg.Spec.InitProvider.ResourceShareArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceShareArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ResourceShareAccepter.
func (mg *ResourceShareAccepter) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ram.aws.upbound.io", "v1beta1", "PrincipalAssociation", "PrincipalAssociationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ShareArn),
			Extract:      resource.ExtractParamPath("resource_share_arn", false),
			Reference:    mg.Spec.ForProvider.ShareArnRef,
			Selector:     mg.Spec.ForProvider.ShareArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ShareArn")
	}
	mg.Spec.ForProvider.ShareArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ShareArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ram.aws.upbound.io", "v1beta1", "PrincipalAssociation", "PrincipalAssociationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ShareArn),
			Extract:      resource.ExtractParamPath("resource_share_arn", false),
			Reference:    mg.Spec.InitProvider.ShareArnRef,
			Selector:     mg.Spec.InitProvider.ShareArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ShareArn")
	}
	mg.Spec.InitProvider.ShareArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ShareArnRef = rsp.ResolvedReference

	return nil
}
