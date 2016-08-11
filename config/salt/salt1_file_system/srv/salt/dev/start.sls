##############################
##### 1. 保证Service处于启动状态
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

######################
###### 2. 部署PHP文件
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

############################
###### 3. 部署php-fpm文件
{% if grains['os'] == 'CentOS' %}
/etc/php.d:
  file.recurse:
    - source: salt://dev/rpm/files/etc/php.d
    - user: {{ pillar['dev']['user']['worker'] }}
    - group: root
    - file_mode: 660
    - dir_mode: 750

/etc/php-fpm.d:
  file.recurse:
    - source: salt://dev/rpm/files/etc/php-fpm.d
    - user: {{ pillar['dev']['user']['worker'] }}
    - group: root
    - file_mode: 660
    - dir_mode: 750

/etc/php-zts.d:
  file.recurse:
    - source: salt://dev/rpm/files/etc/php-zts.d
    - user: {{ pillar['dev']['user']['worker'] }}
    - group: root
    - file_mode: 660
    - dir_mode: 750

/etc/php-fpm.conf:
  file.managed:
    - source: salt://dev/rpm/files/etc/php-fpm.conf
    - user: {{ pillar['dev']['user']['worker'] }}
    - group: root
    - mode: 660

/etc/php.ini:
  file.managed:
    - source: salt://dev/rpm/files/etc/php.ini
    - user: {{ pillar['dev']['user']['worker'] }}
    - group: root
    - mode: 660

{% elif grains['os'] == 'Ubuntu' %}
{% endif %}

############################
###### 3. 部署nginx文件
{% if grains['os'] == 'CentOS' %}
/etc/nginx:
  file.recurse:
    - source: salt://dev/rpm/files/etc/nginx
    - user: {{ pillar['dev']['user']['worker'] }}
    - group: root
    - file_mode: 660
    - dir_mode: 750
{% elif grains['os'] == 'Ubuntu' %}
{% endif %}

############################
###### 3. 部署supervisor文件
/etc/supervisor:
  file.recurse:
    - source: salt://dev/source/files/etc/supervisor
    - user: {{ pillar['dev']['user']['worker'] }}
    - group: root
    - file_mode: 660
    - dir_mode: 750

############################
###### 4. 重新载入php-fpm服务
{% if grains['os'] == 'CentOS' %}
service php-fpm reload:
  cmd.run
{% elif grains['os'] == 'Ubuntu' %}
{% endif %}

#######################################################
###### 5. 重新载入supervisor服务
## php-fpm在前是因为supervisor可能依赖php-fpm提供的web接口
#######################################################
supervisord.reread:
  module.run:
    - user: root

fame_push_notify:0:
  supervisord.running:
    - restart: true
    - user: root

############################
###### 6. 重新载入nginx服务
service nginx reload:
  cmd.run:
    - user: root

#############################
##### 7. 更新计划任务
{% if grains['os'] == 'CentOS' %}
php /data/server/fame/artisan schedule:run:
  cron.present:
    - minute: '*'
    - user: {{ pillar['dev']['user']['worker'] }}
    - identifier: fame_transaction_watch
    - comment: app crontab tasks run everymin
{% elif grains['os'] == 'Ubuntu' %}
{% endif %}
