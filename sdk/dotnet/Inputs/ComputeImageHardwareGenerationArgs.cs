// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class ComputeImageHardwareGenerationArgs : global::Pulumi.ResourceArgs
    {
        [Input("generation2Features")]
        public Input<Inputs.ComputeImageHardwareGenerationGeneration2FeaturesArgs>? Generation2Features { get; set; }

        [Input("legacyFeatures")]
        public Input<Inputs.ComputeImageHardwareGenerationLegacyFeaturesArgs>? LegacyFeatures { get; set; }

        public ComputeImageHardwareGenerationArgs()
        {
        }
        public static new ComputeImageHardwareGenerationArgs Empty => new ComputeImageHardwareGenerationArgs();
    }
}
