/* For license and copyright information please see LEGAL file in repository */

package approuter

// ReqRegisterGuestConnection :
type ReqRegisterGuestConnection struct {
	ConnectionPublicKey  [32]byte `valid:"PublicKey"`
	SuggestedCipherSuite string
	UserAgent            string
}

// RegisterGuestConnection : Make new connection by ECC algorithm.
func (s *Server) RegisterGuestConnection(req *ReqRegisterGuestConnection) error {
	// Make ConnectionData.EncryptionKey by ECC algorithm with :
	// req.ConnectionPublicKey
	// RunningServerData.PublicKeyCryptography.PrivateKey
	return nil
}
