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
    public sealed class YdbDatabaseServerlessServerlessDatabase
    {
        public readonly bool? EnableThrottlingRcuLimit;
        public readonly int? ProvisionedRcuLimit;
        public readonly int? StorageSizeLimit;
        public readonly int? ThrottlingRcuLimit;

        [OutputConstructor]
        private YdbDatabaseServerlessServerlessDatabase(
            bool? enableThrottlingRcuLimit,

            int? provisionedRcuLimit,

            int? storageSizeLimit,

            int? throttlingRcuLimit)
        {
            EnableThrottlingRcuLimit = enableThrottlingRcuLimit;
            ProvisionedRcuLimit = provisionedRcuLimit;
            StorageSizeLimit = storageSizeLimit;
            ThrottlingRcuLimit = throttlingRcuLimit;
        }
    }
}
