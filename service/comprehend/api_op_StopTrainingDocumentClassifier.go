// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops a document classifier training job while in progress. If the training job
// state is TRAINING , the job is marked for termination and put into the
// STOP_REQUESTED state. If the training job completes before it can be stopped, it
// is put into the TRAINED ; otherwise the training job is stopped and put into the
// STOPPED state and the service sends back an HTTP 200 response with an empty HTTP
// body.
func (c *Client) StopTrainingDocumentClassifier(ctx context.Context, params *StopTrainingDocumentClassifierInput, optFns ...func(*Options)) (*StopTrainingDocumentClassifierOutput, error) {
	if params == nil {
		params = &StopTrainingDocumentClassifierInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopTrainingDocumentClassifier", params, optFns, c.addOperationStopTrainingDocumentClassifierMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopTrainingDocumentClassifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopTrainingDocumentClassifierInput struct {

	// The Amazon Resource Name (ARN) that identifies the document classifier
	// currently being trained.
	//
	// This member is required.
	DocumentClassifierArn *string

	noSmithyDocumentSerde
}

type StopTrainingDocumentClassifierOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopTrainingDocumentClassifierMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopTrainingDocumentClassifier{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopTrainingDocumentClassifier{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StopTrainingDocumentClassifier"); err != nil {
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
	if err = addOpStopTrainingDocumentClassifierValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopTrainingDocumentClassifier(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopTrainingDocumentClassifier(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StopTrainingDocumentClassifier",
	}
}
