// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhuborchestrator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Get migration workflow.
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

	// The ID of the migration workflow.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetWorkflowOutput struct {

	// The configuration ID of the application configured in Application Discovery
	// Service.
	AdsApplicationConfigurationId *string

	// The name of the application configured in Application Discovery Service.
	AdsApplicationName *string

	// The Amazon Resource Name (ARN) of the migration workflow.
	Arn *string

	// Get a list of completed steps in the migration workflow.
	CompletedSteps *int32

	// The time at which the migration workflow was created.
	CreationTime *time.Time

	// The description of the migration workflow.
	Description *string

	// The time at which the migration workflow ended.
	EndTime *time.Time

	// The ID of the migration workflow.
	Id *string

	// The time at which the migration workflow was last modified.
	LastModifiedTime *time.Time

	// The time at which the migration workflow was last started.
	LastStartTime *time.Time

	// The time at which the migration workflow was last stopped.
	LastStopTime *time.Time

	// The name of the migration workflow.
	Name *string

	// The status of the migration workflow.
	Status types.MigrationWorkflowStatusEnum

	// The status message of the migration workflow.
	StatusMessage *string

	// The tags added to the migration workflow.
	Tags map[string]string

	// The ID of the template.
	TemplateId *string

	// List of AWS services utilized in a migration workflow.
	Tools []types.Tool

	// The total number of steps in the migration workflow.
	TotalSteps *int32

	// The Amazon S3 bucket where the migration logs are stored.
	WorkflowBucket *string

	// The inputs required for creating the migration workflow.
	WorkflowInputs map[string]types.StepInput

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

func newServiceMetadataMiddleware_opGetWorkflow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetWorkflow",
	}
}
