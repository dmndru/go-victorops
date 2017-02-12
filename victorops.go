package victorops

import "log"

var logger *log.Logger // A logger that can be set by consumers
/*
  Added as a var so that we can change this for testing purposes
*/
const victorOpsAPI = "https://api.victorops.com/"

// Client is http api client
type Client struct {
	config struct {
		id  string
		key string
	}
	debug bool
}

// New initialize api client with api id and key
func New(id, key string, debug bool) *Client {
	s := &Client{}
	s.config.id = id
	s.config.key = key
	s.debug = debug
	return s
}

// SetLogger let's library users supply a logger, so that api debugging
// can be logged along with the application's debugging info.
func SetLogger(l *log.Logger) {
	logger = l
}
