include:
  {% if grains['os'] == 'CentOS' %}
  - install.rpm.rsync
  {% elif grains['os'] == 'Ubuntu' %}

  {% endif %}