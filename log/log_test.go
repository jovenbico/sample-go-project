package log

import (
	"errors"
	"testing"
)

func TestLevel(t *testing.T) {
	Debug.Println("testing debug level")
	Info.Println("testing info level")
	Warn.Println("testing warn level")
	Error.Println("testing error level")

	Error.PrintlnErr(errors.New("testing error{}"))
	Error.PrintlnErr(nil)
}
