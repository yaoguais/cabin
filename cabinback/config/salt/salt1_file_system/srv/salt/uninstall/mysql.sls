include:
  {% if grains['os'] == 'CentOS' %}
  - uninstall.rpm.mysql
  {% elif grains['os'] == 'Ubuntu' %}

  {% endif %}