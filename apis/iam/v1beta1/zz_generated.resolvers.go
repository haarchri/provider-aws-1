// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	common "github.com/crossplane-contrib/provider-upjet-aws/config/common"
	apisresolver "github.com/crossplane-contrib/provider-upjet-aws/internal/apis"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this AccessKey.
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

func (mg *AccessKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.User),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.UserRef,
			Selector:     mg.Spec.ForProvider.UserSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.User")
	}
	mg.Spec.ForProvider.User = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.User),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.UserRef,
			Selector:     mg.Spec.InitProvider.UserSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.User")
	}
	mg.Spec.InitProvider.User = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.UserRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GroupMembership.
func (mg *GroupMembership) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Group", "GroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Group),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.GroupRef,
			Selector:     mg.Spec.ForProvider.GroupSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Group")
	}
	mg.Spec.ForProvider.Group = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Users),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.UserRefs,
			Selector:      mg.Spec.ForProvider.UserSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Users")
	}
	mg.Spec.ForProvider.Users = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.UserRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Group", "GroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Group),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.GroupRef,
			Selector:     mg.Spec.InitProvider.GroupSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Group")
	}
	mg.Spec.InitProvider.Group = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.GroupRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Users),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.UserRefs,
			Selector:      mg.Spec.InitProvider.UserSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Users")
	}
	mg.Spec.InitProvider.Users = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.UserRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this GroupPolicyAttachment.
func (mg *GroupPolicyAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Group", "GroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Group),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.GroupRef,
			Selector:     mg.Spec.ForProvider.GroupSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Group")
	}
	mg.Spec.ForProvider.Group = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Policy", "PolicyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.PolicyArnRef,
			Selector:     mg.Spec.ForProvider.PolicyArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyArn")
	}
	mg.Spec.ForProvider.PolicyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Group", "GroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Group),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.GroupRef,
			Selector:     mg.Spec.InitProvider.GroupSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Group")
	}
	mg.Spec.InitProvider.Group = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.GroupRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Policy", "PolicyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PolicyArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.PolicyArnRef,
			Selector:     mg.Spec.InitProvider.PolicyArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PolicyArn")
	}
	mg.Spec.InitProvider.PolicyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PolicyArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this InstanceProfile.
func (mg *InstanceProfile) ResolveReferences(ctx context.Context, c client.Reader) error {
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
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Role),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RoleRef,
			Selector:     mg.Spec.ForProvider.RoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Role")
	}
	mg.Spec.ForProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Role),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RoleRef,
			Selector:     mg.Spec.InitProvider.RoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Role")
	}
	mg.Spec.InitProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RolePolicy.
func (mg *RolePolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
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
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Role),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.RoleRef,
			Selector:     mg.Spec.ForProvider.RoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Role")
	}
	mg.Spec.ForProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RolePolicyAttachment.
func (mg *RolePolicyAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Policy", "PolicyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.PolicyArnRef,
			Selector:     mg.Spec.ForProvider.PolicyArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyArn")
	}
	mg.Spec.ForProvider.PolicyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Role),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RoleRef,
			Selector:     mg.Spec.ForProvider.RoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Role")
	}
	mg.Spec.ForProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Policy", "PolicyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PolicyArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.PolicyArnRef,
			Selector:     mg.Spec.InitProvider.PolicyArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PolicyArn")
	}
	mg.Spec.InitProvider.PolicyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PolicyArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Role),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RoleRef,
			Selector:     mg.Spec.InitProvider.RoleSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Role")
	}
	mg.Spec.InitProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServiceSpecificCredential.
func (mg *ServiceSpecificCredential) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.UserNameRef,
			Selector:     mg.Spec.ForProvider.UserNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserName")
	}
	mg.Spec.ForProvider.UserName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.UserName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.UserNameRef,
			Selector:     mg.Spec.InitProvider.UserNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.UserName")
	}
	mg.Spec.InitProvider.UserName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.UserNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserGroupMembership.
func (mg *UserGroupMembership) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Group", "GroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Groups),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.GroupRefs,
			Selector:      mg.Spec.ForProvider.GroupSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Groups")
	}
	mg.Spec.ForProvider.Groups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.GroupRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.User),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.UserRef,
			Selector:     mg.Spec.ForProvider.UserSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.User")
	}
	mg.Spec.ForProvider.User = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Group", "GroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Groups),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.GroupRefs,
			Selector:      mg.Spec.InitProvider.GroupSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Groups")
	}
	mg.Spec.InitProvider.Groups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.GroupRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.User),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.UserRef,
			Selector:     mg.Spec.InitProvider.UserSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.User")
	}
	mg.Spec.InitProvider.User = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.UserRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserLoginProfile.
func (mg *UserLoginProfile) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.User),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.UserRef,
			Selector:     mg.Spec.ForProvider.UserSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.User")
	}
	mg.Spec.ForProvider.User = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.User),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.UserRef,
			Selector:     mg.Spec.InitProvider.UserSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.User")
	}
	mg.Spec.InitProvider.User = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.UserRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserPolicyAttachment.
func (mg *UserPolicyAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Policy", "PolicyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.PolicyArnRef,
			Selector:     mg.Spec.ForProvider.PolicyArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyArn")
	}
	mg.Spec.ForProvider.PolicyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.User),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.UserRef,
			Selector:     mg.Spec.ForProvider.UserSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.User")
	}
	mg.Spec.ForProvider.User = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Policy", "PolicyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PolicyArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.PolicyArnRef,
			Selector:     mg.Spec.InitProvider.PolicyArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PolicyArn")
	}
	mg.Spec.InitProvider.PolicyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PolicyArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.User),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.UserRef,
			Selector:     mg.Spec.InitProvider.UserSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.User")
	}
	mg.Spec.InitProvider.User = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.UserRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserSSHKey.
func (mg *UserSSHKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Username),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.UsernameRef,
			Selector:     mg.Spec.ForProvider.UsernameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Username")
	}
	mg.Spec.ForProvider.Username = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UsernameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "User", "UserList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Username),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.UsernameRef,
			Selector:     mg.Spec.InitProvider.UsernameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Username")
	}
	mg.Spec.InitProvider.Username = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.UsernameRef = rsp.ResolvedReference

	return nil
}
