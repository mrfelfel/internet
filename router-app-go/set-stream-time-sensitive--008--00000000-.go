/* For license and copyright information please see LEGAL file in repository */

package approuter

// SetStreamTimeSensitive : Use in VoIP, IPTV, ...
func SetStreamTimeSensitive() {
	// Dropping packets is preferable to waiting for packets delayed due to retransmissions.
	// Developer can ask to complete data for offline usage after first data usage.
}
