// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/agorobets/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about Yandex SmartCaptcha. For more information, see [the official documentation](https://yandex.cloud/docs/smartcaptcha/).
//
// This data source is used to define Captcha that can be used by other resources.
//
// ## Example Usage
//
// {{ tffile "examples/smartcaptcha_captcha/d_smartcaptcha_captcha_1.tf" }}
//
// {{ tffile "examples/smartcaptcha_captcha/d_smartcaptcha_captcha_2.tf" }}
func LookupSmartcaptchaCaptcha(ctx *pulumi.Context, args *LookupSmartcaptchaCaptchaArgs, opts ...pulumi.InvokeOption) (*LookupSmartcaptchaCaptchaResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSmartcaptchaCaptchaResult
	err := ctx.Invoke("yandex:index/getSmartcaptchaCaptcha:getSmartcaptchaCaptcha", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSmartcaptchaCaptcha.
type LookupSmartcaptchaCaptchaArgs struct {
	// ID of the Captcha.
	//
	// > One of `captchaId` or `name` should be specified.
	CaptchaId *string `pulumi:"captchaId"`
	CloudId   *string `pulumi:"cloudId"`
	FolderId  *string `pulumi:"folderId"`
	// Name of the Captcha.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getSmartcaptchaCaptcha.
type LookupSmartcaptchaCaptchaResult struct {
	// List of allowed host names, see [Domain validation](https://yandex.cloud/docs/smartcaptcha/concepts/domain-validation).
	AllowedSites []string `pulumi:"allowedSites"`
	CaptchaId    *string  `pulumi:"captchaId"`
	// Additional task type of the captcha. Possible values are documented below.
	ChallengeType string `pulumi:"challengeType"`
	// Client key of the captcha, see [CAPTCHA keys](https://yandex.cloud/docs/smartcaptcha/concepts/keys).
	ClientKey string `pulumi:"clientKey"`
	CloudId   string `pulumi:"cloudId"`
	// Complexity of the captcha.
	Complexity string `pulumi:"complexity"`
	// The Captcha creation timestamp.
	CreatedAt string `pulumi:"createdAt"`
	// Determines whether captcha is protected from being deleted.
	DeletionProtection bool   `pulumi:"deletionProtection"`
	FolderId           string `pulumi:"folderId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the rule. The name is unique within the captcha. 1-50 characters long.
	Name string `pulumi:"name"`
	// List of variants to use in security_rules. The structure is documented below.
	OverrideVariants []GetSmartcaptchaCaptchaOverrideVariant `pulumi:"overrideVariants"`
	// Additional task type of the captcha.
	PreCheckType string `pulumi:"preCheckType"`
	// List of security rules. The structure is documented below.
	SecurityRules []GetSmartcaptchaCaptchaSecurityRule `pulumi:"securityRules"`
	// JSON with variables to define the captcha appearance. For more details see generated JSON in cloud console.
	StyleJson string `pulumi:"styleJson"`
	Suspend   bool   `pulumi:"suspend"`
	// Turn off host name check, see [Domain validation](https://yandex.cloud/docs/smartcaptcha/concepts/domain-validation).
	TurnOffHostnameCheck bool `pulumi:"turnOffHostnameCheck"`
}

func LookupSmartcaptchaCaptchaOutput(ctx *pulumi.Context, args LookupSmartcaptchaCaptchaOutputArgs, opts ...pulumi.InvokeOption) LookupSmartcaptchaCaptchaResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSmartcaptchaCaptchaResultOutput, error) {
			args := v.(LookupSmartcaptchaCaptchaArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getSmartcaptchaCaptcha:getSmartcaptchaCaptcha", args, LookupSmartcaptchaCaptchaResultOutput{}, options).(LookupSmartcaptchaCaptchaResultOutput), nil
		}).(LookupSmartcaptchaCaptchaResultOutput)
}

// A collection of arguments for invoking getSmartcaptchaCaptcha.
type LookupSmartcaptchaCaptchaOutputArgs struct {
	// ID of the Captcha.
	//
	// > One of `captchaId` or `name` should be specified.
	CaptchaId pulumi.StringPtrInput `pulumi:"captchaId"`
	CloudId   pulumi.StringPtrInput `pulumi:"cloudId"`
	FolderId  pulumi.StringPtrInput `pulumi:"folderId"`
	// Name of the Captcha.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupSmartcaptchaCaptchaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSmartcaptchaCaptchaArgs)(nil)).Elem()
}

// A collection of values returned by getSmartcaptchaCaptcha.
type LookupSmartcaptchaCaptchaResultOutput struct{ *pulumi.OutputState }

func (LookupSmartcaptchaCaptchaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSmartcaptchaCaptchaResult)(nil)).Elem()
}

func (o LookupSmartcaptchaCaptchaResultOutput) ToLookupSmartcaptchaCaptchaResultOutput() LookupSmartcaptchaCaptchaResultOutput {
	return o
}

func (o LookupSmartcaptchaCaptchaResultOutput) ToLookupSmartcaptchaCaptchaResultOutputWithContext(ctx context.Context) LookupSmartcaptchaCaptchaResultOutput {
	return o
}

// List of allowed host names, see [Domain validation](https://yandex.cloud/docs/smartcaptcha/concepts/domain-validation).
func (o LookupSmartcaptchaCaptchaResultOutput) AllowedSites() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSmartcaptchaCaptchaResult) []string { return v.AllowedSites }).(pulumi.StringArrayOutput)
}

func (o LookupSmartcaptchaCaptchaResultOutput) CaptchaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSmartcaptchaCaptchaResult) *string { return v.CaptchaId }).(pulumi.StringPtrOutput)
}

// Additional task type of the captcha. Possible values are documented below.
func (o LookupSmartcaptchaCaptchaResultOutput) ChallengeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSmartcaptchaCaptchaResult) string { return v.ChallengeType }).(pulumi.StringOutput)
}

// Client key of the captcha, see [CAPTCHA keys](https://yandex.cloud/docs/smartcaptcha/concepts/keys).
func (o LookupSmartcaptchaCaptchaResultOutput) ClientKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSmartcaptchaCaptchaResult) string { return v.ClientKey }).(pulumi.StringOutput)
}

func (o LookupSmartcaptchaCaptchaResultOutput) CloudId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSmartcaptchaCaptchaResult) string { return v.CloudId }).(pulumi.StringOutput)
}

// Complexity of the captcha.
func (o LookupSmartcaptchaCaptchaResultOutput) Complexity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSmartcaptchaCaptchaResult) string { return v.Complexity }).(pulumi.StringOutput)
}

// The Captcha creation timestamp.
func (o LookupSmartcaptchaCaptchaResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSmartcaptchaCaptchaResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Determines whether captcha is protected from being deleted.
func (o LookupSmartcaptchaCaptchaResultOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSmartcaptchaCaptchaResult) bool { return v.DeletionProtection }).(pulumi.BoolOutput)
}

func (o LookupSmartcaptchaCaptchaResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSmartcaptchaCaptchaResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSmartcaptchaCaptchaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSmartcaptchaCaptchaResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the rule. The name is unique within the captcha. 1-50 characters long.
func (o LookupSmartcaptchaCaptchaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSmartcaptchaCaptchaResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of variants to use in security_rules. The structure is documented below.
func (o LookupSmartcaptchaCaptchaResultOutput) OverrideVariants() GetSmartcaptchaCaptchaOverrideVariantArrayOutput {
	return o.ApplyT(func(v LookupSmartcaptchaCaptchaResult) []GetSmartcaptchaCaptchaOverrideVariant {
		return v.OverrideVariants
	}).(GetSmartcaptchaCaptchaOverrideVariantArrayOutput)
}

// Additional task type of the captcha.
func (o LookupSmartcaptchaCaptchaResultOutput) PreCheckType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSmartcaptchaCaptchaResult) string { return v.PreCheckType }).(pulumi.StringOutput)
}

// List of security rules. The structure is documented below.
func (o LookupSmartcaptchaCaptchaResultOutput) SecurityRules() GetSmartcaptchaCaptchaSecurityRuleArrayOutput {
	return o.ApplyT(func(v LookupSmartcaptchaCaptchaResult) []GetSmartcaptchaCaptchaSecurityRule { return v.SecurityRules }).(GetSmartcaptchaCaptchaSecurityRuleArrayOutput)
}

// JSON with variables to define the captcha appearance. For more details see generated JSON in cloud console.
func (o LookupSmartcaptchaCaptchaResultOutput) StyleJson() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSmartcaptchaCaptchaResult) string { return v.StyleJson }).(pulumi.StringOutput)
}

func (o LookupSmartcaptchaCaptchaResultOutput) Suspend() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSmartcaptchaCaptchaResult) bool { return v.Suspend }).(pulumi.BoolOutput)
}

// Turn off host name check, see [Domain validation](https://yandex.cloud/docs/smartcaptcha/concepts/domain-validation).
func (o LookupSmartcaptchaCaptchaResultOutput) TurnOffHostnameCheck() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSmartcaptchaCaptchaResult) bool { return v.TurnOffHostnameCheck }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSmartcaptchaCaptchaResultOutput{})
}
