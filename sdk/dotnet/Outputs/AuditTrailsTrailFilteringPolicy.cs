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
    public sealed class AuditTrailsTrailFilteringPolicy
    {
        /// <summary>
        /// Structure describing filtering process for the service-specific data events.
        /// </summary>
        public readonly ImmutableArray<Outputs.AuditTrailsTrailFilteringPolicyDataEventsFilter> DataEventsFilters;
        /// <summary>
        /// Structure describing filtering process for management events.
        /// </summary>
        public readonly Outputs.AuditTrailsTrailFilteringPolicyManagementEventsFilter? ManagementEventsFilter;

        [OutputConstructor]
        private AuditTrailsTrailFilteringPolicy(
            ImmutableArray<Outputs.AuditTrailsTrailFilteringPolicyDataEventsFilter> dataEventsFilters,

            Outputs.AuditTrailsTrailFilteringPolicyManagementEventsFilter? managementEventsFilter)
        {
            DataEventsFilters = dataEventsFilters;
            ManagementEventsFilter = managementEventsFilter;
        }
    }
}
