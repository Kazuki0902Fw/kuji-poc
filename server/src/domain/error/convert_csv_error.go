package error

type ConvertCSVErrorType int

const (
	ConvertCSVErrorTypeNoHeaderCol           ConvertCSVErrorType = 1
	ConvertCSVErrorTypeColLengthNotEnough    ConvertCSVErrorType = 2
	ConvertCSVErrorTypeNoNamedColInHeader    ConvertCSVErrorType = 3
	ConvertCSVErrorTypeFailedConvertColValue ConvertCSVErrorType = 4
)

type ConvertCSVError interface {
	Error() string
	Type() ConvertCSVErrorType
	RelatedProps() map[string]any
}
