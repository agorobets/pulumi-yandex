// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/agorobets/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex Cloud Service Agent.
//
// ## Example Usage
//
// {{ tffile "examples/iam_service_agent/d_iam_service_agent_1.tf" }}
func GetIamServiceAgent(ctx *pulumi.Context, args *GetIamServiceAgentArgs, opts ...pulumi.InvokeOption) (*GetIamServiceAgentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIamServiceAgentResult
	err := ctx.Invoke("yandex:index/getIamServiceAgent:getIamServiceAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIamServiceAgent.
type GetIamServiceAgentArgs struct {
	// ID of the cloud.
	CloudId *string `pulumi:"cloudId"`
	// ID of the service-control microservice.
	MicroserviceId *string `pulumi:"microserviceId"`
	// ID of the service-control service.
	ServiceId *string `pulumi:"serviceId"`
}

// A collection of values returned by getIamServiceAgent.
type GetIamServiceAgentResult struct {
	CloudId string `pulumi:"cloudId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the service-control microservice.
	MicroserviceId string `pulumi:"microserviceId"`
	// ID of the service-control service.
	ServiceId string `pulumi:"serviceId"`
}

func GetIamServiceAgentOutput(ctx *pulumi.Context, args GetIamServiceAgentOutputArgs, opts ...pulumi.InvokeOption) GetIamServiceAgentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetIamServiceAgentResultOutput, error) {
			args := v.(GetIamServiceAgentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getIamServiceAgent:getIamServiceAgent", args, GetIamServiceAgentResultOutput{}, options).(GetIamServiceAgentResultOutput), nil
		}).(GetIamServiceAgentResultOutput)
}

// A collection of arguments for invoking getIamServiceAgent.
type GetIamServiceAgentOutputArgs struct {
	// ID of the cloud.
	CloudId pulumi.StringPtrInput `pulumi:"cloudId"`
	// ID of the service-control microservice.
	MicroserviceId pulumi.StringPtrInput `pulumi:"microserviceId"`
	// ID of the service-control service.
	ServiceId pulumi.StringPtrInput `pulumi:"serviceId"`
}

func (GetIamServiceAgentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIamServiceAgentArgs)(nil)).Elem()
}

// A collection of values returned by getIamServiceAgent.
type GetIamServiceAgentResultOutput struct{ *pulumi.OutputState }

func (GetIamServiceAgentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIamServiceAgentResult)(nil)).Elem()
}

func (o GetIamServiceAgentResultOutput) ToGetIamServiceAgentResultOutput() GetIamServiceAgentResultOutput {
	return o
}

func (o GetIamServiceAgentResultOutput) ToGetIamServiceAgentResultOutputWithContext(ctx context.Context) GetIamServiceAgentResultOutput {
	return o
}

func (o GetIamServiceAgentResultOutput) CloudId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIamServiceAgentResult) string { return v.CloudId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIamServiceAgentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIamServiceAgentResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the service-control microservice.
func (o GetIamServiceAgentResultOutput) MicroserviceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIamServiceAgentResult) string { return v.MicroserviceId }).(pulumi.StringOutput)
}

// ID of the service-control service.
func (o GetIamServiceAgentResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIamServiceAgentResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIamServiceAgentResultOutput{})
}
