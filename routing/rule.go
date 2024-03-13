package routing

// rule types
type RuleType uint8 // see consts.go for detailed definition

// assume payload is string
//
// rules can be operated on unit of ProxyRule.
// the proxyrule MUST NOT modify since its built,
// reducing lock cost.
type ProxyRule struct {
	RuleSets  map[RuleType]Set      // some type of large rule lists' key are partially ordered, so set can work faster
	RuleLists map[RuleType][]string // all other rules, mismatch unless every rule is satisfied
	Proxy     []string              // proxy id(s), if len>1, it should be a load balancer
	Resolve   bool                  // whether resolve a domain
}
