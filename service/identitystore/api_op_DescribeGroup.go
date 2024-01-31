// Code generated by smithy-go-codegen DO NOT EDIT.

package identitystore

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the group metadata and attributes from GroupId in an identity store.
// If you have administrator access to a member account, you can use this API from
// the member account. Read about member accounts (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_access.html)
// in the Organizations User Guide.
func (c *Client) DescribeGroup(ctx context.Context, params *DescribeGroupInput, optFns ...func(*Options)) (*DescribeGroupOutput, error) {
	if params == nil {
		params = &DescribeGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGroup", params, optFns, c.addOperationDescribeGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGroupInput struct {

	// The identifier for a group in the identity store.
	//
	// This member is required.
	GroupId *string

	// The globally unique identifier for the identity store, such as d-1234567890 . In
	// this example, d- is a fixed prefix, and 1234567890 is a randomly generated
	// string that contains numbers and lower case letters. This value is generated at
	// the time that a new identity store is created.
	//
	// This member is required.
	IdentityStoreId *string

	noSmithyDocumentSerde
}

type DescribeGroupOutput struct {

	// The identifier for a group in the identity store.
	//
	// This member is required.
	GroupId *string

	// The globally unique identifier for the identity store.
	//
	// This member is required.
	IdentityStoreId *string

	// A string containing a description of the group.
	Description *string

	// The group’s display name value. The length limit is 1,024 characters. This
	// value can consist of letters, accented characters, symbols, numbers,
	// punctuation, tab, new line, carriage return, space, and nonbreaking space in
	// this attribute. This value is specified at the time that the group is created
	// and stored as an attribute of the group object in the identity store.
	DisplayName *string

	// A list of ExternalId objects that contains the identifiers issued to this
	// resource by an external identity provider.
	ExternalIds []types.ExternalId

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeGroup"); err != nil {
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
	if err = addOpDescribeGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeGroup",
	}
}
