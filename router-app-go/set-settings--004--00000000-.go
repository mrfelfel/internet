/* For license and copyright information please see LEGAL file in repository */

package approuter

// ReqSetSettings : Don't send this request between active stream due may streams data will lost!
type ReqSetSettings struct {
	PacketPayloadSize uint16 // Defaults: 1200 byte. It can't be under 1200 byte. Exclude network or transport header.
	// Packet data compression type e.g. gzip, ...
}

// SetSettings : Change settings set in RegisterConnection
func SetSettings(req *ReqSetSettings) (err error) {
	return nil
}
