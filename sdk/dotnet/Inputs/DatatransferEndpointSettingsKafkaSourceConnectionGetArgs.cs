// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferEndpointSettingsKafkaSourceConnectionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier of the Managed Kafka cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Connection settings of the on-premise Kafka server.
        /// </summary>
        [Input("onPremise")]
        public Input<Inputs.DatatransferEndpointSettingsKafkaSourceConnectionOnPremiseGetArgs>? OnPremise { get; set; }

        public DatatransferEndpointSettingsKafkaSourceConnectionGetArgs()
        {
        }
        public static new DatatransferEndpointSettingsKafkaSourceConnectionGetArgs Empty => new DatatransferEndpointSettingsKafkaSourceConnectionGetArgs();
    }
}
