include:
  {% if grains['os'] == 'CentOS' %}
  - install.rpm.php7
  {% elif grains['os'] == 'Ubuntu' %}

  {% endif %}
