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
    public sealed class GetComputeInstancePlacementPolicyResult
    {
        /// <summary>
        /// List of host affinity rules. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetComputeInstancePlacementPolicyHostAffinityRuleResult> HostAffinityRules;
        /// <summary>
        /// Specifies the id of the Placement Group to assign to the instance.
        /// </summary>
        public readonly string? PlacementGroupId;
        /// <summary>
        /// Specifies the number of partition in the Placement Group with the partition placement strategy.
        /// </summary>
        public readonly int? PlacementGroupPartition;

        [OutputConstructor]
        private GetComputeInstancePlacementPolicyResult(
            ImmutableArray<Outputs.GetComputeInstancePlacementPolicyHostAffinityRuleResult> hostAffinityRules,

            string? placementGroupId,

            int? placementGroupPartition)
        {
            HostAffinityRules = hostAffinityRules;
            PlacementGroupId = placementGroupId;
            PlacementGroupPartition = placementGroupPartition;
        }
    }
}
