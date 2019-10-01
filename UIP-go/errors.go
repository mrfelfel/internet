/* For license and copyright information please see LEGAL file in repository */

package uip

import "./error"

// Declare IP errors number
const (
	payloadEmpty = 00000 + (iota + 1)
	packetTooShort
)

// Declare Errors Details
var (
	PayloadEmpty = error.NewError("Stream data payload can't be empty", payloadEmpty, 0)
	PacketTooShort = error.NewError("IP packet is empty or too short than standard header", packetTooShort, 0)
)
