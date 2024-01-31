// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html)
// . With the latest version, AWS WAF has a single set of endpoints for regional
// and global use. Permanently deletes a GeoMatchSet . You can't delete a
// GeoMatchSet if it's still used in any Rules or if it still includes any
// countries. If you just want to remove a GeoMatchSet from a Rule , use UpdateRule
// . To permanently delete a GeoMatchSet from AWS WAF, perform the following
// steps:
//   - Update the GeoMatchSet to remove any countries. For more information, see
//     UpdateGeoMatchSet .
//   - Use GetChangeToken to get the change token that you provide in the
//     ChangeToken parameter of a DeleteGeoMatchSet request.
//   - Submit a DeleteGeoMatchSet request.
func (c *Client) DeleteGeoMatchSet(ctx context.Context, params *DeleteGeoMatchSetInput, optFns ...func(*Options)) (*DeleteGeoMatchSetOutput, error) {
	if params == nil {
		params = &DeleteGeoMatchSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteGeoMatchSet", params, optFns, c.addOperationDeleteGeoMatchSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteGeoMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteGeoMatchSetInput struct {

	// The value returned by the most recent call to GetChangeToken .
	//
	// This member is required.
	ChangeToken *string

	// The GeoMatchSetID of the GeoMatchSet that you want to delete. GeoMatchSetId is
	// returned by CreateGeoMatchSet and by ListGeoMatchSets .
	//
	// This member is required.
	GeoMatchSetId *string

	noSmithyDocumentSerde
}

type DeleteGeoMatchSetOutput struct {

	// The ChangeToken that you used to submit the DeleteGeoMatchSet request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus .
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteGeoMatchSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteGeoMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteGeoMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteGeoMatchSet"); err != nil {
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
	if err = addOpDeleteGeoMatchSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteGeoMatchSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteGeoMatchSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteGeoMatchSet",
	}
}
