mv /etc/init.d/postgresql /etc/init.d/postgresql-9.5:
  cmd.run

yum -y remove postgresql95-server:
  cmd.run
