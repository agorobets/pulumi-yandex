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
    public sealed class ServerlessContainerMount
    {
        public readonly Outputs.ServerlessContainerMountEphemeralDisk? EphemeralDisk;
        public readonly string? Mode;
        public readonly string MountPointPath;
        public readonly Outputs.ServerlessContainerMountObjectStorage? ObjectStorage;

        [OutputConstructor]
        private ServerlessContainerMount(
            Outputs.ServerlessContainerMountEphemeralDisk? ephemeralDisk,

            string? mode,

            string mountPointPath,

            Outputs.ServerlessContainerMountObjectStorage? objectStorage)
        {
            EphemeralDisk = ephemeralDisk;
            Mode = mode;
            MountPointPath = mountPointPath;
            ObjectStorage = objectStorage;
        }
    }
}
