// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the details of a single prepared statement or a list of up to 256
// prepared statements for the array of prepared statement names that you provide.
// Requires you to have access to the workgroup to which the prepared statements
// belong. If a prepared statement cannot be retrieved for the name specified, the
// statement is listed in UnprocessedPreparedStatementNames .
func (c *Client) BatchGetPreparedStatement(ctx context.Context, params *BatchGetPreparedStatementInput, optFns ...func(*Options)) (*BatchGetPreparedStatementOutput, error) {
	if params == nil {
		params = &BatchGetPreparedStatementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetPreparedStatement", params, optFns, c.addOperationBatchGetPreparedStatementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetPreparedStatementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetPreparedStatementInput struct {

	// A list of prepared statement names to return.
	//
	// This member is required.
	PreparedStatementNames []string

	// The name of the workgroup to which the prepared statements belong.
	//
	// This member is required.
	WorkGroup *string

	noSmithyDocumentSerde
}

type BatchGetPreparedStatementOutput struct {

	// The list of prepared statements returned.
	PreparedStatements []types.PreparedStatement

	// A list of one or more prepared statements that were requested but could not be
	// returned.
	UnprocessedPreparedStatementNames []types.UnprocessedPreparedStatementName

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetPreparedStatementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetPreparedStatement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetPreparedStatement{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetPreparedStatement"); err != nil {
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
	if err = addOpBatchGetPreparedStatementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetPreparedStatement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetPreparedStatement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetPreparedStatement",
	}
}
