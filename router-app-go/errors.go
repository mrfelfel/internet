/* For license and copyright information please see LEGAL file in repository */

package approuter

import "./error"

// Declare ChaparKhane errors number
const (
	streamDataPayloadEmpty = 00000 + (iota + 1)
	serviceNotFound
)

// Declare Errors Details
var (
	StreamDataPayloadEmpty = error.NewError("Stream data payload can't be empty", streamDataPayloadEmpty, 0)
	ServiceNotFound         = error.NewError("Requested Service is out range of services in this version of service", serviceNotFound, 0)
)
