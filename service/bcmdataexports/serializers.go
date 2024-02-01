// Code generated by smithy-go-codegen DO NOT EDIT.

package bcmdataexports

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/bcmdataexports/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"path"
)

type awsAwsjson11_serializeOpCreateExport struct {
}

func (*awsAwsjson11_serializeOpCreateExport) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpCreateExport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateExportInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSBillingAndCostManagementDataExports.CreateExport")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentCreateExportInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDeleteExport struct {
}

func (*awsAwsjson11_serializeOpDeleteExport) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDeleteExport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteExportInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSBillingAndCostManagementDataExports.DeleteExport")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDeleteExportInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpGetExecution struct {
}

func (*awsAwsjson11_serializeOpGetExecution) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpGetExecution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetExecutionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSBillingAndCostManagementDataExports.GetExecution")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentGetExecutionInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpGetExport struct {
}

func (*awsAwsjson11_serializeOpGetExport) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpGetExport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetExportInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSBillingAndCostManagementDataExports.GetExport")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentGetExportInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpGetTable struct {
}

func (*awsAwsjson11_serializeOpGetTable) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpGetTable) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetTableInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSBillingAndCostManagementDataExports.GetTable")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentGetTableInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpListExecutions struct {
}

func (*awsAwsjson11_serializeOpListExecutions) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpListExecutions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListExecutionsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSBillingAndCostManagementDataExports.ListExecutions")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentListExecutionsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpListExports struct {
}

func (*awsAwsjson11_serializeOpListExports) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpListExports) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListExportsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSBillingAndCostManagementDataExports.ListExports")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentListExportsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpListTables struct {
}

func (*awsAwsjson11_serializeOpListTables) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpListTables) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListTablesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSBillingAndCostManagementDataExports.ListTables")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentListTablesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpListTagsForResource struct {
}

func (*awsAwsjson11_serializeOpListTagsForResource) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpListTagsForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListTagsForResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSBillingAndCostManagementDataExports.ListTagsForResource")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentListTagsForResourceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpTagResource struct {
}

func (*awsAwsjson11_serializeOpTagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpTagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*TagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSBillingAndCostManagementDataExports.TagResource")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentTagResourceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpUntagResource struct {
}

func (*awsAwsjson11_serializeOpUntagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpUntagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UntagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSBillingAndCostManagementDataExports.UntagResource")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentUntagResourceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpUpdateExport struct {
}

func (*awsAwsjson11_serializeOpUpdateExport) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpUpdateExport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateExportInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSBillingAndCostManagementDataExports.UpdateExport")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentUpdateExportInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson11_serializeDocumentDataQuery(v *types.DataQuery, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.QueryStatement != nil {
		ok := object.Key("QueryStatement")
		ok.String(*v.QueryStatement)
	}

	if v.TableConfigurations != nil {
		ok := object.Key("TableConfigurations")
		if err := awsAwsjson11_serializeDocumentTableConfigurations(v.TableConfigurations, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeDocumentDestinationConfigurations(v *types.DestinationConfigurations, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.S3Destination != nil {
		ok := object.Key("S3Destination")
		if err := awsAwsjson11_serializeDocumentS3Destination(v.S3Destination, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeDocumentExport(v *types.Export, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DataQuery != nil {
		ok := object.Key("DataQuery")
		if err := awsAwsjson11_serializeDocumentDataQuery(v.DataQuery, ok); err != nil {
			return err
		}
	}

	if v.Description != nil {
		ok := object.Key("Description")
		ok.String(*v.Description)
	}

	if v.DestinationConfigurations != nil {
		ok := object.Key("DestinationConfigurations")
		if err := awsAwsjson11_serializeDocumentDestinationConfigurations(v.DestinationConfigurations, ok); err != nil {
			return err
		}
	}

	if v.ExportArn != nil {
		ok := object.Key("ExportArn")
		ok.String(*v.ExportArn)
	}

	if v.Name != nil {
		ok := object.Key("Name")
		ok.String(*v.Name)
	}

	if v.RefreshCadence != nil {
		ok := object.Key("RefreshCadence")
		if err := awsAwsjson11_serializeDocumentRefreshCadence(v.RefreshCadence, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeDocumentRefreshCadence(v *types.RefreshCadence, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.Frequency) > 0 {
		ok := object.Key("Frequency")
		ok.String(string(v.Frequency))
	}

	return nil
}

func awsAwsjson11_serializeDocumentResourceTag(v *types.ResourceTag, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Key != nil {
		ok := object.Key("Key")
		ok.String(*v.Key)
	}

	if v.Value != nil {
		ok := object.Key("Value")
		ok.String(*v.Value)
	}

	return nil
}

func awsAwsjson11_serializeDocumentResourceTagKeyList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentResourceTagList(v []types.ResourceTag, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentResourceTag(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentS3Destination(v *types.S3Destination, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.S3Bucket != nil {
		ok := object.Key("S3Bucket")
		ok.String(*v.S3Bucket)
	}

	if v.S3OutputConfigurations != nil {
		ok := object.Key("S3OutputConfigurations")
		if err := awsAwsjson11_serializeDocumentS3OutputConfigurations(v.S3OutputConfigurations, ok); err != nil {
			return err
		}
	}

	if v.S3Prefix != nil {
		ok := object.Key("S3Prefix")
		ok.String(*v.S3Prefix)
	}

	if v.S3Region != nil {
		ok := object.Key("S3Region")
		ok.String(*v.S3Region)
	}

	return nil
}

func awsAwsjson11_serializeDocumentS3OutputConfigurations(v *types.S3OutputConfigurations, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.Compression) > 0 {
		ok := object.Key("Compression")
		ok.String(string(v.Compression))
	}

	if len(v.Format) > 0 {
		ok := object.Key("Format")
		ok.String(string(v.Format))
	}

	if len(v.OutputType) > 0 {
		ok := object.Key("OutputType")
		ok.String(string(v.OutputType))
	}

	if len(v.Overwrite) > 0 {
		ok := object.Key("Overwrite")
		ok.String(string(v.Overwrite))
	}

	return nil
}

func awsAwsjson11_serializeDocumentTableConfigurations(v map[string]map[string]string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		if vv := v[key]; vv == nil {
			continue
		}
		if err := awsAwsjson11_serializeDocumentTableProperties(v[key], om); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentTableProperties(v map[string]string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		om.String(v[key])
	}
	return nil
}

func awsAwsjson11_serializeOpDocumentCreateExportInput(v *CreateExportInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Export != nil {
		ok := object.Key("Export")
		if err := awsAwsjson11_serializeDocumentExport(v.Export, ok); err != nil {
			return err
		}
	}

	if v.ResourceTags != nil {
		ok := object.Key("ResourceTags")
		if err := awsAwsjson11_serializeDocumentResourceTagList(v.ResourceTags, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDeleteExportInput(v *DeleteExportInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ExportArn != nil {
		ok := object.Key("ExportArn")
		ok.String(*v.ExportArn)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentGetExecutionInput(v *GetExecutionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ExecutionId != nil {
		ok := object.Key("ExecutionId")
		ok.String(*v.ExecutionId)
	}

	if v.ExportArn != nil {
		ok := object.Key("ExportArn")
		ok.String(*v.ExportArn)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentGetExportInput(v *GetExportInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ExportArn != nil {
		ok := object.Key("ExportArn")
		ok.String(*v.ExportArn)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentGetTableInput(v *GetTableInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.TableName != nil {
		ok := object.Key("TableName")
		ok.String(*v.TableName)
	}

	if v.TableProperties != nil {
		ok := object.Key("TableProperties")
		if err := awsAwsjson11_serializeDocumentTableProperties(v.TableProperties, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentListExecutionsInput(v *ListExecutionsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ExportArn != nil {
		ok := object.Key("ExportArn")
		ok.String(*v.ExportArn)
	}

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentListExportsInput(v *ListExportsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentListTablesInput(v *ListTablesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentListTagsForResourceInput(v *ListTagsForResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.ResourceArn != nil {
		ok := object.Key("ResourceArn")
		ok.String(*v.ResourceArn)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentTagResourceInput(v *TagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceArn != nil {
		ok := object.Key("ResourceArn")
		ok.String(*v.ResourceArn)
	}

	if v.ResourceTags != nil {
		ok := object.Key("ResourceTags")
		if err := awsAwsjson11_serializeDocumentResourceTagList(v.ResourceTags, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentUntagResourceInput(v *UntagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceArn != nil {
		ok := object.Key("ResourceArn")
		ok.String(*v.ResourceArn)
	}

	if v.ResourceTagKeys != nil {
		ok := object.Key("ResourceTagKeys")
		if err := awsAwsjson11_serializeDocumentResourceTagKeyList(v.ResourceTagKeys, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentUpdateExportInput(v *UpdateExportInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Export != nil {
		ok := object.Key("Export")
		if err := awsAwsjson11_serializeDocumentExport(v.Export, ok); err != nil {
			return err
		}
	}

	if v.ExportArn != nil {
		ok := object.Key("ExportArn")
		ok.String(*v.ExportArn)
	}

	return nil
}
