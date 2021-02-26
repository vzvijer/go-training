package paasio

import (
	"io"
	"sync"
)

type PaasReader struct {
	n_bytes int64
	n_reads int
	mu      sync.Mutex
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &PaasReader{}
}

func (pr *PaasReader) Read(p []byte) (int, error) {
	pr.mu.Lock()
	pr.n_bytes += int64(len(p))
	pr.n_reads++
	pr.mu.Unlock()

	return len(p), nil
}

func (pr *PaasReader) ReadCount() (n int64, nops int) {
	return pr.n_bytes, pr.n_reads
}

type PaasWriter struct {
	n_bytes int64
	n_reads int
	mu      sync.Mutex
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &PaasWriter{}
}

func (pw *PaasWriter) Write(p []byte) (int, error) {
	pw.mu.Lock()
	pw.n_bytes += int64(len(p))
	pw.n_reads++
	pw.mu.Unlock()

	return len(p), nil
}

func (pw *PaasWriter) WriteCount() (n int64, nops int) {
	return pw.n_bytes, pw.n_reads
}

type PaasReaderWriter struct {
	PaasReader
	PaasWriter
}

func NewReadWriteCounter(readWriter io.ReadWriter) ReadWriteCounter {
	return &PaasReaderWriter{}
}

func (prw *PaasReaderWriter) Read(p []byte) (int, error) {
	return prw.Read(p)
}

func (prw *PaasReaderWriter) ReadCount() (n int64, nops int) {
	return prw.ReadCount()
}

func (prw *PaasReaderWriter) Write(p []byte) (int, error) {
	return prw.Write(p)
}

func (prw *PaasReaderWriter) WriteCount() (n int64, nops int) {
	return prw.WriteCount()
}
