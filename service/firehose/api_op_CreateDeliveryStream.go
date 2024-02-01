// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Kinesis Data Firehose delivery stream. By default, you can create up
// to 50 delivery streams per Amazon Web Services Region. This is an asynchronous
// operation that immediately returns. The initial status of the delivery stream is
// CREATING . After the delivery stream is created, its status is ACTIVE and it
// now accepts data. If the delivery stream creation fails, the status transitions
// to CREATING_FAILED . Attempts to send data to a delivery stream that is not in
// the ACTIVE state cause an exception. To check the state of a delivery stream,
// use DescribeDeliveryStream . If the status of a delivery stream is
// CREATING_FAILED , this status doesn't change, and you can't invoke
// CreateDeliveryStream again on it. However, you can invoke the
// DeleteDeliveryStream operation to delete it. A Kinesis Data Firehose delivery
// stream can be configured to receive records directly from providers using
// PutRecord or PutRecordBatch , or it can be configured to use an existing Kinesis
// stream as its source. To specify a Kinesis data stream as input, set the
// DeliveryStreamType parameter to KinesisStreamAsSource , and provide the Kinesis
// stream Amazon Resource Name (ARN) and role ARN in the
// KinesisStreamSourceConfiguration parameter. To create a delivery stream with
// server-side encryption (SSE) enabled, include
// DeliveryStreamEncryptionConfigurationInput in your request. This is optional.
// You can also invoke StartDeliveryStreamEncryption to turn on SSE for an
// existing delivery stream that doesn't have SSE enabled. A delivery stream is
// configured with a single destination, such as Amazon Simple Storage Service
// (Amazon S3), Amazon Redshift, Amazon OpenSearch Service, Amazon OpenSearch
// Serverless, Splunk, and any custom HTTP endpoint or HTTP endpoints owned by or
// supported by third-party service providers, including Datadog, Dynatrace,
// LogicMonitor, MongoDB, New Relic, and Sumo Logic. You must specify only one of
// the following destination configuration parameters:
// ExtendedS3DestinationConfiguration , S3DestinationConfiguration ,
// ElasticsearchDestinationConfiguration , RedshiftDestinationConfiguration , or
// SplunkDestinationConfiguration . When you specify S3DestinationConfiguration ,
// you can also provide the following optional values: BufferingHints,
// EncryptionConfiguration , and CompressionFormat . By default, if no
// BufferingHints value is provided, Kinesis Data Firehose buffers data up to 5 MB
// or for 5 minutes, whichever condition is satisfied first. BufferingHints is a
// hint, so there are some cases where the service cannot adhere to these
// conditions strictly. For example, record boundaries might be such that the size
// is a little over or under the configured buffering size. By default, no
// encryption is performed. We strongly recommend that you enable encryption to
// ensure secure data storage in Amazon S3. A few notes about Amazon Redshift as a
// destination:
//   - An Amazon Redshift destination requires an S3 bucket as intermediate
//     location. Kinesis Data Firehose first delivers data to Amazon S3 and then uses
//     COPY syntax to load data into an Amazon Redshift table. This is specified in
//     the RedshiftDestinationConfiguration.S3Configuration parameter.
//   - The compression formats SNAPPY or ZIP cannot be specified in
//     RedshiftDestinationConfiguration.S3Configuration because the Amazon Redshift
//     COPY operation that reads from the S3 bucket doesn't support these compression
//     formats.
//   - We strongly recommend that you use the user name and password you provide
//     exclusively with Kinesis Data Firehose, and that the permissions for the account
//     are restricted for Amazon Redshift INSERT permissions.
//
// Kinesis Data Firehose assumes the IAM role that is configured as part of the
// destination. The role should allow the Kinesis Data Firehose principal to assume
// the role, and the role should have permissions that allow the service to deliver
// the data. For more information, see Grant Kinesis Data Firehose Access to an
// Amazon S3 Destination (https://docs.aws.amazon.com/firehose/latest/dev/controlling-access.html#using-iam-s3)
// in the Amazon Kinesis Data Firehose Developer Guide.
func (c *Client) CreateDeliveryStream(ctx context.Context, params *CreateDeliveryStreamInput, optFns ...func(*Options)) (*CreateDeliveryStreamOutput, error) {
	if params == nil {
		params = &CreateDeliveryStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDeliveryStream", params, optFns, c.addOperationCreateDeliveryStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDeliveryStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDeliveryStreamInput struct {

	// The name of the delivery stream. This name must be unique per Amazon Web
	// Services account in the same Amazon Web Services Region. If the delivery streams
	// are in different accounts or different Regions, you can have multiple delivery
	// streams with the same name.
	//
	// This member is required.
	DeliveryStreamName *string

	// The destination in the Serverless offering for Amazon OpenSearch Service. You
	// can specify only one destination.
	AmazonOpenSearchServerlessDestinationConfiguration *types.AmazonOpenSearchServerlessDestinationConfiguration

	// The destination in Amazon OpenSearch Service. You can specify only one
	// destination.
	AmazonopensearchserviceDestinationConfiguration *types.AmazonopensearchserviceDestinationConfiguration

	// Used to specify the type and Amazon Resource Name (ARN) of the KMS key needed
	// for Server-Side Encryption (SSE).
	DeliveryStreamEncryptionConfigurationInput *types.DeliveryStreamEncryptionConfigurationInput

	// The delivery stream type. This parameter can be one of the following values:
	//   - DirectPut : Provider applications access the delivery stream directly.
	//   - KinesisStreamAsSource : The delivery stream uses a Kinesis data stream as a
	//   source.
	DeliveryStreamType types.DeliveryStreamType

	// The destination in Amazon ES. You can specify only one destination.
	ElasticsearchDestinationConfiguration *types.ElasticsearchDestinationConfiguration

	// The destination in Amazon S3. You can specify only one destination.
	ExtendedS3DestinationConfiguration *types.ExtendedS3DestinationConfiguration

	// Enables configuring Kinesis Firehose to deliver data to any HTTP endpoint
	// destination. You can specify only one destination.
	HttpEndpointDestinationConfiguration *types.HttpEndpointDestinationConfiguration

	// When a Kinesis data stream is used as the source for the delivery stream, a
	// KinesisStreamSourceConfiguration containing the Kinesis data stream Amazon
	// Resource Name (ARN) and the role ARN for the source stream.
	KinesisStreamSourceConfiguration *types.KinesisStreamSourceConfiguration

	// The configuration for the Amazon MSK cluster to be used as the source for a
	// delivery stream.
	MSKSourceConfiguration *types.MSKSourceConfiguration

	// The destination in Amazon Redshift. You can specify only one destination.
	RedshiftDestinationConfiguration *types.RedshiftDestinationConfiguration

	// [Deprecated] The destination in Amazon S3. You can specify only one destination.
	//
	// Deprecated: This member has been deprecated.
	S3DestinationConfiguration *types.S3DestinationConfiguration

	// Configure Snowflake destination
	SnowflakeDestinationConfiguration *types.SnowflakeDestinationConfiguration

	// The destination in Splunk. You can specify only one destination.
	SplunkDestinationConfiguration *types.SplunkDestinationConfiguration

	// A set of tags to assign to the delivery stream. A tag is a key-value pair that
	// you can define and assign to Amazon Web Services resources. Tags are metadata.
	// For example, you can add friendly names and descriptions or other types of
	// information that can help you distinguish the delivery stream. For more
	// information about tags, see Using Cost Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
	// in the Amazon Web Services Billing and Cost Management User Guide. You can
	// specify up to 50 tags when creating a delivery stream.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDeliveryStreamOutput struct {

	// The ARN of the delivery stream.
	DeliveryStreamARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDeliveryStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDeliveryStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDeliveryStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDeliveryStream"); err != nil {
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
	if err = addOpCreateDeliveryStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDeliveryStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDeliveryStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDeliveryStream",
	}
}
