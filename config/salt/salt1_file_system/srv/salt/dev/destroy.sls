include:
  - uninstall.base
  - uninstall.nginx
  - uninstall.php7
  - uninstall.redis
  - uninstall.phpmyadmin
  - uninstall.mongodb
  - uninstall.postgresql
  - uninstall.mysql


user.worker:
  user.absent:
    - name: {{ pillar['dev']['user']['worker'] }}

