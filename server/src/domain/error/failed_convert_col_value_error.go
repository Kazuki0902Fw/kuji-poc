package error

import (
	"fmt"
)

func NewFailedConvertColValueError(
	rowNumber int,
	colNumber int,
	colValue string,
	dstProperty string,
	detail string,
) *failedConvertColValueError {
	return &failedConvertColValueError{
		rowNumber:   rowNumber,
		colNumber:   colNumber,
		colValue:    colValue,
		dstProperty: dstProperty,
		detail:      detail,
	}
}

var _ ConvertCSVError = &failedConvertColValueError{}

type failedConvertColValueError struct {
	rowNumber   int
	colNumber   int
	colValue    string
	dstProperty string
	detail      string
}

func (e *failedConvertColValueError) Error() string {
	return fmt.Sprintf(
		"failed to convert col value: %s in row number: %d, col number: %d for mapping to %s. detail: %s",
		e.colValue, e.rowNumber, e.colNumber, e.dstProperty, e.detail,
	)
}

func (e *failedConvertColValueError) Type() ConvertCSVErrorType {
	return ConvertCSVErrorTypeFailedConvertColValue
}

func (e *failedConvertColValueError) RelatedProps() map[string]any {
	return map[string]any{
		"rowNumber":   e.rowNumber,
		"colNumber":   e.colNumber,
		"colValue":    e.colValue,
		"dstProperty": e.dstProperty,
	}
}
