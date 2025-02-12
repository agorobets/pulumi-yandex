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
    public sealed class DatatransferTransferTransformation
    {
        /// <summary>
        /// A list of transformers. You can specify exactly 1 transformer in each element of list.
        /// </summary>
        public readonly ImmutableArray<Outputs.DatatransferTransferTransformationTransformer> Transformers;

        [OutputConstructor]
        private DatatransferTransferTransformation(ImmutableArray<Outputs.DatatransferTransferTransformationTransformer> transformers)
        {
            Transformers = transformers;
        }
    }
}
