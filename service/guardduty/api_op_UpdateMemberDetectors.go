// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Contains information on member accounts to be updated. There might be regional
// differences because some data sources might not be available in all the Amazon
// Web Services Regions where GuardDuty is presently supported. For more
// information, see Regions and endpoints (https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html)
// .
func (c *Client) UpdateMemberDetectors(ctx context.Context, params *UpdateMemberDetectorsInput, optFns ...func(*Options)) (*UpdateMemberDetectorsOutput, error) {
	if params == nil {
		params = &UpdateMemberDetectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMemberDetectors", params, optFns, c.addOperationUpdateMemberDetectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMemberDetectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMemberDetectorsInput struct {

	// A list of member account IDs to be updated.
	//
	// This member is required.
	AccountIds []string

	// The detector ID of the administrator account.
	//
	// This member is required.
	DetectorId *string

	// Describes which data sources will be updated.
	//
	// Deprecated: This parameter is deprecated, use Features instead
	DataSources *types.DataSourceConfigurations

	// A list of features that will be updated for the specified member accounts.
	Features []types.MemberFeaturesConfiguration

	noSmithyDocumentSerde
}

type UpdateMemberDetectorsOutput struct {

	// A list of member account IDs that were unable to be processed along with an
	// explanation for why they were not processed.
	//
	// This member is required.
	UnprocessedAccounts []types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMemberDetectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateMemberDetectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateMemberDetectors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateMemberDetectors"); err != nil {
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
	if err = addOpUpdateMemberDetectorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMemberDetectors(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateMemberDetectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateMemberDetectors",
	}
}
