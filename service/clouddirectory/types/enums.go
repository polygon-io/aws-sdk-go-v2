// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type BatchReadExceptionType string

// Enum values for BatchReadExceptionType
const (
	BatchReadExceptionTypeValidationException             BatchReadExceptionType = "ValidationException"
	BatchReadExceptionTypeInvalidArnException             BatchReadExceptionType = "InvalidArnException"
	BatchReadExceptionTypeResourceNotFoundException       BatchReadExceptionType = "ResourceNotFoundException"
	BatchReadExceptionTypeInvalidNextTokenException       BatchReadExceptionType = "InvalidNextTokenException"
	BatchReadExceptionTypeAccessDeniedException           BatchReadExceptionType = "AccessDeniedException"
	BatchReadExceptionTypeNotNodeException                BatchReadExceptionType = "NotNodeException"
	BatchReadExceptionTypeFacetValidationException        BatchReadExceptionType = "FacetValidationException"
	BatchReadExceptionTypeCannotListParentOfRootException BatchReadExceptionType = "CannotListParentOfRootException"
	BatchReadExceptionTypeNotIndexException               BatchReadExceptionType = "NotIndexException"
	BatchReadExceptionTypeNotPolicyException              BatchReadExceptionType = "NotPolicyException"
	BatchReadExceptionTypeDirectoryNotEnabledException    BatchReadExceptionType = "DirectoryNotEnabledException"
	BatchReadExceptionTypeLimitExceededException          BatchReadExceptionType = "LimitExceededException"
	BatchReadExceptionTypeInternalServiceException        BatchReadExceptionType = "InternalServiceException"
)

// Values returns all known values for BatchReadExceptionType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BatchReadExceptionType) Values() []BatchReadExceptionType {
	return []BatchReadExceptionType{
		"ValidationException",
		"InvalidArnException",
		"ResourceNotFoundException",
		"InvalidNextTokenException",
		"AccessDeniedException",
		"NotNodeException",
		"FacetValidationException",
		"CannotListParentOfRootException",
		"NotIndexException",
		"NotPolicyException",
		"DirectoryNotEnabledException",
		"LimitExceededException",
		"InternalServiceException",
	}
}

type BatchWriteExceptionType string

// Enum values for BatchWriteExceptionType
const (
	BatchWriteExceptionTypeInternalServiceException         BatchWriteExceptionType = "InternalServiceException"
	BatchWriteExceptionTypeValidationException              BatchWriteExceptionType = "ValidationException"
	BatchWriteExceptionTypeInvalidArnException              BatchWriteExceptionType = "InvalidArnException"
	BatchWriteExceptionTypeLinkNameAlreadyInUseException    BatchWriteExceptionType = "LinkNameAlreadyInUseException"
	BatchWriteExceptionTypeStillContainsLinksException      BatchWriteExceptionType = "StillContainsLinksException"
	BatchWriteExceptionTypeFacetValidationException         BatchWriteExceptionType = "FacetValidationException"
	BatchWriteExceptionTypeObjectNotDetachedException       BatchWriteExceptionType = "ObjectNotDetachedException"
	BatchWriteExceptionTypeResourceNotFoundException        BatchWriteExceptionType = "ResourceNotFoundException"
	BatchWriteExceptionTypeAccessDeniedException            BatchWriteExceptionType = "AccessDeniedException"
	BatchWriteExceptionTypeInvalidAttachmentException       BatchWriteExceptionType = "InvalidAttachmentException"
	BatchWriteExceptionTypeNotIndexException                BatchWriteExceptionType = "NotIndexException"
	BatchWriteExceptionTypeNotNodeException                 BatchWriteExceptionType = "NotNodeException"
	BatchWriteExceptionTypeIndexedAttributeMissingException BatchWriteExceptionType = "IndexedAttributeMissingException"
	BatchWriteExceptionTypeObjectAlreadyDetachedException   BatchWriteExceptionType = "ObjectAlreadyDetachedException"
	BatchWriteExceptionTypeNotPolicyException               BatchWriteExceptionType = "NotPolicyException"
	BatchWriteExceptionTypeDirectoryNotEnabledException     BatchWriteExceptionType = "DirectoryNotEnabledException"
	BatchWriteExceptionTypeLimitExceededException           BatchWriteExceptionType = "LimitExceededException"
	BatchWriteExceptionTypeUnsupportedIndexTypeException    BatchWriteExceptionType = "UnsupportedIndexTypeException"
)

// Values returns all known values for BatchWriteExceptionType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BatchWriteExceptionType) Values() []BatchWriteExceptionType {
	return []BatchWriteExceptionType{
		"InternalServiceException",
		"ValidationException",
		"InvalidArnException",
		"LinkNameAlreadyInUseException",
		"StillContainsLinksException",
		"FacetValidationException",
		"ObjectNotDetachedException",
		"ResourceNotFoundException",
		"AccessDeniedException",
		"InvalidAttachmentException",
		"NotIndexException",
		"NotNodeException",
		"IndexedAttributeMissingException",
		"ObjectAlreadyDetachedException",
		"NotPolicyException",
		"DirectoryNotEnabledException",
		"LimitExceededException",
		"UnsupportedIndexTypeException",
	}
}

type ConsistencyLevel string

// Enum values for ConsistencyLevel
const (
	ConsistencyLevelSerializable ConsistencyLevel = "SERIALIZABLE"
	ConsistencyLevelEventual     ConsistencyLevel = "EVENTUAL"
)

// Values returns all known values for ConsistencyLevel. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConsistencyLevel) Values() []ConsistencyLevel {
	return []ConsistencyLevel{
		"SERIALIZABLE",
		"EVENTUAL",
	}
}

type DirectoryState string

// Enum values for DirectoryState
const (
	DirectoryStateEnabled  DirectoryState = "ENABLED"
	DirectoryStateDisabled DirectoryState = "DISABLED"
	DirectoryStateDeleted  DirectoryState = "DELETED"
)

// Values returns all known values for DirectoryState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DirectoryState) Values() []DirectoryState {
	return []DirectoryState{
		"ENABLED",
		"DISABLED",
		"DELETED",
	}
}

type FacetAttributeType string

// Enum values for FacetAttributeType
const (
	FacetAttributeTypeString   FacetAttributeType = "STRING"
	FacetAttributeTypeBinary   FacetAttributeType = "BINARY"
	FacetAttributeTypeBoolean  FacetAttributeType = "BOOLEAN"
	FacetAttributeTypeNumber   FacetAttributeType = "NUMBER"
	FacetAttributeTypeDatetime FacetAttributeType = "DATETIME"
	FacetAttributeTypeVariant  FacetAttributeType = "VARIANT"
)

// Values returns all known values for FacetAttributeType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FacetAttributeType) Values() []FacetAttributeType {
	return []FacetAttributeType{
		"STRING",
		"BINARY",
		"BOOLEAN",
		"NUMBER",
		"DATETIME",
		"VARIANT",
	}
}

type FacetStyle string

// Enum values for FacetStyle
const (
	FacetStyleStatic  FacetStyle = "STATIC"
	FacetStyleDynamic FacetStyle = "DYNAMIC"
)

// Values returns all known values for FacetStyle. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FacetStyle) Values() []FacetStyle {
	return []FacetStyle{
		"STATIC",
		"DYNAMIC",
	}
}

type ObjectType string

// Enum values for ObjectType
const (
	ObjectTypeNode     ObjectType = "NODE"
	ObjectTypeLeafNode ObjectType = "LEAF_NODE"
	ObjectTypePolicy   ObjectType = "POLICY"
	ObjectTypeIndex    ObjectType = "INDEX"
)

// Values returns all known values for ObjectType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ObjectType) Values() []ObjectType {
	return []ObjectType{
		"NODE",
		"LEAF_NODE",
		"POLICY",
		"INDEX",
	}
}

type RangeMode string

// Enum values for RangeMode
const (
	RangeModeFirst                   RangeMode = "FIRST"
	RangeModeLast                    RangeMode = "LAST"
	RangeModeLastBeforeMissingValues RangeMode = "LAST_BEFORE_MISSING_VALUES"
	RangeModeInclusive               RangeMode = "INCLUSIVE"
	RangeModeExclusive               RangeMode = "EXCLUSIVE"
)

// Values returns all known values for RangeMode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (RangeMode) Values() []RangeMode {
	return []RangeMode{
		"FIRST",
		"LAST",
		"LAST_BEFORE_MISSING_VALUES",
		"INCLUSIVE",
		"EXCLUSIVE",
	}
}

type RequiredAttributeBehavior string

// Enum values for RequiredAttributeBehavior
const (
	RequiredAttributeBehaviorRequiredAlways RequiredAttributeBehavior = "REQUIRED_ALWAYS"
	RequiredAttributeBehaviorNotRequired    RequiredAttributeBehavior = "NOT_REQUIRED"
)

// Values returns all known values for RequiredAttributeBehavior. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (RequiredAttributeBehavior) Values() []RequiredAttributeBehavior {
	return []RequiredAttributeBehavior{
		"REQUIRED_ALWAYS",
		"NOT_REQUIRED",
	}
}

type RuleType string

// Enum values for RuleType
const (
	RuleTypeBinaryLength     RuleType = "BINARY_LENGTH"
	RuleTypeNumberComparison RuleType = "NUMBER_COMPARISON"
	RuleTypeStringFromSet    RuleType = "STRING_FROM_SET"
	RuleTypeStringLength     RuleType = "STRING_LENGTH"
)

// Values returns all known values for RuleType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (RuleType) Values() []RuleType {
	return []RuleType{
		"BINARY_LENGTH",
		"NUMBER_COMPARISON",
		"STRING_FROM_SET",
		"STRING_LENGTH",
	}
}

type UpdateActionType string

// Enum values for UpdateActionType
const (
	UpdateActionTypeCreateOrUpdate UpdateActionType = "CREATE_OR_UPDATE"
	UpdateActionTypeDelete         UpdateActionType = "DELETE"
)

// Values returns all known values for UpdateActionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UpdateActionType) Values() []UpdateActionType {
	return []UpdateActionType{
		"CREATE_OR_UPDATE",
		"DELETE",
	}
}
