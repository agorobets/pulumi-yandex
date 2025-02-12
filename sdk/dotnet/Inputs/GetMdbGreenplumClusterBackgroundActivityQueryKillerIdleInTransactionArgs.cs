// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetMdbGreenplumClusterBackgroundActivityQueryKillerIdleInTransactionInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Flag that indicates whether script is enabled.
        /// </summary>
        [Input("enable", required: true)]
        public Input<bool> Enable { get; set; } = null!;

        [Input("ignoreUsers", required: true)]
        private InputList<string>? _ignoreUsers;

        /// <summary>
        /// List of users to ignore when considering queries to terminate.
        /// </summary>
        public InputList<string> IgnoreUsers
        {
            get => _ignoreUsers ?? (_ignoreUsers = new InputList<string>());
            set => _ignoreUsers = value;
        }

        /// <summary>
        /// Maximum duration for this type of queries (in seconds).
        /// </summary>
        [Input("maxAge", required: true)]
        public Input<int> MaxAge { get; set; } = null!;

        public GetMdbGreenplumClusterBackgroundActivityQueryKillerIdleInTransactionInputArgs()
        {
        }
        public static new GetMdbGreenplumClusterBackgroundActivityQueryKillerIdleInTransactionInputArgs Empty => new GetMdbGreenplumClusterBackgroundActivityQueryKillerIdleInTransactionInputArgs();
    }
}
