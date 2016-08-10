include:
  {% if grains['os'] == 'CentOS' %}
  - uninstall.rpm.mongodb
  {% elif grains['os'] == 'Ubuntu' %}

  {% endif %}