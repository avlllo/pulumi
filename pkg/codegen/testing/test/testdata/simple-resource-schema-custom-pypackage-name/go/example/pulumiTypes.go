// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigMap struct {
	Config *string `pulumi:"config"`
}

// ConfigMapInput is an input type that accepts ConfigMap and ConfigMapOutput values.
// You can construct a concrete instance of `ConfigMapInput` via:
//
//          ConfigMap{ "key": ConfigArgs{...} }
type ConfigMapInput interface {
	pulumi.Input

	ToConfigMapOutput() ConfigMapOutput
	ToConfigMapOutputWithContext(context.Context) ConfigMapOutput
}

type ConfigMapArgs struct {
	Config pulumi.StringPtrInput `pulumi:"config"`
}

func (ConfigMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigMap)(nil)).Elem()
}

func (i ConfigMapArgs) ToConfigMapOutput() ConfigMapOutput {
	return i.ToConfigMapOutputWithContext(context.Background())
}

func (i ConfigMapArgs) ToConfigMapOutputWithContext(ctx context.Context) ConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapOutput)
}

// ConfigMapArrayInput is an input type that accepts ConfigMapArray and ConfigMapArrayOutput values.
// You can construct a concrete instance of `ConfigMapArrayInput` via:
//
//          ConfigMapArray{ ConfigMapArgs{...} }
type ConfigMapArrayInput interface {
	pulumi.Input

	ToConfigMapArrayOutput() ConfigMapArrayOutput
	ToConfigMapArrayOutputWithContext(context.Context) ConfigMapArrayOutput
}

type ConfigMapArray []ConfigMapInput

func (ConfigMapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigMap)(nil)).Elem()
}

func (i ConfigMapArray) ToConfigMapArrayOutput() ConfigMapArrayOutput {
	return i.ToConfigMapArrayOutputWithContext(context.Background())
}

func (i ConfigMapArray) ToConfigMapArrayOutputWithContext(ctx context.Context) ConfigMapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapArrayOutput)
}

type ConfigMapOutput struct{ *pulumi.OutputState }

func (ConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigMap)(nil)).Elem()
}

func (o ConfigMapOutput) ToConfigMapOutput() ConfigMapOutput {
	return o
}

func (o ConfigMapOutput) ToConfigMapOutputWithContext(ctx context.Context) ConfigMapOutput {
	return o
}

func (o ConfigMapOutput) Config() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigMap) *string { return v.Config }).(pulumi.StringPtrOutput)
}

type ConfigMapArrayOutput struct{ *pulumi.OutputState }

func (ConfigMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigMap)(nil)).Elem()
}

func (o ConfigMapArrayOutput) ToConfigMapArrayOutput() ConfigMapArrayOutput {
	return o
}

func (o ConfigMapArrayOutput) ToConfigMapArrayOutputWithContext(ctx context.Context) ConfigMapArrayOutput {
	return o
}

func (o ConfigMapArrayOutput) Index(i pulumi.IntInput) ConfigMapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigMap {
		return vs[0].([]ConfigMap)[vs[1].(int)]
	}).(ConfigMapOutput)
}

type Object struct {
	Bar     *string     `pulumi:"bar"`
	Configs []ConfigMap `pulumi:"configs"`
	Foo     *Resource   `pulumi:"foo"`
	// List of lists of other objects
	Others [][]SomeOtherObject `pulumi:"others"`
	// Mapping from string to list of some other object
	StillOthers map[string][]SomeOtherObject `pulumi:"stillOthers"`
}

// ObjectInput is an input type that accepts ObjectArgs and ObjectOutput values.
// You can construct a concrete instance of `ObjectInput` via:
//
//          ObjectArgs{...}
type ObjectInput interface {
	pulumi.Input

	ToObjectOutput() ObjectOutput
	ToObjectOutputWithContext(context.Context) ObjectOutput
}

type ObjectArgs struct {
	Bar     pulumi.StringPtrInput `pulumi:"bar"`
	Configs ConfigMapArrayInput   `pulumi:"configs"`
	Foo     ResourceInput         `pulumi:"foo"`
	// List of lists of other objects
	Others SomeOtherObjectArrayArrayInput `pulumi:"others"`
	// Mapping from string to list of some other object
	StillOthers SomeOtherObjectArrayMapInput `pulumi:"stillOthers"`
}

func (ObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Object)(nil)).Elem()
}

func (i ObjectArgs) ToObjectOutput() ObjectOutput {
	return i.ToObjectOutputWithContext(context.Background())
}

func (i ObjectArgs) ToObjectOutputWithContext(ctx context.Context) ObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectOutput)
}

type ObjectOutput struct{ *pulumi.OutputState }

func (ObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Object)(nil)).Elem()
}

func (o ObjectOutput) ToObjectOutput() ObjectOutput {
	return o
}

func (o ObjectOutput) ToObjectOutputWithContext(ctx context.Context) ObjectOutput {
	return o
}

func (o ObjectOutput) Bar() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Object) *string { return v.Bar }).(pulumi.StringPtrOutput)
}

func (o ObjectOutput) Configs() ConfigMapArrayOutput {
	return o.ApplyT(func(v Object) []ConfigMap { return v.Configs }).(ConfigMapArrayOutput)
}

func (o ObjectOutput) Foo() ResourceOutput {
	return o.ApplyT(func(v Object) *Resource { return v.Foo }).(ResourceOutput)
}

// List of lists of other objects
func (o ObjectOutput) Others() SomeOtherObjectArrayArrayOutput {
	return o.ApplyT(func(v Object) [][]SomeOtherObject { return v.Others }).(SomeOtherObjectArrayArrayOutput)
}

// Mapping from string to list of some other object
func (o ObjectOutput) StillOthers() SomeOtherObjectArrayMapOutput {
	return o.ApplyT(func(v Object) map[string][]SomeOtherObject { return v.StillOthers }).(SomeOtherObjectArrayMapOutput)
}

type OtherResourceOutputType struct {
	Foo *string `pulumi:"foo"`
}

// OtherResourceOutputTypeInput is an input type that accepts OtherResourceOutputTypeArgs and OtherResourceOutputTypeOutput values.
// You can construct a concrete instance of `OtherResourceOutputTypeInput` via:
//
//          OtherResourceOutputTypeArgs{...}
type OtherResourceOutputTypeInput interface {
	pulumi.Input

	ToOtherResourceOutputTypeOutput() OtherResourceOutputTypeOutput
	ToOtherResourceOutputTypeOutputWithContext(context.Context) OtherResourceOutputTypeOutput
}

type OtherResourceOutputTypeArgs struct {
	Foo pulumi.StringPtrInput `pulumi:"foo"`
}

func (OtherResourceOutputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OtherResourceOutputType)(nil)).Elem()
}

func (i OtherResourceOutputTypeArgs) ToOtherResourceOutputTypeOutput() OtherResourceOutputTypeOutput {
	return i.ToOtherResourceOutputTypeOutputWithContext(context.Background())
}

func (i OtherResourceOutputTypeArgs) ToOtherResourceOutputTypeOutputWithContext(ctx context.Context) OtherResourceOutputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OtherResourceOutputTypeOutput)
}

type OtherResourceOutputTypeOutput struct{ *pulumi.OutputState }

func (OtherResourceOutputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OtherResourceOutputType)(nil)).Elem()
}

func (o OtherResourceOutputTypeOutput) ToOtherResourceOutputTypeOutput() OtherResourceOutputTypeOutput {
	return o
}

func (o OtherResourceOutputTypeOutput) ToOtherResourceOutputTypeOutputWithContext(ctx context.Context) OtherResourceOutputTypeOutput {
	return o
}

func (o OtherResourceOutputTypeOutput) Foo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OtherResourceOutputType) *string { return v.Foo }).(pulumi.StringPtrOutput)
}

type SomeOtherObject struct {
	Baz *string `pulumi:"baz"`
}

// SomeOtherObjectInput is an input type that accepts SomeOtherObjectArgs and SomeOtherObjectOutput values.
// You can construct a concrete instance of `SomeOtherObjectInput` via:
//
//          SomeOtherObjectArgs{...}
type SomeOtherObjectInput interface {
	pulumi.Input

	ToSomeOtherObjectOutput() SomeOtherObjectOutput
	ToSomeOtherObjectOutputWithContext(context.Context) SomeOtherObjectOutput
}

type SomeOtherObjectArgs struct {
	Baz pulumi.StringPtrInput `pulumi:"baz"`
}

func (SomeOtherObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SomeOtherObject)(nil)).Elem()
}

func (i SomeOtherObjectArgs) ToSomeOtherObjectOutput() SomeOtherObjectOutput {
	return i.ToSomeOtherObjectOutputWithContext(context.Background())
}

func (i SomeOtherObjectArgs) ToSomeOtherObjectOutputWithContext(ctx context.Context) SomeOtherObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SomeOtherObjectOutput)
}

// SomeOtherObjectArrayInput is an input type that accepts SomeOtherObjectArray and SomeOtherObjectArrayOutput values.
// You can construct a concrete instance of `SomeOtherObjectArrayInput` via:
//
//          SomeOtherObjectArray{ SomeOtherObjectArgs{...} }
type SomeOtherObjectArrayInput interface {
	pulumi.Input

	ToSomeOtherObjectArrayOutput() SomeOtherObjectArrayOutput
	ToSomeOtherObjectArrayOutputWithContext(context.Context) SomeOtherObjectArrayOutput
}

type SomeOtherObjectArray []SomeOtherObjectInput

func (SomeOtherObjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SomeOtherObject)(nil)).Elem()
}

func (i SomeOtherObjectArray) ToSomeOtherObjectArrayOutput() SomeOtherObjectArrayOutput {
	return i.ToSomeOtherObjectArrayOutputWithContext(context.Background())
}

func (i SomeOtherObjectArray) ToSomeOtherObjectArrayOutputWithContext(ctx context.Context) SomeOtherObjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SomeOtherObjectArrayOutput)
}

type SomeOtherObjectOutput struct{ *pulumi.OutputState }

func (SomeOtherObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SomeOtherObject)(nil)).Elem()
}

func (o SomeOtherObjectOutput) ToSomeOtherObjectOutput() SomeOtherObjectOutput {
	return o
}

func (o SomeOtherObjectOutput) ToSomeOtherObjectOutputWithContext(ctx context.Context) SomeOtherObjectOutput {
	return o
}

func (o SomeOtherObjectOutput) Baz() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SomeOtherObject) *string { return v.Baz }).(pulumi.StringPtrOutput)
}

type SomeOtherObjectArrayOutput struct{ *pulumi.OutputState }

func (SomeOtherObjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SomeOtherObject)(nil)).Elem()
}

func (o SomeOtherObjectArrayOutput) ToSomeOtherObjectArrayOutput() SomeOtherObjectArrayOutput {
	return o
}

func (o SomeOtherObjectArrayOutput) ToSomeOtherObjectArrayOutputWithContext(ctx context.Context) SomeOtherObjectArrayOutput {
	return o
}

func (o SomeOtherObjectArrayOutput) Index(i pulumi.IntInput) SomeOtherObjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SomeOtherObject {
		return vs[0].([]SomeOtherObject)[vs[1].(int)]
	}).(SomeOtherObjectOutput)
}

type SomeOtherObjectArrayArray []SomeOtherObjectArrayInput

func (SomeOtherObjectArrayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[][]SomeOtherObject)(nil)).Elem()
}

func (i SomeOtherObjectArrayArray) ToSomeOtherObjectArrayArrayOutput() SomeOtherObjectArrayArrayOutput {
	return i.ToSomeOtherObjectArrayArrayOutputWithContext(context.Background())
}

func (i SomeOtherObjectArrayArray) ToSomeOtherObjectArrayArrayOutputWithContext(ctx context.Context) SomeOtherObjectArrayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SomeOtherObjectArrayArrayOutput)
}

// SomeOtherObjectArrayArrayInput is an input type that accepts SomeOtherObjectArrayArray and SomeOtherObjectArrayArrayOutput values.
// You can construct a concrete instance of `SomeOtherObjectArrayArrayInput` via:
//
//          SomeOtherObjectArrayArray{ SomeOtherObjectArray{ SomeOtherObjectArgs{...} } }
type SomeOtherObjectArrayArrayInput interface {
	pulumi.Input

	ToSomeOtherObjectArrayArrayOutput() SomeOtherObjectArrayArrayOutput
	ToSomeOtherObjectArrayArrayOutputWithContext(context.Context) SomeOtherObjectArrayArrayOutput
}

type SomeOtherObjectArrayArrayOutput struct{ *pulumi.OutputState }

func (SomeOtherObjectArrayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[][]SomeOtherObject)(nil)).Elem()
}

func (o SomeOtherObjectArrayArrayOutput) ToSomeOtherObjectArrayArrayOutput() SomeOtherObjectArrayArrayOutput {
	return o
}

func (o SomeOtherObjectArrayArrayOutput) ToSomeOtherObjectArrayArrayOutputWithContext(ctx context.Context) SomeOtherObjectArrayArrayOutput {
	return o
}

func (o SomeOtherObjectArrayArrayOutput) Index(i pulumi.IntInput) SomeOtherObjectArrayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) []SomeOtherObject {
		return vs[0].([][]SomeOtherObject)[vs[1].(int)]
	}).(SomeOtherObjectArrayOutput)
}

type SomeOtherObjectArrayMap map[string]SomeOtherObjectArrayInput

func (SomeOtherObjectArrayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string][]SomeOtherObject)(nil)).Elem()
}

func (i SomeOtherObjectArrayMap) ToSomeOtherObjectArrayMapOutput() SomeOtherObjectArrayMapOutput {
	return i.ToSomeOtherObjectArrayMapOutputWithContext(context.Background())
}

func (i SomeOtherObjectArrayMap) ToSomeOtherObjectArrayMapOutputWithContext(ctx context.Context) SomeOtherObjectArrayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SomeOtherObjectArrayMapOutput)
}

// SomeOtherObjectArrayMapInput is an input type that accepts SomeOtherObjectArrayMap and SomeOtherObjectArrayMapOutput values.
// You can construct a concrete instance of `SomeOtherObjectArrayMapInput` via:
//
//          SomeOtherObjectArrayMap{ "key": SomeOtherObjectArray{ SomeOtherObjectArgs{...} } }
type SomeOtherObjectArrayMapInput interface {
	pulumi.Input

	ToSomeOtherObjectArrayMapOutput() SomeOtherObjectArrayMapOutput
	ToSomeOtherObjectArrayMapOutputWithContext(context.Context) SomeOtherObjectArrayMapOutput
}

type SomeOtherObjectArrayMapOutput struct{ *pulumi.OutputState }

func (SomeOtherObjectArrayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string][]SomeOtherObject)(nil)).Elem()
}

func (o SomeOtherObjectArrayMapOutput) ToSomeOtherObjectArrayMapOutput() SomeOtherObjectArrayMapOutput {
	return o
}

func (o SomeOtherObjectArrayMapOutput) ToSomeOtherObjectArrayMapOutputWithContext(ctx context.Context) SomeOtherObjectArrayMapOutput {
	return o
}

func (o SomeOtherObjectArrayMapOutput) MapIndex(k pulumi.StringInput) SomeOtherObjectArrayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) []SomeOtherObject {
		return vs[0].(map[string][]SomeOtherObject)[vs[1].(string)]
	}).(SomeOtherObjectArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigMapInput)(nil)).Elem(), ConfigMapArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigMapArrayInput)(nil)).Elem(), ConfigMapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectInput)(nil)).Elem(), ObjectArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OtherResourceOutputTypeInput)(nil)).Elem(), OtherResourceOutputTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SomeOtherObjectInput)(nil)).Elem(), SomeOtherObjectArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SomeOtherObjectArrayInput)(nil)).Elem(), SomeOtherObjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SomeOtherObjectArrayArrayInput)(nil)).Elem(), SomeOtherObjectArrayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SomeOtherObjectArrayMapInput)(nil)).Elem(), SomeOtherObjectArrayMap{})
	pulumi.RegisterOutputType(ConfigMapOutput{})
	pulumi.RegisterOutputType(ConfigMapArrayOutput{})
	pulumi.RegisterOutputType(ObjectOutput{})
	pulumi.RegisterOutputType(OtherResourceOutputTypeOutput{})
	pulumi.RegisterOutputType(SomeOtherObjectOutput{})
	pulumi.RegisterOutputType(SomeOtherObjectArrayOutput{})
	pulumi.RegisterOutputType(SomeOtherObjectArrayArrayOutput{})
	pulumi.RegisterOutputType(SomeOtherObjectArrayMapOutput{})
}
