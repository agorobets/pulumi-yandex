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
    public sealed class DatatransferTransferTransformationTransformerMaskField
    {
        /// <summary>
        /// List of strings that specify the name of the column for data masking (a regular expression).
        /// </summary>
        public readonly ImmutableArray<string> Columns;
        /// <summary>
        /// Mask function.
        /// </summary>
        public readonly Outputs.DatatransferTransferTransformationTransformerMaskFieldFunction? Function;
        /// <summary>
        /// Table filter.
        /// </summary>
        public readonly Outputs.DatatransferTransferTransformationTransformerMaskFieldTables? Tables;

        [OutputConstructor]
        private DatatransferTransferTransformationTransformerMaskField(
            ImmutableArray<string> columns,

            Outputs.DatatransferTransferTransformationTransformerMaskFieldFunction? function,

            Outputs.DatatransferTransferTransformationTransformerMaskFieldTables? tables)
        {
            Columns = columns;
            Function = function;
            Tables = tables;
        }
    }
}
