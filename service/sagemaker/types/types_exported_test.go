// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
)

func ExampleAutoMLProblemTypeConfig_outputUsage() {
	var union types.AutoMLProblemTypeConfig
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AutoMLProblemTypeConfigMemberImageClassificationJobConfig:
		_ = v.Value // Value is types.ImageClassificationJobConfig

	case *types.AutoMLProblemTypeConfigMemberTabularJobConfig:
		_ = v.Value // Value is types.TabularJobConfig

	case *types.AutoMLProblemTypeConfigMemberTextClassificationJobConfig:
		_ = v.Value // Value is types.TextClassificationJobConfig

	case *types.AutoMLProblemTypeConfigMemberTextGenerationJobConfig:
		_ = v.Value // Value is types.TextGenerationJobConfig

	case *types.AutoMLProblemTypeConfigMemberTimeSeriesForecastingJobConfig:
		_ = v.Value // Value is types.TimeSeriesForecastingJobConfig

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.TextGenerationJobConfig
var _ *types.TextClassificationJobConfig
var _ *types.TimeSeriesForecastingJobConfig
var _ *types.ImageClassificationJobConfig
var _ *types.TabularJobConfig

func ExampleAutoMLProblemTypeResolvedAttributes_outputUsage() {
	var union types.AutoMLProblemTypeResolvedAttributes
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AutoMLProblemTypeResolvedAttributesMemberTabularResolvedAttributes:
		_ = v.Value // Value is types.TabularResolvedAttributes

	case *types.AutoMLProblemTypeResolvedAttributesMemberTextGenerationResolvedAttributes:
		_ = v.Value // Value is types.TextGenerationResolvedAttributes

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.TextGenerationResolvedAttributes
var _ *types.TabularResolvedAttributes

func ExampleCollectionConfig_outputUsage() {
	var union types.CollectionConfig
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.CollectionConfigMemberVectorConfig:
		_ = v.Value // Value is types.VectorConfig

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.VectorConfig

func ExampleCustomFileSystem_outputUsage() {
	var union types.CustomFileSystem
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.CustomFileSystemMemberEFSFileSystem:
		_ = v.Value // Value is types.EFSFileSystem

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.EFSFileSystem

func ExampleCustomFileSystemConfig_outputUsage() {
	var union types.CustomFileSystemConfig
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.CustomFileSystemConfigMemberEFSFileSystemConfig:
		_ = v.Value // Value is types.EFSFileSystemConfig

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.EFSFileSystemConfig

func ExampleMetricSpecification_outputUsage() {
	var union types.MetricSpecification
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.MetricSpecificationMemberCustomized:
		_ = v.Value // Value is types.CustomizedMetricSpecification

	case *types.MetricSpecificationMemberPredefined:
		_ = v.Value // Value is types.PredefinedMetricSpecification

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.PredefinedMetricSpecification
var _ *types.CustomizedMetricSpecification

func ExampleScalingPolicy_outputUsage() {
	var union types.ScalingPolicy
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ScalingPolicyMemberTargetTracking:
		_ = v.Value // Value is types.TargetTrackingScalingPolicyConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.TargetTrackingScalingPolicyConfiguration

func ExampleTrialComponentParameterValue_outputUsage() {
	var union types.TrialComponentParameterValue
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.TrialComponentParameterValueMemberNumberValue:
		_ = v.Value // Value is float64

	case *types.TrialComponentParameterValueMemberStringValue:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *float64
