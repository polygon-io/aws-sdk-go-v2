// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/detective/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the investigation results of an investigation for a behavior graph.
func (c *Client) GetInvestigation(ctx context.Context, params *GetInvestigationInput, optFns ...func(*Options)) (*GetInvestigationOutput, error) {
	if params == nil {
		params = &GetInvestigationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInvestigation", params, optFns, c.addOperationGetInvestigationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInvestigationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInvestigationInput struct {

	// The ARN of the behavior graph.
	//
	// This member is required.
	GraphArn *string

	// The investigation ID of the investigation report.
	//
	// This member is required.
	InvestigationId *string

	noSmithyDocumentSerde
}

type GetInvestigationOutput struct {

	// The UTC time stamp of the creation time of the investigation report.
	CreatedTime *time.Time

	// The unique Amazon Resource Name (ARN) of the IAM user and IAM role.
	EntityArn *string

	// Type of entity. For example, Amazon Web Services accounts, such as IAM user and
	// role.
	EntityType types.EntityType

	// The ARN of the behavior graph.
	GraphArn *string

	// The investigation ID of the investigation report.
	InvestigationId *string

	// The data and time when the investigation began. The value is an UTC ISO8601
	// formatted string. For example, 2021-08-18T16:35:56.284Z.
	ScopeEndTime *time.Time

	// The start date and time for the scope time set to generate the investigation
	// report.
	ScopeStartTime *time.Time

	// Severity based on the likelihood and impact of the indicators of compromise
	// discovered in the investigation.
	Severity types.Severity

	// The current state of the investigation. An archived investigation indicates you
	// have completed reviewing the investigation.
	State types.State

	// Status based on the completion status of the investigation.
	Status types.Status

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetInvestigationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetInvestigation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetInvestigation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetInvestigation"); err != nil {
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
	if err = addOpGetInvestigationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInvestigation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetInvestigation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetInvestigation",
	}
}
