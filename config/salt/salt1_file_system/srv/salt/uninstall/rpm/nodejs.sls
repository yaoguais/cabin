rm -f /etc/yum.repos.d/nodesource-el.repo:
  cmd.run

yum -y remove nodejs:
  cmd.run