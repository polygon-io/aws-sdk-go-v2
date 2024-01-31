// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates an Firewall Manager administrator account. The account must
// be a member of the organization that was onboarded to Firewall Manager by
// AssociateAdminAccount . Only the organization's management account can create an
// Firewall Manager administrator account. When you create an Firewall Manager
// administrator account, the service checks to see if the account is already a
// delegated administrator within Organizations. If the account isn't a delegated
// administrator, Firewall Manager calls Organizations to delegate the account
// within Organizations. For more information about administrator accounts within
// Organizations, see Managing the Amazon Web Services Accounts in Your
// Organization (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts.html)
// .
func (c *Client) PutAdminAccount(ctx context.Context, params *PutAdminAccountInput, optFns ...func(*Options)) (*PutAdminAccountOutput, error) {
	if params == nil {
		params = &PutAdminAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAdminAccount", params, optFns, c.addOperationPutAdminAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAdminAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAdminAccountInput struct {

	// The Amazon Web Services account ID to add as an Firewall Manager administrator
	// account. The account must be a member of the organization that was onboarded to
	// Firewall Manager by AssociateAdminAccount . For more information about
	// Organizations, see Managing the Amazon Web Services Accounts in Your
	// Organization (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts.html)
	// .
	//
	// This member is required.
	AdminAccount *string

	// Configures the resources that the specified Firewall Manager administrator can
	// manage. As a best practice, set the administrative scope according to the
	// principles of least privilege. Only grant the administrator the specific
	// resources or permissions that they need to perform the duties of their role.
	AdminScope *types.AdminScope

	noSmithyDocumentSerde
}

type PutAdminAccountOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAdminAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAdminAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAdminAccount{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutAdminAccount"); err != nil {
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
	if err = addOpPutAdminAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAdminAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAdminAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutAdminAccount",
	}
}
