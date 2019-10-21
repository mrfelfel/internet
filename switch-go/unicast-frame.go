/* For license and copyright information please see LEGAL file in repository */

package frameswitch

// UnicastFrame will send frame to the specific port
func UnicastFrame(incomeFrame *Frame, incomePort uint8) {
	// First check receive port is same with declaration one in frame!
	if incomeFrame.GetSwitchPortNum(incomeFrame.GetNextHop()) != incomePort {
		// Not allowed rule to send frame from other port than incomePort
		// Send response or just ignore frame
		return
	}

	incomeFrame.IncrementNextHop()

	// Send frame to desire port queue
	// send(incomeFrame, incomeFrame.GetNextHop())
}
