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
    public sealed class SwsWafProfileExclusionRuleCondition
    {
        public readonly Outputs.SwsWafProfileExclusionRuleConditionAuthority? Authority;
        public readonly ImmutableArray<Outputs.SwsWafProfileExclusionRuleConditionHeader> Headers;
        public readonly Outputs.SwsWafProfileExclusionRuleConditionHttpMethod? HttpMethod;
        public readonly Outputs.SwsWafProfileExclusionRuleConditionRequestUri? RequestUri;
        public readonly Outputs.SwsWafProfileExclusionRuleConditionSourceIp? SourceIp;

        [OutputConstructor]
        private SwsWafProfileExclusionRuleCondition(
            Outputs.SwsWafProfileExclusionRuleConditionAuthority? authority,

            ImmutableArray<Outputs.SwsWafProfileExclusionRuleConditionHeader> headers,

            Outputs.SwsWafProfileExclusionRuleConditionHttpMethod? httpMethod,

            Outputs.SwsWafProfileExclusionRuleConditionRequestUri? requestUri,

            Outputs.SwsWafProfileExclusionRuleConditionSourceIp? sourceIp)
        {
            Authority = authority;
            Headers = headers;
            HttpMethod = httpMethod;
            RequestUri = requestUri;
            SourceIp = sourceIp;
        }
    }
}
