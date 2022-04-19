// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Foo struct {
	pulumi.ResourceState
}

// NewFoo registers a new resource with the given unique name, arguments, and options.
func NewFoo(ctx *pulumi.Context,
	name string, args *FooArgs, opts ...pulumi.ResourceOption) (*Foo, error) {
	if args == nil {
		args = &FooArgs{}
	}

	var resource Foo
	err := ctx.RegisterRemoteComponentResource("example::Foo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type fooArgs struct {
}

// The set of arguments for constructing a Foo resource.
type FooArgs struct {
}

func (FooArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fooArgs)(nil)).Elem()
}

func (r *Foo) GetKubeconfig(ctx *pulumi.Context, args *FooGetKubeconfigArgs) (pulumi.StringOutput, error) {
	out, err := ctx.Call("example::Foo/getKubeconfig", args, fooGetKubeconfigResultOutput{}, r)
	if err != nil {
		return pulumi.StringOutput{}, err
	}
	return out.(fooGetKubeconfigResultOutput).Kubeconfig(), nil
}

type fooGetKubeconfigArgs struct {
	ProfileName *string `pulumi:"profileName"`
	RoleArn     *string `pulumi:"roleArn"`
}

// The set of arguments for the GetKubeconfig method of the Foo resource.
type FooGetKubeconfigArgs struct {
	ProfileName pulumi.StringPtrInput
	RoleArn     pulumi.StringPtrInput
}

func (FooGetKubeconfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fooGetKubeconfigArgs)(nil)).Elem()
}

type fooGetKubeconfigResult struct {
	Kubeconfig string `pulumi:"kubeconfig"`
}

type fooGetKubeconfigResultOutput struct{ *pulumi.OutputState }

func (fooGetKubeconfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*fooGetKubeconfigResult)(nil)).Elem()
}

func (o fooGetKubeconfigResultOutput) Kubeconfig() pulumi.StringOutput {
	return o.ApplyT(func(v fooGetKubeconfigResult) string { return v.Kubeconfig }).(pulumi.StringOutput)
}

type FooInput interface {
	pulumi.Input

	ToFooOutput() FooOutput
	ToFooOutputWithContext(ctx context.Context) FooOutput
}

func (*Foo) ElementType() reflect.Type {
	return reflect.TypeOf((**Foo)(nil)).Elem()
}

func (i *Foo) ToFooOutput() FooOutput {
	return i.ToFooOutputWithContext(context.Background())
}

func (i *Foo) ToFooOutputWithContext(ctx context.Context) FooOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FooOutput)
}

type FooOutput struct{ *pulumi.OutputState }

func (FooOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Foo)(nil)).Elem()
}

func (o FooOutput) ToFooOutput() FooOutput {
	return o
}

func (o FooOutput) ToFooOutputWithContext(ctx context.Context) FooOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FooInput)(nil)).Elem(), &Foo{})
	pulumi.RegisterOutputType(FooOutput{})
	pulumi.RegisterOutputType(fooGetKubeconfigResultOutput{})
}
