package error

import (
	"fmt"

	"kujicole/domain/model"
)

func NewNoHeaderColError(inputCSVFormatID model.ID) *noHeaderColError {
	return &noHeaderColError{
		inputCSVFormatID: inputCSVFormatID,
	}
}

var _ ConvertCSVError = &noHeaderColError{}

type noHeaderColError struct {
	inputCSVFormatID model.ID
}

func (e *noHeaderColError) Error() string {
	return fmt.Sprintf("No header col in csv. but inputCSVFormatID: %s is needs header col", e.inputCSVFormatID)
}

func (e *noHeaderColError) Type() ConvertCSVErrorType {
	return ConvertCSVErrorTypeNoHeaderCol
}

func (e *noHeaderColError) RelatedProps() map[string]any {
	return map[string]any{
		"inputCSVFormatID": e.inputCSVFormatID,
	}
}
