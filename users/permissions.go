package users

// Permission represents a permission
type Permission int

// These constants define the bitflags which represent the permissions a user may have
const (
	PermissionAdministrator = 0x01
	PermissionModerator     = 0x02
)

// HasPermission checks whether or not the current user has at least one of the given permissions
func (user *User) HasPermission(permissions ...Permission) bool {
	for _, permission := range permissions {
		if user.Permissions&int(permission) == int(permission) {
			return true
		}
	}
	return false
}

// GrantPermission gives the current user the given permission
func (user *User) GrantPermission(permission Permission) {
	user.Permissions |= int(permission)
}

// RevokePermission revokes the given permission from the current user
func (user *User) RevokePermission(permission Permission) {
	user.Permissions &= ^int(permission)
}
