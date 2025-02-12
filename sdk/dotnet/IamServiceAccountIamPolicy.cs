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
    /// **IAM policy for a service account**
    /// 
    /// When managing IAM roles, you can treat a service account either as a resource or as an identity. This resource is used to add IAM policy bindings to a service account resource to configure permissions that define who can edit the service account.
    /// 
    /// There are three different resources that help you manage your IAM policy for a service account. Each of these resources is used for a different use case:
    /// 
    /// * yandex_iam_service_account_iam_policy: Authoritative. Sets the IAM policy for the service account and replaces any existing policy already attached.
    /// * yandex_iam_service_account_iam_binding: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service account are preserved.
    /// * yandex_iam_service_account_iam_member: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role of the service account are preserved.
    /// 
    /// &gt; `yandex.IamServiceAccountIamPolicy` **cannot** be used in conjunction with `yandex.IamServiceAccountIamBinding` and `yandex.IamServiceAccountIamMember` or they will conflict over what your policy should be.
    /// 
    /// &gt; `yandex.IamServiceAccountIamBinding` resources **can be** used in conjunction with `yandex.IamServiceAccountIamMember` resources **only if** they do not grant privileges to the same role.
    /// 
    /// ## Example Usage
    /// 
    /// {{ tffile "examples/iam_service_account_iam_policy/r_iam_service_account_iam_policy_1.tf" }}
    /// 
    /// ## Import
    /// 
    /// Service account IAM policy resources can be imported using the service account ID.
    /// 
    /// ```sh
    /// $ pulumi import yandex:index/iamServiceAccountIamPolicy:IamServiceAccountIamPolicy admin-account-iam service_account_id
    /// ```
    /// </summary>
    [YandexResourceType("yandex:index/iamServiceAccountIamPolicy:IamServiceAccountIamPolicy")]
    public partial class IamServiceAccountIamPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The policy data generated by a `yandex.getIamPolicy` data source.
        /// </summary>
        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;

        /// <summary>
        /// The service account ID to apply a policy to.
        /// </summary>
        [Output("serviceAccountId")]
        public Output<string> ServiceAccountId { get; private set; } = null!;


        /// <summary>
        /// Create a IamServiceAccountIamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamServiceAccountIamPolicy(string name, IamServiceAccountIamPolicyArgs args, CustomResourceOptions? options = null)
            : base("yandex:index/iamServiceAccountIamPolicy:IamServiceAccountIamPolicy", name, args ?? new IamServiceAccountIamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IamServiceAccountIamPolicy(string name, Input<string> id, IamServiceAccountIamPolicyState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/iamServiceAccountIamPolicy:IamServiceAccountIamPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IamServiceAccountIamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamServiceAccountIamPolicy Get(string name, Input<string> id, IamServiceAccountIamPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new IamServiceAccountIamPolicy(name, id, state, options);
        }
    }

    public sealed class IamServiceAccountIamPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy data generated by a `yandex.getIamPolicy` data source.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        /// <summary>
        /// The service account ID to apply a policy to.
        /// </summary>
        [Input("serviceAccountId", required: true)]
        public Input<string> ServiceAccountId { get; set; } = null!;

        public IamServiceAccountIamPolicyArgs()
        {
        }
        public static new IamServiceAccountIamPolicyArgs Empty => new IamServiceAccountIamPolicyArgs();
    }

    public sealed class IamServiceAccountIamPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy data generated by a `yandex.getIamPolicy` data source.
        /// </summary>
        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        /// <summary>
        /// The service account ID to apply a policy to.
        /// </summary>
        [Input("serviceAccountId")]
        public Input<string>? ServiceAccountId { get; set; }

        public IamServiceAccountIamPolicyState()
        {
        }
        public static new IamServiceAccountIamPolicyState Empty => new IamServiceAccountIamPolicyState();
    }
}
