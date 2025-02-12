// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    /// <summary>
    /// Manages a gateway within the Yandex Cloud. For more information, see [the official documentation](https://yandex.cloud/docs/vpc/concepts/gateways).
    /// 
    /// * How-to Guides
    ///   * [Cloud Networking](https://yandex.cloud/docs/vpc/)
    /// 
    /// ## Example Usage
    /// 
    /// {{ tffile "examples/vpc_gateway/r_vpc_gateway_1.tf" }}
    /// 
    /// ## Import
    /// 
    /// A gateway can be imported using the `id` of the resource, e.g.
    /// 
    /// ```sh
    /// $ pulumi import yandex:index/vpcGateway:VpcGateway default gateway_id
    /// ```
    /// </summary>
    [YandexResourceType("yandex:index/vpcGateway:VpcGateway")]
    public partial class VpcGateway : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Creation timestamp of the key.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
        /// </summary>
        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        /// <summary>
        /// Labels to apply to this VPC Gateway. A list of key/value pairs.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Name of the VPC Gateway. Provided by the client when the VPC Gateway is created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Shared egress gateway configuration. Currently empty.
        /// </summary>
        [Output("sharedEgressGateway")]
        public Output<Outputs.VpcGatewaySharedEgressGateway?> SharedEgressGateway { get; private set; } = null!;


        /// <summary>
        /// Create a VpcGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcGateway(string name, VpcGatewayArgs? args = null, CustomResourceOptions? options = null)
            : base("yandex:index/vpcGateway:VpcGateway", name, args ?? new VpcGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcGateway(string name, Input<string> id, VpcGatewayState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/vpcGateway:VpcGateway", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/agorobets/pulumi-yandex/releases",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VpcGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcGateway Get(string name, Input<string> id, VpcGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcGateway(name, id, state, options);
        }
    }

    public sealed class VpcGatewayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this VPC Gateway. A list of key/value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the VPC Gateway. Provided by the client when the VPC Gateway is created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Shared egress gateway configuration. Currently empty.
        /// </summary>
        [Input("sharedEgressGateway")]
        public Input<Inputs.VpcGatewaySharedEgressGatewayArgs>? SharedEgressGateway { get; set; }

        public VpcGatewayArgs()
        {
        }
        public static new VpcGatewayArgs Empty => new VpcGatewayArgs();
    }

    public sealed class VpcGatewayState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation timestamp of the key.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this VPC Gateway. A list of key/value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the VPC Gateway. Provided by the client when the VPC Gateway is created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Shared egress gateway configuration. Currently empty.
        /// </summary>
        [Input("sharedEgressGateway")]
        public Input<Inputs.VpcGatewaySharedEgressGatewayGetArgs>? SharedEgressGateway { get; set; }

        public VpcGatewayState()
        {
        }
        public static new VpcGatewayState Empty => new VpcGatewayState();
    }
}
