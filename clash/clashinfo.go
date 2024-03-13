package clash

/*
Clash basic config

must included fields:
	"version": string // daemon's version code
	"ua": string // daemon's user-agent when updating proxy providers

	"port": uin16 // http proxy port
	"socks-port": uint16 // socks5 port, may implement sokcs4a
	"mixed-port": uint16 // port handling both http and socks proxy
	"allow-lan": bool // allow visit from other addr than localhost
	"log": LogLevel

may included fields:
	"redir-port": uint16 // redirect transparent proxy port
	"tproxy-port": uint16 // linux tproxy port
	"tun": map[string]any // tun device config

library implemented fields:
	"ipv6": true
	"log-level": string // log-level
	"mode": "rule"|"direct"|"global"
	"bind-address": "*"

TODO: tun device config
*/
type ClashInfo map[string]any
