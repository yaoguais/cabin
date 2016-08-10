include:
  {% if grains['os'] == 'CentOS' %}
  - install.rpm.nginx
  {% elif grains['os'] == 'Ubuntu' %}

  {% endif %}
