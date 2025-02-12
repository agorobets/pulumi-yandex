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
    public sealed class GetMdbGreenplumClusterBackgroundActivityAnalyzeAndVacuumResult
    {
        /// <summary>
        /// Maximum duration of the `ANALYZE` operation, in seconds. The default value is `36000`. As soon as this period expires, the `ANALYZE` operation will be forced to terminate.
        /// </summary>
        public readonly int AnalyzeTimeout;
        /// <summary>
        /// Time of day in 'HH:MM' format when scripts should run.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Maximum duration of the `VACUUM` operation, in seconds. The default value is `36000`. As soon as this period expires, the `VACUUM` operation will be forced to terminate.
        /// </summary>
        public readonly int VacuumTimeout;

        [OutputConstructor]
        private GetMdbGreenplumClusterBackgroundActivityAnalyzeAndVacuumResult(
            int analyzeTimeout,

            string startTime,

            int vacuumTimeout)
        {
            AnalyzeTimeout = analyzeTimeout;
            StartTime = startTime;
            VacuumTimeout = vacuumTimeout;
        }
    }
}
