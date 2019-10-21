/* For license and copyright information please see LEGAL file in repository */

package frameswitch

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
	if len(*f) < FrameLen {
		// return FrameTooShort
	}

	return nil
}

// SwitchFrame detect and send frame!
func (f *Frame) SwitchFrame(incomePort uint8) {
	if f.GetTotalHop() == 255 {
		BroadcastFrame(f, incomePort)
	} else {
		UnicastFrame(f, incomePort)
	}
}

// GetNextHop will return NextHop in memory safe way!
func (f *Frame) GetNextHop() (NextHop uint8) {
	return uint8((*f)[0])
}

// IncrementNextHop will increment NextHop number in frame!
func (f *Frame) IncrementNextHop() {
	(*f)[0] = byte(f.GetNextHop() + 1)
}

// GetTotalHop will return TotalHop in memory safe way!
func (f *Frame) GetTotalHop() (TotalHop uint8) {
	return uint8((*f)[1])
}

// GetNextHeader will return NextHeader in memory safe way!
func (f *Frame) GetNextHeader() (NextHeader uint8) {
	return uint8((*f)[2])
}

// GetSwitchPortNum will return SwitchPortNum of i in memory safe way!
func (f *Frame) GetSwitchPortNum(i uint8) (SwitchPortNum uint8) {
	return uint8((*f)[i+3])
}

// SetSwitchPortNum will set SwitchPortNum of i!
func (f *Frame) SetSwitchPortNum(i uint8, SwitchPortNum uint8) {
	(*f)[i+3] = byte(SwitchPortNum)
}

// GetPayload will return Payload in memory safe way!
func (f *Frame) GetPayload() (Payload []byte) {
	return (*f)[f.GetTotalHop()+3:]
}
