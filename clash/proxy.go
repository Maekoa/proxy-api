package clash

/*
proxy, added by parse config file and feeding

must included fields:

	"name": string //  name of proxy
	"type": string //  e.g. ShadowSocks, Trojan
	"udp": bool //  whether the proxy supports udp

library managed field:

	"history": []map[string]{
	    "time": time.Time,
	    "delay": int,
	    "meanDelay": int
	} // a list of latency
*/
type ProxyNode map[string]any
