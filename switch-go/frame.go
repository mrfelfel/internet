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
// All switch logic is in this function!!
func (f *Frame) SwitchFrame(incomePort uint8) {
	// The reasons of use SetSwitchPortNum method here:
	// - UnicastFrame : To be sure receive port is same with declaration one in frame we replace it always!
	// - BroadcastFrame : To improve performance, previous switch just send frame without declare port, we must declare it now!
	// - Rule&Security : To be sure physical network port is same on sender and receiver switch, we must set it again here!
	f.SetSwitchPortNum(f.GetNextHop(), incomePort)

	f.IncrementNextHop()

	if f.GetTotalHop() == 0 {
		// BroadcastFrame use to broadcast a frame to all ports!
		// Due to frame must have at least 2 hop so we use unused TotalHop==0 for multicast farmes to all ports!
		// So both TotalHop==0x00 & TotalHop==0xff have 256 SwitchPortNum space in frame header!
		// Frame must have all Switch0PortNum to Switch254PortNum with 0 byte data in header otherwise frame payload rewrite by switch!
		var i uint8
		for i = 0; i <= 255; i++ {
			// Send frame to desire port queue
			// send(incomeFrame, i)
		}
	} else {
		// UnicastFrame will send frame to the specific port
		// Send frame to desire port queue
		// send(f, f.GetNextHop())
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
// BEWARE! To maximize usage of total hop number we assume 0x00 as 1 ... 0xff as 256 hop number!
// Min TotalHop number is 2 so we use 0 for broadcasting frame and 1 for ...!
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
	return (*f)[f.GetTotalHop()+4:]
}
