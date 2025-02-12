// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class ServerlessContainerMountArgs : global::Pulumi.ResourceArgs
    {
        [Input("ephemeralDisk")]
        public Input<Inputs.ServerlessContainerMountEphemeralDiskArgs>? EphemeralDisk { get; set; }

        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("mountPointPath", required: true)]
        public Input<string> MountPointPath { get; set; } = null!;

        [Input("objectStorage")]
        public Input<Inputs.ServerlessContainerMountObjectStorageArgs>? ObjectStorage { get; set; }

        public ServerlessContainerMountArgs()
        {
        }
        public static new ServerlessContainerMountArgs Empty => new ServerlessContainerMountArgs();
    }
}
