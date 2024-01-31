// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// List item for keys to delete.
type DeleteKeyRequestListItem struct {

	// The key of the key value pair to be deleted.
	//
	// This member is required.
	Key *string

	noSmithyDocumentSerde
}

// A key value pair.
type ListKeysResponseListItem struct {

	// The key of the key value pair.
	//
	// This member is required.
	Key *string

	// The value of the key value pair.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// List item for key value pair to put.
type PutKeyRequestListItem struct {

	// The key of the key value pair list item to put.
	//
	// This member is required.
	Key *string

	// The value for the key value pair to put.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
