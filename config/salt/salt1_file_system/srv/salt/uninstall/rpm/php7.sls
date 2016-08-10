rm -f /etc/pki/rpm-gpg/RPM-GPG-KEY-webtatic-andy-php7:
  cmd.run

rm -f /etc/yum.repos.d/webtatic-php7.repo:
  cmd.run

yum -y remove php70w:
  cmd.run