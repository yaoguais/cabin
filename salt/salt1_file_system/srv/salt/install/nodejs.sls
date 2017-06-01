include:
  {% if grains['os'] == 'CentOS' %}
  - install.rpm.nodejs
  {% elif grains['os'] == 'Ubuntu' %}

  {% endif %}

npm install -g bower:
  cmd.run
