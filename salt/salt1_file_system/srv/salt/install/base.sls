include:
  {% if grains['os'] == 'CentOS' %}
  - install.rpm.base
  {% elif grains['os'] == 'Ubuntu' %}

  {% endif %}
