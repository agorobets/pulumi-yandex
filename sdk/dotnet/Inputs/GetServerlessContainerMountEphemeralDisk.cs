// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetServerlessContainerMountEphemeralDiskArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Optional block size of the ephemeral disk in KB
        /// </summary>
        [Input("blockSizeKb", required: true)]
        public int BlockSizeKb { get; set; }

        /// <summary>
        /// Size of the ephemeral disk in GB
        /// </summary>
        [Input("sizeGb", required: true)]
        public int SizeGb { get; set; }

        public GetServerlessContainerMountEphemeralDiskArgs()
        {
        }
        public static new GetServerlessContainerMountEphemeralDiskArgs Empty => new GetServerlessContainerMountEphemeralDiskArgs();
    }
}
