package gomiddlewares

import (
	"io"
	"os"
)

func ExampleSetOutput_singleWriter() {
	// Open the log file for writing.
	logFile, _ := os.Create("logfile")

	// Set the output to the log file.
	SetOutput(logFile)
}

func ExampleSetOutput_multiWriter() {
	// Open the log file for writing.
	logFile, _ := os.Create("logfile")

	// Create a MultiWriter that writes to both os.Stdout and the log file.
	mw := io.MultiWriter(os.Stdout, logFile)

	// Set the output to the MultiWriter.
	SetOutput(mw)
}
