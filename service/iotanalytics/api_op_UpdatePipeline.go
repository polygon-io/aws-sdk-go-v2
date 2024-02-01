// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the settings of a pipeline. You must specify both a channel and a
// datastore activity and, optionally, as many as 23 additional activities in the
// pipelineActivities array.
func (c *Client) UpdatePipeline(ctx context.Context, params *UpdatePipelineInput, optFns ...func(*Options)) (*UpdatePipelineOutput, error) {
	if params == nil {
		params = &UpdatePipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePipeline", params, optFns, c.addOperationUpdatePipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePipelineInput struct {

	// A list of PipelineActivity objects. Activities perform transformations on your
	// messages, such as removing, renaming or adding message attributes; filtering
	// messages based on attribute values; invoking your Lambda functions on messages
	// for advanced processing; or performing mathematical transformations to normalize
	// device data. The list can be 2-25 PipelineActivity objects and must contain
	// both a channel and a datastore activity. Each entry in the list must contain
	// only one activity. For example: pipelineActivities = [ { "channel": { ... } },
	// { "lambda": { ... } }, ... ]
	//
	// This member is required.
	PipelineActivities []types.PipelineActivity

	// The name of the pipeline to update.
	//
	// This member is required.
	PipelineName *string

	noSmithyDocumentSerde
}

type UpdatePipelineOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePipeline"); err != nil {
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
	if err = addOpUpdatePipelineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePipeline(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePipeline",
	}
}
