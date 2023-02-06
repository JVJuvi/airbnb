package codetype

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleMember Role = "member"
	RoleGuest  Role = "guest"
)

func (r Role) IsValid() bool {
	switch r {
	case RoleAdmin, RoleMember, RoleGuest:
		return true
	default:
		return false
	}
}

func (r Role) IsAdmin() bool {
	if r == RoleAdmin {
		return true
	}

	return false
}
