package webperms

type User interface {
	GetGroups() []Group
	GetUserPerms() Permissions
}


type Group struct {
	Name string
	Perms Permissions
}


type Permissions map[string]bool

func (p *Permissions) Set(name string, allow bool) {
	p[name] = bool
}

func (p *Permissions) Get(name string) bool {
	perm, present := p[name]
	if present {
		return perm
	}
	return false
}


func (u *User) HasPerm(permName string) bool {
	var perm bool
	var present bool
	// check user perms
	perm, present = u.GetUserPerms().Allow[permName]
	if present && perm {
		return true
	}
	// check group perms
	for i, g := range u.getGroups() {
		perm, present = g.Perms.Allow[permName]
		if present && perm {
			return true
		}
	}
	// so far 
	return false
}
