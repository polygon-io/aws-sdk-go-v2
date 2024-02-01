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

// Gets information about a workflow.
func (c *Client) GetWorkflow(ctx context.Context, params *GetWorkflowInput, optFns ...func(*Options)) (*GetWorkflowOutput, error) {
	if params == nil {
		params = &GetWorkflowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWorkflow", params, optFns, c.addOperationGetWorkflowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWorkflowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkflowInput struct {

	// The workflow's ID.
	//
	// This member is required.
	Id *string

	// The export format for the workflow.
	Export []types.WorkflowExport

	// The workflow's type.
	Type types.WorkflowType

	noSmithyDocumentSerde
}

type GetWorkflowOutput struct {

	// The computational accelerator specified to run the workflow.
	Accelerators types.Accelerators

	// The workflow's ARN.
	Arn *string

	// When the workflow was created.
	CreationTime *time.Time

	// The workflow's definition.
	Definition *string

	// The workflow's description.
	Description *string

	// The workflow's digest.
	Digest *string

	// The workflow's engine.
	Engine types.WorkflowEngine

	// The workflow's ID.
	Id *string

	// The path of the main definition file for the workflow.
	Main *string

	// Gets metadata for workflow.
	Metadata map[string]string

	// The workflow's name.
	Name *string

	// The workflow's parameter template.
	ParameterTemplate map[string]types.WorkflowParameter

	// The workflow's status.
	Status types.WorkflowStatus

	// The workflow's status message.
	StatusMessage *string

	// The workflow's storage capacity in gigabytes.
	StorageCapacity *int32

	// The workflow's tags.
	Tags map[string]string

	// The workflow's type.
	Type types.WorkflowType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWorkflowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetWorkflow"); err != nil {
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
	if err = addEndpointPrefix_opGetWorkflowMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetWorkflowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorkflow(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetWorkflowMiddleware struct {
}

func (*endpointPrefix_opGetWorkflowMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetWorkflowMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "workflows-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetWorkflowMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetWorkflowMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// GetWorkflowAPIClient is a client that implements the GetWorkflow operation.
type GetWorkflowAPIClient interface {
	GetWorkflow(context.Context, *GetWorkflowInput, ...func(*Options)) (*GetWorkflowOutput, error)
}

var _ GetWorkflowAPIClient = (*Client)(nil)

// WorkflowActiveWaiterOptions are waiter options for WorkflowActiveWaiter
type WorkflowActiveWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// WorkflowActiveWaiter will use default minimum delay of 3 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, WorkflowActiveWaiter will use default max delay of 30 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *GetWorkflowInput, *GetWorkflowOutput, error) (bool, error)
}

// WorkflowActiveWaiter defines the waiters for WorkflowActive
type WorkflowActiveWaiter struct {
	client GetWorkflowAPIClient

	options WorkflowActiveWaiterOptions
}

// NewWorkflowActiveWaiter constructs a WorkflowActiveWaiter.
func NewWorkflowActiveWaiter(client GetWorkflowAPIClient, optFns ...func(*WorkflowActiveWaiterOptions)) *WorkflowActiveWaiter {
	options := WorkflowActiveWaiterOptions{}
	options.MinDelay = 3 * time.Second
	options.MaxDelay = 30 * time.Second
	options.Retryable = workflowActiveStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &WorkflowActiveWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for WorkflowActive waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *WorkflowActiveWaiter) Wait(ctx context.Context, params *GetWorkflowInput, maxWaitDur time.Duration, optFns ...func(*WorkflowActiveWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for WorkflowActive waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *WorkflowActiveWaiter) WaitForOutput(ctx context.Context, params *GetWorkflowInput, maxWaitDur time.Duration, optFns ...func(*WorkflowActiveWaiterOptions)) (*GetWorkflowOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 30 * time.Second
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

		out, err := w.client.GetWorkflow(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for WorkflowActive waiter")
}

func workflowActiveStateRetryable(ctx context.Context, input *GetWorkflowInput, output *GetWorkflowOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ACTIVE"
		value, ok := pathValue.(types.WorkflowStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.WorkflowStatus value, got %T", pathValue)
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

		expectedValue := "CREATING"
		value, ok := pathValue.(types.WorkflowStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.WorkflowStatus value, got %T", pathValue)
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

		expectedValue := "UPDATING"
		value, ok := pathValue.(types.WorkflowStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.WorkflowStatus value, got %T", pathValue)
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

		expectedValue := "FAILED"
		value, ok := pathValue.(types.WorkflowStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.WorkflowStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetWorkflow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetWorkflow",
	}
}
