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
    public sealed class GetSwsWafProfileCoreRuleSetResult
    {
        /// <summary>
        /// Anomaly score. Enter an integer within the range of 2 and 10000. The higher this value, the more likely it is that the request that satisfies the rule is an attack. See [Rules](https://yandex.cloud/en/docs/smartwebsecurity/concepts/waf#anomaly) for more details.
        /// </summary>
        public readonly int InboundAnomalyScore;
        /// <summary>
        /// Paranoia level. Enter an integer within the range of 1 and 4. Paranoia level classifies rules according to their aggression. The higher the paranoia level, the better your protection, but also the higher the probability of WAF false positives. See [Rules](https://yandex.cloud/en/docs/smartwebsecurity/concepts/waf#paranoia) for more details. NOTE: this option has no effect on enabling or disabling rules, it is used only as recommendation for user to enable all rules with paranoia_level &lt;= this value.
        /// </summary>
        public readonly int ParanoiaLevel;
        public readonly ImmutableArray<Outputs.GetSwsWafProfileCoreRuleSetRuleSetResult> RuleSets;

        [OutputConstructor]
        private GetSwsWafProfileCoreRuleSetResult(
            int inboundAnomalyScore,

            int paranoiaLevel,

            ImmutableArray<Outputs.GetSwsWafProfileCoreRuleSetRuleSetResult> ruleSets)
        {
            InboundAnomalyScore = inboundAnomalyScore;
            ParanoiaLevel = paranoiaLevel;
            RuleSets = ruleSets;
        }
    }
}
