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
)

// Starts the asynchronous analysis of an input document for relationships between
// detected items such as key-value pairs, tables, and selection elements.
// StartDocumentAnalysis can analyze text in documents that are in JPEG, PNG, TIFF,
// and PDF format. The documents are stored in an Amazon S3 bucket. Use
// DocumentLocation to specify the bucket name and file name of the document.
// StartDocumentAnalysis returns a job identifier ( JobId ) that you use to get the
// results of the operation. When text analysis is finished, Amazon Textract
// publishes a completion status to the Amazon Simple Notification Service (Amazon
// SNS) topic that you specify in NotificationChannel . To get the results of the
// text analysis operation, first check that the status value published to the
// Amazon SNS topic is SUCCEEDED . If so, call GetDocumentAnalysis , and pass the
// job identifier ( JobId ) from the initial call to StartDocumentAnalysis . For
// more information, see Document Text Analysis (https://docs.aws.amazon.com/textract/latest/dg/how-it-works-analyzing.html)
// .
func (c *Client) StartDocumentAnalysis(ctx context.Context, params *StartDocumentAnalysisInput, optFns ...func(*Options)) (*StartDocumentAnalysisOutput, error) {
	if params == nil {
		params = &StartDocumentAnalysisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDocumentAnalysis", params, optFns, c.addOperationStartDocumentAnalysisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDocumentAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDocumentAnalysisInput struct {

	// The location of the document to be processed.
	//
	// This member is required.
	DocumentLocation *types.DocumentLocation

	// A list of the types of analysis to perform. Add TABLES to the list to return
	// information about the tables that are detected in the input document. Add FORMS
	// to return detected form data. To perform both types of analysis, add TABLES and
	// FORMS to FeatureTypes . All lines and words detected in the document are
	// included in the response (including text that isn't related to the value of
	// FeatureTypes ).
	//
	// This member is required.
	FeatureTypes []types.FeatureType

	// Specifies the adapter to be used when analyzing a document.
	AdaptersConfig *types.AdaptersConfig

	// The idempotent token that you use to identify the start request. If you use the
	// same token with multiple StartDocumentAnalysis requests, the same JobId is
	// returned. Use ClientRequestToken to prevent the same job from being
	// accidentally started more than once. For more information, see Calling Amazon
	// Textract Asynchronous Operations (https://docs.aws.amazon.com/textract/latest/dg/api-async.html)
	// .
	ClientRequestToken *string

	// An identifier that you specify that's included in the completion notification
	// published to the Amazon SNS topic. For example, you can use JobTag to identify
	// the type of document that the completion notification corresponds to (such as a
	// tax form or a receipt).
	JobTag *string

	// The KMS key used to encrypt the inference results. This can be in either Key ID
	// or Key Alias format. When a KMS key is provided, the KMS key will be used for
	// server-side encryption of the objects in the customer bucket. When this
	// parameter is not enabled, the result will be encrypted server side,using SSE-S3.
	KMSKeyId *string

	// The Amazon SNS topic ARN that you want Amazon Textract to publish the
	// completion status of the operation to.
	NotificationChannel *types.NotificationChannel

	// Sets if the output will go to a customer defined bucket. By default, Amazon
	// Textract will save the results internally to be accessed by the
	// GetDocumentAnalysis operation.
	OutputConfig *types.OutputConfig

	//
	QueriesConfig *types.QueriesConfig

	noSmithyDocumentSerde
}

type StartDocumentAnalysisOutput struct {

	// The identifier for the document text detection job. Use JobId to identify the
	// job in a subsequent call to GetDocumentAnalysis . A JobId value is only valid
	// for 7 days.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartDocumentAnalysisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartDocumentAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartDocumentAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartDocumentAnalysis"); err != nil {
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
	if err = addOpStartDocumentAnalysisValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartDocumentAnalysis(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartDocumentAnalysis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartDocumentAnalysis",
	}
}
