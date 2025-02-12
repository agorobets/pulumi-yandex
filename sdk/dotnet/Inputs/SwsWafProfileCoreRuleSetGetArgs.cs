// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class SwsWafProfileCoreRuleSetGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Anomaly score. Enter an integer within the range of 2 and 10000. The higher this value, the more likely it is that the request that satisfies the rule is an attack. See [Rules](https://yandex.cloud/en/docs/smartwebsecurity/concepts/waf#anomaly) for more details.
        /// </summary>
        [Input("inboundAnomalyScore")]
        public Input<int>? InboundAnomalyScore { get; set; }

        /// <summary>
        /// Paranoia level. Enter an integer within the range of 1 and 4. Paranoia level classifies rules according to their aggression. The higher the paranoia level, the better your protection, but also the higher the probability of WAF false positives. See [Rules](https://yandex.cloud/en/docs/smartwebsecurity/concepts/waf#paranoia) for more details. NOTE: this option has no effect on enabling or disabling rules, it is used only as recommendation for user to enable all rules with paranoia_level &lt;= this value.
        /// </summary>
        [Input("paranoiaLevel")]
        public Input<int>? ParanoiaLevel { get; set; }

        [Input("ruleSet", required: true)]
        public Input<Inputs.SwsWafProfileCoreRuleSetRuleSetGetArgs> RuleSet { get; set; } = null!;

        public SwsWafProfileCoreRuleSetGetArgs()
        {
        }
        public static new SwsWafProfileCoreRuleSetGetArgs Empty => new SwsWafProfileCoreRuleSetGetArgs();
    }
}
