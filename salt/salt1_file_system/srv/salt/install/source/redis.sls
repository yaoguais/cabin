/tmp/redis_install.sh:
  file.managed:
    - source: salt://install/source/scripts/redis_install.sh
    - user: root
    - group: root
    - mode: 755
  cmd.run:
    - user: root
    - shell: /bin/bash 

/etc/redis.conf:
  file.managed:
    - source: salt://install/source/files/etc/redis.conf
    - user: root
    - group: root
    - mode: 660

/etc/init.d/redis:
  file.managed:
    - source: salt://install/source/files/etc/init.d/redis
    - user: root
    - group: root
    - mode: 755
