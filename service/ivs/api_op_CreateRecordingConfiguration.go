// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ivs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new recording configuration, used to enable recording to Amazon S3.
// Known issue: In the us-east-1 region, if you use the Amazon Web Services CLI to
// create a recording configuration, it returns success even if the S3 bucket is in
// a different region. In this case, the state of the recording configuration is
// CREATE_FAILED (instead of ACTIVE ). (In other regions, the CLI correctly returns
// failure if the bucket is in a different region.) Workaround: Ensure that your S3
// bucket is in the same region as the recording configuration. If you create a
// recording configuration in a different region as your S3 bucket, delete that
// recording configuration and create a new one with an S3 bucket from the correct
// region.
func (c *Client) CreateRecordingConfiguration(ctx context.Context, params *CreateRecordingConfigurationInput, optFns ...func(*Options)) (*CreateRecordingConfigurationOutput, error) {
	if params == nil {
		params = &CreateRecordingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRecordingConfiguration", params, optFns, c.addOperationCreateRecordingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRecordingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRecordingConfigurationInput struct {

	// A complex type that contains a destination configuration for where recorded
	// video will be stored.
	//
	// This member is required.
	DestinationConfiguration *types.DestinationConfiguration

	// Recording-configuration name. The value does not need to be unique.
	Name *string

	// If a broadcast disconnects and then reconnects within the specified interval,
	// the multiple streams will be considered a single broadcast and merged together.
	// Default: 0.
	RecordingReconnectWindowSeconds int32

	// Object that describes which renditions should be recorded for a stream.
	RenditionConfiguration *types.RenditionConfiguration

	// Array of 1-50 maps, each of the form string:string (key:value) . See Tagging
	// Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// for more information, including restrictions that apply to tags and "Tag naming
	// limits and requirements"; Amazon IVS has no service-specific constraints beyond
	// what is documented there.
	Tags map[string]string

	// A complex type that allows you to enable/disable the recording of thumbnails
	// for a live session and modify the interval at which thumbnails are generated for
	// the live session.
	ThumbnailConfiguration *types.ThumbnailConfiguration

	noSmithyDocumentSerde
}

type CreateRecordingConfigurationOutput struct {

	//
	RecordingConfiguration *types.RecordingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRecordingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRecordingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRecordingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRecordingConfiguration"); err != nil {
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
	if err = addOpCreateRecordingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRecordingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRecordingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRecordingConfiguration",
	}
}
