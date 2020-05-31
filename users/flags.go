package users

// These constants define the bitflags a user may have
const (
	Flagdministrator = 0x01
	FlagModerator    = 0x02
	FlagBlacklisted  = 0x04
)

// HasFlag checks whether or not the current user has at least one of the given flags
func (user *User) HasFlag(flags ...int) bool {
	for _, flag := range flags {
		if user.Flags&flag == flag {
			return true
		}
	}
	return false
}

// AssignFlag gives the current user the given flag
func (user *User) AssignFlag(flag int) {
	user.Flags |= flag
}

// RevokeFlag revokes the given flag from the current user
func (user *User) RevokeFlag(flag int) {
	user.Flags &= ^flag
}
