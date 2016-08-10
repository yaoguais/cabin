include:
  {% if grains['os'] == 'CentOS' %}
  - install.rpm.mysql
  {% elif grains['os'] == 'Ubuntu' %}

  {% endif %}
