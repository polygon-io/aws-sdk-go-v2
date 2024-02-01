// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the celebrity recognition results for a Amazon Rekognition Video analysis
// started by StartCelebrityRecognition . Celebrity recognition in a video is an
// asynchronous operation. Analysis is started by a call to
// StartCelebrityRecognition which returns a job identifier ( JobId ). When the
// celebrity recognition operation finishes, Amazon Rekognition Video publishes a
// completion status to the Amazon Simple Notification Service topic registered in
// the initial call to StartCelebrityRecognition . To get the results of the
// celebrity recognition analysis, first check that the status value published to
// the Amazon SNS topic is SUCCEEDED . If so, call GetCelebrityDetection and pass
// the job identifier ( JobId ) from the initial call to StartCelebrityDetection .
// For more information, see Working With Stored Videos in the Amazon Rekognition
// Developer Guide. GetCelebrityRecognition returns detected celebrities and the
// time(s) they are detected in an array ( Celebrities ) of CelebrityRecognition
// objects. Each CelebrityRecognition contains information about the celebrity in
// a CelebrityDetail object and the time, Timestamp , the celebrity was detected.
// This CelebrityDetail object stores information about the detected celebrity's
// face attributes, a face bounding box, known gender, the celebrity's name, and a
// confidence estimate. GetCelebrityRecognition only returns the default facial
// attributes ( BoundingBox , Confidence , Landmarks , Pose , and Quality ). The
// BoundingBox field only applies to the detected face instance. The other facial
// attributes listed in the Face object of the following response syntax are not
// returned. For more information, see FaceDetail in the Amazon Rekognition
// Developer Guide. By default, the Celebrities array is sorted by time
// (milliseconds from the start of the video). You can also sort the array by
// celebrity by specifying the value ID in the SortBy input parameter. The
// CelebrityDetail object includes the celebrity identifer and additional
// information urls. If you don't store the additional information urls, you can
// get them later by calling GetCelebrityInfo with the celebrity identifer. No
// information is returned for faces not recognized as celebrities. Use MaxResults
// parameter to limit the number of labels returned. If there are more results than
// specified in MaxResults , the value of NextToken in the operation response
// contains a pagination token for getting the next set of results. To get the next
// page of results, call GetCelebrityDetection and populate the NextToken request
// parameter with the token value returned from the previous call to
// GetCelebrityRecognition .
func (c *Client) GetCelebrityRecognition(ctx context.Context, params *GetCelebrityRecognitionInput, optFns ...func(*Options)) (*GetCelebrityRecognitionOutput, error) {
	if params == nil {
		params = &GetCelebrityRecognitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCelebrityRecognition", params, optFns, c.addOperationGetCelebrityRecognitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCelebrityRecognitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCelebrityRecognitionInput struct {

	// Job identifier for the required celebrity recognition analysis. You can get the
	// job identifer from a call to StartCelebrityRecognition .
	//
	// This member is required.
	JobId *string

	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000. If you specify a value greater than 1000, a maximum of 1000
	// results is returned. The default value is 1000.
	MaxResults *int32

	// If the previous response was incomplete (because there is more recognized
	// celebrities to retrieve), Amazon Rekognition Video returns a pagination token in
	// the response. You can use this pagination token to retrieve the next set of
	// celebrities.
	NextToken *string

	// Sort to use for celebrities returned in Celebrities field. Specify ID to sort
	// by the celebrity identifier, specify TIMESTAMP to sort by the time the
	// celebrity was recognized.
	SortBy types.CelebrityRecognitionSortBy

	noSmithyDocumentSerde
}

type GetCelebrityRecognitionOutput struct {

	// Array of celebrities recognized in the video.
	Celebrities []types.CelebrityRecognition

	// Job identifier for the celebrity recognition operation for which you want to
	// obtain results. The job identifer is returned by an initial call to
	// StartCelebrityRecognition.
	JobId *string

	// The current status of the celebrity recognition job.
	JobStatus types.VideoJobStatus

	// A job identifier specified in the call to StartCelebrityRecognition and
	// returned in the job completion notification sent to your Amazon Simple
	// Notification Service topic.
	JobTag *string

	// If the response is truncated, Amazon Rekognition Video returns this token that
	// you can use in the subsequent request to retrieve the next set of celebrities.
	NextToken *string

	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string

	// Video file stored in an Amazon S3 bucket. Amazon Rekognition video start
	// operations such as StartLabelDetection use Video to specify a video for
	// analysis. The supported file formats are .mp4, .mov and .avi.
	Video *types.Video

	// Information about a video that Amazon Rekognition Video analyzed. Videometadata
	// is returned in every page of paginated responses from a Amazon Rekognition Video
	// operation.
	VideoMetadata *types.VideoMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCelebrityRecognitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCelebrityRecognition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCelebrityRecognition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCelebrityRecognition"); err != nil {
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
	if err = addOpGetCelebrityRecognitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCelebrityRecognition(options.Region), middleware.Before); err != nil {
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

// GetCelebrityRecognitionAPIClient is a client that implements the
// GetCelebrityRecognition operation.
type GetCelebrityRecognitionAPIClient interface {
	GetCelebrityRecognition(context.Context, *GetCelebrityRecognitionInput, ...func(*Options)) (*GetCelebrityRecognitionOutput, error)
}

var _ GetCelebrityRecognitionAPIClient = (*Client)(nil)

// GetCelebrityRecognitionPaginatorOptions is the paginator options for
// GetCelebrityRecognition
type GetCelebrityRecognitionPaginatorOptions struct {
	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000. If you specify a value greater than 1000, a maximum of 1000
	// results is returned. The default value is 1000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetCelebrityRecognitionPaginator is a paginator for GetCelebrityRecognition
type GetCelebrityRecognitionPaginator struct {
	options   GetCelebrityRecognitionPaginatorOptions
	client    GetCelebrityRecognitionAPIClient
	params    *GetCelebrityRecognitionInput
	nextToken *string
	firstPage bool
}

// NewGetCelebrityRecognitionPaginator returns a new
// GetCelebrityRecognitionPaginator
func NewGetCelebrityRecognitionPaginator(client GetCelebrityRecognitionAPIClient, params *GetCelebrityRecognitionInput, optFns ...func(*GetCelebrityRecognitionPaginatorOptions)) *GetCelebrityRecognitionPaginator {
	if params == nil {
		params = &GetCelebrityRecognitionInput{}
	}

	options := GetCelebrityRecognitionPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetCelebrityRecognitionPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetCelebrityRecognitionPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetCelebrityRecognition page.
func (p *GetCelebrityRecognitionPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetCelebrityRecognitionOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetCelebrityRecognition(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetCelebrityRecognition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCelebrityRecognition",
	}
}
