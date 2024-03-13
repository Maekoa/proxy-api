package clash

/*
connection

must included fields:
    "host": string // http/tls host
    "network": string // l4 protocol, e.g. tcp udp
    "sourceIP": string // source ip
    "destinationIP": string
    "sourcePort": string
    "destinationPort": string
    "upload": int // total uploaded data in Bit
    "download": int // total downloaded data in Bit
    "start": time.Time // start time
    "proxy": string // node id
    "rule": RuleType
    "rulePayload": string

suggested fields:
    "processPath": string // dialer process of connection
*/
// connection
type Connection map[string]any
