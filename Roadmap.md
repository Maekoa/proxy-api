# Roadmap

- [ ] Clash API
- [ ] V2Ray API

## Clash API
channel:
- log /logs get

tick
- traffic /traffic get

one-off
- version /version get
- (optional) flush fake ip /cache/fakeip/flush post
- config /configs 
  - get (only basic config)
  - put {"path": "", "payload": ""} (reload config from)
  - patch {...config}
- proxies /proxies get
- proxy /proxies/:name
  - get
  - put {"name":""} (select name node)
- delay /proxies/:name/delay get ?timeout(ms) ?url 
- rules /rules get
- connections /connections
  - get
  - delete
- close a connection /connections/:id delete
- proxy providers/group /providers/proxies
  - get
  - \*cover providers put
  - \*patch providers patch
- proxy provider/group /providers/proxies/:name 
  - get
  - put (update provider)
  - \*delete (delete provider)
- provider/group health check /providers/proxies/:name/healthcheck get
- rule providers /providers/rules 
  - get
  - \*cover providers put
  - \*patch providers patch
- rule provider /providers/rules/:name
  - get
  - put (update provider)
  - \*delete (delete provider)