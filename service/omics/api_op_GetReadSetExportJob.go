// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Gets information about a read set export job.
func (c *Client) GetReadSetExportJob(ctx context.Context, params *GetReadSetExportJobInput, optFns ...func(*Options)) (*GetReadSetExportJobOutput, error) {
	if params == nil {
		params = &GetReadSetExportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReadSetExportJob", params, optFns, c.addOperationGetReadSetExportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReadSetExportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetReadSetExportJobInput struct {

	// The job's ID.
	//
	// This member is required.
	Id *string

	// The job's sequence store ID.
	//
	// This member is required.
	SequenceStoreId *string

	noSmithyDocumentSerde
}

type GetReadSetExportJobOutput struct {

	// When the job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The job's destination in Amazon S3.
	//
	// This member is required.
	Destination *string

	// The job's ID.
	//
	// This member is required.
	Id *string

	// The job's sequence store ID.
	//
	// This member is required.
	SequenceStoreId *string

	// The job's status.
	//
	// This member is required.
	Status types.ReadSetExportJobStatus

	// When the job completed.
	CompletionTime *time.Time

	// The job's read sets.
	ReadSets []types.ExportReadSetDetail

	// The job's status message.
	StatusMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetReadSetExportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetReadSetExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetReadSetExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetReadSetExportJob"); err != nil {
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
	if err = addEndpointPrefix_opGetReadSetExportJobMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetReadSetExportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReadSetExportJob(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetReadSetExportJobMiddleware struct {
}

func (*endpointPrefix_opGetReadSetExportJobMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetReadSetExportJobMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "control-storage-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetReadSetExportJobMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetReadSetExportJobMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// GetReadSetExportJobAPIClient is a client that implements the
// GetReadSetExportJob operation.
type GetReadSetExportJobAPIClient interface {
	GetReadSetExportJob(context.Context, *GetReadSetExportJobInput, ...func(*Options)) (*GetReadSetExportJobOutput, error)
}

var _ GetReadSetExportJobAPIClient = (*Client)(nil)

// ReadSetExportJobCompletedWaiterOptions are waiter options for
// ReadSetExportJobCompletedWaiter
type ReadSetExportJobCompletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ReadSetExportJobCompletedWaiter will use default minimum delay of 30 seconds.
	// Note that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, ReadSetExportJobCompletedWaiter will use default max delay of 600
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetReadSetExportJobInput, *GetReadSetExportJobOutput, error) (bool, error)
}

// ReadSetExportJobCompletedWaiter defines the waiters for
// ReadSetExportJobCompleted
type ReadSetExportJobCompletedWaiter struct {
	client GetReadSetExportJobAPIClient

	options ReadSetExportJobCompletedWaiterOptions
}

// NewReadSetExportJobCompletedWaiter constructs a ReadSetExportJobCompletedWaiter.
func NewReadSetExportJobCompletedWaiter(client GetReadSetExportJobAPIClient, optFns ...func(*ReadSetExportJobCompletedWaiterOptions)) *ReadSetExportJobCompletedWaiter {
	options := ReadSetExportJobCompletedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 600 * time.Second
	options.Retryable = readSetExportJobCompletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ReadSetExportJobCompletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ReadSetExportJobCompleted waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *ReadSetExportJobCompletedWaiter) Wait(ctx context.Context, params *GetReadSetExportJobInput, maxWaitDur time.Duration, optFns ...func(*ReadSetExportJobCompletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ReadSetExportJobCompleted waiter
// and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ReadSetExportJobCompletedWaiter) WaitForOutput(ctx context.Context, params *GetReadSetExportJobInput, maxWaitDur time.Duration, optFns ...func(*ReadSetExportJobCompletedWaiterOptions)) (*GetReadSetExportJobOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 600 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetReadSetExportJob(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for ReadSetExportJobCompleted waiter")
}

func readSetExportJobCompletedStateRetryable(ctx context.Context, input *GetReadSetExportJobInput, output *GetReadSetExportJobOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "COMPLETED"
		value, ok := pathValue.(types.ReadSetExportJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ReadSetExportJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "SUBMITTED"
		value, ok := pathValue.(types.ReadSetExportJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ReadSetExportJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "IN_PROGRESS"
		value, ok := pathValue.(types.ReadSetExportJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ReadSetExportJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CANCELLING"
		value, ok := pathValue.(types.ReadSetExportJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ReadSetExportJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CANCELLED"
		value, ok := pathValue.(types.ReadSetExportJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ReadSetExportJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		value, ok := pathValue.(types.ReadSetExportJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ReadSetExportJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "COMPLETED_WITH_FAILURES"
		value, ok := pathValue.(types.ReadSetExportJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ReadSetExportJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetReadSetExportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetReadSetExportJob",
	}
}
