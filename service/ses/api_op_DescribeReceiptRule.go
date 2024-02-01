// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the details of the specified receipt rule. For information about
// setting up receipt rules, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.html)
// . You can execute this operation no more than once per second.
func (c *Client) DescribeReceiptRule(ctx context.Context, params *DescribeReceiptRuleInput, optFns ...func(*Options)) (*DescribeReceiptRuleOutput, error) {
	if params == nil {
		params = &DescribeReceiptRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReceiptRule", params, optFns, c.addOperationDescribeReceiptRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReceiptRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to return the details of a receipt rule. You use receipt
// rules to receive email with Amazon SES. For more information, see the Amazon
// SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-concepts.html)
// .
type DescribeReceiptRuleInput struct {

	// The name of the receipt rule.
	//
	// This member is required.
	RuleName *string

	// The name of the receipt rule set that the receipt rule belongs to.
	//
	// This member is required.
	RuleSetName *string

	noSmithyDocumentSerde
}

// Represents the details of a receipt rule.
type DescribeReceiptRuleOutput struct {

	// A data structure that contains the specified receipt rule's name, actions,
	// recipients, domains, enabled status, scan status, and Transport Layer Security
	// (TLS) policy.
	Rule *types.ReceiptRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReceiptRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeReceiptRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeReceiptRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeReceiptRule"); err != nil {
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
	if err = addOpDescribeReceiptRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReceiptRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeReceiptRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeReceiptRule",
	}
}
