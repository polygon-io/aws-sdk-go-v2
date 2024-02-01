// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the enrollment (opt in) status of an account to the Compute Optimizer
// service. If the account is the management account of an organization, this
// action also confirms the enrollment status of member accounts of the
// organization. Use the GetEnrollmentStatusesForOrganization action to get
// detailed information about the enrollment status of member accounts of an
// organization.
func (c *Client) GetEnrollmentStatus(ctx context.Context, params *GetEnrollmentStatusInput, optFns ...func(*Options)) (*GetEnrollmentStatusOutput, error) {
	if params == nil {
		params = &GetEnrollmentStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEnrollmentStatus", params, optFns, c.addOperationGetEnrollmentStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEnrollmentStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEnrollmentStatusInput struct {
	noSmithyDocumentSerde
}

type GetEnrollmentStatusOutput struct {

	// The Unix epoch timestamp, in seconds, of when the account enrollment status was
	// last updated.
	LastUpdatedTimestamp *time.Time

	// Confirms the enrollment status of member accounts of the organization, if the
	// account is a management account of an organization.
	MemberAccountsEnrolled bool

	// The count of organization member accounts that are opted in to the service, if
	// your account is an organization management account.
	NumberOfMemberAccountsOptedIn *int32

	// The enrollment status of the account.
	Status types.Status

	// The reason for the enrollment status of the account. For example, an account
	// might show a status of Pending because member accounts of an organization
	// require more time to be enrolled in the service.
	StatusReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEnrollmentStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetEnrollmentStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetEnrollmentStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetEnrollmentStatus"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEnrollmentStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEnrollmentStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetEnrollmentStatus",
	}
}
