package webperms

type Group interface {
	Name() string
	SetName(name string)
	GetPerm(perm string) bool
	SetPerm(perm string, allow bool)
}


type BasicGroup struct {
	name string
	perms Permissions
}


func NewGroup(name string) {
	return BasicGroup{
		name: name
		perms: NewPermissions()
	}
}


func (g *BasicGroup) Name() string {
	return g.name
}

func (g *BasicGroup) SetName(name string) {
	g.name = name
}

func (g *BasicGroup) GetPerm(perm string) bool {
	return g.perms.Get(perm)
}

func (g *BasicGroup) SetPerm(perm string, allow bool) {
	g.perms.Set(perm, allow)
}
