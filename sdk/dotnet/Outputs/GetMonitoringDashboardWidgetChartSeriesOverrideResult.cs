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
    public sealed class GetMonitoringDashboardWidgetChartSeriesOverrideResult
    {
        /// <summary>
        /// Name of the Dashboard.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Override settings.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMonitoringDashboardWidgetChartSeriesOverrideSettingResult> Settings;
        /// <summary>
        /// Series index.
        /// </summary>
        public readonly string TargetIndex;

        [OutputConstructor]
        private GetMonitoringDashboardWidgetChartSeriesOverrideResult(
            string name,

            ImmutableArray<Outputs.GetMonitoringDashboardWidgetChartSeriesOverrideSettingResult> settings,

            string targetIndex)
        {
            Name = name;
            Settings = settings;
            TargetIndex = targetIndex;
        }
    }
}
