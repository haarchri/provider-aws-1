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

	// ResolveReferences of this Cluster.
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.KMSKeyArnRef,
			Selector:     mg.Spec.ForProvider.KMSKeyArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyArn")
	}
	mg.Spec.ForProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.SecurityGroupIDRefs,
			Selector:      mg.Spec.ForProvider.SecurityGroupIDSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIDRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("memorydb.aws.upbound.io", "v1beta1", "SubnetGroup", "SubnetGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetGroupName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.SubnetGroupNameRef,
			Selector:     mg.Spec.ForProvider.SubnetGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetGroupName")
	}
	mg.Spec.ForProvider.SubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KMSKeyArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.KMSKeyArnRef,
			Selector:     mg.Spec.InitProvider.KMSKeyArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KMSKeyArn")
	}
	mg.Spec.InitProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KMSKeyArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SecurityGroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.SecurityGroupIDRefs,
			Selector:      mg.Spec.InitProvider.SecurityGroupIDSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecurityGroupIds")
	}
	mg.Spec.InitProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SecurityGroupIDRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("memorydb.aws.upbound.io", "v1beta1", "SubnetGroup", "SubnetGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubnetGroupName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.SubnetGroupNameRef,
			Selector:     mg.Spec.InitProvider.SubnetGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetGroupName")
	}
	mg.Spec.InitProvider.SubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubnetGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Snapshot.
func (mg *Snapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("memorydb.aws.upbound.io", "v1beta1", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ClusterNameRef,
			Selector:     mg.Spec.ForProvider.ClusterNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.KMSKeyArnRef,
			Selector:     mg.Spec.ForProvider.KMSKeyArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyArn")
	}
	mg.Spec.ForProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("memorydb.aws.upbound.io", "v1beta1", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ClusterNameRef,
			Selector:     mg.Spec.InitProvider.ClusterNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterName")
	}
	mg.Spec.InitProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KMSKeyArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.KMSKeyArnRef,
			Selector:     mg.Spec.InitProvider.KMSKeyArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KMSKeyArn")
	}
	mg.Spec.InitProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KMSKeyArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubnetGroup.
func (mg *SubnetGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.SubnetIDRefs,
			Selector:      mg.Spec.ForProvider.SubnetIDSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetIds")
	}
	mg.Spec.ForProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetIDRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SubnetIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.SubnetIDRefs,
			Selector:      mg.Spec.InitProvider.SubnetIDSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetIds")
	}
	mg.Spec.InitProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SubnetIDRefs = mrsp.ResolvedReferences

	return nil
}
