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
    /// Allows creation and management of a single binding within IAM policy for an existing Yandex Resource Manager cloud.
    /// 
    /// ## Example Usage
    /// 
    /// {{ tffile "examples/resourcemanager_cloud_iam_binding/r_resourcemanager_cloud_iam_binding_1.tf" }}
    /// 
    /// ## Import
    /// 
    /// IAM binding imports use space-delimited identifiers; first the resource in question and then the role. These bindings can be imported using the `cloud_id` and role, e.g.
    /// 
    /// ```sh
    /// $ pulumi import yandex:index/resourcemanagerCloudIamBinding:ResourcemanagerCloudIamBinding viewer "cloud_id viewer"
    /// ```
    /// </summary>
    [YandexResourceType("yandex:index/resourcemanagerCloudIamBinding:ResourcemanagerCloudIamBinding")]
    public partial class ResourcemanagerCloudIamBinding : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the cloud to attach the policy to.
        /// </summary>
        [Output("cloudId")]
        public Output<string> CloudId { get; private set; } = null!;

        /// <summary>
        /// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following values:
        /// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
        /// * **serviceAccount:{service_account_id}**: A unique service account ID.
        /// * **federatedUser:{federated_user_id}**: A unique federated user ID.
        /// * **federatedUser:{federated_user_id}:**: A unique saml federation user account ID.
        /// * **group:{group_id}**: A unique group ID.
        /// * **system:group:federation:{federation_id}:users**: All users in federation.
        /// * **system:group:organization:{organization_id}:users**: All users in organization.
        /// * **system:allAuthenticatedUsers**: All authenticated users.
        /// * **system:allUsers**: All users, including unauthenticated ones.
        /// 
        /// Note: for more information about system groups, see the [documentation](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group).
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The role that should be assigned. Only one `yandex.ResourcemanagerCloudIamBinding` can be used per role.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        [Output("sleepAfter")]
        public Output<int?> SleepAfter { get; private set; } = null!;


        /// <summary>
        /// Create a ResourcemanagerCloudIamBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourcemanagerCloudIamBinding(string name, ResourcemanagerCloudIamBindingArgs args, CustomResourceOptions? options = null)
            : base("yandex:index/resourcemanagerCloudIamBinding:ResourcemanagerCloudIamBinding", name, args ?? new ResourcemanagerCloudIamBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourcemanagerCloudIamBinding(string name, Input<string> id, ResourcemanagerCloudIamBindingState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/resourcemanagerCloudIamBinding:ResourcemanagerCloudIamBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResourcemanagerCloudIamBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourcemanagerCloudIamBinding Get(string name, Input<string> id, ResourcemanagerCloudIamBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new ResourcemanagerCloudIamBinding(name, id, state, options);
        }
    }

    public sealed class ResourcemanagerCloudIamBindingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the cloud to attach the policy to.
        /// </summary>
        [Input("cloudId", required: true)]
        public Input<string> CloudId { get; set; } = null!;

        [Input("members", required: true)]
        private InputList<string>? _members;

        /// <summary>
        /// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following values:
        /// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
        /// * **serviceAccount:{service_account_id}**: A unique service account ID.
        /// * **federatedUser:{federated_user_id}**: A unique federated user ID.
        /// * **federatedUser:{federated_user_id}:**: A unique saml federation user account ID.
        /// * **group:{group_id}**: A unique group ID.
        /// * **system:group:federation:{federation_id}:users**: All users in federation.
        /// * **system:group:organization:{organization_id}:users**: All users in organization.
        /// * **system:allAuthenticatedUsers**: All authenticated users.
        /// * **system:allUsers**: All users, including unauthenticated ones.
        /// 
        /// Note: for more information about system groups, see the [documentation](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group).
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The role that should be assigned. Only one `yandex.ResourcemanagerCloudIamBinding` can be used per role.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        [Input("sleepAfter")]
        public Input<int>? SleepAfter { get; set; }

        public ResourcemanagerCloudIamBindingArgs()
        {
        }
        public static new ResourcemanagerCloudIamBindingArgs Empty => new ResourcemanagerCloudIamBindingArgs();
    }

    public sealed class ResourcemanagerCloudIamBindingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the cloud to attach the policy to.
        /// </summary>
        [Input("cloudId")]
        public Input<string>? CloudId { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following values:
        /// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
        /// * **serviceAccount:{service_account_id}**: A unique service account ID.
        /// * **federatedUser:{federated_user_id}**: A unique federated user ID.
        /// * **federatedUser:{federated_user_id}:**: A unique saml federation user account ID.
        /// * **group:{group_id}**: A unique group ID.
        /// * **system:group:federation:{federation_id}:users**: All users in federation.
        /// * **system:group:organization:{organization_id}:users**: All users in organization.
        /// * **system:allAuthenticatedUsers**: All authenticated users.
        /// * **system:allUsers**: All users, including unauthenticated ones.
        /// 
        /// Note: for more information about system groups, see the [documentation](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group).
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The role that should be assigned. Only one `yandex.ResourcemanagerCloudIamBinding` can be used per role.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("sleepAfter")]
        public Input<int>? SleepAfter { get; set; }

        public ResourcemanagerCloudIamBindingState()
        {
        }
        public static new ResourcemanagerCloudIamBindingState Empty => new ResourcemanagerCloudIamBindingState();
    }
}
