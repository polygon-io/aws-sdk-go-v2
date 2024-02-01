// Code generated by smithy-go-codegen DO NOT EDIT.

package databrew

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databrew/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves detailed information about the ruleset.
func (c *Client) DescribeRuleset(ctx context.Context, params *DescribeRulesetInput, optFns ...func(*Options)) (*DescribeRulesetOutput, error) {
	if params == nil {
		params = &DescribeRulesetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRuleset", params, optFns, c.addOperationDescribeRulesetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRulesetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRulesetInput struct {

	// The name of the ruleset to be described.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DescribeRulesetOutput struct {

	// The name of the ruleset.
	//
	// This member is required.
	Name *string

	// The date and time that the ruleset was created.
	CreateDate *time.Time

	// The Amazon Resource Name (ARN) of the user who created the ruleset.
	CreatedBy *string

	// The description of the ruleset.
	Description *string

	// The Amazon Resource Name (ARN) of the user who last modified the ruleset.
	LastModifiedBy *string

	// The modification date and time of the ruleset.
	LastModifiedDate *time.Time

	// The Amazon Resource Name (ARN) for the ruleset.
	ResourceArn *string

	// A list of rules that are defined with the ruleset. A rule includes one or more
	// checks to be validated on a DataBrew dataset.
	Rules []types.Rule

	// Metadata tags that have been applied to the ruleset.
	Tags map[string]string

	// The Amazon Resource Name (ARN) of a resource (dataset) that the ruleset is
	// associated with.
	TargetArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRulesetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRuleset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRuleset{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeRuleset"); err != nil {
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
	if err = addOpDescribeRulesetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRuleset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeRuleset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeRuleset",
	}
}
