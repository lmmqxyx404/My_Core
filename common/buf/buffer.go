package buf

import (
	"io"

	"github.com/lmmqxyx404/my_core/common/bytespool"
)

const (
	// Size of a regular buffer.
	Size = 8192
)

var pool = bytespool.GetPool(Size)

// Buffer is a recyclable allocation of a byte array. Buffer.Release() recycles
// the buffer into an internal buffer pool, in order to recreate a buffer more
// quickly.
type Buffer struct {
	v         []byte
	start     int32
	end       int32
	unmanaged bool
	// todo: add udp
	// UDP       *net.Destination
}

// New creates a Buffer with 0 length and 8K capacity.
func New() *Buffer {
	buf := pool.Get().([]byte)
	if cap(buf) >= Size {
		buf = buf[:Size]
	} else {
		buf = make([]byte, Size)
	}

	return &Buffer{
		v: buf,
	}
}

// ReadFullFrom reads exact size of bytes from given reader, or until error occurs.
func (b *Buffer) ReadFullFrom(reader io.Reader, size int32) (int64, error) {
	end := b.end + size
	if end > int32(len(b.v)) {
		v := end
		return 0, newError("out of bound: ", v)
	}
	n, err := io.ReadFull(reader, b.v[b.end:end])
	b.end += int32(n)
	return int64(n), err
}

// IsEmpty returns true if the buffer is empty.
func (b *Buffer) IsEmpty() bool {
	return b.Len() == 0
}

// Len returns the length of the buffer content.
func (b *Buffer) Len() int32 {
	if b == nil {
		return 0
	}
	return b.end - b.start
}

// Read implements io.Reader.Read().
func (b *Buffer) Read(data []byte) (int, error) {
	if b.Len() == 0 {
		return 0, io.EOF
	}
	nBytes := copy(data, b.v[b.start:b.end])
	if int32(nBytes) == b.Len() {
		b.Clear()
	} else {
		b.start += int32(nBytes)
	}
	return nBytes, nil
}
