package models

type AdminUserModel struct {
	ID       int64  `db:"id" dt:"index:1; order:desc"`
	Username string `db:"username" dt:"index:2; order:asc"`
	Password string `db:"password"`
	Role     int64  `db:"role" dt:"index:3; sortable:false"`
	_        string `dt:"table:admin_users"`
}

type AdminRoleModel struct {
	ID         int64  `db:"id" json:"id" dt:"index:1"`
	Name       string `db:"name" json:"name" dt:"index:2"`
	Privileges string `db:"privileges" json:"privileges" dt:"index:3"`
	_          string `dt:"table:admin_roles"`
}

type AdminPrivilegeModel struct {
	ID        int64  `db:"id" json:"id" dt:"index:1"`
	Name      string `db:"name" json:"name" dt:"index:2"`
	Group     string `db:"group" json:"group" dt:"index:3"`
	Privilege string `db:"privilege" json:"privilege" dt:"index:4"`
	_         string `dt:"table:admin_privileges"`
}
