// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class ComputeImageHardwareGenerationLegacyFeaturesArgs : global::Pulumi.ResourceArgs
    {
        [Input("pciTopology")]
        public Input<string>? PciTopology { get; set; }

        public ComputeImageHardwareGenerationLegacyFeaturesArgs()
        {
        }
        public static new ComputeImageHardwareGenerationLegacyFeaturesArgs Empty => new ComputeImageHardwareGenerationLegacyFeaturesArgs();
    }
}
