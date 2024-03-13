/*
*
Clash uses chains of proxy group, enabling flexible switch of both ruleset
and proxy, though webui at present don't support direct modify on rules.

However, a typical proxy daemon usually supports flattened configuration on
inbounds, outbounds, and routing rules.

We provide a support to clash-like proxy groups via providers, thus supporting
existing webui and software, without sacrificing flexibility.

We agree that
  - there will be a default local-stored config file, loads when daemon start
  - provider MUST provide daemon-specific raw data of proxy data, and MAY provide
    daemon-specific config fields.
  - there is several equivalent inbounds set by clash basic config
  - rules are managed by this library, while routing is implemented by daemon

this project will manage
- groups and providers
- rule sets
-
-
*
*/
package clash

import (
	"github.com/go-chi/chi/v5"
	"github.com/maekoa/proxy-api/routing"
)

// assume all inbounds shares the same rule of outbound routing rules
//
// MUST NOT close old connections when updating, (only ReloadConfig MAY close them)
type ClashCompatible interface {
	Logger() chan Event      // channel of (loglevel, string)
	Transferred() (int, int) // (up, down) in Bit
	Config() ClashInfo       // clash-compatible json format
	RefreshListener(ClashInfo) error
	RestartWith(string) error // load from local file path (optional implement on url)

	// proxy node

	Proxies() map[string]ProxyNode                // `map[proxy id](proxy content)`
	TestProxy(id string, url string) (int, error) // return latency in ms, url has been checked
	RemoveProxies(proxies []string)               // MUST NOT close related connection
	Feed(string) (map[string]ProxyNode, error)    // load proxies from proxy provider, and returns added node

	// routing

	SetDefaultOutbound(id string) error // set default outbound proxy
	Routings() []*routing.ProxyRule
	AddRoutings([]*routing.ProxyRule) error
	RemoveRoutings([]*routing.ProxyRule) error // clear all the rules

	// connections
	Connections() map[string]Connection
	CloseConnections(ids []string) error

	AddExtra(*chi.Mux) // you can manually add path to api handler
}

// map is the basic of so-called "inherit", and therefore the best choice to implement extensive data
