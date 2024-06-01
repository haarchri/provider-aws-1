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
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this DiskAttachment.
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

func (mg *DiskAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("lightsail.aws.upbound.io", "v1beta1", "Disk", "DiskList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DiskName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DiskNameRef,
			Selector:     mg.Spec.ForProvider.DiskNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DiskName")
	}
	mg.Spec.ForProvider.DiskName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DiskNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lightsail.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.InstanceNameRef,
			Selector:     mg.Spec.ForProvider.InstanceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceName")
	}
	mg.Spec.ForProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lightsail.aws.upbound.io", "v1beta1", "Disk", "DiskList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DiskName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.DiskNameRef,
			Selector:     mg.Spec.InitProvider.DiskNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DiskName")
	}
	mg.Spec.InitProvider.DiskName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DiskNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lightsail.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.InstanceNameRef,
			Selector:     mg.Spec.InitProvider.InstanceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceName")
	}
	mg.Spec.InitProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DomainEntry.
func (mg *DomainEntry) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("lightsail.aws.upbound.io", "v1beta1", "Domain", "DomainList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DomainName),
			Extract:      resource.ExtractParamPath("domain_name", false),
			Reference:    mg.Spec.ForProvider.DomainNameRef,
			Selector:     mg.Spec.ForProvider.DomainNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DomainName")
	}
	mg.Spec.ForProvider.DomainName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DomainNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this InstancePublicPorts.
func (mg *InstancePublicPorts) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("lightsail.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.InstanceNameRef,
			Selector:     mg.Spec.ForProvider.InstanceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceName")
	}
	mg.Spec.ForProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lightsail.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.InstanceNameRef,
			Selector:     mg.Spec.InitProvider.InstanceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceName")
	}
	mg.Spec.InitProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LBAttachment.
func (mg *LBAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("lightsail.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.InstanceNameRef,
			Selector:     mg.Spec.ForProvider.InstanceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceName")
	}
	mg.Spec.ForProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lightsail.aws.upbound.io", "v1beta1", "LB", "LBList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LBName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LBNameRef,
			Selector:     mg.Spec.ForProvider.LBNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LBName")
	}
	mg.Spec.ForProvider.LBName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LBNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lightsail.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.InstanceNameRef,
			Selector:     mg.Spec.InitProvider.InstanceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceName")
	}
	mg.Spec.InitProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lightsail.aws.upbound.io", "v1beta1", "LB", "LBList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LBName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.LBNameRef,
			Selector:     mg.Spec.InitProvider.LBNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LBName")
	}
	mg.Spec.InitProvider.LBName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LBNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LBCertificate.
func (mg *LBCertificate) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("lightsail.aws.upbound.io", "v1beta1", "LB", "LBList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LBName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.LBNameRef,
			Selector:     mg.Spec.ForProvider.LBNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LBName")
	}
	mg.Spec.ForProvider.LBName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LBNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StaticIPAttachment.
func (mg *StaticIPAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("lightsail.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.InstanceNameRef,
			Selector:     mg.Spec.ForProvider.InstanceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceName")
	}
	mg.Spec.ForProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lightsail.aws.upbound.io", "v1beta1", "StaticIP", "StaticIPList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StaticIPName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.StaticIPNameRef,
			Selector:     mg.Spec.ForProvider.StaticIPNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StaticIPName")
	}
	mg.Spec.ForProvider.StaticIPName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StaticIPNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lightsail.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.InstanceNameRef,
			Selector:     mg.Spec.InitProvider.InstanceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceName")
	}
	mg.Spec.InitProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lightsail.aws.upbound.io", "v1beta1", "StaticIP", "StaticIPList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StaticIPName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.StaticIPNameRef,
			Selector:     mg.Spec.InitProvider.StaticIPNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.StaticIPName")
	}
	mg.Spec.InitProvider.StaticIPName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.StaticIPNameRef = rsp.ResolvedReference

	return nil
}
