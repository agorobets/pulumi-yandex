// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class SwsWafProfileRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines is it rule blocking or not.
        /// </summary>
        [Input("isBlocking")]
        public Input<bool>? IsBlocking { get; set; }

        /// <summary>
        /// Determines is it rule enabled or not.
        /// </summary>
        [Input("isEnabled")]
        public Input<bool>? IsEnabled { get; set; }

        /// <summary>
        /// Rule ID.
        /// </summary>
        [Input("ruleId", required: true)]
        public Input<string> RuleId { get; set; } = null!;

        public SwsWafProfileRuleGetArgs()
        {
        }
        public static new SwsWafProfileRuleGetArgs Empty => new SwsWafProfileRuleGetArgs();
    }
}
