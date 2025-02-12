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
    public sealed class GetMonitoringDashboardParametrizationParameterLabelValueResult
    {
        /// <summary>
        /// Default value.
        /// </summary>
        public readonly ImmutableArray<string> DefaultValues;
        /// <summary>
        /// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
        /// 
        /// &gt; One of `dashboard_id` or `name` should be specified.
        /// </summary>
        public readonly string FolderId;
        /// <summary>
        /// Label key to list label values.
        /// </summary>
        public readonly string LabelKey;
        /// <summary>
        /// Specifies the multiselectable values of parameter.
        /// </summary>
        public readonly bool Multiselectable;
        /// <summary>
        /// Selectors to select metric label values.
        /// </summary>
        public readonly string Selectors;

        [OutputConstructor]
        private GetMonitoringDashboardParametrizationParameterLabelValueResult(
            ImmutableArray<string> defaultValues,

            string folderId,

            string labelKey,

            bool multiselectable,

            string selectors)
        {
            DefaultValues = defaultValues;
            FolderId = folderId;
            LabelKey = labelKey;
            Multiselectable = multiselectable;
            Selectors = selectors;
        }
    }
}
