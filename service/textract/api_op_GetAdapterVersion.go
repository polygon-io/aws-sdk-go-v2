// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets configuration information for the specified adapter version, including:
// AdapterId, AdapterVersion, FeatureTypes, Status, StatusMessage, DatasetConfig,
// KMSKeyId, OutputConfig, Tags and EvaluationMetrics.
func (c *Client) GetAdapterVersion(ctx context.Context, params *GetAdapterVersionInput, optFns ...func(*Options)) (*GetAdapterVersionOutput, error) {
	if params == nil {
		params = &GetAdapterVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAdapterVersion", params, optFns, c.addOperationGetAdapterVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAdapterVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAdapterVersionInput struct {

	// A string specifying a unique ID for the adapter version you want to retrieve
	// information for.
	//
	// This member is required.
	AdapterId *string

	// A string specifying the adapter version you want to retrieve information for.
	//
	// This member is required.
	AdapterVersion *string

	noSmithyDocumentSerde
}

type GetAdapterVersionOutput struct {

	// A string containing a unique ID for the adapter version being retrieved.
	AdapterId *string

	// A string containing the adapter version that has been retrieved.
	AdapterVersion *string

	// The time that the adapter version was created.
	CreationTime *time.Time

	// Specifies a dataset used to train a new adapter version. Takes a
	// ManifestS3Objec as the value.
	DatasetConfig *types.AdapterVersionDatasetConfig

	// The evaluation metrics (F1 score, Precision, and Recall) for the requested
	// version, grouped by baseline metrics and adapter version.
	EvaluationMetrics []types.AdapterVersionEvaluationMetric

	// List of the targeted feature types for the requested adapter version.
	FeatureTypes []types.FeatureType

	// The identifier for your AWS Key Management Service key (AWS KMS key). Used to
	// encrypt your documents.
	KMSKeyId *string

	// Sets whether or not your output will go to a user created bucket. Used to set
	// the name of the bucket, and the prefix on the output file. OutputConfig is an
	// optional parameter which lets you adjust where your output will be placed. By
	// default, Amazon Textract will store the results internally and can only be
	// accessed by the Get API operations. With OutputConfig enabled, you can set the
	// name of the bucket the output will be sent to the file prefix of the results
	// where you can download your results. Additionally, you can set the KMSKeyID
	// parameter to a customer master key (CMK) to encrypt your output. Without this
	// parameter set Amazon Textract will encrypt server-side using the AWS managed CMK
	// for Amazon S3. Decryption of Customer Content is necessary for processing of the
	// documents by Amazon Textract. If your account is opted out under an AI services
	// opt out policy then all unencrypted Customer Content is immediately and
	// permanently deleted after the Customer Content has been processed by the
	// service. No copy of of the output is retained by Amazon Textract. For
	// information about how to opt out, see Managing AI services opt-out policy.  (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_ai-opt-out.html)
	// For more information on data privacy, see the Data Privacy FAQ (https://aws.amazon.com/compliance/data-privacy-faq/)
	// .
	OutputConfig *types.OutputConfig

	// The status of the adapter version that has been requested.
	Status types.AdapterVersionStatus

	// A message that describes the status of the requested adapter version.
	StatusMessage *string

	// A set of tags (key-value pairs) that are associated with the adapter version.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAdapterVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAdapterVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAdapterVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAdapterVersion"); err != nil {
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
	if err = addOpGetAdapterVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAdapterVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAdapterVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAdapterVersion",
	}
}
