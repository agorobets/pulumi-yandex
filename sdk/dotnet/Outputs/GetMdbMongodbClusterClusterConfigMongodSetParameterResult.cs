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
    public sealed class GetMdbMongodbClusterClusterConfigMongodSetParameterResult
    {
        /// <summary>
        /// (Optional) Enables the auditing of authorization successes. Can be either true or false. For more information, see the [auditAuthorizationSuccess](https://www.mongodb.com/docs/manual/reference/parameters/#mongodb-parameter-param.auditAuthorizationSuccess) description in the official documentation. Available only in enterprise edition.
        /// </summary>
        public readonly bool? AuditAuthorizationSuccess;
        public readonly bool? EnableFlowControl;
        public readonly int? MinSnapshotHistoryWindowInSeconds;

        [OutputConstructor]
        private GetMdbMongodbClusterClusterConfigMongodSetParameterResult(
            bool? auditAuthorizationSuccess,

            bool? enableFlowControl,

            int? minSnapshotHistoryWindowInSeconds)
        {
            AuditAuthorizationSuccess = auditAuthorizationSuccess;
            EnableFlowControl = enableFlowControl;
            MinSnapshotHistoryWindowInSeconds = minSnapshotHistoryWindowInSeconds;
        }
    }
}
