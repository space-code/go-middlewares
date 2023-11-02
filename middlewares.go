package gomiddlewares

import (
	"io"
	"os"
)

var output io.Writer

func init() {
	output = os.Stdout
}

func SetOutput(w io.Writer) {
	output = w
}
