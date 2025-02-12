// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetMonitoringDashboard
    {
        /// <summary>
        /// Get information about a Yandex Monitoring dashboard.
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/monitoring_dashboard/d_monitoring_dashboard_1.tf" }}
        /// </summary>
        public static Task<GetMonitoringDashboardResult> InvokeAsync(GetMonitoringDashboardArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMonitoringDashboardResult>("yandex:index/getMonitoringDashboard:getMonitoringDashboard", args ?? new GetMonitoringDashboardArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Yandex Monitoring dashboard.
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/monitoring_dashboard/d_monitoring_dashboard_1.tf" }}
        /// </summary>
        public static Output<GetMonitoringDashboardResult> Invoke(GetMonitoringDashboardInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMonitoringDashboardResult>("yandex:index/getMonitoringDashboard:getMonitoringDashboard", args ?? new GetMonitoringDashboardInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Yandex Monitoring dashboard.
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/monitoring_dashboard/d_monitoring_dashboard_1.tf" }}
        /// </summary>
        public static Output<GetMonitoringDashboardResult> Invoke(GetMonitoringDashboardInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMonitoringDashboardResult>("yandex:index/getMonitoringDashboard:getMonitoringDashboard", args ?? new GetMonitoringDashboardInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMonitoringDashboardArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Dashboard ID.
        /// </summary>
        [Input("dashboardId")]
        public string? DashboardId { get; set; }

        /// <summary>
        /// Chart description in dashboard (not enabled in UI).
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
        /// 
        /// &gt; One of `dashboard_id` or `name` should be specified.
        /// </summary>
        [Input("folderId")]
        public string? FolderId { get; set; }

        /// <summary>
        /// Name of the Dashboard.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetMonitoringDashboardArgs()
        {
        }
        public static new GetMonitoringDashboardArgs Empty => new GetMonitoringDashboardArgs();
    }

    public sealed class GetMonitoringDashboardInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Dashboard ID.
        /// </summary>
        [Input("dashboardId")]
        public Input<string>? DashboardId { get; set; }

        /// <summary>
        /// Chart description in dashboard (not enabled in UI).
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
        /// 
        /// &gt; One of `dashboard_id` or `name` should be specified.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// Name of the Dashboard.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetMonitoringDashboardInvokeArgs()
        {
        }
        public static new GetMonitoringDashboardInvokeArgs Empty => new GetMonitoringDashboardInvokeArgs();
    }


    [OutputType]
    public sealed class GetMonitoringDashboardResult
    {
        public readonly string? DashboardId;
        /// <summary>
        /// Chart description in dashboard (not enabled in UI).
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Labels folder ID.
        /// </summary>
        public readonly string FolderId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A set of key/value label pairs to assign to the Dashboard.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Series name or empty.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Dashboard parametrization
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMonitoringDashboardParametrizationResult> Parametrizations;
        /// <summary>
        /// Title or empty.
        /// </summary>
        public readonly string Title;
        /// <summary>
        /// Widgets
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMonitoringDashboardWidgetResult> Widgets;

        [OutputConstructor]
        private GetMonitoringDashboardResult(
            string? dashboardId,

            string? description,

            string folderId,

            string id,

            ImmutableDictionary<string, string> labels,

            string? name,

            ImmutableArray<Outputs.GetMonitoringDashboardParametrizationResult> parametrizations,

            string title,

            ImmutableArray<Outputs.GetMonitoringDashboardWidgetResult> widgets)
        {
            DashboardId = dashboardId;
            Description = description;
            FolderId = folderId;
            Id = id;
            Labels = labels;
            Name = name;
            Parametrizations = parametrizations;
            Title = title;
            Widgets = widgets;
        }
    }
}
