/* For license and copyright information please see LEGAL file in repository */

package approuter

import (
	uip "../uip-go"
	srpc "./sRPC"
)

// DefaultServer use as default to start a server.
var DefaultServer = NewServer()

// generator will overwrite these handlers!
func init() {
	DefaultServer.Handlers.UIPHandler = handleUIP
	DefaultServer.Handlers.SCPHandler = handleSCP
}

// handleUIP use to handle UIP packet but
// We strongly suggest you use generator to improve performance and get all library functionality!
func handleUIP(s *Server, osIncomePacket []byte) {
	var err error
	var sd *StreamData
	var SCP = new(scp.Packet)
	var ok bool
	var cd *ConnectionData

	// Parse ConnectionID on packet header
	err = SCP.ParseConnectionID(payload)
	if err != nil {
		// Send response or just ignore packet
		return
	}

	// Find ConnectionData from ConnectionPool by ConnectionID
	cd, ok = s.ConnectionPool[SCP.ConnectionID]
	// If it is first time that ConnectionID used
	if !ok {
		err = s.HandleServerServices(sd)
		// If it is not register guest connection service, so
		if err != nil {
			// get Connection detail from .BuildInformation.AuthorizationServer
			// Call related SDK
			return
		}
		return
	}

	// Continue parse packet header!
	err = SCP.ParsePacket(payload, cd.FrameSize, cd.EncryptionKey, cd.CipherSuite)
	if err != nil {
		// Send response or just ignore packet
		return
	}

	// Check internal services in server-services files
	// Check sd.SCPHeader.StreamID == 1 || sd.SCPHeader.PacketID == 1 in sdk for client!
	if SCP.StreamID == 0 || SCP.PacketID == 0 {
		// response with sRPC protocol manner
		s.HandleServerServices(sd)
		return
	}

	sd, ok = cd.StreamPool[SCP.StreamID]
	if !ok {
		// Send response or just ignore packet
		return
	}

	sd.ConnectionData = cd
	sd.SCP = SCP
	sd.Payload = SCP.Payload

	// Check TimeSensitive or last packet of stream here.
	// Check sd.SCPHeader.PacketID == 4294967295 in sdk for client
	if SCP.PacketID == 4294967294 || sd.TimeSensitive {
		// Increment request count for rate limiting
		sd.ConnectionData.RequestCount++

		// Send last part of that stream!
		// Client said stream had been finished and server must continue process.
		// call app protocol indicate in manifest like sRPC!
		handleSRPC(s, sd)
		return
	}

	// add payload to StreamPool
}

// sRPC is our experimental SabzCity remote procedure call Protocol!!
func handleSRPC(s *Server, sd *StreamData) {
	var ok bool
	var sf ServiceFunc

	sd.SRPC = new(srpc.Packet)

	// Parse ConnectionID on packet header
	sd.Err = sd.SRPC.ParsePacket(sd.Payload)
	if sd.Err != nil {
		return
	}

	sd.ServiceID = sd.SRPC.ServiceID

	sf, ok = s.Handlers.Services[sd.ServiceID]
	if !ok {
		sd.Err = ServiceNotFound
		// handle sd.Err
		// Send response or just ignore packet
		return
	}
	sf(sd)
}

// handleHTTP use to handle http packets. It can use for architectures like restful, ...
// Protocol Standard
// http2 : https://httpwg.org/specs/rfc7540.html
func handleHTTP(s *Server, sd *StreamData) {
	// Ready data for logics & do some logic
	// - Route by URL
	// - Encode||Decode body by mime type header

	// Add Server Header to response : "ChaparKhane" || SCP means "SabzCityPlatform".

	// If project don't have any logic that support data on e.g. HTTP (restful, ...) we reject request with related error.
}
