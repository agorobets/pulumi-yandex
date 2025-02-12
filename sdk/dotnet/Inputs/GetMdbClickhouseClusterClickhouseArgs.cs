// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetMdbClickhouseClusterClickhouseInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Main ClickHouse cluster configuration. The structure is documented below.
        /// </summary>
        [Input("config", required: true)]
        public Input<Inputs.GetMdbClickhouseClusterClickhouseConfigInputArgs> Config { get; set; } = null!;

        /// <summary>
        /// Resources allocated to hosts of the shard. The resources specified for the shard takes precedence over the resources specified for the cluster. The structure is documented below.
        /// </summary>
        [Input("resources", required: true)]
        public Input<Inputs.GetMdbClickhouseClusterClickhouseResourcesInputArgs> Resources { get; set; } = null!;

        public GetMdbClickhouseClusterClickhouseInputArgs()
        {
        }
        public static new GetMdbClickhouseClusterClickhouseInputArgs Empty => new GetMdbClickhouseClusterClickhouseInputArgs();
    }
}
