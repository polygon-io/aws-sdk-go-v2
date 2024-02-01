// Code generated by smithy-go-codegen DO NOT EDIT.

package entityresolution

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a MatchingWorkflow object which stores the configuration of the data
// processing job to be run. It is important to note that there should not be a
// pre-existing MatchingWorkflow with the same name. To modify an existing
// workflow, utilize the UpdateMatchingWorkflow API.
func (c *Client) CreateMatchingWorkflow(ctx context.Context, params *CreateMatchingWorkflowInput, optFns ...func(*Options)) (*CreateMatchingWorkflowOutput, error) {
	if params == nil {
		params = &CreateMatchingWorkflowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMatchingWorkflow", params, optFns, c.addOperationCreateMatchingWorkflowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMatchingWorkflowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMatchingWorkflowInput struct {

	// A list of InputSource objects, which have the fields InputSourceARN and
	// SchemaName .
	//
	// This member is required.
	InputSourceConfig []types.InputSource

	// A list of OutputSource objects, each of which contains fields OutputS3Path ,
	// ApplyNormalization , and Output .
	//
	// This member is required.
	OutputSourceConfig []types.OutputSource

	// An object which defines the resolutionType and the ruleBasedProperties .
	//
	// This member is required.
	ResolutionTechniques *types.ResolutionTechniques

	// The Amazon Resource Name (ARN) of the IAM role. Entity Resolution assumes this
	// role to create resources on your behalf as part of workflow execution.
	//
	// This member is required.
	RoleArn *string

	// The name of the workflow. There can't be multiple MatchingWorkflows with the
	// same name.
	//
	// This member is required.
	WorkflowName *string

	// A description of the workflow.
	Description *string

	// An object which defines an incremental run type and has only incrementalRunType
	// as a field.
	IncrementalRunConfig *types.IncrementalRunConfig

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateMatchingWorkflowOutput struct {

	// A list of InputSource objects, which have the fields InputSourceARN and
	// SchemaName .
	//
	// This member is required.
	InputSourceConfig []types.InputSource

	// A list of OutputSource objects, each of which contains fields OutputS3Path ,
	// ApplyNormalization , and Output .
	//
	// This member is required.
	OutputSourceConfig []types.OutputSource

	// An object which defines the resolutionType and the ruleBasedProperties .
	//
	// This member is required.
	ResolutionTechniques *types.ResolutionTechniques

	// The Amazon Resource Name (ARN) of the IAM role. Entity Resolution assumes this
	// role to create resources on your behalf as part of workflow execution.
	//
	// This member is required.
	RoleArn *string

	// The ARN (Amazon Resource Name) that Entity Resolution generated for the
	// MatchingWorkflow .
	//
	// This member is required.
	WorkflowArn *string

	// The name of the workflow.
	//
	// This member is required.
	WorkflowName *string

	// A description of the workflow.
	Description *string

	// An object which defines an incremental run type and has only incrementalRunType
	// as a field.
	IncrementalRunConfig *types.IncrementalRunConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMatchingWorkflowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMatchingWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMatchingWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateMatchingWorkflow"); err != nil {
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
	if err = addOpCreateMatchingWorkflowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMatchingWorkflow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMatchingWorkflow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateMatchingWorkflow",
	}
}
