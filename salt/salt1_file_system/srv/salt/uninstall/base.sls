include:
  {% if grains['os'] == 'CentOS' %}
  - uninstall.rpm.base
  {% elif grains['os'] == 'Ubuntu' %}

  {% endif %}