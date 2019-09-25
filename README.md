# CaptainHook

Webhook event bridge and transformer. 

## Docs
 
- [Chaining](docs/chaining.md)


## TODO

- [ ] Improve failure logging & logging in general. (I.E when template post fails at dest, 
    when datamap is access where property doesn't exist). 
- [ ] Increase on unit testing.
- [ ] implement integration tests.
- [ ] Authentication 
- [ ] Brainstorm new destinations (Protocols, services) and create list of future enhancements. 
- [ ] Documentation on using & developing.

### Improvements

- on hookengine implement a fallback, where if destination is not a uri searches for hooks by name. rules can be 
simplified from https://<span/>localhost:... , to _hookname_
- new type of rule where a goroutine call can be made.
- new type of rule where a shell call can be made. 
