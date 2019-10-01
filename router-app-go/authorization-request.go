/* For license and copyright information please see LEGAL file in repository */

package approuter

// AuthorizeWhich will authorize by ConnectionData.AccessControl.Which
func (sd *StreamData) AuthorizeWhich() {
	// Check requested user have enough access
	var notAuthorize bool
	for _, service := range sd.ConnectionData.AccessControl.Which {
		if service == sd.ServiceID {
			notAuthorize = false
			break
		} else {
			notAuthorize = true
		}
	}
	if notAuthorize == true {
		// sd.Err =
		return
	}
}

// AuthorizeWhen will authorize by ConnectionData.AccessControl.When
func (sd *StreamData) AuthorizeWhen() {}

// AuthorizeWhere will authorize by ConnectionData.AccessControl.Where
func (sd *StreamData) AuthorizeWhere() {
	var notAuthorize bool
	for _, ip := range sd.ConnectionData.AccessControl.Where {
		// TODO : ip may contain zero padding!! org may restricted user to isp not subnet nor even device!!
		if ip == sd.UIP.SourceIPAddress {
			notAuthorize = false
			break
		} else {
			notAuthorize = true
		}
	}
	if notAuthorize == true {
		// sd.Err =
		return
	}
}
