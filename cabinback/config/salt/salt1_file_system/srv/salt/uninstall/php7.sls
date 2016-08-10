include:
  {% if grains['os'] == 'CentOS' %}
  - uninstall.rpm.php7
  {% elif grains['os'] == 'Ubuntu' %}

  {% endif %}