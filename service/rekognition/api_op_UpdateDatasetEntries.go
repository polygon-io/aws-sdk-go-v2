// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation applies only to Amazon Rekognition Custom Labels. Adds or
// updates one or more entries (images) in a dataset. An entry is a JSON Line which
// contains the information for a single image, including the image location,
// assigned labels, and object location bounding boxes. For more information, see
// Image-Level labels in manifest files and Object localization in manifest files
// in the Amazon Rekognition Custom Labels Developer Guide. If the source-ref
// field in the JSON line references an existing image, the existing image in the
// dataset is updated. If source-ref field doesn't reference an existing image,
// the image is added as a new image to the dataset. You specify the changes that
// you want to make in the Changes input parameter. There isn't a limit to the
// number JSON Lines that you can change, but the size of Changes must be less
// than 5MB. UpdateDatasetEntries returns immediatly, but the dataset update might
// take a while to complete. Use DescribeDataset to check the current status. The
// dataset updated successfully if the value of Status is UPDATE_COMPLETE . To
// check if any non-terminal errors occured, call ListDatasetEntries and check for
// the presence of errors lists in the JSON Lines. Dataset update fails if a
// terminal error occurs ( Status = UPDATE_FAILED ). Currently, you can't access
// the terminal error information from the Amazon Rekognition Custom Labels SDK.
// This operation requires permissions to perform the
// rekognition:UpdateDatasetEntries action.
func (c *Client) UpdateDatasetEntries(ctx context.Context, params *UpdateDatasetEntriesInput, optFns ...func(*Options)) (*UpdateDatasetEntriesOutput, error) {
	if params == nil {
		params = &UpdateDatasetEntriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDatasetEntries", params, optFns, c.addOperationUpdateDatasetEntriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDatasetEntriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDatasetEntriesInput struct {

	// The changes that you want to make to the dataset.
	//
	// This member is required.
	Changes *types.DatasetChanges

	// The Amazon Resource Name (ARN) of the dataset that you want to update.
	//
	// This member is required.
	DatasetArn *string

	noSmithyDocumentSerde
}

type UpdateDatasetEntriesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDatasetEntriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDatasetEntries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDatasetEntries{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDatasetEntries"); err != nil {
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
	if err = addOpUpdateDatasetEntriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDatasetEntries(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDatasetEntries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDatasetEntries",
	}
}
