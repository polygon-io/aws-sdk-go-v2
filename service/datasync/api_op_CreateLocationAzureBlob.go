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

// Creates an endpoint for a Microsoft Azure Blob Storage container that DataSync
// can use as a transfer source or destination. Before you begin, make sure you
// know how DataSync accesses Azure Blob Storage (https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-access)
// and works with access tiers (https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-access-tiers)
// and blob types (https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#blob-types)
// . You also need a DataSync agent (https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-creating-agent)
// that can connect to your container.
func (c *Client) CreateLocationAzureBlob(ctx context.Context, params *CreateLocationAzureBlobInput, optFns ...func(*Options)) (*CreateLocationAzureBlobOutput, error) {
	if params == nil {
		params = &CreateLocationAzureBlobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLocationAzureBlob", params, optFns, c.addOperationCreateLocationAzureBlobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLocationAzureBlobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLocationAzureBlobInput struct {

	// Specifies the Amazon Resource Name (ARN) of the DataSync agent that can connect
	// with your Azure Blob Storage container. You can specify more than one agent. For
	// more information, see Using multiple agents for your transfer (https://docs.aws.amazon.com/datasync/latest/userguide/multiple-agents.html)
	// .
	//
	// This member is required.
	AgentArns []string

	// Specifies the authentication method DataSync uses to access your Azure Blob
	// Storage. DataSync can access blob storage using a shared access signature (SAS).
	//
	// This member is required.
	AuthenticationType types.AzureBlobAuthenticationType

	// Specifies the URL of the Azure Blob Storage container involved in your transfer.
	//
	// This member is required.
	ContainerUrl *string

	// Specifies the access tier that you want your objects or files transferred into.
	// This only applies when using the location as a transfer destination. For more
	// information, see Access tiers (https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-access-tiers)
	// .
	AccessTier types.AzureAccessTier

	// Specifies the type of blob that you want your objects or files to be when
	// transferring them into Azure Blob Storage. Currently, DataSync only supports
	// moving data into Azure Blob Storage as block blobs. For more information on blob
	// types, see the Azure Blob Storage documentation (https://learn.microsoft.com/en-us/rest/api/storageservices/understanding-block-blobs--append-blobs--and-page-blobs)
	// .
	BlobType types.AzureBlobType

	// Specifies the SAS configuration that allows DataSync to access your Azure Blob
	// Storage.
	SasConfiguration *types.AzureBlobSasConfiguration

	// Specifies path segments if you want to limit your transfer to a virtual
	// directory in your container (for example, /my/images ).
	Subdirectory *string

	// Specifies labels that help you categorize, filter, and search for your Amazon
	// Web Services resources. We recommend creating at least a name tag for your
	// transfer location.
	Tags []types.TagListEntry

	noSmithyDocumentSerde
}

type CreateLocationAzureBlobOutput struct {

	// The ARN of the Azure Blob Storage transfer location that you created.
	LocationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLocationAzureBlobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLocationAzureBlob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLocationAzureBlob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLocationAzureBlob"); err != nil {
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
	if err = addOpCreateLocationAzureBlobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocationAzureBlob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLocationAzureBlob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLocationAzureBlob",
	}
}
