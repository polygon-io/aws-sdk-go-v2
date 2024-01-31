// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Detaches the specified customer managed policy from the specified PermissionSet .
func (c *Client) DetachCustomerManagedPolicyReferenceFromPermissionSet(ctx context.Context, params *DetachCustomerManagedPolicyReferenceFromPermissionSetInput, optFns ...func(*Options)) (*DetachCustomerManagedPolicyReferenceFromPermissionSetOutput, error) {
	if params == nil {
		params = &DetachCustomerManagedPolicyReferenceFromPermissionSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachCustomerManagedPolicyReferenceFromPermissionSet", params, optFns, c.addOperationDetachCustomerManagedPolicyReferenceFromPermissionSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachCustomerManagedPolicyReferenceFromPermissionSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachCustomerManagedPolicyReferenceFromPermissionSetInput struct {

	// Specifies the name and path of a customer managed policy. You must have an IAM
	// policy that matches the name and path in each Amazon Web Services account where
	// you want to deploy your permission set.
	//
	// This member is required.
	CustomerManagedPolicyReference *types.CustomerManagedPolicyReference

	// The ARN of the IAM Identity Center instance under which the operation will be
	// executed.
	//
	// This member is required.
	InstanceArn *string

	// The ARN of the PermissionSet .
	//
	// This member is required.
	PermissionSetArn *string

	noSmithyDocumentSerde
}

type DetachCustomerManagedPolicyReferenceFromPermissionSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetachCustomerManagedPolicyReferenceFromPermissionSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetachCustomerManagedPolicyReferenceFromPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetachCustomerManagedPolicyReferenceFromPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DetachCustomerManagedPolicyReferenceFromPermissionSet"); err != nil {
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
	if err = addOpDetachCustomerManagedPolicyReferenceFromPermissionSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetachCustomerManagedPolicyReferenceFromPermissionSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetachCustomerManagedPolicyReferenceFromPermissionSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DetachCustomerManagedPolicyReferenceFromPermissionSet",
	}
}
