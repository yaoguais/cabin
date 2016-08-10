include:
  {% if grains['os'] == 'CentOS' %}
  - uninstall.rpm.postgresql
  {% elif grains['os'] == 'Ubuntu' %}

  {% endif %}