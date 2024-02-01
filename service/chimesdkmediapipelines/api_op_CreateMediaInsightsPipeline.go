// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmediapipelines

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a media insights pipeline.
func (c *Client) CreateMediaInsightsPipeline(ctx context.Context, params *CreateMediaInsightsPipelineInput, optFns ...func(*Options)) (*CreateMediaInsightsPipelineOutput, error) {
	if params == nil {
		params = &CreateMediaInsightsPipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMediaInsightsPipeline", params, optFns, c.addOperationCreateMediaInsightsPipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMediaInsightsPipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMediaInsightsPipelineInput struct {

	// The ARN of the pipeline's configuration.
	//
	// This member is required.
	MediaInsightsPipelineConfigurationArn *string

	// The unique identifier for the media insights pipeline request.
	ClientRequestToken *string

	// The runtime configuration for the Kinesis video recording stream source.
	KinesisVideoStreamRecordingSourceRuntimeConfiguration *types.KinesisVideoStreamRecordingSourceRuntimeConfiguration

	// The runtime configuration for the Kinesis video stream source of the media
	// insights pipeline.
	KinesisVideoStreamSourceRuntimeConfiguration *types.KinesisVideoStreamSourceRuntimeConfiguration

	// The runtime metadata for the media insights pipeline. Consists of a key-value
	// map of strings.
	MediaInsightsRuntimeMetadata map[string]string

	// The runtime configuration for the S3 recording sink. If specified, the settings
	// in this structure override any settings in S3RecordingSinkConfiguration .
	S3RecordingSinkRuntimeConfiguration *types.S3RecordingSinkRuntimeConfiguration

	// The tags assigned to the media insights pipeline.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateMediaInsightsPipelineOutput struct {

	// The media insights pipeline object.
	//
	// This member is required.
	MediaInsightsPipeline *types.MediaInsightsPipeline

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMediaInsightsPipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMediaInsightsPipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMediaInsightsPipeline{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateMediaInsightsPipeline"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateMediaInsightsPipelineMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateMediaInsightsPipelineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMediaInsightsPipeline(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpCreateMediaInsightsPipeline struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateMediaInsightsPipeline) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateMediaInsightsPipeline) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateMediaInsightsPipelineInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateMediaInsightsPipelineInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateMediaInsightsPipelineMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateMediaInsightsPipeline{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateMediaInsightsPipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateMediaInsightsPipeline",
	}
}
