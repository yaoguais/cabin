include:
  {% if grains['os'] == 'CentOS' %}
  - install.rpm.postgresql
  {% elif grains['os'] == 'Ubuntu' %}

  {% endif %}
