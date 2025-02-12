// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetOrganizationmanagerUserSshKey
    {
        /// <summary>
        /// Get information about a Yandex Cloud User SSH Key.
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/organizationmanager_user_ssh_key/d_organizationmanager_user_ssh_key_1.tf" }}
        /// </summary>
        public static Task<GetOrganizationmanagerUserSshKeyResult> InvokeAsync(GetOrganizationmanagerUserSshKeyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationmanagerUserSshKeyResult>("yandex:index/getOrganizationmanagerUserSshKey:getOrganizationmanagerUserSshKey", args ?? new GetOrganizationmanagerUserSshKeyArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Yandex Cloud User SSH Key.
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/organizationmanager_user_ssh_key/d_organizationmanager_user_ssh_key_1.tf" }}
        /// </summary>
        public static Output<GetOrganizationmanagerUserSshKeyResult> Invoke(GetOrganizationmanagerUserSshKeyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationmanagerUserSshKeyResult>("yandex:index/getOrganizationmanagerUserSshKey:getOrganizationmanagerUserSshKey", args ?? new GetOrganizationmanagerUserSshKeyInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Yandex Cloud User SSH Key.
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/organizationmanager_user_ssh_key/d_organizationmanager_user_ssh_key_1.tf" }}
        /// </summary>
        public static Output<GetOrganizationmanagerUserSshKeyResult> Invoke(GetOrganizationmanagerUserSshKeyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationmanagerUserSshKeyResult>("yandex:index/getOrganizationmanagerUserSshKey:getOrganizationmanagerUserSshKey", args ?? new GetOrganizationmanagerUserSshKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrganizationmanagerUserSshKeyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Data of the user ssh key.
        /// </summary>
        [Input("data")]
        public string? Data { get; set; }

        /// <summary>
        /// User ssh key will be no longer valid after expiration timestamp.
        /// </summary>
        [Input("expiresAt")]
        public string? ExpiresAt { get; set; }

        /// <summary>
        /// Name of the user ssh key.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Organization that the user ssh key belongs to.
        /// </summary>
        [Input("organizationId")]
        public string? OrganizationId { get; set; }

        /// <summary>
        /// Subject that the user ssh key belongs to.
        /// </summary>
        [Input("subjectId")]
        public string? SubjectId { get; set; }

        /// <summary>
        /// ID of the user ssh key.
        /// </summary>
        [Input("userSshKeyId", required: true)]
        public string UserSshKeyId { get; set; } = null!;

        public GetOrganizationmanagerUserSshKeyArgs()
        {
        }
        public static new GetOrganizationmanagerUserSshKeyArgs Empty => new GetOrganizationmanagerUserSshKeyArgs();
    }

    public sealed class GetOrganizationmanagerUserSshKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Data of the user ssh key.
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// User ssh key will be no longer valid after expiration timestamp.
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// Name of the user ssh key.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Organization that the user ssh key belongs to.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// Subject that the user ssh key belongs to.
        /// </summary>
        [Input("subjectId")]
        public Input<string>? SubjectId { get; set; }

        /// <summary>
        /// ID of the user ssh key.
        /// </summary>
        [Input("userSshKeyId", required: true)]
        public Input<string> UserSshKeyId { get; set; } = null!;

        public GetOrganizationmanagerUserSshKeyInvokeArgs()
        {
        }
        public static new GetOrganizationmanagerUserSshKeyInvokeArgs Empty => new GetOrganizationmanagerUserSshKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrganizationmanagerUserSshKeyResult
    {
        /// <summary>
        /// User ssh key creation timestamp.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Data of the user ssh key.
        /// </summary>
        public readonly string? Data;
        /// <summary>
        /// User ssh key will be no longer valid after expiration timestamp.
        /// </summary>
        public readonly string? ExpiresAt;
        /// <summary>
        /// Auto generated fingerprint of the user ssh key.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the user ssh key.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Organization that the user ssh key belongs to.
        /// </summary>
        public readonly string? OrganizationId;
        /// <summary>
        /// Subject that the user ssh key belongs to.
        /// </summary>
        public readonly string? SubjectId;
        public readonly string UserSshKeyId;

        [OutputConstructor]
        private GetOrganizationmanagerUserSshKeyResult(
            string createdAt,

            string? data,

            string? expiresAt,

            string fingerprint,

            string id,

            string? name,

            string? organizationId,

            string? subjectId,

            string userSshKeyId)
        {
            CreatedAt = createdAt;
            Data = data;
            ExpiresAt = expiresAt;
            Fingerprint = fingerprint;
            Id = id;
            Name = name;
            OrganizationId = organizationId;
            SubjectId = subjectId;
            UserSshKeyId = userSshKeyId;
        }
    }
}
