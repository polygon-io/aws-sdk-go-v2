// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation renews a domain for the specified number of years. The cost of
// renewing your domain is billed to your Amazon Web Services account. We recommend
// that you renew your domain several weeks before the expiration date. Some TLD
// registries delete domains before the expiration date if you haven't renewed far
// enough in advance. For more information about renewing domain registration, see
// Renewing Registration for a Domain (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-renew.html)
// in the Amazon Route 53 Developer Guide.
func (c *Client) RenewDomain(ctx context.Context, params *RenewDomainInput, optFns ...func(*Options)) (*RenewDomainOutput, error) {
	if params == nil {
		params = &RenewDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RenewDomain", params, optFns, c.addOperationRenewDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RenewDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A RenewDomain request includes the number of years that you want to renew for
// and the current expiration year.
type RenewDomainInput struct {

	// The year when the registration for the domain is set to expire. This value must
	// match the current expiration date for the domain.
	//
	// This member is required.
	CurrentExpiryYear int32

	// The name of the domain that you want to renew.
	//
	// This member is required.
	DomainName *string

	// The number of years that you want to renew the domain for. The maximum number
	// of years depends on the top-level domain. For the range of valid values for your
	// domain, see Domains that You Can Register with Amazon Route 53 (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html)
	// in the Amazon Route 53 Developer Guide. Default: 1
	DurationInYears *int32

	noSmithyDocumentSerde
}

type RenewDomainOutput struct {

	// Identifier for tracking the progress of the request. To query the operation
	// status, use GetOperationDetail (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html)
	// .
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRenewDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRenewDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRenewDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RenewDomain"); err != nil {
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
	if err = addOpRenewDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRenewDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRenewDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RenewDomain",
	}
}
