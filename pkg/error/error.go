package error

type Error struct {
	Code        uint64 // {xxx} {yyy} xxx- standart error code, yyy - cpecific parameters
	Description string
	//Locale      language.Tag
}

func New() *Error {
	return &Error{}
}

func (err *Error) SetCode(code uint64) *Error {
	err.Code = code
	return err
}

func (err *Error) SetMessage(description string) *Error {
	err.Description = description
	return err
}

func (err *Error) GetHttpCode() int {
	return err.convertErrToHttpCode()
}

func (err *Error) convertErrToHttpCode() int {
	return int(err.Code % 1000)
}
