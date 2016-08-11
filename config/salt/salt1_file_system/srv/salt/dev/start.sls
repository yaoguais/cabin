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

supervisor.service:
  service.running:
    - name: supervisor
    - enable: True

/etc/supervisor:
  file.recurse:
    - source: salt://dev/source/files/etc/supervisor
    - user: root
    - group: root
    - file_mode: 660
    - dir_mode: 660

{% if grains['os'] == 'CentOS' %}
/etc/nginx:
  file.recurse:
    - source: salt://dev/rpm/files/etc/nginx
    - user: root
    - group: root
    - file_mode: 660
    - dir_mode: 660
{% elif grains['os'] == 'Ubuntu' %}
{% endif %}

service nginx reload:
  cmd.run:
    - user: root

{% if grains['os'] == 'CentOS' %}
/etc/php.d:
  file.recurse:
    - source: salt://dev/rpm/files/etc/php.d
    - user: root
    - group: root
    - file_mode: 660
    - dir_mode: 660

/etc/php-fpm.d:
  file.recurse:
    - source: salt://dev/rpm/files/etc/php-fpm.d
    - user: root
    - group: root
    - file_mode: 660
    - dir_mode: 660

/etc/php-zts.d:
  file.recurse:
    - source: salt://dev/rpm/files/etc/php-zts.d
    - user: root
    - group: root
    - file_mode: 660
    - dir_mode: 660

/etc/php-fpm.conf:
  file.managed:
    - source: salt://dev/rpm/files/etc/php-fpm.conf
    - user: root
    - group: root
    - mode: 644

/etc/php.ini:
  file.managed:
    - source: salt://dev/rpm/files/etc/php.ini
    - user: root
    - group: root
    - mode: 644

service php-fpm reload:
  cmd.run

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
