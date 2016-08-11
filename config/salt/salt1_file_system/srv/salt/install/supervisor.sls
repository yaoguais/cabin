include:
  - install.source.supervisor

{% if grains['os'] == 'CentOS' %}
/etc/init.d/supervisor:
  file.managed:
    - source: salt://install/rpm/files/etc/init.d/supervisor
    - user: root
    - group: root
    - mode: 755
{% elif grains['os'] == 'Ubuntu' %}
{% endif %}