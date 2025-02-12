// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class SwsSecurityProfileSecurityRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional description of the rule. 0-512 characters long.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// This mode allows you to test your security profile or a single rule.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// Name of the rule. The name is unique within the security profile. 1-50 characters long.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Determines the priority for checking the incoming traffic.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Rule actions, see [Rule actions](https://yandex.cloud/en/docs/smartwebsecurity/concepts/rules#rule-action). The structure is documented below.
        /// </summary>
        [Input("ruleCondition")]
        public Input<Inputs.SwsSecurityProfileSecurityRuleRuleConditionGetArgs>? RuleCondition { get; set; }

        /// <summary>
        /// Smart Protection rule, see [Smart Protection rules](https://yandex.cloud/en/docs/smartwebsecurity/concepts/rules#smart-protection-rules). The structure is documented below.
        /// </summary>
        [Input("smartProtection")]
        public Input<Inputs.SwsSecurityProfileSecurityRuleSmartProtectionGetArgs>? SmartProtection { get; set; }

        /// <summary>
        /// Web Application Firewall (WAF) rule, see [WAF rules](https://yandex.cloud/en/docs/smartwebsecurity/concepts/rules#waf-rules). The structure is documented below.
        /// 
        /// &gt; Exactly one rule specifier: `smart_protection` or `rule_condition` or `waf` should be specified.
        /// </summary>
        [Input("waf")]
        public Input<Inputs.SwsSecurityProfileSecurityRuleWafGetArgs>? Waf { get; set; }

        public SwsSecurityProfileSecurityRuleGetArgs()
        {
        }
        public static new SwsSecurityProfileSecurityRuleGetArgs Empty => new SwsSecurityProfileSecurityRuleGetArgs();
    }
}
