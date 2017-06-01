pgdg95:
  pkgrepo.managed:
    - humanname: PostgreSQL 9.5 $releasever - $basearch
    - baseurl: https://download.postgresql.org/pub/repos/yum/9.5/redhat/rhel-$releasever-$basearch
    - gpgcheck: 1
    - gpgkey: file:///etc/pki/rpm-gpg/RPM-GPG-KEY-PGDG-95

/etc/pki/rpm-gpg/RPM-GPG-KEY-PGDG-95:
  file.managed:
    - source: salt://install/rpm/files/etc/pki/rpm-gpg/RPM-GPG-KEY-PGDG-95
    - user: root
    - group: root
    - mode: 644

postgresql.packages:
  pkg.installed:
    - pkgs:
      - postgresql95-server

if [ -f /etc/init.d/postgresql-9.5 ]; then mv /etc/init.d/postgresql-9.5 /etc/init.d/postgresql; fi:
  cmd.run


service postgresql initdb:
  cmd.run


# postgresql 初始化数据库,首先需要切换到postgres用户登录服务器
# sudo su - postgres
# 然后使用 psql 命令登录
# 使用 \password postgres 设置postgres用户的密码, 这里设置成111111
# 创建用户admin
# CREATE USER admin WITH PASSWORD '111111';
# 创建数据库并授权
# CREATE DATABASE fame OWNER admin;
# GRANT ALL PRIVILEGES ON DATABASE fame to admin;
# 输入\l查看数据库信息
# 输入\q退出服务器
# 修改 /var/lib/pgsql/9.5/data/pg_hba.conf 中的3行,将ident修改为trust
# local   all             all                                     trust
# IPv4 local connections:
# host    all             all             127.0.0.1/32            trust
# IPv6 local connections:
# host    all             all             ::1/128                 stust
# 重启服务器
# 然后输入 psql -d fame -U admin -W 登录(登录后不能切换数据库)
#
