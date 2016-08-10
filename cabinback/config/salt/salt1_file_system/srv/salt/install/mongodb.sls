include:
  {% if grains['os'] == 'CentOS' %}
  - install.rpm.mongodb
  {% elif grains['os'] == 'Ubuntu' %}

  {% endif %}
