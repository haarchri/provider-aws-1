// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	apisresolver "github.com/crossplane-contrib/provider-upjet-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *DelegatedAdministrator) ResolveReferences( // ResolveReferences of this DelegatedAdministrator.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("organizations.aws.upbound.io", "v1beta1", "Account", "AccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccountID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.AccountIDRef,
			Selector:     mg.Spec.ForProvider.AccountIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccountID")
	}
	mg.Spec.ForProvider.AccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccountIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("organizations.aws.upbound.io", "v1beta1", "Account", "AccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AccountID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.AccountIDRef,
			Selector:     mg.Spec.InitProvider.AccountIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AccountID")
	}
	mg.Spec.InitProvider.AccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AccountIDRef = rsp.ResolvedReference

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
		m, l, err = apisresolver.GetManagedResource("organizations.aws.upbound.io", "v1beta1", "Policy", "PolicyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.PolicyIDRef,
			Selector:     mg.Spec.ForProvider.PolicyIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyID")
	}
	mg.Spec.ForProvider.PolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("organizations.aws.upbound.io", "v1beta1", "Policy", "PolicyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PolicyID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.PolicyIDRef,
			Selector:     mg.Spec.InitProvider.PolicyIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PolicyID")
	}
	mg.Spec.InitProvider.PolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PolicyIDRef = rsp.ResolvedReference

	return nil
}
