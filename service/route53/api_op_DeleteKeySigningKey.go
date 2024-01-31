// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a key-signing key (KSK). Before you can delete a KSK, you must
// deactivate it. The KSK must be deactivated before you can delete it regardless
// of whether the hosted zone is enabled for DNSSEC signing. You can use
// DeactivateKeySigningKey (https://docs.aws.amazon.com/Route53/latest/APIReference/API_DeactivateKeySigningKey.html)
// to deactivate the key before you delete it. Use GetDNSSEC (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetDNSSEC.html)
// to verify that the KSK is in an INACTIVE status.
func (c *Client) DeleteKeySigningKey(ctx context.Context, params *DeleteKeySigningKeyInput, optFns ...func(*Options)) (*DeleteKeySigningKeyOutput, error) {
	if params == nil {
		params = &DeleteKeySigningKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteKeySigningKey", params, optFns, c.addOperationDeleteKeySigningKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteKeySigningKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteKeySigningKeyInput struct {

	// A unique string used to identify a hosted zone.
	//
	// This member is required.
	HostedZoneId *string

	// A string used to identify a key-signing key (KSK).
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DeleteKeySigningKeyOutput struct {

	// A complex type that describes change information about changes made to your
	// hosted zone.
	//
	// This member is required.
	ChangeInfo *types.ChangeInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteKeySigningKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteKeySigningKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteKeySigningKey{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteKeySigningKey"); err != nil {
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
	if err = addOpDeleteKeySigningKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteKeySigningKey(options.Region), middleware.Before); err != nil {
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
	if err = addSanitizeURLMiddleware(stack); err != nil {
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

func newServiceMetadataMiddleware_opDeleteKeySigningKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteKeySigningKey",
	}
}
