/* For license and copyright information please see LEGAL file in repository */

package wireswitch

// Frame use for frame methods!
type Frame []byte

const (
	// FrameLen is minimum frame length
	// 3 Byte header + 72 Byte min payload
	FrameLen = 75
)

// CheckFrame will check frame for any bad situation!
// Always check frame before use any other frame methods otherwise panic may occur!
func (f *Frame) CheckFrame() (err error) {
	return nil
}