base.packages:
  pkg.installed:
    - pkgs:
      - wget
      - ntpdate

cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime:
  cmd.run

ntpdate us.pool.ntp.org:
  cmd.run

yum remove git:
  cmd.run

yum -y install http://mirrors.neusoft.edu.cn/repoforge/redhat/el6/en/x86_64/extras/RPMS/perl-Git-1.7.12.4-1.el6.rfx.x86_64.rpm http://mirrors.neusoft.edu.cn/repoforge/redhat/el6/en/x86_64/extras/RPMS/git-1.7.12.4-1.el6.rfx.x86_64.rpm:
  cmd.run


