// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a new version to the existing schema. Returns an error if new version of
// schema does not meet the compatibility requirements of the schema set. This API
// will not create a new schema set and will return a 404 error if the schema set
// is not already present in the Schema Registry. If this is the first schema
// definition to be registered in the Schema Registry, this API will store the
// schema version and return immediately. Otherwise, this call has the potential to
// run longer than other operations due to compatibility modes. You can call the
// GetSchemaVersion API with the SchemaVersionId to check compatibility modes. If
// the same schema definition is already stored in Schema Registry as a version,
// the schema ID of the existing schema is returned to the caller.
func (c *Client) RegisterSchemaVersion(ctx context.Context, params *RegisterSchemaVersionInput, optFns ...func(*Options)) (*RegisterSchemaVersionOutput, error) {
	if params == nil {
		params = &RegisterSchemaVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterSchemaVersion", params, optFns, c.addOperationRegisterSchemaVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterSchemaVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterSchemaVersionInput struct {

	// The schema definition using the DataFormat setting for the SchemaName .
	//
	// This member is required.
	SchemaDefinition *string

	// This is a wrapper structure to contain schema identity fields. The structure
	// contains:
	//   - SchemaId$SchemaArn: The Amazon Resource Name (ARN) of the schema. Either
	//   SchemaArn or SchemaName and RegistryName has to be provided.
	//   - SchemaId$SchemaName: The name of the schema. Either SchemaArn or SchemaName
	//   and RegistryName has to be provided.
	//
	// This member is required.
	SchemaId *types.SchemaId

	noSmithyDocumentSerde
}

type RegisterSchemaVersionOutput struct {

	// The unique ID that represents the version of this schema.
	SchemaVersionId *string

	// The status of the schema version.
	Status types.SchemaVersionStatus

	// The version of this schema (for sync flow only, in case this is the first
	// version).
	VersionNumber *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterSchemaVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterSchemaVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterSchemaVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterSchemaVersion"); err != nil {
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
	if err = addOpRegisterSchemaVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterSchemaVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterSchemaVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterSchemaVersion",
	}
}
