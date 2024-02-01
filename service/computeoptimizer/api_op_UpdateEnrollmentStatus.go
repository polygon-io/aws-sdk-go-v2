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
)

// Updates the enrollment (opt in and opt out) status of an account to the Compute
// Optimizer service. If the account is a management account of an organization,
// this action can also be used to enroll member accounts of the organization. You
// must have the appropriate permissions to opt in to Compute Optimizer, to view
// its recommendations, and to opt out. For more information, see Controlling
// access with Amazon Web Services Identity and Access Management (https://docs.aws.amazon.com/compute-optimizer/latest/ug/security-iam.html)
// in the Compute Optimizer User Guide. When you opt in, Compute Optimizer
// automatically creates a service-linked role in your account to access its data.
// For more information, see Using Service-Linked Roles for Compute Optimizer (https://docs.aws.amazon.com/compute-optimizer/latest/ug/using-service-linked-roles.html)
// in the Compute Optimizer User Guide.
func (c *Client) UpdateEnrollmentStatus(ctx context.Context, params *UpdateEnrollmentStatusInput, optFns ...func(*Options)) (*UpdateEnrollmentStatusOutput, error) {
	if params == nil {
		params = &UpdateEnrollmentStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEnrollmentStatus", params, optFns, c.addOperationUpdateEnrollmentStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEnrollmentStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEnrollmentStatusInput struct {

	// The new enrollment status of the account. The following status options are
	// available:
	//   - Active - Opts in your account to the Compute Optimizer service. Compute
	//   Optimizer begins analyzing the configuration and utilization metrics of your
	//   Amazon Web Services resources after you opt in. For more information, see
	//   Metrics analyzed by Compute Optimizer (https://docs.aws.amazon.com/compute-optimizer/latest/ug/metrics.html)
	//   in the Compute Optimizer User Guide.
	//   - Inactive - Opts out your account from the Compute Optimizer service. Your
	//   account's recommendations and related metrics data will be deleted from Compute
	//   Optimizer after you opt out.
	// The Pending and Failed options cannot be used to update the enrollment status
	// of an account. They are returned in the response of a request to update the
	// enrollment status of an account.
	//
	// This member is required.
	Status types.Status

	// Indicates whether to enroll member accounts of the organization if the account
	// is the management account of an organization.
	IncludeMemberAccounts bool

	noSmithyDocumentSerde
}

type UpdateEnrollmentStatusOutput struct {

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

func (c *Client) addOperationUpdateEnrollmentStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateEnrollmentStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateEnrollmentStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateEnrollmentStatus"); err != nil {
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
	if err = addOpUpdateEnrollmentStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEnrollmentStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEnrollmentStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateEnrollmentStatus",
	}
}
