/* For license and copyright information please see LEGAL file in repository */

package uip

// Packet use for packet methods!
type Packet []byte

const (
	// PacketLen is minimum packet length of UIP packet
	// 448bit header + 128bit min payload
	PacketLen = 72
)

// CheckPacket will check packet for any bad situation!
// Always check packet before use any other packet methods otherwise panic occur!
func (p *Packet) CheckPacket() (err error) {
}

// GetDestinationUIP will return DestinationUIP in memory unsafe way!
func (p *Packet) GetDestinationUIP() (DestinationUIP [16]byte) {
}

// GetSourceUIP will return SourceUIP in memory unsafe way!
func (p *Packet) GetSourceUIP() (SourceUIP [16]byte) {
}

// GetConnectionID will return ConnectionID in memory unsafe way!
func (p *Packet) GetConnectionID() (ConnectionID [16]byte) {
	// UIP don't support ConnectionID == 0!
	// Hack situation || DDoS attack || ...
}

// GetStreamID will return StreamID in memory unsafe way!
func (p *Packet) GetStreamID() (StreamID uint32) {
}

// GetPacketID will return PacketID in memory unsafe way!
func (p *Packet) GetPacketID() (PacketID [16]byte) {
}
