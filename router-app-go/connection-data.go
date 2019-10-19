/* For license and copyright information please see LEGAL file in repository */

package approuter

// ConnectionData can use by any type users itself or delegate to other users to act as the owner!
// Each user in each device need unique connection to another party.
type ConnectionData struct {
	SessionID         [16]byte
	ConnectionID      [16]byte // Connection UUID, First bit indicate guest or registered user
	Status            uint8    // 0:close 1:open 2:rate-limited 3:closing 4:opening 5:
	Weight            uint8    // 16 queue for priority weight of the connections exist.
	OwnerID           [16]byte // Can't change after creation. Guest=ConnectionPublicKey
	OwnerType         uint8    // 0:Guest, 1:Person, 2:Org, 3:App, ...
	DelegateUserID    [16]byte // Can't change after first set. Guest={1}
	PeerUIPAddress    [16]byte
	CipherSuite       uint16   // Selected algorithms //https://en.wikipedia.org/wiki/Cipher_suite
	FrameSize         uint16   // Default frame size is 128bit due cipher block size
	EncryptionKey     [32]byte // 256bit encryption key
	PacketPayloadSize uint16   // Defaults: 1200 byte. It can't be under 1200 byte. Exclude network or transport header.
	MaxBandwidth      uint64   // use to tell the peer to slow down or packets will be drops in OS queues!
	RequestCount      uint64   // Use for PayAsGo strategy.
	BytesSent         uint64   // Counts the bytes of payload data sent.
	PacketsSent       uint64   // Counts packets sent.
	BytesReceived     uint64   // Counts the bytes of payload data Receive.
	PacketsReceived   uint64   // Counts packets Receive.
	AccessControl     AccessControl
	StreamPool        map[uint32]*StreamData // StreamID
}

// AccessControl : Use ABAC features for AccessControl fields.
// Must store arrays in sort for easy read and comparison
type AccessControl struct {
	// Remove Useless Inner interval in When key in AccessControl.
	// e.g. 150000/160000 and 151010/153030 the second one is useless!
	// Iso8601 Time intervals <start>/<end> ["hhmmss/hhmmss", "hhmmss/hhmmss"]!!!
	// Just use GMT0!!!
	When  []uint64
	Where [][16]byte // ["UIP", "UIP"]!
	Which []uint32   // ["ServiceID", "ServiceID"] Just in specific AppID
	How   []string   //
	What  [][16]byte // ["RecordUUID", "RecordUUID"]
	If    []string   //
}

// NewConnection use to make new connection and initialize inner maps!
func NewConnection() *ConnectionData {
	return &ConnectionData{
		StreamPool: make(map[uint32]*StreamData),
	}
}

// RegisterConnection use to register new connection in server connection pool!!
func (s *Server) RegisterConnection(cd *ConnectionData) {
	s.ConnectionPool[cd.ConnectionID] = cd
}

// UnRegisterConnection use to un-register exiting connection in server connection pool!!
func (s *Server) UnRegisterConnection(cd *ConnectionData) {
	delete(s.ConnectionPool, cd.ConnectionID)
}
