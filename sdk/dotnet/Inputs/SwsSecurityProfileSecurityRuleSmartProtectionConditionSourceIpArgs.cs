// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class SwsSecurityProfileSecurityRuleSmartProtectionConditionSourceIpArgs : global::Pulumi.ResourceArgs
    {
        [Input("geoIpMatch")]
        public Input<Inputs.SwsSecurityProfileSecurityRuleSmartProtectionConditionSourceIpGeoIpMatchArgs>? GeoIpMatch { get; set; }

        [Input("geoIpNotMatch")]
        public Input<Inputs.SwsSecurityProfileSecurityRuleSmartProtectionConditionSourceIpGeoIpNotMatchArgs>? GeoIpNotMatch { get; set; }

        [Input("ipRangesMatch")]
        public Input<Inputs.SwsSecurityProfileSecurityRuleSmartProtectionConditionSourceIpIpRangesMatchArgs>? IpRangesMatch { get; set; }

        [Input("ipRangesNotMatch")]
        public Input<Inputs.SwsSecurityProfileSecurityRuleSmartProtectionConditionSourceIpIpRangesNotMatchArgs>? IpRangesNotMatch { get; set; }

        public SwsSecurityProfileSecurityRuleSmartProtectionConditionSourceIpArgs()
        {
        }
        public static new SwsSecurityProfileSecurityRuleSmartProtectionConditionSourceIpArgs Empty => new SwsSecurityProfileSecurityRuleSmartProtectionConditionSourceIpArgs();
    }
}
