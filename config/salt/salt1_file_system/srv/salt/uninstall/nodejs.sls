include:
  {% if grains['os'] == 'CentOS' %}
  - uninstall.rpm.nodejs
  {% elif grains['os'] == 'Ubuntu' %}

  {% endif %}