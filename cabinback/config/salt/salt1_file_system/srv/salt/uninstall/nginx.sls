include:
  {% if grains['os'] == 'CentOS' %}
  - uninstall.rpm.nginx
  {% elif grains['os'] == 'Ubuntu' %}

  {% endif %}