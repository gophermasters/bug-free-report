package errors

import "github.com/theGOURL/warning"

func CloseConnError(err error){
	warning.PRINT_DEFAULT_ERRORS(err,"unable to close database connection");
}