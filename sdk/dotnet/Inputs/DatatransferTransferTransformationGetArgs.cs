// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferTransferTransformationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("transformers")]
        private InputList<Inputs.DatatransferTransferTransformationTransformerGetArgs>? _transformers;

        /// <summary>
        /// A list of transformers. You can specify exactly 1 transformer in each element of list.
        /// </summary>
        public InputList<Inputs.DatatransferTransferTransformationTransformerGetArgs> Transformers
        {
            get => _transformers ?? (_transformers = new InputList<Inputs.DatatransferTransferTransformationTransformerGetArgs>());
            set => _transformers = value;
        }

        public DatatransferTransferTransformationGetArgs()
        {
        }
        public static new DatatransferTransferTransformationGetArgs Empty => new DatatransferTransferTransformationGetArgs();
    }
}
