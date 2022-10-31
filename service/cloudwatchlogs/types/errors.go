// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The event was already logged.
type DataAlreadyAcceptedException struct {
	Message *string

	ExpectedSequenceToken *string

	noSmithyDocumentSerde
}

func (e *DataAlreadyAcceptedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DataAlreadyAcceptedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DataAlreadyAcceptedException) ErrorCode() string             { return "DataAlreadyAcceptedException" }
func (e *DataAlreadyAcceptedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation is not valid on the specified resource.
type InvalidOperationException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidOperationException) ErrorCode() string             { return "InvalidOperationException" }
func (e *InvalidOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A parameter is specified incorrectly.
type InvalidParameterException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string             { return "InvalidParameterException" }
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The sequence token is not valid. You can get the correct sequence token in the
// expectedSequenceToken field in the InvalidSequenceTokenException message.
type InvalidSequenceTokenException struct {
	Message *string

	ExpectedSequenceToken *string

	noSmithyDocumentSerde
}

func (e *InvalidSequenceTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSequenceTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSequenceTokenException) ErrorCode() string             { return "InvalidSequenceTokenException" }
func (e *InvalidSequenceTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have reached the maximum number of resources that can be created.
type LimitExceededException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The query string is not valid. Details about this error are displayed in a
// QueryCompileError object. For more information, see QueryCompileError
// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_QueryCompileError.html).
// For more information about valid query syntax, see CloudWatch Logs Insights
// Query Syntax
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
type MalformedQueryException struct {
	Message *string

	QueryCompileError *QueryCompileError

	noSmithyDocumentSerde
}

func (e *MalformedQueryException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MalformedQueryException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MalformedQueryException) ErrorCode() string             { return "MalformedQueryException" }
func (e *MalformedQueryException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Multiple concurrent requests to update the same resource were in conflict.
type OperationAbortedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *OperationAbortedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OperationAbortedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OperationAbortedException) ErrorCode() string             { return "OperationAbortedException" }
func (e *OperationAbortedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource already exists.
type ResourceAlreadyExistsException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ResourceAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceAlreadyExistsException) ErrorCode() string             { return "ResourceAlreadyExistsException" }
func (e *ResourceAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource does not exist.
type ResourceNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service cannot complete the request.
type ServiceUnavailableException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ServiceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailableException) ErrorCode() string             { return "ServiceUnavailableException" }
func (e *ServiceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// A resource can have no more than 50 tags.
type TooManyTagsException struct {
	Message *string

	ResourceName *string

	noSmithyDocumentSerde
}

func (e *TooManyTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTagsException) ErrorCode() string             { return "TooManyTagsException" }
func (e *TooManyTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The most likely cause is an invalid Amazon Web Services access key ID or secret
// key.
type UnrecognizedClientException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UnrecognizedClientException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnrecognizedClientException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnrecognizedClientException) ErrorCode() string             { return "UnrecognizedClientException" }
func (e *UnrecognizedClientException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
