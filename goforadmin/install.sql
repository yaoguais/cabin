CREATE DATABASE project;
USE project;

CREATE TABLE IF NOT EXISTS admin_users (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  username varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  password varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  role bigint(20) unsigned NOT NULL DEFAULT '0', /*角色: 1-超级管理员 2-管理员*/
  PRIMARY KEY (id),
  UNIQUE KEY (username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS admin_roles (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '', /*角色名称*/
  privileges text COLLATE utf8mb4_unicode_ci NOT NULL, /*权限ID列表,用逗号分隔*/
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS admin_privileges (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '', /*权限名称*/
  `group` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '', /*权限组*/
  /*权限, "POST requestURI"或者"GET requestURI"*/
  privilege text COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

/*账号:admin 密码:111111 */
/*账号:test  密码:111111 */
INSERT INTO admin_users(id, username, password, role) VALUES
  ('1', 'admin', '$2a$10$Gu3S.gTUXGz.GtcJgCaih.f256n1c/whZb7oKWOMQhKN5R8dwWRAu', '1'),
  ('2', 'manager', '$2a$10$Gu3S.gTUXGz.GtcJgCaih.f256n1c/whZb7oKWOMQhKN5R8dwWRAu', '2');

INSERT INTO admin_roles(id, name, privileges) VALUES
  ('1', '超级管理员', ''),
  ('2', '管理员', '1,2');

INSERT INTO admin_privileges(id, name, `group`, privilege) VALUES
  ('1', '查看首页', '权限模块', 'GET /'),
  ('2', '退出登录', '权限模块', 'GET /logout');
