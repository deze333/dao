package dao

//------------------------------------------------------------
// Initialization
//------------------------------------------------------------

// Initializes the package. Provide map with the following values:
// "id": internal server ID, use "_" for default single server
// "server": ie "localhost"
// (not implemented) "port": optional port number
// "options": ie "connect=direct"
// "mode": ie "monotonic"
// "log": empty for no log, any value will activate logging
func Initialize(params map[string]string) (err error) {

	// Connect to the server
	err = connectToServer(params)
	return
}

func DeInitialize() {
	disconnectServers()
}
