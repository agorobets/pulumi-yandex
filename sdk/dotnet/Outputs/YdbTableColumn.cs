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
    public sealed class YdbTableColumn
    {
        /// <summary>
        /// Column group
        /// </summary>
        public readonly string? Family;
        /// <summary>
        /// Column family name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A column cannot have the NULL data type. ( Default: false )
        /// </summary>
        public readonly bool? NotNull;
        /// <summary>
        /// Column data type. YQL data types are used.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private YdbTableColumn(
            string? family,

            string name,

            bool? notNull,

            string type)
        {
            Family = family;
            Name = name;
            NotNull = notNull;
            Type = type;
        }
    }
}
