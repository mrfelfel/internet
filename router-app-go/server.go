/* For license and copyright information please see LEGAL file in repository */

package approuter

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server represents an ChaparKhane server needed data to serving as server.
type Server struct {
	Status                int            // 0:stop 1:running
	GracefulStop          chan os.Signal // GracefulStop is a channel of os.Signals that we will watch for -SIGTERM
	Network               Network
	Handlers              Handlers
	PublicKeyCryptography AppPublicKeyCryptography
	ConnectionPool        map[[16]byte]*ConnectionData // --TODO-- Just one process can give new ConnectionID due concurrency problems! or lock needs
	Assets                *Assets                      // Data in Assets dependency(folder) of repo
}

// Network :
type Network struct {
	UIPRange [14]byte
	MTU      uint16 // Maximum Transmission Unit. network+transport+application
}

// Handlers use to store related data use in handle packet from start to end!
type Handlers struct {
	UIPHandler PacketHandler
	SCPHandler PacketHandler
	Services   map[uint32]ServiceFunc
}

// PacketHandler use to register packet handler!
type PacketHandler func(*Server, []byte)

// ServiceFunc use to register server public services (APIs) just for DefaultServer and not use in generator!
type ServiceFunc func(*StreamData)

// AppPublicKeyCryptography : Public-key for related domain.
type AppPublicKeyCryptography struct {
	PublicKey  [32]byte // Use new algorithm like 256bit ECC(256bit) instead of RSA(4096bit)
	PrivateKey [32]byte // Use new algorithm like 256bit ECC(256bit) instead of RSA(4096bit)
}

// NewServer will make new server object
func NewServer() *Server {
	var s = Server{
		GracefulStop:   make(chan os.Signal),
		ConnectionPool: make(map[[16]byte]*ConnectionData),
	}

	s.Handlers.Services = make(map[uint32]ServiceFunc)

	// make public & private key and store them
	s.PublicKeyCryptography.PublicKey = [32]byte{}
	s.PublicKeyCryptography.PrivateKey = [32]byte{}

	return &s
}

// RegisterPublicKey use to register public key in apis.sabz.city
func (s *Server) RegisterPublicKey() (err error) {
	return nil
}

// RegisterUIP use to get new UIP & MTU from OS router!
func (s *Server) RegisterUIP() (err error) {
	// send PublicKey to router and get IP if user granted. otherwise log error.
	s.Network.UIPRange = [14]byte{}

	// Get MTU from router
	s.Network.MTU = 1200

	// Because ChaparKhane is server based application must have IP access.
	// otherwise close server app and return err

	return nil
}

// Start will start the server.
func (s *Server) Start() (err error) {
	// Tell others server start!
	s.Status = 1

	// watch for SIGTERM and SIGINT from the operating system, and notify the app on the channel
	signal.Notify(s.GracefulStop, syscall.SIGTERM)
	signal.Notify(s.GracefulStop, syscall.SIGINT)
	go func() {
		// wait for our os signal to stop the app
		// on the graceful stop channel
		// this goroutine will block until we get an OS signal
		var sig = <-s.GracefulStop
		fmt.Printf("caught sig: %+v", sig)

		// sleep for 60 seconds to waiting for app to finish,
		fmt.Println("Waiting for server to finish, will take 60 seconds")
		time.Sleep(60 * time.Second)

		s.Shutdown()

		os.Exit(s.Status)
	}()

	// Get UserGivenPermission from OS

	// Make & Register publicKey
	if err = s.RegisterPublicKey(); err != nil {
		return err
	}

	// Get IP & MTU from OS router. Register income packet service.
	if err = s.RegisterIP(); err != nil {
		return err
	}

	// register s.Handlers.UIPHandler for income packet handler

	return nil
}

// Shutdown use to graceful stop server!!
func (s *Server) Shutdown() {
	// ... Do business Logic for shutdown
	// Shutdown works by:
	// first closing open listener for income packet and refuse all new packet,
	// then closing all idle connections,
	// and then waiting indefinitely for connections to return to idle
	// and then shut down

	// Send signal to DNS & Certificate server to revoke app data.

	// it must change to 0 otherwise it means app can't close normally
	s.Status = 1
}
