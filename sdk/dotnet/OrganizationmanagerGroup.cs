// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    /// <summary>
    /// Allows management of a single Group within an existing Yandex Cloud Organization. For more information, see [the official documentation](https://yandex.cloud/docs/organization/manage-groups).
    /// 
    /// ## Example Usage
    /// 
    /// {{ tffile "examples/organizationmanager_group/r_organizationmanager_group_1.tf" }}
    /// 
    /// ## Import
    /// 
    /// A Yandex Cloud Organization Manager Group can be imported using the `id` of the resource, e.g.:
    /// 
    /// ```sh
    /// $ pulumi import yandex:index/organizationmanagerGroup:OrganizationmanagerGroup group "group_id"
    /// ```
    /// </summary>
    [YandexResourceType("yandex:index/organizationmanagerGroup:OrganizationmanagerGroup")]
    public partial class OrganizationmanagerGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// (Computed) The SAML Federation creation timestamp.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The description of the Group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the Group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The organization to attach this Group to.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationmanagerGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationmanagerGroup(string name, OrganizationmanagerGroupArgs args, CustomResourceOptions? options = null)
            : base("yandex:index/organizationmanagerGroup:OrganizationmanagerGroup", name, args ?? new OrganizationmanagerGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationmanagerGroup(string name, Input<string> id, OrganizationmanagerGroupState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/organizationmanagerGroup:OrganizationmanagerGroup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/agorobets/pulumi-yandex/releases",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OrganizationmanagerGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationmanagerGroup Get(string name, Input<string> id, OrganizationmanagerGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new OrganizationmanagerGroup(name, id, state, options);
        }
    }

    public sealed class OrganizationmanagerGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization to attach this Group to.
        /// </summary>
        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        public OrganizationmanagerGroupArgs()
        {
        }
        public static new OrganizationmanagerGroupArgs Empty => new OrganizationmanagerGroupArgs();
    }

    public sealed class OrganizationmanagerGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Computed) The SAML Federation creation timestamp.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The description of the Group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization to attach this Group to.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        public OrganizationmanagerGroupState()
        {
        }
        public static new OrganizationmanagerGroupState Empty => new OrganizationmanagerGroupState();
    }
}
