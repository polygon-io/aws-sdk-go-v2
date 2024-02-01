// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an endpoint for a Network File System (NFS) file server that DataSync
// can use for a data transfer. For more information, see Configuring transfers to
// or from an NFS file server (https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html)
// . If you're copying data to or from an Snowcone device, you can also use
// CreateLocationNfs to create your transfer location. For more information, see
// Configuring transfers with Snowcone (https://docs.aws.amazon.com/datasync/latest/userguide/nfs-on-snowcone.html)
// .
func (c *Client) CreateLocationNfs(ctx context.Context, params *CreateLocationNfsInput, optFns ...func(*Options)) (*CreateLocationNfsOutput, error) {
	if params == nil {
		params = &CreateLocationNfsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLocationNfs", params, optFns, c.addOperationCreateLocationNfsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLocationNfsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateLocationNfsRequest
type CreateLocationNfsInput struct {

	// Specifies the Amazon Resource Name (ARN) of the DataSync agent that want to
	// connect to your NFS file server. You can specify more than one agent. For more
	// information, see Using multiple agents for transfers (https://docs.aws.amazon.com/datasync/latest/userguide/multiple-agents.html)
	// .
	//
	// This member is required.
	OnPremConfig *types.OnPremConfig

	// Specifies the Domain Name System (DNS) name or IP version 4 address of the NFS
	// file server that your DataSync agent connects to.
	//
	// This member is required.
	ServerHostname *string

	// Specifies the export path in your NFS file server that you want DataSync to
	// mount. This path (or a subdirectory of the path) is where DataSync transfers
	// data to or from. For information on configuring an export for DataSync, see
	// Accessing NFS file servers (https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html#accessing-nfs)
	// .
	//
	// This member is required.
	Subdirectory *string

	// Specifies the options that DataSync can use to mount your NFS file server.
	MountOptions *types.NfsMountOptions

	// Specifies labels that help you categorize, filter, and search for your Amazon
	// Web Services resources. We recommend creating at least a name tag for your
	// location.
	Tags []types.TagListEntry

	noSmithyDocumentSerde
}

// CreateLocationNfsResponse
type CreateLocationNfsOutput struct {

	// The ARN of the transfer location that you created for your NFS file server.
	LocationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLocationNfsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLocationNfs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLocationNfs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLocationNfs"); err != nil {
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
	if err = addOpCreateLocationNfsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocationNfs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLocationNfs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLocationNfs",
	}
}
