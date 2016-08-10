rm -f /etc/pki/rpm-gpg/RPM-GPG-KEY-mysql:
  cmd.run

yum -y remove mysql-community-server mysql-community-client:
  cmd.run
