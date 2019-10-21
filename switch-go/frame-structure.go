/* For license and copyright information please see LEGAL file in repository */

package frameswitch

// frameStructure is represent statless switching protocol structure!
// It is just to show protocol in better way, we never use this type!
// StartDelimiter, EndDelimiter & CheckSequence may add to this structure due to physical layer rules!
// Ports number can be mutable due to physical links limits. Endpoint must beware of this aspect!
// up-to 256 switch port number(up-to Switch255PortNum) can add to header!
// Each switch device in any location of link can be wire or wireless with any Energy||Frequency specs (e.g. Fiber, WiFi, LAN, Bluetooth, ...)
type frameStructure struct {
	NextHop        uint8  // Indicate switch port number use on next switch device!
	TotalHop       uint8  // Total hop numbers and also indicate payload location (Payload = rawFrame[frame.TotalHop+3:])! 255 use for multicast farmes to all ports!
	NextHeader     uint8  // Indicate upper layer protocol equal EtherType!
	Switch0PortNum uint8  // Source Port Number.
	Switch1PortNum uint8  // Any other than Switch0PortNum can be Destination Port Number.
	Payload        []byte // up to 8192 Byte or 8KB. Enough to stream 1.5Mbps video call in each 40ms frames (1.5/8*1024/1000*40=7.68KB)
}
