/* For license and copyright information please see LEGAL file in repository */

package approuter

// CloseConnection : Use to close connection and drop all incomplete data & stream.
// Connection can be open again by use TransferConnection() if related service exist in platform!
func (s *Server) CloseConnection() {}
