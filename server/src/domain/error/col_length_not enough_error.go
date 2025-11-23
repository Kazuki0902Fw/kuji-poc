package error

import (
	"fmt"
)

func NewColLengthNotEnoughError(colNumber int, needsColNumber int, dstProperty string) *colLengthNotEnoughError {
	return &colLengthNotEnoughError{
		colNumber:      colNumber,
		needsColNumber: needsColNumber,
		dstProperty:    dstProperty,
	}
}

var _ ConvertCSVError = &colLengthNotEnoughError{}

type colLengthNotEnoughError struct {
	colNumber      int
	needsColNumber int
	dstProperty    string
}

func (e *colLengthNotEnoughError) Error() string {
	return fmt.Sprintf(
		"col length is not enough in csv at col number: %d. needs %d col for mappint to %s",
		e.colNumber, e.needsColNumber, e.dstProperty,
	)
}

func (e *colLengthNotEnoughError) Type() ConvertCSVErrorType {
	return ConvertCSVErrorTypeColLengthNotEnough
}

func (e *colLengthNotEnoughError) RelatedProps() map[string]any {
	return map[string]any{
		"colNumber":      e.colNumber,
		"needsColNumber": e.needsColNumber,
		"dstProperty":    e.dstProperty,
	}
}
