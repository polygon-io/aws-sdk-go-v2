// Code generated by smithy-go-codegen DO NOT EDIT.

package pipes

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/pipes/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreatePipe struct {
}

func (*validateOpCreatePipe) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreatePipe) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreatePipeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreatePipeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeletePipe struct {
}

func (*validateOpDeletePipe) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeletePipe) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeletePipeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeletePipeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribePipe struct {
}

func (*validateOpDescribePipe) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribePipe) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribePipeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribePipeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartPipe struct {
}

func (*validateOpStartPipe) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartPipe) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartPipeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartPipeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopPipe struct {
}

func (*validateOpStopPipe) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopPipe) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopPipeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopPipeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdatePipe struct {
}

func (*validateOpUpdatePipe) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdatePipe) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdatePipeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdatePipeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreatePipeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreatePipe{}, middleware.After)
}

func addOpDeletePipeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeletePipe{}, middleware.After)
}

func addOpDescribePipeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribePipe{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpStartPipeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartPipe{}, middleware.After)
}

func addOpStopPipeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopPipe{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdatePipeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdatePipe{}, middleware.After)
}

func validateAwsVpcConfiguration(v *types.AwsVpcConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AwsVpcConfiguration"}
	if v.Subnets == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Subnets"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateBatchContainerOverrides(v *types.BatchContainerOverrides) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchContainerOverrides"}
	if v.ResourceRequirements != nil {
		if err := validateBatchResourceRequirementsList(v.ResourceRequirements); err != nil {
			invalidParams.AddNested("ResourceRequirements", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateBatchResourceRequirement(v *types.BatchResourceRequirement) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchResourceRequirement"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateBatchResourceRequirementsList(v []types.BatchResourceRequirement) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchResourceRequirementsList"}
	for i := range v {
		if err := validateBatchResourceRequirement(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCapacityProviderStrategy(v []types.CapacityProviderStrategyItem) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CapacityProviderStrategy"}
	for i := range v {
		if err := validateCapacityProviderStrategyItem(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCapacityProviderStrategyItem(v *types.CapacityProviderStrategyItem) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CapacityProviderStrategyItem"}
	if v.CapacityProvider == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CapacityProvider"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCloudwatchLogsLogDestinationParameters(v *types.CloudwatchLogsLogDestinationParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CloudwatchLogsLogDestinationParameters"}
	if v.LogGroupArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LogGroupArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEcsContainerOverride(v *types.EcsContainerOverride) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EcsContainerOverride"}
	if v.EnvironmentFiles != nil {
		if err := validateEcsEnvironmentFileList(v.EnvironmentFiles); err != nil {
			invalidParams.AddNested("EnvironmentFiles", err.(smithy.InvalidParamsError))
		}
	}
	if v.ResourceRequirements != nil {
		if err := validateEcsResourceRequirementsList(v.ResourceRequirements); err != nil {
			invalidParams.AddNested("ResourceRequirements", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEcsContainerOverrideList(v []types.EcsContainerOverride) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EcsContainerOverrideList"}
	for i := range v {
		if err := validateEcsContainerOverride(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEcsEnvironmentFile(v *types.EcsEnvironmentFile) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EcsEnvironmentFile"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEcsEnvironmentFileList(v []types.EcsEnvironmentFile) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EcsEnvironmentFileList"}
	for i := range v {
		if err := validateEcsEnvironmentFile(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEcsEphemeralStorage(v *types.EcsEphemeralStorage) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EcsEphemeralStorage"}
	if v.SizeInGiB == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SizeInGiB"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEcsResourceRequirement(v *types.EcsResourceRequirement) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EcsResourceRequirement"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEcsResourceRequirementsList(v []types.EcsResourceRequirement) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EcsResourceRequirementsList"}
	for i := range v {
		if err := validateEcsResourceRequirement(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEcsTaskOverride(v *types.EcsTaskOverride) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EcsTaskOverride"}
	if v.ContainerOverrides != nil {
		if err := validateEcsContainerOverrideList(v.ContainerOverrides); err != nil {
			invalidParams.AddNested("ContainerOverrides", err.(smithy.InvalidParamsError))
		}
	}
	if v.EphemeralStorage != nil {
		if err := validateEcsEphemeralStorage(v.EphemeralStorage); err != nil {
			invalidParams.AddNested("EphemeralStorage", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFirehoseLogDestinationParameters(v *types.FirehoseLogDestinationParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FirehoseLogDestinationParameters"}
	if v.DeliveryStreamArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeliveryStreamArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNetworkConfiguration(v *types.NetworkConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NetworkConfiguration"}
	if v.AwsvpcConfiguration != nil {
		if err := validateAwsVpcConfiguration(v.AwsvpcConfiguration); err != nil {
			invalidParams.AddNested("AwsvpcConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePipeLogConfigurationParameters(v *types.PipeLogConfigurationParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PipeLogConfigurationParameters"}
	if v.S3LogDestination != nil {
		if err := validateS3LogDestinationParameters(v.S3LogDestination); err != nil {
			invalidParams.AddNested("S3LogDestination", err.(smithy.InvalidParamsError))
		}
	}
	if v.FirehoseLogDestination != nil {
		if err := validateFirehoseLogDestinationParameters(v.FirehoseLogDestination); err != nil {
			invalidParams.AddNested("FirehoseLogDestination", err.(smithy.InvalidParamsError))
		}
	}
	if v.CloudwatchLogsLogDestination != nil {
		if err := validateCloudwatchLogsLogDestinationParameters(v.CloudwatchLogsLogDestination); err != nil {
			invalidParams.AddNested("CloudwatchLogsLogDestination", err.(smithy.InvalidParamsError))
		}
	}
	if len(v.Level) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Level"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePipeSourceActiveMQBrokerParameters(v *types.PipeSourceActiveMQBrokerParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PipeSourceActiveMQBrokerParameters"}
	if v.Credentials == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Credentials"))
	}
	if v.QueueName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePipeSourceDynamoDBStreamParameters(v *types.PipeSourceDynamoDBStreamParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PipeSourceDynamoDBStreamParameters"}
	if len(v.StartingPosition) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("StartingPosition"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePipeSourceKinesisStreamParameters(v *types.PipeSourceKinesisStreamParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PipeSourceKinesisStreamParameters"}
	if len(v.StartingPosition) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("StartingPosition"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePipeSourceManagedStreamingKafkaParameters(v *types.PipeSourceManagedStreamingKafkaParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PipeSourceManagedStreamingKafkaParameters"}
	if v.TopicName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TopicName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePipeSourceParameters(v *types.PipeSourceParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PipeSourceParameters"}
	if v.KinesisStreamParameters != nil {
		if err := validatePipeSourceKinesisStreamParameters(v.KinesisStreamParameters); err != nil {
			invalidParams.AddNested("KinesisStreamParameters", err.(smithy.InvalidParamsError))
		}
	}
	if v.DynamoDBStreamParameters != nil {
		if err := validatePipeSourceDynamoDBStreamParameters(v.DynamoDBStreamParameters); err != nil {
			invalidParams.AddNested("DynamoDBStreamParameters", err.(smithy.InvalidParamsError))
		}
	}
	if v.ActiveMQBrokerParameters != nil {
		if err := validatePipeSourceActiveMQBrokerParameters(v.ActiveMQBrokerParameters); err != nil {
			invalidParams.AddNested("ActiveMQBrokerParameters", err.(smithy.InvalidParamsError))
		}
	}
	if v.RabbitMQBrokerParameters != nil {
		if err := validatePipeSourceRabbitMQBrokerParameters(v.RabbitMQBrokerParameters); err != nil {
			invalidParams.AddNested("RabbitMQBrokerParameters", err.(smithy.InvalidParamsError))
		}
	}
	if v.ManagedStreamingKafkaParameters != nil {
		if err := validatePipeSourceManagedStreamingKafkaParameters(v.ManagedStreamingKafkaParameters); err != nil {
			invalidParams.AddNested("ManagedStreamingKafkaParameters", err.(smithy.InvalidParamsError))
		}
	}
	if v.SelfManagedKafkaParameters != nil {
		if err := validatePipeSourceSelfManagedKafkaParameters(v.SelfManagedKafkaParameters); err != nil {
			invalidParams.AddNested("SelfManagedKafkaParameters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePipeSourceRabbitMQBrokerParameters(v *types.PipeSourceRabbitMQBrokerParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PipeSourceRabbitMQBrokerParameters"}
	if v.Credentials == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Credentials"))
	}
	if v.QueueName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePipeSourceSelfManagedKafkaParameters(v *types.PipeSourceSelfManagedKafkaParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PipeSourceSelfManagedKafkaParameters"}
	if v.TopicName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TopicName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePipeTargetBatchJobParameters(v *types.PipeTargetBatchJobParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PipeTargetBatchJobParameters"}
	if v.JobDefinition == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobDefinition"))
	}
	if v.JobName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobName"))
	}
	if v.ContainerOverrides != nil {
		if err := validateBatchContainerOverrides(v.ContainerOverrides); err != nil {
			invalidParams.AddNested("ContainerOverrides", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePipeTargetEcsTaskParameters(v *types.PipeTargetEcsTaskParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PipeTargetEcsTaskParameters"}
	if v.TaskDefinitionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskDefinitionArn"))
	}
	if v.NetworkConfiguration != nil {
		if err := validateNetworkConfiguration(v.NetworkConfiguration); err != nil {
			invalidParams.AddNested("NetworkConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.CapacityProviderStrategy != nil {
		if err := validateCapacityProviderStrategy(v.CapacityProviderStrategy); err != nil {
			invalidParams.AddNested("CapacityProviderStrategy", err.(smithy.InvalidParamsError))
		}
	}
	if v.Overrides != nil {
		if err := validateEcsTaskOverride(v.Overrides); err != nil {
			invalidParams.AddNested("Overrides", err.(smithy.InvalidParamsError))
		}
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePipeTargetKinesisStreamParameters(v *types.PipeTargetKinesisStreamParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PipeTargetKinesisStreamParameters"}
	if v.PartitionKey == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PartitionKey"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePipeTargetParameters(v *types.PipeTargetParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PipeTargetParameters"}
	if v.KinesisStreamParameters != nil {
		if err := validatePipeTargetKinesisStreamParameters(v.KinesisStreamParameters); err != nil {
			invalidParams.AddNested("KinesisStreamParameters", err.(smithy.InvalidParamsError))
		}
	}
	if v.EcsTaskParameters != nil {
		if err := validatePipeTargetEcsTaskParameters(v.EcsTaskParameters); err != nil {
			invalidParams.AddNested("EcsTaskParameters", err.(smithy.InvalidParamsError))
		}
	}
	if v.BatchJobParameters != nil {
		if err := validatePipeTargetBatchJobParameters(v.BatchJobParameters); err != nil {
			invalidParams.AddNested("BatchJobParameters", err.(smithy.InvalidParamsError))
		}
	}
	if v.RedshiftDataParameters != nil {
		if err := validatePipeTargetRedshiftDataParameters(v.RedshiftDataParameters); err != nil {
			invalidParams.AddNested("RedshiftDataParameters", err.(smithy.InvalidParamsError))
		}
	}
	if v.SageMakerPipelineParameters != nil {
		if err := validatePipeTargetSageMakerPipelineParameters(v.SageMakerPipelineParameters); err != nil {
			invalidParams.AddNested("SageMakerPipelineParameters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePipeTargetRedshiftDataParameters(v *types.PipeTargetRedshiftDataParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PipeTargetRedshiftDataParameters"}
	if v.Database == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Database"))
	}
	if v.Sqls == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Sqls"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePipeTargetSageMakerPipelineParameters(v *types.PipeTargetSageMakerPipelineParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PipeTargetSageMakerPipelineParameters"}
	if v.PipelineParameterList != nil {
		if err := validateSageMakerPipelineParameterList(v.PipelineParameterList); err != nil {
			invalidParams.AddNested("PipelineParameterList", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3LogDestinationParameters(v *types.S3LogDestinationParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3LogDestinationParameters"}
	if v.BucketName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BucketName"))
	}
	if v.BucketOwner == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BucketOwner"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSageMakerPipelineParameter(v *types.SageMakerPipelineParameter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SageMakerPipelineParameter"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSageMakerPipelineParameterList(v []types.SageMakerPipelineParameter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SageMakerPipelineParameterList"}
	for i := range v {
		if err := validateSageMakerPipelineParameter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdatePipeSourceActiveMQBrokerParameters(v *types.UpdatePipeSourceActiveMQBrokerParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdatePipeSourceActiveMQBrokerParameters"}
	if v.Credentials == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Credentials"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdatePipeSourceParameters(v *types.UpdatePipeSourceParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdatePipeSourceParameters"}
	if v.ActiveMQBrokerParameters != nil {
		if err := validateUpdatePipeSourceActiveMQBrokerParameters(v.ActiveMQBrokerParameters); err != nil {
			invalidParams.AddNested("ActiveMQBrokerParameters", err.(smithy.InvalidParamsError))
		}
	}
	if v.RabbitMQBrokerParameters != nil {
		if err := validateUpdatePipeSourceRabbitMQBrokerParameters(v.RabbitMQBrokerParameters); err != nil {
			invalidParams.AddNested("RabbitMQBrokerParameters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdatePipeSourceRabbitMQBrokerParameters(v *types.UpdatePipeSourceRabbitMQBrokerParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdatePipeSourceRabbitMQBrokerParameters"}
	if v.Credentials == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Credentials"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreatePipeInput(v *CreatePipeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreatePipeInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Source == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Source"))
	}
	if v.SourceParameters != nil {
		if err := validatePipeSourceParameters(v.SourceParameters); err != nil {
			invalidParams.AddNested("SourceParameters", err.(smithy.InvalidParamsError))
		}
	}
	if v.Target == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Target"))
	}
	if v.TargetParameters != nil {
		if err := validatePipeTargetParameters(v.TargetParameters); err != nil {
			invalidParams.AddNested("TargetParameters", err.(smithy.InvalidParamsError))
		}
	}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if v.LogConfiguration != nil {
		if err := validatePipeLogConfigurationParameters(v.LogConfiguration); err != nil {
			invalidParams.AddNested("LogConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeletePipeInput(v *DeletePipeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeletePipeInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribePipeInput(v *DescribePipeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribePipeInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartPipeInput(v *StartPipeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartPipeInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopPipeInput(v *StopPipeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopPipeInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdatePipeInput(v *UpdatePipeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdatePipeInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.SourceParameters != nil {
		if err := validateUpdatePipeSourceParameters(v.SourceParameters); err != nil {
			invalidParams.AddNested("SourceParameters", err.(smithy.InvalidParamsError))
		}
	}
	if v.TargetParameters != nil {
		if err := validatePipeTargetParameters(v.TargetParameters); err != nil {
			invalidParams.AddNested("TargetParameters", err.(smithy.InvalidParamsError))
		}
	}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if v.LogConfiguration != nil {
		if err := validatePipeLogConfigurationParameters(v.LogConfiguration); err != nil {
			invalidParams.AddNested("LogConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
