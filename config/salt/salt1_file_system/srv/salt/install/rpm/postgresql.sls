pgdg95:
  pkgrepo.managed:
    - humanname: PostgreSQL 9.5 $releasever - $basearch
    - baseurl: https://download.postgresql.org/pub/repos/yum/9.5/redhat/rhel-$releasever-$basearch
    - gpgcheck: 1
    - gpgkey: file:///etc/pki/rpm-gpg/RPM-GPG-KEY-PGDG-95

/etc/pki/rpm-gpg/RPM-GPG-KEY-PGDG-95:
  file.managed:
    - source: salt://install/rpm/files/etc/pki/rpm-gpg/RPM-GPG-KEY-PGDG-95
    - user: root
    - group: root
    - mode: 644

postgresql.packages:
  pkg.installed:
    - pkgs:
      - postgresql95-server

if [ -f /etc/init.d/postgresql-9.5 ]; then mv /etc/init.d/postgresql-9.5 /etc/init.d/postgresql; fi:
  cmd.run


service postgresql initdb:
  cmd.run