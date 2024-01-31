// Code generated by smithy-go-codegen DO NOT EDIT.

package finspace

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/finspace/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves details of the dataview.
func (c *Client) GetKxDataview(ctx context.Context, params *GetKxDataviewInput, optFns ...func(*Options)) (*GetKxDataviewOutput, error) {
	if params == nil {
		params = &GetKxDataviewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetKxDataview", params, optFns, c.addOperationGetKxDataviewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetKxDataviewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetKxDataviewInput struct {

	// The name of the database where you created the dataview.
	//
	// This member is required.
	DatabaseName *string

	// A unique identifier for the dataview.
	//
	// This member is required.
	DataviewName *string

	// A unique identifier for the kdb environment, from where you want to retrieve
	// the dataview details.
	//
	// This member is required.
	EnvironmentId *string

	noSmithyDocumentSerde
}

type GetKxDataviewOutput struct {

	// The current active changeset versions of the database on the given dataview.
	ActiveVersions []types.KxDataviewActiveVersion

	// The option to specify whether you want to apply all the future additions and
	// corrections automatically to the dataview when new changesets are ingested. The
	// default value is false.
	AutoUpdate bool

	// The identifier of the availability zones.
	AvailabilityZoneId *string

	// The number of availability zones you want to assign per cluster. This can be
	// one of the following
	//   - SINGLE – Assigns one availability zone per cluster.
	//   - MULTI – Assigns all the availability zones per cluster.
	AzMode types.KxAzMode

	// A unique identifier of the changeset that you want to use to ingest data.
	ChangesetId *string

	// The timestamp at which the dataview was created in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	CreatedTimestamp *time.Time

	// The name of the database where you created the dataview.
	DatabaseName *string

	// A unique identifier for the dataview.
	DataviewName *string

	// A description of the dataview.
	Description *string

	// A unique identifier for the kdb environment, from where you want to retrieve
	// the dataview details.
	EnvironmentId *string

	// The last time that the dataview was updated in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	LastModifiedTimestamp *time.Time

	// The configuration that contains the database path of the data that you want to
	// place on each selected volume. Each segment must have a unique database path for
	// each volume. If you do not explicitly specify any database path for a volume,
	// they are accessible from the cluster through the default S3/object store
	// segment.
	SegmentConfigurations []types.KxDataviewSegmentConfiguration

	// The status of dataview creation.
	//   - CREATING – The dataview creation is in progress.
	//   - UPDATING – The dataview is in the process of being updated.
	//   - ACTIVE – The dataview is active.
	Status types.KxDataviewStatus

	// The error message when a failed state occurs.
	StatusReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetKxDataviewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetKxDataview{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetKxDataview{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetKxDataview"); err != nil {
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
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
	if err = addOpGetKxDataviewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetKxDataview(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetKxDataview(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetKxDataview",
	}
}
