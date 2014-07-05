
package webperms

type Permissions interface {
	Get(perm string) bool
	Set(perm string, allow bool)
}


type BasicPermissions map[string]bool


func NewPermissions() Permissions {
	return map[string]bool
}


func (p *BasicPermissions) Set(name string, allow bool) {
	p[name] = bool
}


func (p *BasicPermissions) Get(name string) bool {
	perm, present := p[name]
	if present {
		return perm
	}
	return false
}
