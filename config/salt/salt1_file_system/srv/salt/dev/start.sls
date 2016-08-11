nginx.service:
  service.running:
    - name: nginx
    - enable: True

php-fpm.service:
  service.running:
    - name: php-fpm
    - enable: True

redis.service:
  service.running:
    - name: redis
    - enable: True

postgresql.service:
  service.running:
    - name: postgresql
    - enable: True

mysqld.service:
  service.running:
    - name: mysqld
    - enable: True

mongod.service:
  service.running:
    - name: mongod
    - enable: True


{% if grains['os'] == 'CentOS' %}
/etc/nginx:
  file.directory:
    - user: root
    - group: root
    - mode: 660
{% elif grains['os'] == 'Ubuntu' %}
{% endif %}

/tmp/dev_install.sh:
  file.managed:
    - source: salt://dev/source/scripts/dev_install.sh
    - user: root
    - group: root
    - mode: 755
  cmd.run:
    - user: root
    - shell: /bin/bash
    - env:
      - WORKER_USER: {{ pillar['dev']['user']['worker'] }}
