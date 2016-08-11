include:
  {% if grains['os'] == 'CentOS' %}
  - install.rpm.php7
  {% elif grains['os'] == 'Ubuntu' %}

  {% endif %}


# install composer
if [ ! -f /usr/local/bin/composer ]; then cd /tmp;curl -sS https://getcomposer.org/installer | php; mv composer.phar /usr/local/bin/composer; fi:
  cmd.run:
    - user: root