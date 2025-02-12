// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetOrganizationmanagerGroup
    {
        /// <summary>
        /// Get information about a Yandex Cloud Organization Manager Group. For more information, see [the official documentation](https://yandex.cloud/docs/organization/manage-groups).
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/organizationmanager_group/d_organizationmanager_group_1.tf" }}
        /// </summary>
        public static Task<GetOrganizationmanagerGroupResult> InvokeAsync(GetOrganizationmanagerGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationmanagerGroupResult>("yandex:index/getOrganizationmanagerGroup:getOrganizationmanagerGroup", args ?? new GetOrganizationmanagerGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Yandex Cloud Organization Manager Group. For more information, see [the official documentation](https://yandex.cloud/docs/organization/manage-groups).
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/organizationmanager_group/d_organizationmanager_group_1.tf" }}
        /// </summary>
        public static Output<GetOrganizationmanagerGroupResult> Invoke(GetOrganizationmanagerGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationmanagerGroupResult>("yandex:index/getOrganizationmanagerGroup:getOrganizationmanagerGroup", args ?? new GetOrganizationmanagerGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Yandex Cloud Organization Manager Group. For more information, see [the official documentation](https://yandex.cloud/docs/organization/manage-groups).
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/organizationmanager_group/d_organizationmanager_group_1.tf" }}
        /// </summary>
        public static Output<GetOrganizationmanagerGroupResult> Invoke(GetOrganizationmanagerGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationmanagerGroupResult>("yandex:index/getOrganizationmanagerGroup:getOrganizationmanagerGroup", args ?? new GetOrganizationmanagerGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrganizationmanagerGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of a Group.
        /// </summary>
        [Input("groupId")]
        public string? GroupId { get; set; }

        /// <summary>
        /// Name of a Group.
        /// 
        /// &gt; One of `group_id` or `name` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Organization that the Group belongs to. If value is omitted, the default provider organization is used.
        /// </summary>
        [Input("organizationId")]
        public string? OrganizationId { get; set; }

        public GetOrganizationmanagerGroupArgs()
        {
        }
        public static new GetOrganizationmanagerGroupArgs Empty => new GetOrganizationmanagerGroupArgs();
    }

    public sealed class GetOrganizationmanagerGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of a Group.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// Name of a Group.
        /// 
        /// &gt; One of `group_id` or `name` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Organization that the Group belongs to. If value is omitted, the default provider organization is used.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        public GetOrganizationmanagerGroupInvokeArgs()
        {
        }
        public static new GetOrganizationmanagerGroupInvokeArgs Empty => new GetOrganizationmanagerGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrganizationmanagerGroupResult
    {
        /// <summary>
        /// The Group creation timestamp.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The description of the Group.
        /// </summary>
        public readonly string Description;
        public readonly string GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of members of the Group. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetOrganizationmanagerGroupMemberResult> Members;
        public readonly string Name;
        public readonly string? OrganizationId;

        [OutputConstructor]
        private GetOrganizationmanagerGroupResult(
            string createdAt,

            string description,

            string groupId,

            string id,

            ImmutableArray<Outputs.GetOrganizationmanagerGroupMemberResult> members,

            string name,

            string? organizationId)
        {
            CreatedAt = createdAt;
            Description = description;
            GroupId = groupId;
            Id = id;
            Members = members;
            Name = name;
            OrganizationId = organizationId;
        }
    }
}
