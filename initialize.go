package dao

//------------------------------------------------------------
// Initialization
//------------------------------------------------------------

// Creates DB server connection.
// Provide map with the following values:
// "id": internal server ID, use "_" for default single server
// "server": ie "localhost"
// (not implemented) "port": optional port number
// "options": ie "connect=direct"
// "mode": ie "monotonic"
// "log": empty for no log, any string value will activate logging
func ServerConnect(params map[string]string) (err error) {

	err = connectToServer(params)
	return
}

// Gracefully closes servers connections.
func ServersDisconnect() {
	disconnectServers()
}
