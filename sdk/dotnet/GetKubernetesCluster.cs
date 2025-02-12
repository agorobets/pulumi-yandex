// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetKubernetesCluster
    {
        /// <summary>
        /// Get information about a Yandex Cloud Managed Kubernetes Cluster. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-kubernetes/concepts/#kubernetes-cluster).
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/kubernetes_cluster/d_kubernetes_cluster_1.tf" }}
        /// </summary>
        public static Task<GetKubernetesClusterResult> InvokeAsync(GetKubernetesClusterArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKubernetesClusterResult>("yandex:index/getKubernetesCluster:getKubernetesCluster", args ?? new GetKubernetesClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Yandex Cloud Managed Kubernetes Cluster. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-kubernetes/concepts/#kubernetes-cluster).
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/kubernetes_cluster/d_kubernetes_cluster_1.tf" }}
        /// </summary>
        public static Output<GetKubernetesClusterResult> Invoke(GetKubernetesClusterInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKubernetesClusterResult>("yandex:index/getKubernetesCluster:getKubernetesCluster", args ?? new GetKubernetesClusterInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Yandex Cloud Managed Kubernetes Cluster. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-kubernetes/concepts/#kubernetes-cluster).
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/kubernetes_cluster/d_kubernetes_cluster_1.tf" }}
        /// </summary>
        public static Output<GetKubernetesClusterResult> Invoke(GetKubernetesClusterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetKubernetesClusterResult>("yandex:index/getKubernetesCluster:getKubernetesCluster", args ?? new GetKubernetesClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKubernetesClusterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of a specific Kubernetes cluster.
        /// </summary>
        [Input("clusterId")]
        public string? ClusterId { get; set; }

        /// <summary>
        /// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public string? FolderId { get; set; }

        /// <summary>
        /// Name of a specific Kubernetes cluster.
        /// 
        /// &gt; One of `cluster_id` or `name` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetKubernetesClusterArgs()
        {
        }
        public static new GetKubernetesClusterArgs Empty => new GetKubernetesClusterArgs();
    }

    public sealed class GetKubernetesClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of a specific Kubernetes cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// Name of a specific Kubernetes cluster.
        /// 
        /// &gt; One of `cluster_id` or `name` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetKubernetesClusterInvokeArgs()
        {
        }
        public static new GetKubernetesClusterInvokeArgs Empty => new GetKubernetesClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetKubernetesClusterResult
    {
        public readonly string ClusterId;
        /// <summary>
        /// IP range for allocating pod addresses.
        /// </summary>
        public readonly string ClusterIpv4Range;
        /// <summary>
        /// Identical to cluster_ipv4_range but for the IPv6 protocol.
        /// </summary>
        public readonly string ClusterIpv6Range;
        /// <summary>
        /// The Kubernetes cluster creation timestamp.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// A description of the Kubernetes cluster.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (Optional) ID of the folder default Log group of which should be used to collect logs.
        /// </summary>
        public readonly string FolderId;
        /// <summary>
        /// Health of the Kubernetes cluster.
        /// </summary>
        public readonly string Health;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// cluster KMS provider parameters.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesClusterKmsProviderResult> KmsProviders;
        /// <summary>
        /// A set of key/value label pairs to assign to the Kubernetes cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// (Optional) ID of the Yandex Cloud Logging [Log group](https://cloud.yandex.com/docs/logging/concepts/log-group).
        /// </summary>
        public readonly string LogGroupId;
        /// <summary>
        /// Kubernetes master configuration options. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesClusterMasterResult> Masters;
        public readonly string Name;
        /// <summary>
        /// The ID of the cluster network.
        /// </summary>
        public readonly string NetworkId;
        /// <summary>
        /// (Optional) Network Implementation options. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesClusterNetworkImplementationResult> NetworkImplementations;
        /// <summary>
        /// Network policy provider for the cluster, if present. Possible values: `CALICO`.
        /// </summary>
        public readonly string NetworkPolicyProvider;
        /// <summary>
        /// Size of the masks that are assigned to each node in the cluster.
        /// </summary>
        public readonly int NodeIpv4CidrMaskSize;
        /// <summary>
        /// Service account to be used by the worker nodes of the Kubernetes cluster to access Container Registry or to push node logs and metrics.
        /// </summary>
        public readonly string NodeServiceAccountId;
        /// <summary>
        /// Cluster release channel.
        /// </summary>
        public readonly string ReleaseChannel;
        /// <summary>
        /// Service account to be used for provisioning Compute Cloud and VPC resources for Kubernetes cluster. Selected service account should have `edit` role on the folder where the Kubernetes cluster will be located and on the folder where selected network resides.
        /// </summary>
        public readonly string ServiceAccountId;
        /// <summary>
        /// IP range Kubernetes services Kubernetes cluster IP addresses will be allocated from.
        /// </summary>
        public readonly string ServiceIpv4Range;
        /// <summary>
        /// Identical to service_ipv4_range but for the IPv6 protocol.
        /// </summary>
        public readonly string ServiceIpv6Range;
        /// <summary>
        /// Status of the Kubernetes cluster.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetKubernetesClusterResult(
            string clusterId,

            string clusterIpv4Range,

            string clusterIpv6Range,

            string createdAt,

            string description,

            string folderId,

            string health,

            string id,

            ImmutableArray<Outputs.GetKubernetesClusterKmsProviderResult> kmsProviders,

            ImmutableDictionary<string, string> labels,

            string logGroupId,

            ImmutableArray<Outputs.GetKubernetesClusterMasterResult> masters,

            string name,

            string networkId,

            ImmutableArray<Outputs.GetKubernetesClusterNetworkImplementationResult> networkImplementations,

            string networkPolicyProvider,

            int nodeIpv4CidrMaskSize,

            string nodeServiceAccountId,

            string releaseChannel,

            string serviceAccountId,

            string serviceIpv4Range,

            string serviceIpv6Range,

            string status)
        {
            ClusterId = clusterId;
            ClusterIpv4Range = clusterIpv4Range;
            ClusterIpv6Range = clusterIpv6Range;
            CreatedAt = createdAt;
            Description = description;
            FolderId = folderId;
            Health = health;
            Id = id;
            KmsProviders = kmsProviders;
            Labels = labels;
            LogGroupId = logGroupId;
            Masters = masters;
            Name = name;
            NetworkId = networkId;
            NetworkImplementations = networkImplementations;
            NetworkPolicyProvider = networkPolicyProvider;
            NodeIpv4CidrMaskSize = nodeIpv4CidrMaskSize;
            NodeServiceAccountId = nodeServiceAccountId;
            ReleaseChannel = releaseChannel;
            ServiceAccountId = serviceAccountId;
            ServiceIpv4Range = serviceIpv4Range;
            ServiceIpv6Range = serviceIpv6Range;
            Status = status;
        }
    }
}
