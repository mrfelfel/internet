/* For license and copyright information please see LEGAL file in repository */

package approuter

// Manifest use to store service manifest data
var Manifest AppManifest

// AppManifest : Object to store application service information.
type AppManifest struct {
	AppID            [16]byte // Application ID
	Name             string
	Domain           string
	Icon             string
	Description      string
	TermsOfService   string
	TAGS             []string // Use to categorized apps e.g. Music, GPS, ...
	Contact          AppContact
	Licence          AppLicence
	Services         Services
	BuildInformation AppBuildInformation
	MinHardwareSpec  AppMinHardwareSpec
	// 4 type of devices trend now:
	// Out of XP scope - Cloud Computing - ::/1 (Cheapest scenario)
	// Adjacent XP - Fog Computing - ::/32
	// Inside XP scope - Edge Computing - ::/64
	// End user device - ::/128
	MaxNetworkScalability uint8 // If ::/32 it mean also ::/1 but not ::/64!!
}

// AppContact :
type AppContact struct {
	Name  string
	URL   string
	Email string
}

// AppLicence :
type AppLicence struct {
	Name string
	URL  string
}

// Services use to store some information about APIs services
type Services struct {
	TotalServices uint32
	LastID        uint32
}

// AppBuildInformation :
type AppBuildInformation struct {
	OS                   []string // PersiaOS, Linux, Windows, ...
	NetworkLayer         []string // IPv4, IPv6, ... (Now support just IPv6)
	TransportProtocols   []string // SCP, TCP, UDP, ... (Now support just SCP)
	ApplicationProtocols []string // sRPC, HTTP, ...
	UseAI                bool     // false: Open the lot of security concerns but use more resource.
	PublicApp            bool     // false: chaparkhane not accept guest user creation.
	AuthorizationServer  string   // Domain name that have sRPC needed store connection data. default is "ConnectionAuth.sabz.city"
	GuestDailyRateLimit  uint32   // Max open stream per day.
	PersonDailyRateLimit uint32   // Max open stream per day.
	OrgDailyRateLimit    uint32   // Max open stream per day. 0 means no limit and good for PayAsGo strategy!
}

// AppPermission :
type AppPermission struct {
	IP                bool // Internet. Charge user if network available!
	Notification      bool
	RunInBackground   bool
	UseIPInBackground bool
	Camera            bool
	Location          bool // GPS, ...
	Speaker           bool
	Microphone        bool
	BodySensor        bool
	Bluetooth         bool
	USB               bool
	NFC               bool
	FingerPrint       bool
	Vibrate           bool
}

// AppMinHardwareSpec : Minimum hardware specification for each instance of application.
type AppMinHardwareSpec struct {
	CPU              uint64 // Hz
	RAM              uint64 // Byte
	GPU              uint64 // Hz
	NetworkBandwidth uint64 // Byte
	HDD              uint64 // Byte, Hard disk drive as storage engine. Act as object+block storage.
	SSD              uint64 // Byte, Solid state drive as storage engine. Act as object+block storage.
}
