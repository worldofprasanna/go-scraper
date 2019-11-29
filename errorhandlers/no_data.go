package errorhandlers

// NoDataFound - Custom Error struct to denote there is no data
type NoDataFound struct {
	msg string
	userFriendlyMsg string
}

//NewNoDataFound - Create new NoDataFound error
func NewNoDataFound(errorMsg string) error {
	return &NoDataFound {
		msg: errorMsg,
	}
}

func (e *NoDataFound) Error() string {
	return e.msg
}

// UserFriendlyMsg - Display the user friendly error message
func (e *NoDataFound) UserFriendlyMsg() string {
	return e.msg
}