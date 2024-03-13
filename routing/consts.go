package routing

const (
	MATCH  RuleType = iota // set default
	DOMAIN                 // MUST be puny_coded
	DOMAIN_SUFFIX
	DOMAIN_KEYWORD
	DOMAIN_REGEX
	GEOSITE

	IP_CIDR
	IPv6_CIDR
	GEOIP
)

// optional types, if not implemented, than retrurn mismatch
const (
	PORT RuleType = iota + 128 // range(#d-#d) or single port(#d)
	SRC_IP_CIDR
	SRC_IPv6_CIDR
	SRC_PORT
	L4_PROTO   // tcp,udp
	APP_PROTO  // tls, https, bt, ...
	INBOUND_ID // multiple inbound
	PROCESS    // process

	EXTEND RuleType = 255
)

var RuleTypeName = map[RuleType]string{
	MATCH:          "Match",
	DOMAIN:         "Domain",
	DOMAIN_SUFFIX:  "Domain Suffix",
	DOMAIN_KEYWORD: "Domain Keyword",
	DOMAIN_REGEX:   "Domain Regex",
	GEOSITE:        "GeoSite",

	IP_CIDR:       "Destination IP CIDR",
	IPv6_CIDR:     "Destination IPv6 CIDR",
	GEOIP:         "GeoIP",
	PORT:          "Destination Port",
	SRC_IP_CIDR:   "Source IP CIDR",
	SRC_IPv6_CIDR: "Source IPv6 CIDR",
	SRC_PORT:      "Source Port",
	L4_PROTO:      "Layer-4 Protocol",     // tcp,udp
	APP_PROTO:     "Application Protocol", // tls, https, bt, ...
	INBOUND_ID:    "Inbound ID",           // multiple inbound
	PROCESS:       "Process Name",

	EXTEND: "Extension",
}

func (p RuleType) String() string {
	return RuleTypeName[p]
}
