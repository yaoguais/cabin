include:
  - install.base
  - install.nginx
  - install.php7
  - install.redis
  - install.phpmyadmin
  - install.mongodb
  - install.postgresql
  - install.mysql

user.worker:
  user.present:
    - name: {{ pillar['dev']['user']['worker'] }}
    - shell: /bin/bash
    - createhome: False

