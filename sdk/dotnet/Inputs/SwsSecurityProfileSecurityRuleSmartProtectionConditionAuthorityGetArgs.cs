// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class SwsSecurityProfileSecurityRuleSmartProtectionConditionAuthorityGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("authorities")]
        private InputList<Inputs.SwsSecurityProfileSecurityRuleSmartProtectionConditionAuthorityAuthorityGetArgs>? _authorities;
        public InputList<Inputs.SwsSecurityProfileSecurityRuleSmartProtectionConditionAuthorityAuthorityGetArgs> Authorities
        {
            get => _authorities ?? (_authorities = new InputList<Inputs.SwsSecurityProfileSecurityRuleSmartProtectionConditionAuthorityAuthorityGetArgs>());
            set => _authorities = value;
        }

        public SwsSecurityProfileSecurityRuleSmartProtectionConditionAuthorityGetArgs()
        {
        }
        public static new SwsSecurityProfileSecurityRuleSmartProtectionConditionAuthorityGetArgs Empty => new SwsSecurityProfileSecurityRuleSmartProtectionConditionAuthorityGetArgs();
    }
}
