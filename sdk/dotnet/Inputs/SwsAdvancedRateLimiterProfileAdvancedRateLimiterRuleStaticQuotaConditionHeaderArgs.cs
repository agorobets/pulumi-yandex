// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionHeaderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the ARL profile. The name is unique within the folder. 1-50 characters long.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// String value of the key.
        /// </summary>
        [Input("value", required: true)]
        public Input<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionHeaderValueArgs> Value { get; set; } = null!;

        public SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionHeaderArgs()
        {
        }
        public static new SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionHeaderArgs Empty => new SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionHeaderArgs();
    }
}
