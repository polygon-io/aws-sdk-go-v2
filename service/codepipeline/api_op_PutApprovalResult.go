// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides the response to a manual approval request to CodePipeline. Valid
// responses include Approved and Rejected.
func (c *Client) PutApprovalResult(ctx context.Context, params *PutApprovalResultInput, optFns ...func(*Options)) (*PutApprovalResultOutput, error) {
	if params == nil {
		params = &PutApprovalResultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutApprovalResult", params, optFns, c.addOperationPutApprovalResultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutApprovalResultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a PutApprovalResult action.
type PutApprovalResultInput struct {

	// The name of the action for which approval is requested.
	//
	// This member is required.
	ActionName *string

	// The name of the pipeline that contains the action.
	//
	// This member is required.
	PipelineName *string

	// Represents information about the result of the approval request.
	//
	// This member is required.
	Result *types.ApprovalResult

	// The name of the stage that contains the action.
	//
	// This member is required.
	StageName *string

	// The system-generated token used to identify a unique approval request. The
	// token for each open approval request can be obtained using the GetPipelineState
	// action. It is used to validate that the approval request corresponding to this
	// token is still valid.
	//
	// This member is required.
	Token *string

	noSmithyDocumentSerde
}

// Represents the output of a PutApprovalResult action.
type PutApprovalResultOutput struct {

	// The timestamp showing when the approval or rejection was submitted.
	ApprovedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutApprovalResultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutApprovalResult{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutApprovalResult{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutApprovalResult"); err != nil {
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
	if err = addOpPutApprovalResultValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutApprovalResult(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutApprovalResult(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutApprovalResult",
	}
}
