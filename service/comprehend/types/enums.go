// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AugmentedManifestsDocumentTypeFormat string

// Enum values for AugmentedManifestsDocumentTypeFormat
const (
	AugmentedManifestsDocumentTypeFormatPlainTextDocument      AugmentedManifestsDocumentTypeFormat = "PLAIN_TEXT_DOCUMENT"
	AugmentedManifestsDocumentTypeFormatSemiStructuredDocument AugmentedManifestsDocumentTypeFormat = "SEMI_STRUCTURED_DOCUMENT"
)

// Values returns all known values for AugmentedManifestsDocumentTypeFormat. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (AugmentedManifestsDocumentTypeFormat) Values() []AugmentedManifestsDocumentTypeFormat {
	return []AugmentedManifestsDocumentTypeFormat{
		"PLAIN_TEXT_DOCUMENT",
		"SEMI_STRUCTURED_DOCUMENT",
	}
}

type BlockType string

// Enum values for BlockType
const (
	BlockTypeLine BlockType = "LINE"
	BlockTypeWord BlockType = "WORD"
)

// Values returns all known values for BlockType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (BlockType) Values() []BlockType {
	return []BlockType{
		"LINE",
		"WORD",
	}
}

type DatasetDataFormat string

// Enum values for DatasetDataFormat
const (
	DatasetDataFormatComprehendCsv     DatasetDataFormat = "COMPREHEND_CSV"
	DatasetDataFormatAugmentedManifest DatasetDataFormat = "AUGMENTED_MANIFEST"
)

// Values returns all known values for DatasetDataFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DatasetDataFormat) Values() []DatasetDataFormat {
	return []DatasetDataFormat{
		"COMPREHEND_CSV",
		"AUGMENTED_MANIFEST",
	}
}

type DatasetStatus string

// Enum values for DatasetStatus
const (
	DatasetStatusCreating  DatasetStatus = "CREATING"
	DatasetStatusCompleted DatasetStatus = "COMPLETED"
	DatasetStatusFailed    DatasetStatus = "FAILED"
)

// Values returns all known values for DatasetStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DatasetStatus) Values() []DatasetStatus {
	return []DatasetStatus{
		"CREATING",
		"COMPLETED",
		"FAILED",
	}
}

type DatasetType string

// Enum values for DatasetType
const (
	DatasetTypeTrain DatasetType = "TRAIN"
	DatasetTypeTest  DatasetType = "TEST"
)

// Values returns all known values for DatasetType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DatasetType) Values() []DatasetType {
	return []DatasetType{
		"TRAIN",
		"TEST",
	}
}

type DocumentClassifierDataFormat string

// Enum values for DocumentClassifierDataFormat
const (
	DocumentClassifierDataFormatComprehendCsv     DocumentClassifierDataFormat = "COMPREHEND_CSV"
	DocumentClassifierDataFormatAugmentedManifest DocumentClassifierDataFormat = "AUGMENTED_MANIFEST"
)

// Values returns all known values for DocumentClassifierDataFormat. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (DocumentClassifierDataFormat) Values() []DocumentClassifierDataFormat {
	return []DocumentClassifierDataFormat{
		"COMPREHEND_CSV",
		"AUGMENTED_MANIFEST",
	}
}

type DocumentClassifierDocumentTypeFormat string

// Enum values for DocumentClassifierDocumentTypeFormat
const (
	DocumentClassifierDocumentTypeFormatPlainTextDocument      DocumentClassifierDocumentTypeFormat = "PLAIN_TEXT_DOCUMENT"
	DocumentClassifierDocumentTypeFormatSemiStructuredDocument DocumentClassifierDocumentTypeFormat = "SEMI_STRUCTURED_DOCUMENT"
)

// Values returns all known values for DocumentClassifierDocumentTypeFormat. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (DocumentClassifierDocumentTypeFormat) Values() []DocumentClassifierDocumentTypeFormat {
	return []DocumentClassifierDocumentTypeFormat{
		"PLAIN_TEXT_DOCUMENT",
		"SEMI_STRUCTURED_DOCUMENT",
	}
}

type DocumentClassifierMode string

// Enum values for DocumentClassifierMode
const (
	DocumentClassifierModeMultiClass DocumentClassifierMode = "MULTI_CLASS"
	DocumentClassifierModeMultiLabel DocumentClassifierMode = "MULTI_LABEL"
)

// Values returns all known values for DocumentClassifierMode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DocumentClassifierMode) Values() []DocumentClassifierMode {
	return []DocumentClassifierMode{
		"MULTI_CLASS",
		"MULTI_LABEL",
	}
}

type DocumentReadAction string

// Enum values for DocumentReadAction
const (
	DocumentReadActionTextractDetectDocumentText DocumentReadAction = "TEXTRACT_DETECT_DOCUMENT_TEXT"
	DocumentReadActionTextractAnalyzeDocument    DocumentReadAction = "TEXTRACT_ANALYZE_DOCUMENT"
)

// Values returns all known values for DocumentReadAction. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DocumentReadAction) Values() []DocumentReadAction {
	return []DocumentReadAction{
		"TEXTRACT_DETECT_DOCUMENT_TEXT",
		"TEXTRACT_ANALYZE_DOCUMENT",
	}
}

type DocumentReadFeatureTypes string

// Enum values for DocumentReadFeatureTypes
const (
	DocumentReadFeatureTypesTables DocumentReadFeatureTypes = "TABLES"
	DocumentReadFeatureTypesForms  DocumentReadFeatureTypes = "FORMS"
)

// Values returns all known values for DocumentReadFeatureTypes. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DocumentReadFeatureTypes) Values() []DocumentReadFeatureTypes {
	return []DocumentReadFeatureTypes{
		"TABLES",
		"FORMS",
	}
}

type DocumentReadMode string

// Enum values for DocumentReadMode
const (
	DocumentReadModeServiceDefault          DocumentReadMode = "SERVICE_DEFAULT"
	DocumentReadModeForceDocumentReadAction DocumentReadMode = "FORCE_DOCUMENT_READ_ACTION"
)

// Values returns all known values for DocumentReadMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DocumentReadMode) Values() []DocumentReadMode {
	return []DocumentReadMode{
		"SERVICE_DEFAULT",
		"FORCE_DOCUMENT_READ_ACTION",
	}
}

type DocumentType string

// Enum values for DocumentType
const (
	DocumentTypeNativePdf                      DocumentType = "NATIVE_PDF"
	DocumentTypeScannedPdf                     DocumentType = "SCANNED_PDF"
	DocumentTypeMsWord                         DocumentType = "MS_WORD"
	DocumentTypeImage                          DocumentType = "IMAGE"
	DocumentTypePlainText                      DocumentType = "PLAIN_TEXT"
	DocumentTypeTextractDetectDocumentTextJson DocumentType = "TEXTRACT_DETECT_DOCUMENT_TEXT_JSON"
	DocumentTypeTextractAnalyzeDocumentJson    DocumentType = "TEXTRACT_ANALYZE_DOCUMENT_JSON"
)

// Values returns all known values for DocumentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DocumentType) Values() []DocumentType {
	return []DocumentType{
		"NATIVE_PDF",
		"SCANNED_PDF",
		"MS_WORD",
		"IMAGE",
		"PLAIN_TEXT",
		"TEXTRACT_DETECT_DOCUMENT_TEXT_JSON",
		"TEXTRACT_ANALYZE_DOCUMENT_JSON",
	}
}

type EndpointStatus string

// Enum values for EndpointStatus
const (
	EndpointStatusCreating  EndpointStatus = "CREATING"
	EndpointStatusDeleting  EndpointStatus = "DELETING"
	EndpointStatusFailed    EndpointStatus = "FAILED"
	EndpointStatusInService EndpointStatus = "IN_SERVICE"
	EndpointStatusUpdating  EndpointStatus = "UPDATING"
)

// Values returns all known values for EndpointStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EndpointStatus) Values() []EndpointStatus {
	return []EndpointStatus{
		"CREATING",
		"DELETING",
		"FAILED",
		"IN_SERVICE",
		"UPDATING",
	}
}

type EntityRecognizerDataFormat string

// Enum values for EntityRecognizerDataFormat
const (
	EntityRecognizerDataFormatComprehendCsv     EntityRecognizerDataFormat = "COMPREHEND_CSV"
	EntityRecognizerDataFormatAugmentedManifest EntityRecognizerDataFormat = "AUGMENTED_MANIFEST"
)

// Values returns all known values for EntityRecognizerDataFormat. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (EntityRecognizerDataFormat) Values() []EntityRecognizerDataFormat {
	return []EntityRecognizerDataFormat{
		"COMPREHEND_CSV",
		"AUGMENTED_MANIFEST",
	}
}

type EntityType string

// Enum values for EntityType
const (
	EntityTypePerson         EntityType = "PERSON"
	EntityTypeLocation       EntityType = "LOCATION"
	EntityTypeOrganization   EntityType = "ORGANIZATION"
	EntityTypeCommercialItem EntityType = "COMMERCIAL_ITEM"
	EntityTypeEvent          EntityType = "EVENT"
	EntityTypeDate           EntityType = "DATE"
	EntityTypeQuantity       EntityType = "QUANTITY"
	EntityTypeTitle          EntityType = "TITLE"
	EntityTypeOther          EntityType = "OTHER"
)

// Values returns all known values for EntityType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EntityType) Values() []EntityType {
	return []EntityType{
		"PERSON",
		"LOCATION",
		"ORGANIZATION",
		"COMMERCIAL_ITEM",
		"EVENT",
		"DATE",
		"QUANTITY",
		"TITLE",
		"OTHER",
	}
}

type FlywheelIterationStatus string

// Enum values for FlywheelIterationStatus
const (
	FlywheelIterationStatusTraining      FlywheelIterationStatus = "TRAINING"
	FlywheelIterationStatusEvaluating    FlywheelIterationStatus = "EVALUATING"
	FlywheelIterationStatusCompleted     FlywheelIterationStatus = "COMPLETED"
	FlywheelIterationStatusFailed        FlywheelIterationStatus = "FAILED"
	FlywheelIterationStatusStopRequested FlywheelIterationStatus = "STOP_REQUESTED"
	FlywheelIterationStatusStopped       FlywheelIterationStatus = "STOPPED"
)

// Values returns all known values for FlywheelIterationStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FlywheelIterationStatus) Values() []FlywheelIterationStatus {
	return []FlywheelIterationStatus{
		"TRAINING",
		"EVALUATING",
		"COMPLETED",
		"FAILED",
		"STOP_REQUESTED",
		"STOPPED",
	}
}

type FlywheelStatus string

// Enum values for FlywheelStatus
const (
	FlywheelStatusCreating FlywheelStatus = "CREATING"
	FlywheelStatusActive   FlywheelStatus = "ACTIVE"
	FlywheelStatusUpdating FlywheelStatus = "UPDATING"
	FlywheelStatusDeleting FlywheelStatus = "DELETING"
	FlywheelStatusFailed   FlywheelStatus = "FAILED"
)

// Values returns all known values for FlywheelStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FlywheelStatus) Values() []FlywheelStatus {
	return []FlywheelStatus{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"FAILED",
	}
}

type InputFormat string

// Enum values for InputFormat
const (
	InputFormatOneDocPerFile InputFormat = "ONE_DOC_PER_FILE"
	InputFormatOneDocPerLine InputFormat = "ONE_DOC_PER_LINE"
)

// Values returns all known values for InputFormat. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (InputFormat) Values() []InputFormat {
	return []InputFormat{
		"ONE_DOC_PER_FILE",
		"ONE_DOC_PER_LINE",
	}
}

type InvalidRequestDetailReason string

// Enum values for InvalidRequestDetailReason
const (
	InvalidRequestDetailReasonDocumentSizeExceeded InvalidRequestDetailReason = "DOCUMENT_SIZE_EXCEEDED"
	InvalidRequestDetailReasonUnsupportedDocType   InvalidRequestDetailReason = "UNSUPPORTED_DOC_TYPE"
	InvalidRequestDetailReasonPageLimitExceeded    InvalidRequestDetailReason = "PAGE_LIMIT_EXCEEDED"
	InvalidRequestDetailReasonTextractAccessDenied InvalidRequestDetailReason = "TEXTRACT_ACCESS_DENIED"
)

// Values returns all known values for InvalidRequestDetailReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (InvalidRequestDetailReason) Values() []InvalidRequestDetailReason {
	return []InvalidRequestDetailReason{
		"DOCUMENT_SIZE_EXCEEDED",
		"UNSUPPORTED_DOC_TYPE",
		"PAGE_LIMIT_EXCEEDED",
		"TEXTRACT_ACCESS_DENIED",
	}
}

type InvalidRequestReason string

// Enum values for InvalidRequestReason
const (
	InvalidRequestReasonInvalidDocument InvalidRequestReason = "INVALID_DOCUMENT"
)

// Values returns all known values for InvalidRequestReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InvalidRequestReason) Values() []InvalidRequestReason {
	return []InvalidRequestReason{
		"INVALID_DOCUMENT",
	}
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusSubmitted     JobStatus = "SUBMITTED"
	JobStatusInProgress    JobStatus = "IN_PROGRESS"
	JobStatusCompleted     JobStatus = "COMPLETED"
	JobStatusFailed        JobStatus = "FAILED"
	JobStatusStopRequested JobStatus = "STOP_REQUESTED"
	JobStatusStopped       JobStatus = "STOPPED"
)

// Values returns all known values for JobStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (JobStatus) Values() []JobStatus {
	return []JobStatus{
		"SUBMITTED",
		"IN_PROGRESS",
		"COMPLETED",
		"FAILED",
		"STOP_REQUESTED",
		"STOPPED",
	}
}

type LanguageCode string

// Enum values for LanguageCode
const (
	LanguageCodeEn   LanguageCode = "en"
	LanguageCodeEs   LanguageCode = "es"
	LanguageCodeFr   LanguageCode = "fr"
	LanguageCodeDe   LanguageCode = "de"
	LanguageCodeIt   LanguageCode = "it"
	LanguageCodePt   LanguageCode = "pt"
	LanguageCodeAr   LanguageCode = "ar"
	LanguageCodeHi   LanguageCode = "hi"
	LanguageCodeJa   LanguageCode = "ja"
	LanguageCodeKo   LanguageCode = "ko"
	LanguageCodeZh   LanguageCode = "zh"
	LanguageCodeZhTw LanguageCode = "zh-TW"
)

// Values returns all known values for LanguageCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LanguageCode) Values() []LanguageCode {
	return []LanguageCode{
		"en",
		"es",
		"fr",
		"de",
		"it",
		"pt",
		"ar",
		"hi",
		"ja",
		"ko",
		"zh",
		"zh-TW",
	}
}

type ModelStatus string

// Enum values for ModelStatus
const (
	ModelStatusSubmitted          ModelStatus = "SUBMITTED"
	ModelStatusTraining           ModelStatus = "TRAINING"
	ModelStatusDeleting           ModelStatus = "DELETING"
	ModelStatusStopRequested      ModelStatus = "STOP_REQUESTED"
	ModelStatusStopped            ModelStatus = "STOPPED"
	ModelStatusInError            ModelStatus = "IN_ERROR"
	ModelStatusTrained            ModelStatus = "TRAINED"
	ModelStatusTrainedWithWarning ModelStatus = "TRAINED_WITH_WARNING"
)

// Values returns all known values for ModelStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ModelStatus) Values() []ModelStatus {
	return []ModelStatus{
		"SUBMITTED",
		"TRAINING",
		"DELETING",
		"STOP_REQUESTED",
		"STOPPED",
		"IN_ERROR",
		"TRAINED",
		"TRAINED_WITH_WARNING",
	}
}

type ModelType string

// Enum values for ModelType
const (
	ModelTypeDocumentClassifier ModelType = "DOCUMENT_CLASSIFIER"
	ModelTypeEntityRecognizer   ModelType = "ENTITY_RECOGNIZER"
)

// Values returns all known values for ModelType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ModelType) Values() []ModelType {
	return []ModelType{
		"DOCUMENT_CLASSIFIER",
		"ENTITY_RECOGNIZER",
	}
}

type PageBasedErrorCode string

// Enum values for PageBasedErrorCode
const (
	PageBasedErrorCodeTextractBadPage                       PageBasedErrorCode = "TEXTRACT_BAD_PAGE"
	PageBasedErrorCodeTextractProvisionedThroughputExceeded PageBasedErrorCode = "TEXTRACT_PROVISIONED_THROUGHPUT_EXCEEDED"
	PageBasedErrorCodePageCharactersExceeded                PageBasedErrorCode = "PAGE_CHARACTERS_EXCEEDED"
	PageBasedErrorCodePageSizeExceeded                      PageBasedErrorCode = "PAGE_SIZE_EXCEEDED"
	PageBasedErrorCodeInternalServerError                   PageBasedErrorCode = "INTERNAL_SERVER_ERROR"
)

// Values returns all known values for PageBasedErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PageBasedErrorCode) Values() []PageBasedErrorCode {
	return []PageBasedErrorCode{
		"TEXTRACT_BAD_PAGE",
		"TEXTRACT_PROVISIONED_THROUGHPUT_EXCEEDED",
		"PAGE_CHARACTERS_EXCEEDED",
		"PAGE_SIZE_EXCEEDED",
		"INTERNAL_SERVER_ERROR",
	}
}

type PageBasedWarningCode string

// Enum values for PageBasedWarningCode
const (
	PageBasedWarningCodeInferencingPlaintextWithNativeTrainedModel         PageBasedWarningCode = "INFERENCING_PLAINTEXT_WITH_NATIVE_TRAINED_MODEL"
	PageBasedWarningCodeInferencingNativeDocumentWithPlaintextTrainedModel PageBasedWarningCode = "INFERENCING_NATIVE_DOCUMENT_WITH_PLAINTEXT_TRAINED_MODEL"
)

// Values returns all known values for PageBasedWarningCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PageBasedWarningCode) Values() []PageBasedWarningCode {
	return []PageBasedWarningCode{
		"INFERENCING_PLAINTEXT_WITH_NATIVE_TRAINED_MODEL",
		"INFERENCING_NATIVE_DOCUMENT_WITH_PLAINTEXT_TRAINED_MODEL",
	}
}

type PartOfSpeechTagType string

// Enum values for PartOfSpeechTagType
const (
	PartOfSpeechTagTypeAdj   PartOfSpeechTagType = "ADJ"
	PartOfSpeechTagTypeAdp   PartOfSpeechTagType = "ADP"
	PartOfSpeechTagTypeAdv   PartOfSpeechTagType = "ADV"
	PartOfSpeechTagTypeAux   PartOfSpeechTagType = "AUX"
	PartOfSpeechTagTypeConj  PartOfSpeechTagType = "CONJ"
	PartOfSpeechTagTypeCconj PartOfSpeechTagType = "CCONJ"
	PartOfSpeechTagTypeDet   PartOfSpeechTagType = "DET"
	PartOfSpeechTagTypeIntj  PartOfSpeechTagType = "INTJ"
	PartOfSpeechTagTypeNoun  PartOfSpeechTagType = "NOUN"
	PartOfSpeechTagTypeNum   PartOfSpeechTagType = "NUM"
	PartOfSpeechTagTypeO     PartOfSpeechTagType = "O"
	PartOfSpeechTagTypePart  PartOfSpeechTagType = "PART"
	PartOfSpeechTagTypePron  PartOfSpeechTagType = "PRON"
	PartOfSpeechTagTypePropn PartOfSpeechTagType = "PROPN"
	PartOfSpeechTagTypePunct PartOfSpeechTagType = "PUNCT"
	PartOfSpeechTagTypeSconj PartOfSpeechTagType = "SCONJ"
	PartOfSpeechTagTypeSym   PartOfSpeechTagType = "SYM"
	PartOfSpeechTagTypeVerb  PartOfSpeechTagType = "VERB"
)

// Values returns all known values for PartOfSpeechTagType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PartOfSpeechTagType) Values() []PartOfSpeechTagType {
	return []PartOfSpeechTagType{
		"ADJ",
		"ADP",
		"ADV",
		"AUX",
		"CONJ",
		"CCONJ",
		"DET",
		"INTJ",
		"NOUN",
		"NUM",
		"O",
		"PART",
		"PRON",
		"PROPN",
		"PUNCT",
		"SCONJ",
		"SYM",
		"VERB",
	}
}

type PiiEntitiesDetectionMaskMode string

// Enum values for PiiEntitiesDetectionMaskMode
const (
	PiiEntitiesDetectionMaskModeMask                     PiiEntitiesDetectionMaskMode = "MASK"
	PiiEntitiesDetectionMaskModeReplaceWithPiiEntityType PiiEntitiesDetectionMaskMode = "REPLACE_WITH_PII_ENTITY_TYPE"
)

// Values returns all known values for PiiEntitiesDetectionMaskMode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (PiiEntitiesDetectionMaskMode) Values() []PiiEntitiesDetectionMaskMode {
	return []PiiEntitiesDetectionMaskMode{
		"MASK",
		"REPLACE_WITH_PII_ENTITY_TYPE",
	}
}

type PiiEntitiesDetectionMode string

// Enum values for PiiEntitiesDetectionMode
const (
	PiiEntitiesDetectionModeOnlyRedaction PiiEntitiesDetectionMode = "ONLY_REDACTION"
	PiiEntitiesDetectionModeOnlyOffsets   PiiEntitiesDetectionMode = "ONLY_OFFSETS"
)

// Values returns all known values for PiiEntitiesDetectionMode. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (PiiEntitiesDetectionMode) Values() []PiiEntitiesDetectionMode {
	return []PiiEntitiesDetectionMode{
		"ONLY_REDACTION",
		"ONLY_OFFSETS",
	}
}

type PiiEntityType string

// Enum values for PiiEntityType
const (
	PiiEntityTypeBankAccountNumber                   PiiEntityType = "BANK_ACCOUNT_NUMBER"
	PiiEntityTypeBankRouting                         PiiEntityType = "BANK_ROUTING"
	PiiEntityTypeCreditDebitNumber                   PiiEntityType = "CREDIT_DEBIT_NUMBER"
	PiiEntityTypeCreditDebitCvv                      PiiEntityType = "CREDIT_DEBIT_CVV"
	PiiEntityTypeCreditDebitExpiry                   PiiEntityType = "CREDIT_DEBIT_EXPIRY"
	PiiEntityTypePin                                 PiiEntityType = "PIN"
	PiiEntityTypeEmail                               PiiEntityType = "EMAIL"
	PiiEntityTypeAddress                             PiiEntityType = "ADDRESS"
	PiiEntityTypeName                                PiiEntityType = "NAME"
	PiiEntityTypePhone                               PiiEntityType = "PHONE"
	PiiEntityTypeSsn                                 PiiEntityType = "SSN"
	PiiEntityTypeDateTime                            PiiEntityType = "DATE_TIME"
	PiiEntityTypePassportNumber                      PiiEntityType = "PASSPORT_NUMBER"
	PiiEntityTypeDriverId                            PiiEntityType = "DRIVER_ID"
	PiiEntityTypeUrl                                 PiiEntityType = "URL"
	PiiEntityTypeAge                                 PiiEntityType = "AGE"
	PiiEntityTypeUsername                            PiiEntityType = "USERNAME"
	PiiEntityTypePassword                            PiiEntityType = "PASSWORD"
	PiiEntityTypeAwsAccessKey                        PiiEntityType = "AWS_ACCESS_KEY"
	PiiEntityTypeAwsSecretKey                        PiiEntityType = "AWS_SECRET_KEY"
	PiiEntityTypeIpAddress                           PiiEntityType = "IP_ADDRESS"
	PiiEntityTypeMacAddress                          PiiEntityType = "MAC_ADDRESS"
	PiiEntityTypeAll                                 PiiEntityType = "ALL"
	PiiEntityTypeLicensePlate                        PiiEntityType = "LICENSE_PLATE"
	PiiEntityTypeVehicleIdentificationNumber         PiiEntityType = "VEHICLE_IDENTIFICATION_NUMBER"
	PiiEntityTypeUkNationalInsuranceNumber           PiiEntityType = "UK_NATIONAL_INSURANCE_NUMBER"
	PiiEntityTypeCaSocialInsuranceNumber             PiiEntityType = "CA_SOCIAL_INSURANCE_NUMBER"
	PiiEntityTypeUsIndividualTaxIdentificationNumber PiiEntityType = "US_INDIVIDUAL_TAX_IDENTIFICATION_NUMBER"
	PiiEntityTypeUkUniqueTaxpayerReferenceNumber     PiiEntityType = "UK_UNIQUE_TAXPAYER_REFERENCE_NUMBER"
	PiiEntityTypeInPermanentAccountNumber            PiiEntityType = "IN_PERMANENT_ACCOUNT_NUMBER"
	PiiEntityTypeInNrega                             PiiEntityType = "IN_NREGA"
	PiiEntityTypeInternationalBankAccountNumber      PiiEntityType = "INTERNATIONAL_BANK_ACCOUNT_NUMBER"
	PiiEntityTypeSwiftCode                           PiiEntityType = "SWIFT_CODE"
	PiiEntityTypeUkNationalHealthServiceNumber       PiiEntityType = "UK_NATIONAL_HEALTH_SERVICE_NUMBER"
	PiiEntityTypeCaHealthNumber                      PiiEntityType = "CA_HEALTH_NUMBER"
	PiiEntityTypeInAadhaar                           PiiEntityType = "IN_AADHAAR"
	PiiEntityTypeInVoterNumber                       PiiEntityType = "IN_VOTER_NUMBER"
)

// Values returns all known values for PiiEntityType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PiiEntityType) Values() []PiiEntityType {
	return []PiiEntityType{
		"BANK_ACCOUNT_NUMBER",
		"BANK_ROUTING",
		"CREDIT_DEBIT_NUMBER",
		"CREDIT_DEBIT_CVV",
		"CREDIT_DEBIT_EXPIRY",
		"PIN",
		"EMAIL",
		"ADDRESS",
		"NAME",
		"PHONE",
		"SSN",
		"DATE_TIME",
		"PASSPORT_NUMBER",
		"DRIVER_ID",
		"URL",
		"AGE",
		"USERNAME",
		"PASSWORD",
		"AWS_ACCESS_KEY",
		"AWS_SECRET_KEY",
		"IP_ADDRESS",
		"MAC_ADDRESS",
		"ALL",
		"LICENSE_PLATE",
		"VEHICLE_IDENTIFICATION_NUMBER",
		"UK_NATIONAL_INSURANCE_NUMBER",
		"CA_SOCIAL_INSURANCE_NUMBER",
		"US_INDIVIDUAL_TAX_IDENTIFICATION_NUMBER",
		"UK_UNIQUE_TAXPAYER_REFERENCE_NUMBER",
		"IN_PERMANENT_ACCOUNT_NUMBER",
		"IN_NREGA",
		"INTERNATIONAL_BANK_ACCOUNT_NUMBER",
		"SWIFT_CODE",
		"UK_NATIONAL_HEALTH_SERVICE_NUMBER",
		"CA_HEALTH_NUMBER",
		"IN_AADHAAR",
		"IN_VOTER_NUMBER",
	}
}

type RelationshipType string

// Enum values for RelationshipType
const (
	RelationshipTypeChild RelationshipType = "CHILD"
)

// Values returns all known values for RelationshipType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RelationshipType) Values() []RelationshipType {
	return []RelationshipType{
		"CHILD",
	}
}

type SentimentType string

// Enum values for SentimentType
const (
	SentimentTypePositive SentimentType = "POSITIVE"
	SentimentTypeNegative SentimentType = "NEGATIVE"
	SentimentTypeNeutral  SentimentType = "NEUTRAL"
	SentimentTypeMixed    SentimentType = "MIXED"
)

// Values returns all known values for SentimentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SentimentType) Values() []SentimentType {
	return []SentimentType{
		"POSITIVE",
		"NEGATIVE",
		"NEUTRAL",
		"MIXED",
	}
}

type Split string

// Enum values for Split
const (
	SplitTrain Split = "TRAIN"
	SplitTest  Split = "TEST"
)

// Values returns all known values for Split. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Split) Values() []Split {
	return []Split{
		"TRAIN",
		"TEST",
	}
}

type SyntaxLanguageCode string

// Enum values for SyntaxLanguageCode
const (
	SyntaxLanguageCodeEn SyntaxLanguageCode = "en"
	SyntaxLanguageCodeEs SyntaxLanguageCode = "es"
	SyntaxLanguageCodeFr SyntaxLanguageCode = "fr"
	SyntaxLanguageCodeDe SyntaxLanguageCode = "de"
	SyntaxLanguageCodeIt SyntaxLanguageCode = "it"
	SyntaxLanguageCodePt SyntaxLanguageCode = "pt"
)

// Values returns all known values for SyntaxLanguageCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SyntaxLanguageCode) Values() []SyntaxLanguageCode {
	return []SyntaxLanguageCode{
		"en",
		"es",
		"fr",
		"de",
		"it",
		"pt",
	}
}

type TargetedSentimentEntityType string

// Enum values for TargetedSentimentEntityType
const (
	TargetedSentimentEntityTypePerson         TargetedSentimentEntityType = "PERSON"
	TargetedSentimentEntityTypeLocation       TargetedSentimentEntityType = "LOCATION"
	TargetedSentimentEntityTypeOrganization   TargetedSentimentEntityType = "ORGANIZATION"
	TargetedSentimentEntityTypeFacility       TargetedSentimentEntityType = "FACILITY"
	TargetedSentimentEntityTypeBrand          TargetedSentimentEntityType = "BRAND"
	TargetedSentimentEntityTypeCommercialItem TargetedSentimentEntityType = "COMMERCIAL_ITEM"
	TargetedSentimentEntityTypeMovie          TargetedSentimentEntityType = "MOVIE"
	TargetedSentimentEntityTypeMusic          TargetedSentimentEntityType = "MUSIC"
	TargetedSentimentEntityTypeBook           TargetedSentimentEntityType = "BOOK"
	TargetedSentimentEntityTypeSoftware       TargetedSentimentEntityType = "SOFTWARE"
	TargetedSentimentEntityTypeGame           TargetedSentimentEntityType = "GAME"
	TargetedSentimentEntityTypePersonalTitle  TargetedSentimentEntityType = "PERSONAL_TITLE"
	TargetedSentimentEntityTypeEvent          TargetedSentimentEntityType = "EVENT"
	TargetedSentimentEntityTypeDate           TargetedSentimentEntityType = "DATE"
	TargetedSentimentEntityTypeQuantity       TargetedSentimentEntityType = "QUANTITY"
	TargetedSentimentEntityTypeAttribute      TargetedSentimentEntityType = "ATTRIBUTE"
	TargetedSentimentEntityTypeOther          TargetedSentimentEntityType = "OTHER"
)

// Values returns all known values for TargetedSentimentEntityType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (TargetedSentimentEntityType) Values() []TargetedSentimentEntityType {
	return []TargetedSentimentEntityType{
		"PERSON",
		"LOCATION",
		"ORGANIZATION",
		"FACILITY",
		"BRAND",
		"COMMERCIAL_ITEM",
		"MOVIE",
		"MUSIC",
		"BOOK",
		"SOFTWARE",
		"GAME",
		"PERSONAL_TITLE",
		"EVENT",
		"DATE",
		"QUANTITY",
		"ATTRIBUTE",
		"OTHER",
	}
}

type ToxicContentType string

// Enum values for ToxicContentType
const (
	ToxicContentTypeGraphic           ToxicContentType = "GRAPHIC"
	ToxicContentTypeHarassmentOrAbuse ToxicContentType = "HARASSMENT_OR_ABUSE"
	ToxicContentTypeHateSpeech        ToxicContentType = "HATE_SPEECH"
	ToxicContentTypeInsult            ToxicContentType = "INSULT"
	ToxicContentTypeProfanity         ToxicContentType = "PROFANITY"
	ToxicContentTypeSexual            ToxicContentType = "SEXUAL"
	ToxicContentTypeViolenceOrThreat  ToxicContentType = "VIOLENCE_OR_THREAT"
)

// Values returns all known values for ToxicContentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ToxicContentType) Values() []ToxicContentType {
	return []ToxicContentType{
		"GRAPHIC",
		"HARASSMENT_OR_ABUSE",
		"HATE_SPEECH",
		"INSULT",
		"PROFANITY",
		"SEXUAL",
		"VIOLENCE_OR_THREAT",
	}
}
