// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class SwsSecurityProfileSecurityRuleWafConditionHeaderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the security profile. The name is unique within the folder. 1-50 characters long.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("value", required: true)]
        public Input<Inputs.SwsSecurityProfileSecurityRuleWafConditionHeaderValueArgs> Value { get; set; } = null!;

        public SwsSecurityProfileSecurityRuleWafConditionHeaderArgs()
        {
        }
        public static new SwsSecurityProfileSecurityRuleWafConditionHeaderArgs Empty => new SwsSecurityProfileSecurityRuleWafConditionHeaderArgs();
    }
}
