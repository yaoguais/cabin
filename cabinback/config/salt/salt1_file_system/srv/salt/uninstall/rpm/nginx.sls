yum -y remove nginx18:
  cmd.run

rm -f /etc/pki/rpm-gpg/RPM-GPG-KEY-webtatic-andy-nginx:
  cmd.run

rm -f /etc/yum.repos.d/webtatic-nginx.repo:
  cmd.run