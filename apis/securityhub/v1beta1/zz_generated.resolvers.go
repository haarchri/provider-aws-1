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

func (mg *InviteAccepter) ResolveReferences( // ResolveReferences of this InviteAccepter.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("securityhub.aws.upbound.io", "v1beta1", "Member", "MemberList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MasterID),
			Extract:      resource.ExtractParamPath("master_id", true),
			Reference:    mg.Spec.ForProvider.MasterIDRef,
			Selector:     mg.Spec.ForProvider.MasterIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MasterID")
	}
	mg.Spec.ForProvider.MasterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MasterIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("securityhub.aws.upbound.io", "v1beta1", "Member", "MemberList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MasterID),
			Extract:      resource.ExtractParamPath("master_id", true),
			Reference:    mg.Spec.InitProvider.MasterIDRef,
			Selector:     mg.Spec.InitProvider.MasterIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.MasterID")
	}
	mg.Spec.InitProvider.MasterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.MasterIDRef = rsp.ResolvedReference

	return nil
}
