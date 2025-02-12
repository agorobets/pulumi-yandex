// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Outputs
{

    [OutputType]
    public sealed class GetKubernetesNodeGroupInstanceTemplateResult
    {
        /// <summary>
        /// The specifications for boot disks that will be attached to the instance. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateBootDiskResult> BootDisks;
        /// <summary>
        /// Container network configuration. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateContainerNetworkResult> ContainerNetworks;
        /// <summary>
        /// Container runtime configuration. The structure is documented below.
        /// </summary>
        public readonly Outputs.GetKubernetesNodeGroupInstanceTemplateContainerRuntimeResult ContainerRuntime;
        /// <summary>
        /// GPU settings. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateGpuSettingResult> GpuSettings;
        /// <summary>
        /// Labels assigned to compute nodes (instances), created by the Node Group.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The set of metadata `key:value` pairs assigned to this instance template. This includes custom metadata and predefined keys.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// Name of a specific Kubernetes node group.
        /// 
        /// &gt; One of `node_group_id` or `name` should be specified.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A public address that can be used to access the internet over NAT.
        /// </summary>
        public readonly bool Nat;
        /// <summary>
        /// Type of network acceleration. Values: `standard`, `software_accelerated`.
        /// </summary>
        public readonly string NetworkAccelerationType;
        /// <summary>
        /// An array with the network interfaces that will be attached to the instance. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateNetworkInterfaceResult> NetworkInterfaces;
        /// <summary>
        /// (Optional) The placement policy configuration. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplatePlacementPolicyResult> PlacementPolicies;
        /// <summary>
        /// The ID of the hardware platform configuration for the instance.
        /// </summary>
        public readonly string PlatformId;
        public readonly ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateResourceResult> Resources;
        /// <summary>
        /// The scheduling policy for the instances in node group. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateSchedulingPolicyResult> SchedulingPolicies;

        [OutputConstructor]
        private GetKubernetesNodeGroupInstanceTemplateResult(
            ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateBootDiskResult> bootDisks,

            ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateContainerNetworkResult> containerNetworks,

            Outputs.GetKubernetesNodeGroupInstanceTemplateContainerRuntimeResult containerRuntime,

            ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateGpuSettingResult> gpuSettings,

            ImmutableDictionary<string, string> labels,

            ImmutableDictionary<string, string> metadata,

            string name,

            bool nat,

            string networkAccelerationType,

            ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateNetworkInterfaceResult> networkInterfaces,

            ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplatePlacementPolicyResult> placementPolicies,

            string platformId,

            ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateResourceResult> resources,

            ImmutableArray<Outputs.GetKubernetesNodeGroupInstanceTemplateSchedulingPolicyResult> schedulingPolicies)
        {
            BootDisks = bootDisks;
            ContainerNetworks = containerNetworks;
            ContainerRuntime = containerRuntime;
            GpuSettings = gpuSettings;
            Labels = labels;
            Metadata = metadata;
            Name = name;
            Nat = nat;
            NetworkAccelerationType = networkAccelerationType;
            NetworkInterfaces = networkInterfaces;
            PlacementPolicies = placementPolicies;
            PlatformId = platformId;
            Resources = resources;
            SchedulingPolicies = schedulingPolicies;
        }
    }
}
