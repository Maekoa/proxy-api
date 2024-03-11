package clash

import "github.com/go-chi/chi/v5"

// assume all inbounds shares the same rule of outbound routing rules
//
// NEVER close old connections when updating, (only ReloadConfig MAY close connections)
type ClashCompatible interface {
	Logger() chan Event      // channel of (loglevel, string)
	Speed() (int, int)       // (up, down) in Bit per second
	Transferred() (int, int) // (up, down) in Bit
	Version() string         // version of clash

	// config in json format
	//
	// we suppose implement as https://github.com/MetaCubeX/mihomo/blob/meta-main/hub/route/configs.go
	Config() string                // clash-compatible json format
	SetConfig(string) error        // patches original config, in json, do not close old connections
	ReloadConfig(string) error     // covers original config, in json, do not close old connections
	ReloadConfigFile(string) error // load from local file path (optional implement on url)

	// proxy node

	// `map[proxy id](proxy content)`
	//
	// MUST provide id called "direct" and "reject"
	Proxies() map[string]ProxyNode
	TestProxy(id string) (int, error) // return latency in ms

	// optional: load balancer
	BalancerImplemented() bool         // return false if not implemented
	Balancers() map[string][]ProxyNode // id should share same field with proxy node
	AddBalancer([]ProxyNode) (id string, err error)
	RemoveBalancer(id string) error

	// routing
	SetDefaultOutbound(id string) error // set default outbound proxy
	Routings() []*Rule
	AddRoutings([]*Rule) error // should remove added rule if error occurs
	ClearRoutings() error      // clear all the rules

	// connections
	Connections() map[string]Connection
	CloseConnection(id string) error
	CloseAll() error // close all connections

	AddExtra(*chi.Mux) // you can manually add path to api handler
}

// log
type Event struct {
	LogLevel int
	Payload  string
}

const (
	DEBUG int = iota
	INFO
	WARNING
	ERROR
	SILENT
)

// proxy
type ProxyNode interface {
	Name() string       // name of proxy
	ConfigJSON() string // generate clash compatible config json

	Type() string // e.g. ShadowSocks, Trojan
	UDP() bool    // whether the node supports udp
}

// rule
type Rule struct {
	Type  string   // rule type
	Value string   // rule value
	Proxy []string // proxy balancer id list
	And   *Rule    // rule that should also meet, implement is optional
} // for example DOMAIN,ad.com,REJECT

// connection
type Connection interface {
	Host() string                  // `json: host`
	Protocol() string              // `json: "network"` l4 protocol, e.g. tcp udp
	Src() (ip string, port string) // `json: source(IP|Port)` source
	Dst() (ip string, port string) // `json: destination(IP|Port)` destination
	Transferred() (int, int)       // (up, down) in Bit
	Since() string                 // time string in format 2024-01-01T00:00:00.0000000+08:00
	Match() *Rule                  // rule that matchs
}
