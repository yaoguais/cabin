include:
  {% if grains['os'] == 'CentOS' %}
  - install.rpm.nodejs
  {% elif grains['os'] == 'Ubuntu' %}

  {% endif %}
