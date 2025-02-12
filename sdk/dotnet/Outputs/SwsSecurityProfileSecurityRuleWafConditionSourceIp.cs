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
    public sealed class SwsSecurityProfileSecurityRuleWafConditionSourceIp
    {
        public readonly Outputs.SwsSecurityProfileSecurityRuleWafConditionSourceIpGeoIpMatch? GeoIpMatch;
        public readonly Outputs.SwsSecurityProfileSecurityRuleWafConditionSourceIpGeoIpNotMatch? GeoIpNotMatch;
        public readonly Outputs.SwsSecurityProfileSecurityRuleWafConditionSourceIpIpRangesMatch? IpRangesMatch;
        public readonly Outputs.SwsSecurityProfileSecurityRuleWafConditionSourceIpIpRangesNotMatch? IpRangesNotMatch;

        [OutputConstructor]
        private SwsSecurityProfileSecurityRuleWafConditionSourceIp(
            Outputs.SwsSecurityProfileSecurityRuleWafConditionSourceIpGeoIpMatch? geoIpMatch,

            Outputs.SwsSecurityProfileSecurityRuleWafConditionSourceIpGeoIpNotMatch? geoIpNotMatch,

            Outputs.SwsSecurityProfileSecurityRuleWafConditionSourceIpIpRangesMatch? ipRangesMatch,

            Outputs.SwsSecurityProfileSecurityRuleWafConditionSourceIpIpRangesNotMatch? ipRangesNotMatch)
        {
            GeoIpMatch = geoIpMatch;
            GeoIpNotMatch = geoIpNotMatch;
            IpRangesMatch = ipRangesMatch;
            IpRangesNotMatch = ipRangesNotMatch;
        }
    }
}
