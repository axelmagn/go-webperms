
package webperms


type User interface {
	Groups() []Group
	Perms() Permissions
}


// Can evaluates whether a user has access to a given permission
func (u *User) Can(perm string) bool {
	var allow bool
	// check user permissions
	if u.Perms().Get(perm) {
		return true
	}
	// check group permissions
	for i, g := range u.Perms() {
		if g.GetPerm(perm) {
			return true
		}
	}
}
