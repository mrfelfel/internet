/* For license and copyright information please see LEGAL file in repository */

package uip

// packetStructure is represent protocol structure!
// It is just to show protocol in better way, we never use this type!
type packetStructure struct {
	DestinationXP          [4]byte  // Mutable
	DestinationRouter      [4]byte  // Mutable
	DestinationOS          [4]byte  // Mutable
	DestinationAppID       [2]byte  // Mutable
	DestinationAppProccess [2]byte  // Mutable
	SourceSessionID        [16]byte // Due to privacy concerns each peer must have unique SessionID and don't use same ID! Peers can change or revoke it!
	// Always encrypted (except ==0 for register guest connection)
	// Use even number for client(who start connection) to start a stream e.g. 0,2,4,6,8,10,...
	// Use odd number for server(who receive connection) to start a stream e.g. 1,3,5,7,9,11,...
	// Use 0 for client to send internal service & 1 for server to send internal service to each other.
	// Reset everyday and start from 2 for sending new stream
	StreamID uint32
	// Always encrypted (except ==0 for register guest connection)
	// Use even number for client(who start connection) to send data e.g. 0,2,4,6,8,10,...
	// Use odd number for server(who receive connection) to send data e.g. 1,3,5,7,9,11,...
	// Use 0 for client to send internal service & 1 for server to send internal service to each other.
	// Max 4.72TB can transmit in single stream with 1.18KB payload size as limit to 1.5KB of ethernet frames!
	PacketID uint32
	// Payload can be any OSI application layer model and store in StreamData in order by PacketID.
	// Each instance of running end apps must use specific application protocol e.g. HTTP||sRPC||...
	// Encrypted in App to App connection but not in RegisterConnection service.
	// payload size respect data-link protocol like [Ethernet](https://en.wikipedia.org/wiki/Ethernet_frame) that is 1500 bytes(octets)
	Payload       []byte
	CheckSequence [4]byte
}
