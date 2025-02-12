// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetContainerRegistryIpPermission
    {
        /// <summary>
        /// Get information about a Yandex Container Registry IP Permission. For more information, see [the official documentation](https://yandex.cloud/docs/container-registry/operations/registry/registry-access)
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/container_registry_ip_permission/d_container_registry_ip_permission_1.tf" }}
        /// </summary>
        public static Task<GetContainerRegistryIpPermissionResult> InvokeAsync(GetContainerRegistryIpPermissionArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContainerRegistryIpPermissionResult>("yandex:index/getContainerRegistryIpPermission:getContainerRegistryIpPermission", args ?? new GetContainerRegistryIpPermissionArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Yandex Container Registry IP Permission. For more information, see [the official documentation](https://yandex.cloud/docs/container-registry/operations/registry/registry-access)
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/container_registry_ip_permission/d_container_registry_ip_permission_1.tf" }}
        /// </summary>
        public static Output<GetContainerRegistryIpPermissionResult> Invoke(GetContainerRegistryIpPermissionInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContainerRegistryIpPermissionResult>("yandex:index/getContainerRegistryIpPermission:getContainerRegistryIpPermission", args ?? new GetContainerRegistryIpPermissionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Yandex Container Registry IP Permission. For more information, see [the official documentation](https://yandex.cloud/docs/container-registry/operations/registry/registry-access)
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/container_registry_ip_permission/d_container_registry_ip_permission_1.tf" }}
        /// </summary>
        public static Output<GetContainerRegistryIpPermissionResult> Invoke(GetContainerRegistryIpPermissionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetContainerRegistryIpPermissionResult>("yandex:index/getContainerRegistryIpPermission:getContainerRegistryIpPermission", args ?? new GetContainerRegistryIpPermissionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContainerRegistryIpPermissionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of a specific Container Registry.
        /// </summary>
        [Input("registryId")]
        public string? RegistryId { get; set; }

        /// <summary>
        /// The Name of specific Container Registry.
        /// 
        /// &gt; Either `registry_id` or `registry_name` must be specified.
        /// </summary>
        [Input("registryName")]
        public string? RegistryName { get; set; }

        public GetContainerRegistryIpPermissionArgs()
        {
        }
        public static new GetContainerRegistryIpPermissionArgs Empty => new GetContainerRegistryIpPermissionArgs();
    }

    public sealed class GetContainerRegistryIpPermissionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of a specific Container Registry.
        /// </summary>
        [Input("registryId")]
        public Input<string>? RegistryId { get; set; }

        /// <summary>
        /// The Name of specific Container Registry.
        /// 
        /// &gt; Either `registry_id` or `registry_name` must be specified.
        /// </summary>
        [Input("registryName")]
        public Input<string>? RegistryName { get; set; }

        public GetContainerRegistryIpPermissionInvokeArgs()
        {
        }
        public static new GetContainerRegistryIpPermissionInvokeArgs Empty => new GetContainerRegistryIpPermissionInvokeArgs();
    }


    [OutputType]
    public sealed class GetContainerRegistryIpPermissionResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of configured CIDRs, from which pull is allowed.
        /// </summary>
        public readonly ImmutableArray<string> Pulls;
        /// <summary>
        /// List of configured CIDRs, from which push is allowed.
        /// </summary>
        public readonly ImmutableArray<string> Pushes;
        public readonly string RegistryId;
        public readonly string RegistryName;

        [OutputConstructor]
        private GetContainerRegistryIpPermissionResult(
            string id,

            ImmutableArray<string> pulls,

            ImmutableArray<string> pushes,

            string registryId,

            string registryName)
        {
            Id = id;
            Pulls = pulls;
            Pushes = pushes;
            RegistryId = registryId;
            RegistryName = registryName;
        }
    }
}
