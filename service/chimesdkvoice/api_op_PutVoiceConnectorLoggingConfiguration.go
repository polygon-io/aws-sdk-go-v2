// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a Voice Connector's logging configuration.
func (c *Client) PutVoiceConnectorLoggingConfiguration(ctx context.Context, params *PutVoiceConnectorLoggingConfigurationInput, optFns ...func(*Options)) (*PutVoiceConnectorLoggingConfigurationOutput, error) {
	if params == nil {
		params = &PutVoiceConnectorLoggingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutVoiceConnectorLoggingConfiguration", params, optFns, c.addOperationPutVoiceConnectorLoggingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutVoiceConnectorLoggingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutVoiceConnectorLoggingConfigurationInput struct {

	// The logging configuration being updated.
	//
	// This member is required.
	LoggingConfiguration *types.LoggingConfiguration

	// The Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	noSmithyDocumentSerde
}

type PutVoiceConnectorLoggingConfigurationOutput struct {

	// The updated logging configuration.
	LoggingConfiguration *types.LoggingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutVoiceConnectorLoggingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutVoiceConnectorLoggingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutVoiceConnectorLoggingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutVoiceConnectorLoggingConfiguration"); err != nil {
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
	if err = addOpPutVoiceConnectorLoggingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutVoiceConnectorLoggingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutVoiceConnectorLoggingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutVoiceConnectorLoggingConfiguration",
	}
}
