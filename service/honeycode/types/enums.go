// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Format string

// Enum values for Format
const (
	FormatAuto       Format = "AUTO"
	FormatNumber     Format = "NUMBER"
	FormatCurrency   Format = "CURRENCY"
	FormatDate       Format = "DATE"
	FormatTime       Format = "TIME"
	FormatDateTime   Format = "DATE_TIME"
	FormatPercentage Format = "PERCENTAGE"
	FormatText       Format = "TEXT"
	FormatAccounting Format = "ACCOUNTING"
	FormatContact    Format = "CONTACT"
	FormatRowlink    Format = "ROWLINK"
)

// Values returns all known values for Format. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Format) Values() []Format {
	return []Format{
		"AUTO",
		"NUMBER",
		"CURRENCY",
		"DATE",
		"TIME",
		"DATE_TIME",
		"PERCENTAGE",
		"TEXT",
		"ACCOUNTING",
		"CONTACT",
		"ROWLINK",
	}
}

type ImportDataCharacterEncoding string

// Enum values for ImportDataCharacterEncoding
const (
	ImportDataCharacterEncodingUtf8     ImportDataCharacterEncoding = "UTF-8"
	ImportDataCharacterEncodingUsAscii  ImportDataCharacterEncoding = "US-ASCII"
	ImportDataCharacterEncodingIso88591 ImportDataCharacterEncoding = "ISO-8859-1"
	ImportDataCharacterEncodingUtf16be  ImportDataCharacterEncoding = "UTF-16BE"
	ImportDataCharacterEncodingUtf16le  ImportDataCharacterEncoding = "UTF-16LE"
	ImportDataCharacterEncodingUtf16    ImportDataCharacterEncoding = "UTF-16"
)

// Values returns all known values for ImportDataCharacterEncoding. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ImportDataCharacterEncoding) Values() []ImportDataCharacterEncoding {
	return []ImportDataCharacterEncoding{
		"UTF-8",
		"US-ASCII",
		"ISO-8859-1",
		"UTF-16BE",
		"UTF-16LE",
		"UTF-16",
	}
}

type ImportSourceDataFormat string

// Enum values for ImportSourceDataFormat
const (
	ImportSourceDataFormatDelimitedText ImportSourceDataFormat = "DELIMITED_TEXT"
)

// Values returns all known values for ImportSourceDataFormat. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImportSourceDataFormat) Values() []ImportSourceDataFormat {
	return []ImportSourceDataFormat{
		"DELIMITED_TEXT",
	}
}

type TableDataImportJobStatus string

// Enum values for TableDataImportJobStatus
const (
	TableDataImportJobStatusSubmitted  TableDataImportJobStatus = "SUBMITTED"
	TableDataImportJobStatusInProgress TableDataImportJobStatus = "IN_PROGRESS"
	TableDataImportJobStatusCompleted  TableDataImportJobStatus = "COMPLETED"
	TableDataImportJobStatusFailed     TableDataImportJobStatus = "FAILED"
)

// Values returns all known values for TableDataImportJobStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TableDataImportJobStatus) Values() []TableDataImportJobStatus {
	return []TableDataImportJobStatus{
		"SUBMITTED",
		"IN_PROGRESS",
		"COMPLETED",
		"FAILED",
	}
}

type UpsertAction string

// Enum values for UpsertAction
const (
	UpsertActionUpdated  UpsertAction = "UPDATED"
	UpsertActionAppended UpsertAction = "APPENDED"
)

// Values returns all known values for UpsertAction. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (UpsertAction) Values() []UpsertAction {
	return []UpsertAction{
		"UPDATED",
		"APPENDED",
	}
}
