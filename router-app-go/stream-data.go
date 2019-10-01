/* For license and copyright information please see LEGAL file in repository */

package approuter

import (
	uip "../uip-go"
	http "./HTTP"
	srpc "./sRPC"
)

// StreamData : Can use for send or receive data on specific StreamID.
// It can pass to logic layer to give data access to developer!
// Data flow can be up to down (parse raw income data) or down to up (encode app data with respect MTU)
type StreamData struct {
	ConnectionData *ConnectionData
	StreamID       uint32 // Odd number for server, Even number for Peer.
	Status         uint8  // 0:close 1:open 2:rate-limited 3:closing 4:opening 5:BrokenPacket
	Weight         uint8  // 16 queue for priority weight of the streams exist.
	TimeSensitive  bool   // If true we must call related service in each received packet. VoIP, IPTV, ...
	ServiceID      uint32 // 4294967296 service is more enough on any platform!
	Payload        []byte // Income||Outcome data buffer. Divide to n packet to respect network MTU!

	// Network layer
	UIP uip.Packet // Always last send or received packet data

	// Application layer. Just one is not nil!
	SRPC *srpc.Packet
	HTTP *http.Packet

	Err error
}
