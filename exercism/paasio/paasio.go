package paasio

import (
	"errors"
	"io"
)

type PaasReader struct {
	n_bytes int64
	n_reads int
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return PaasReader{}
}

func (pr PaasReader) Read(p []byte) (int, error) {
	return 0, errors.New("Something")
}

func (pr PaasReader) ReadCount() (n int64, nops int) {
	return pr.n_bytes, pr.n_reads
}

type PaasWriter struct {
	n_bytes int64
	n_reads int
}

func (pw PaasWriter) Write(p []byte) (int, error) {
	return 0, errors.New("Something")
}

func (pw PaasWriter) WriteCount() (n int64, nops int) {
	return pw.n_bytes, pw.n_reads
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return PaasWriter{}
}

type PaasReaderWriter struct {
	PaasReader
	PaasWriter
}

func NewReadWriteCounter() PaasReaderWriter {
	return PaasReaderWriter{}
}
