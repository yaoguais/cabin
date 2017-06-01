include:
  {% if grains['os'] == 'CentOS' %}
  - uninstall.rpm.php7
  {% elif grains['os'] == 'Ubuntu' %}

  {% endif %}

rm -f /usr/local/bin/composer:
  cmd.run:
    - user: root

