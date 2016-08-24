rsync.packages:
  pkg.installed:
    - pkgs:
      - rsync

/etc/rsyncd:
  file.recurse:
    - source: salt://install/rpm/files/etc/rsyncd
    - user: root
    - group: root
    - file_mode: 660
    - dir_mode: 750

/etc/init.d/rsyncd:
  file.managed:
    - source: salt://install/rpm/files/etc/init.d/rsyncd
    - user: root
    - group: root
    - mode: 755