// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disables Health from working with Organizations. To call this operation, you
// must sign in to the organization's management account. For more information, see
// Aggregating Health events (https://docs.aws.amazon.com/health/latest/ug/aggregate-events.html)
// in the Health User Guide. This operation doesn't remove the service-linked role
// from the management account in your organization. You must use the IAM console,
// API, or Command Line Interface (CLI) to remove the service-linked role. For more
// information, see Deleting a Service-Linked Role (https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html#delete-service-linked-role)
// in the IAM User Guide. You can also disable the organizational feature by using
// the Organizations DisableAWSServiceAccess (https://docs.aws.amazon.com/organizations/latest/APIReference/API_DisableAWSServiceAccess.html)
// API operation. After you call this operation, Health stops aggregating events
// for all other Amazon Web Services accounts in your organization. If you call the
// Health API operations for organizational view, Health returns an error. Health
// continues to aggregate health events for your Amazon Web Services account.
func (c *Client) DisableHealthServiceAccessForOrganization(ctx context.Context, params *DisableHealthServiceAccessForOrganizationInput, optFns ...func(*Options)) (*DisableHealthServiceAccessForOrganizationOutput, error) {
	if params == nil {
		params = &DisableHealthServiceAccessForOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableHealthServiceAccessForOrganization", params, optFns, c.addOperationDisableHealthServiceAccessForOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableHealthServiceAccessForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableHealthServiceAccessForOrganizationInput struct {
	noSmithyDocumentSerde
}

type DisableHealthServiceAccessForOrganizationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisableHealthServiceAccessForOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisableHealthServiceAccessForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisableHealthServiceAccessForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisableHealthServiceAccessForOrganization"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableHealthServiceAccessForOrganization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisableHealthServiceAccessForOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisableHealthServiceAccessForOrganization",
	}
}
