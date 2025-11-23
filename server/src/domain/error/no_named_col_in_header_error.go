package error

import (
	"fmt"
)

func NewNoNamedColInHeaderError(
	colName string,
	inputCSVFormatID string,
	dstProperty string,
) *noNamedColInHeaderError {
	return &noNamedColInHeaderError{
		colName:          colName,
		inputCSVFormatID: inputCSVFormatID,
		dstProperty:      dstProperty,
	}
}

var _ ConvertCSVError = &noNamedColInHeaderError{}

type noNamedColInHeaderError struct {
	colName          string
	inputCSVFormatID string
	dstProperty      string
}

func (e *noNamedColInHeaderError) Error() string {
	return fmt.Sprintf(
		"column name: %s is not found in csv. but inputCSVFormatID: %s is needs col for mappint to %s",
		e.colName, e.inputCSVFormatID, e.dstProperty,
	)
}

func (e *noNamedColInHeaderError) Type() ConvertCSVErrorType {
	return ConvertCSVErrorTypeNoNamedColInHeader
}

func (e *noNamedColInHeaderError) RelatedProps() map[string]any {
	return map[string]any{
		"colName":          e.colName,
		"inputCSVFormatID": e.inputCSVFormatID,
		"dstProperty":      e.dstProperty,
	}
}
