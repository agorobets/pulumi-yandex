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
    /// Allows management of [Yandex Cloud API Gateway](https://cloud.yandex.com/docs/api-gateway/).
    /// 
    /// ## Example Usage
    /// 
    /// {{ tffile "examples/api_gateway/r_api_gateway_1.tf" }}
    /// </summary>
    [YandexResourceType("yandex:index/apiGateway:ApiGateway")]
    public partial class ApiGateway : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Canary release settings of gateway.
        /// * `canary.0.weight` - Percentage of requests, which will be processed by canary release.
        /// * `canary.0.variables` - A list of values for variables in gateway specification of canary release.
        /// </summary>
        [Output("canary")]
        public Output<Outputs.ApiGatewayCanary?> Canary { get; private set; } = null!;

        /// <summary>
        /// Gateway connectivity. If specified the gateway will be attached to specified network.
        /// * `connectivity.0.network_id` - Network the gateway will have access to. It's essential to specify network with subnets in all availability zones.
        /// </summary>
        [Output("connectivity")]
        public Output<Outputs.ApiGatewayConnectivity?> Connectivity { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp of the Yandex Cloud API Gateway.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Set of custom domains to be attached to Yandex API Gateway.
        /// </summary>
        [Output("customDomains")]
        public Output<ImmutableArray<Outputs.ApiGatewayCustomDomain>> CustomDomains { get; private set; } = null!;

        /// <summary>
        /// Description of the Yandex Cloud API Gateway.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Default domain for the Yandex API Gateway. Generated at creation time.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Execution timeout in seconds for the Yandex Cloud API Gateway.
        /// </summary>
        [Output("executionTimeout")]
        public Output<string> ExecutionTimeout { get; private set; } = null!;

        /// <summary>
        /// Folder ID for the Yandex Cloud API Gateway. If it is not provided, the default provider folder is used.
        /// </summary>
        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        /// <summary>
        /// A set of key/value label pairs to assign to the Yandex Cloud API Gateway.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Log entries are written to specified log group
        /// </summary>
        [Output("logGroupId")]
        public Output<string> LogGroupId { get; private set; } = null!;

        /// <summary>
        /// Options for logging from Yandex Cloud API Gateway.
        /// </summary>
        [Output("logOptions")]
        public Output<Outputs.ApiGatewayLogOptions?> LogOptions { get; private set; } = null!;

        /// <summary>
        /// Yandex Cloud API Gateway name used to define API Gateway.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// OpenAPI specification for Yandex API Gateway.
        /// </summary>
        [Output("spec")]
        public Output<string> Spec { get; private set; } = null!;

        /// <summary>
        /// Status of the Yandex API Gateway.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// (**DEPRECATED**, use `custom_domains` instead) Set of user domains attached to Yandex API Gateway.
        /// </summary>
        [Output("userDomains")]
        public Output<ImmutableArray<string>> UserDomains { get; private set; } = null!;

        /// <summary>
        /// A set of values for variables in gateway specification.
        /// </summary>
        [Output("variables")]
        public Output<ImmutableDictionary<string, string>?> Variables { get; private set; } = null!;


        /// <summary>
        /// Create a ApiGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiGateway(string name, ApiGatewayArgs args, CustomResourceOptions? options = null)
            : base("yandex:index/apiGateway:ApiGateway", name, args ?? new ApiGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiGateway(string name, Input<string> id, ApiGatewayState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/apiGateway:ApiGateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApiGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiGateway Get(string name, Input<string> id, ApiGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new ApiGateway(name, id, state, options);
        }
    }

    public sealed class ApiGatewayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Canary release settings of gateway.
        /// * `canary.0.weight` - Percentage of requests, which will be processed by canary release.
        /// * `canary.0.variables` - A list of values for variables in gateway specification of canary release.
        /// </summary>
        [Input("canary")]
        public Input<Inputs.ApiGatewayCanaryArgs>? Canary { get; set; }

        /// <summary>
        /// Gateway connectivity. If specified the gateway will be attached to specified network.
        /// * `connectivity.0.network_id` - Network the gateway will have access to. It's essential to specify network with subnets in all availability zones.
        /// </summary>
        [Input("connectivity")]
        public Input<Inputs.ApiGatewayConnectivityArgs>? Connectivity { get; set; }

        [Input("customDomains")]
        private InputList<Inputs.ApiGatewayCustomDomainArgs>? _customDomains;

        /// <summary>
        /// Set of custom domains to be attached to Yandex API Gateway.
        /// </summary>
        public InputList<Inputs.ApiGatewayCustomDomainArgs> CustomDomains
        {
            get => _customDomains ?? (_customDomains = new InputList<Inputs.ApiGatewayCustomDomainArgs>());
            set => _customDomains = value;
        }

        /// <summary>
        /// Description of the Yandex Cloud API Gateway.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Execution timeout in seconds for the Yandex Cloud API Gateway.
        /// </summary>
        [Input("executionTimeout")]
        public Input<string>? ExecutionTimeout { get; set; }

        /// <summary>
        /// Folder ID for the Yandex Cloud API Gateway. If it is not provided, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the Yandex Cloud API Gateway.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Options for logging from Yandex Cloud API Gateway.
        /// </summary>
        [Input("logOptions")]
        public Input<Inputs.ApiGatewayLogOptionsArgs>? LogOptions { get; set; }

        /// <summary>
        /// Yandex Cloud API Gateway name used to define API Gateway.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// OpenAPI specification for Yandex API Gateway.
        /// </summary>
        [Input("spec", required: true)]
        public Input<string> Spec { get; set; } = null!;

        [Input("variables")]
        private InputMap<string>? _variables;

        /// <summary>
        /// A set of values for variables in gateway specification.
        /// </summary>
        public InputMap<string> Variables
        {
            get => _variables ?? (_variables = new InputMap<string>());
            set => _variables = value;
        }

        public ApiGatewayArgs()
        {
        }
        public static new ApiGatewayArgs Empty => new ApiGatewayArgs();
    }

    public sealed class ApiGatewayState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Canary release settings of gateway.
        /// * `canary.0.weight` - Percentage of requests, which will be processed by canary release.
        /// * `canary.0.variables` - A list of values for variables in gateway specification of canary release.
        /// </summary>
        [Input("canary")]
        public Input<Inputs.ApiGatewayCanaryGetArgs>? Canary { get; set; }

        /// <summary>
        /// Gateway connectivity. If specified the gateway will be attached to specified network.
        /// * `connectivity.0.network_id` - Network the gateway will have access to. It's essential to specify network with subnets in all availability zones.
        /// </summary>
        [Input("connectivity")]
        public Input<Inputs.ApiGatewayConnectivityGetArgs>? Connectivity { get; set; }

        /// <summary>
        /// Creation timestamp of the Yandex Cloud API Gateway.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("customDomains")]
        private InputList<Inputs.ApiGatewayCustomDomainGetArgs>? _customDomains;

        /// <summary>
        /// Set of custom domains to be attached to Yandex API Gateway.
        /// </summary>
        public InputList<Inputs.ApiGatewayCustomDomainGetArgs> CustomDomains
        {
            get => _customDomains ?? (_customDomains = new InputList<Inputs.ApiGatewayCustomDomainGetArgs>());
            set => _customDomains = value;
        }

        /// <summary>
        /// Description of the Yandex Cloud API Gateway.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Default domain for the Yandex API Gateway. Generated at creation time.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Execution timeout in seconds for the Yandex Cloud API Gateway.
        /// </summary>
        [Input("executionTimeout")]
        public Input<string>? ExecutionTimeout { get; set; }

        /// <summary>
        /// Folder ID for the Yandex Cloud API Gateway. If it is not provided, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the Yandex Cloud API Gateway.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Log entries are written to specified log group
        /// </summary>
        [Input("logGroupId")]
        public Input<string>? LogGroupId { get; set; }

        /// <summary>
        /// Options for logging from Yandex Cloud API Gateway.
        /// </summary>
        [Input("logOptions")]
        public Input<Inputs.ApiGatewayLogOptionsGetArgs>? LogOptions { get; set; }

        /// <summary>
        /// Yandex Cloud API Gateway name used to define API Gateway.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// OpenAPI specification for Yandex API Gateway.
        /// </summary>
        [Input("spec")]
        public Input<string>? Spec { get; set; }

        /// <summary>
        /// Status of the Yandex API Gateway.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("userDomains")]
        private InputList<string>? _userDomains;

        /// <summary>
        /// (**DEPRECATED**, use `custom_domains` instead) Set of user domains attached to Yandex API Gateway.
        /// </summary>
        [Obsolete(@"The 'user_domains' field has been deprecated. Please use 'custom_domains' instead.")]
        public InputList<string> UserDomains
        {
            get => _userDomains ?? (_userDomains = new InputList<string>());
            set => _userDomains = value;
        }

        [Input("variables")]
        private InputMap<string>? _variables;

        /// <summary>
        /// A set of values for variables in gateway specification.
        /// </summary>
        public InputMap<string> Variables
        {
            get => _variables ?? (_variables = new InputMap<string>());
            set => _variables = value;
        }

        public ApiGatewayState()
        {
        }
        public static new ApiGatewayState Empty => new ApiGatewayState();
    }
}
