package services

import (
	"strconv"

	"github.com/Yaoguais/gadmin/core/models"
	. "github.com/Yaoguais/gadmin/lib/db"
	"github.com/Yaoguais/gadmin/lib/log"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var AdminService = &adminService{}

type adminService struct {
}

func (*adminService) GetAdminUserByUid(uid int64) *models.AdminUserModel {
	m := models.AdminUserModel{}
	err := Db.Get(&m, "SELECT * FROM admin_users WHERE id = ?", uid)

	if err != nil {
		log.Info("get admin user failed", err)
		return nil
	}

	return &m
}

func (*adminService) DeleteAdminUserByUid(uid int64) error {
	_, err := Db.Exec("DELETE FROM admin_users WHERE id = ?", uid)

	return err
}

func (*adminService) DeleteAdminRoleById(id int64) error {
	_, err := Db.Exec("DELETE FROM admin_roles WHERE id = ?", id)

	return err
}

func (*adminService) DeleteAdminPrivilegeById(id int64) error {
	_, err := Db.Exec("DELETE FROM admin_privileges WHERE id = ?", id)

	return err
}

func (*adminService) GetAdminUserByUsername(username string) *models.AdminUserModel {
	m := models.AdminUserModel{}
	err := Db.Get(&m, "SELECT * FROM admin_users WHERE username=?", username)

	if err != nil {
		log.Info("get admin user by username failed", err)
		return nil
	}

	return &m
}

func (*adminService) GetAdminUsers() []models.AdminUserModel {
	m := []models.AdminUserModel{}
	err := Db.Select(&m, "SELECT * FROM admin_users")

	if err != nil {
		log.Info("get admin users failed", err)
		return nil
	}

	return m
}

func (*adminService) GetAdminRoles() []models.AdminRoleModel {
	m := []models.AdminRoleModel{}
	err := Db.Select(&m, "SELECT * FROM admin_roles")

	if err != nil {
		log.Info("get admin roles failed", err)
		return nil
	}

	return m
}

func (*adminService) GetAdminPrivileges() []models.AdminPrivilegeModel {
	m := []models.AdminPrivilegeModel{}
	err := Db.Select(&m, "SELECT * FROM admin_privileges order by `group`")

	if err != nil {
		log.Info("get admin privileges failed", err)
		return nil
	}

	return m
}

func (*adminService) GetAdminRoleById(id int64) *models.AdminRoleModel {
	m := models.AdminRoleModel{}
	err := Db.Get(&m, "SELECT * FROM admin_roles WHERE id = ?", id)

	if err != nil {
		log.Info("get admin role failed", err)
		return nil
	}

	return &m
}

func (*adminService) GetAdminPrivilegeById(id int64) *models.AdminPrivilegeModel {
	m := models.AdminPrivilegeModel{}
	err := Db.Get(&m, "SELECT * FROM admin_privileges WHERE id = ?", id)

	if err != nil {
		log.Info("get admin user failed", err)
		return nil
	}

	return &m
}

func (*adminService) MakePassword(password string) string {
	pwdByte := []byte(password)
	encodePwd, err := bcrypt.GenerateFromPassword(pwdByte, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(encodePwd)
}

func (*adminService) CheckPassword(savedPassword string, inputPassword string) bool {
	savedPwdByte := []byte(savedPassword)
	inputPwdByte := []byte(inputPassword)

	return bcrypt.CompareHashAndPassword(savedPwdByte, inputPwdByte) == nil
}

func (*adminService) SaveUserSession(user *models.AdminUserModel, s *sessions.Session) {
	s.Values["id"] = strconv.FormatInt(user.ID, 10)
	s.Options.HttpOnly = true
	s.Options.MaxAge = 30 * 86400
}

func (as *adminService) UpdateUser(u *models.AdminUserModel) error {
	sql := "UPDATE admin_users SET "
	var params []interface{}
	if len(u.Username) > 0 {
		sql += "username = ?,"
		params = append(params, u.Username)
	}

	if len(u.Password) > 0 {
		sql += "password = ?,"
		params = append(params, as.MakePassword(u.Password))
	}

	if u.Role > 0 {
		sql += "role = ?,"
		params = append(params, u.Role)
	}

	sql = sql[0 : len(sql)-1]

	sql += " where id = ?"
	params = append(params, u.ID)

	_, err := Db.DB.Exec(sql, params...)

	return err
}
