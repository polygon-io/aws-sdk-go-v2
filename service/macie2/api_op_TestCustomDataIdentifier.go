// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Tests a custom data identifier.
func (c *Client) TestCustomDataIdentifier(ctx context.Context, params *TestCustomDataIdentifierInput, optFns ...func(*Options)) (*TestCustomDataIdentifierOutput, error) {
	if params == nil {
		params = &TestCustomDataIdentifierInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestCustomDataIdentifier", params, optFns, c.addOperationTestCustomDataIdentifierMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestCustomDataIdentifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestCustomDataIdentifierInput struct {

	// The regular expression (regex) that defines the pattern to match. The
	// expression can contain as many as 512 characters.
	//
	// This member is required.
	Regex *string

	// The sample text to inspect by using the custom data identifier. The text can
	// contain as many as 1,000 characters.
	//
	// This member is required.
	SampleText *string

	// An array that lists specific character sequences (ignore words) to exclude from
	// the results. If the text matched by the regular expression contains any string
	// in this array, Amazon Macie ignores it. The array can contain as many as 10
	// ignore words. Each ignore word can contain 4-90 UTF-8 characters. Ignore words
	// are case sensitive.
	IgnoreWords []string

	// An array that lists specific character sequences (keywords), one of which must
	// precede and be within proximity (maximumMatchDistance) of the regular expression
	// to match. The array can contain as many as 50 keywords. Each keyword can contain
	// 3-90 UTF-8 characters. Keywords aren't case sensitive.
	Keywords []string

	// The maximum number of characters that can exist between the end of at least one
	// complete character sequence specified by the keywords array and the end of the
	// text that matches the regex pattern. If a complete keyword precedes all the text
	// that matches the pattern and the keyword is within the specified distance,
	// Amazon Macie includes the result. The distance can be 1-300 characters. The
	// default value is 50.
	MaximumMatchDistance *int32

	noSmithyDocumentSerde
}

type TestCustomDataIdentifierOutput struct {

	// The number of occurrences of sample text that matched the criteria specified by
	// the custom data identifier.
	MatchCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTestCustomDataIdentifierMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpTestCustomDataIdentifier{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpTestCustomDataIdentifier{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "TestCustomDataIdentifier"); err != nil {
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
	if err = addOpTestCustomDataIdentifierValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTestCustomDataIdentifier(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTestCustomDataIdentifier(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TestCustomDataIdentifier",
	}
}
