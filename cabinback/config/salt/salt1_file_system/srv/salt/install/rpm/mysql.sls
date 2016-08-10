mysql-connectors-community:
  pkgrepo.managed:
    - humanname: MySQL Connectors Community
    - baseurl: http://repo.mysql.com/yum/mysql-connectors-community/el/6/$basearch/
    - gpgcheck: 1
    - gpgkey: file:///etc/pki/rpm-gpg/RPM-GPG-KEY-mysql

mysql-tools-community:
  pkgrepo.managed:
    - humanname: MySQL Tools Community
    - baseurl: http://repo.mysql.com/yum/mysql-tools-community/el/6/$basearch/
    - gpgcheck: 1
    - gpgkey: file:///etc/pki/rpm-gpg/RPM-GPG-KEY-mysql

mysql57-community:
  pkgrepo.managed:
    - humanname: MySQL 5.7 Community Server 
    - baseurl: http://repo.mysql.com/yum/mysql-5.7-community/el/6/$basearch/
    - gpgcheck: 1
    - gpgkey: file:///etc/pki/rpm-gpg/RPM-GPG-KEY-mysql

/etc/pki/rpm-gpg/RPM-GPG-KEY-mysql:
  file.managed:
    - source: salt://install/rpm/files/etc/pki/rpm-gpg/RPM-GPG-KEY-mysql
    - user: root
    - group: root
    - mode: 644

mysql.packages:
  pkg.installed:
    - pkgs:
      - mysql-community-server
      - mysql-community-client

