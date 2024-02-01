// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels a Gremlin query. See Gremlin query cancellation (https://docs.aws.amazon.com/neptune/latest/userguide/gremlin-api-status-cancel.html)
// for more information. When invoking this operation in a Neptune cluster that has
// IAM authentication enabled, the IAM user or role making the request must have a
// policy attached that allows the neptune-db:CancelQuery (https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#cancelquery)
// IAM action in that cluster.
func (c *Client) CancelGremlinQuery(ctx context.Context, params *CancelGremlinQueryInput, optFns ...func(*Options)) (*CancelGremlinQueryOutput, error) {
	if params == nil {
		params = &CancelGremlinQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelGremlinQuery", params, optFns, c.addOperationCancelGremlinQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelGremlinQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelGremlinQueryInput struct {

	// The unique identifier that identifies the query to be canceled.
	//
	// This member is required.
	QueryId *string

	noSmithyDocumentSerde
}

type CancelGremlinQueryOutput struct {

	// The status of the cancelation
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelGremlinQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelGremlinQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelGremlinQuery{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CancelGremlinQuery"); err != nil {
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
	if err = addOpCancelGremlinQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelGremlinQuery(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelGremlinQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CancelGremlinQuery",
	}
}
