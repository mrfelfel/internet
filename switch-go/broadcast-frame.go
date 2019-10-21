/* For license and copyright information please see LEGAL file in repository */

package frameswitch

// BroadcastFrame use to broadcast a frame to all ports!
// Frame must have all Switch0PortNum to Switch254PortNum with 0 byte data in header otherwise frame payload rewrite by switch!
func BroadcastFrame(incomeFrame *Frame, incomePort uint8) {
	// To be sure receive port is same with declaration one in frame we replace it always!
	incomeFrame.SetSwitchPortNum(incomeFrame.GetNextHop(), incomePort)

	incomeFrame.IncrementNextHop()

	var i uint8
	for i = 0; i <= 255; i++ {
		incomeFrame.SetSwitchPortNum(incomeFrame.GetNextHop(), i)
		// Send frame to desire port queue
		// send(incomeFrame, i)
	}
}
