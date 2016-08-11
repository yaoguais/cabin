base.packages:
  pkg.installed:
    - pkgs:
      - wget
      - vim
      - ntpdate
      - strace

cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime:
  cmd.run

ntpdate us.pool.ntp.org:
  cmd.run

/tmp/git_install.sh:
  file.managed:
    - source: salt://install/rpm/scripts/git_install.sh
    - user: root
    - group: root
    - mode: 755
  cmd.run:
    - user: root
    - shell: /bin/bash

/etc/environment:
  file.managed:
    - source: salt://install/rpm/files/etc/environment
    - user: root
    - group: root
    - mode: 644

