webtatic-php7:
  pkgrepo.managed:
    - humanname: Webtatic Repository EL6
    - mirrorlist: https://mirror.webtatic.com/yum/el6/$basearch/mirrorlist
    - gpgcheck: 1
    - gpgkey: file:///etc/pki/rpm-gpg/RPM-GPG-KEY-webtatic-andy-php7

/etc/pki/rpm-gpg/RPM-GPG-KEY-webtatic-andy-php7:
  file.managed:
    - source: salt://install/rpm/files/etc/pki/rpm-gpg/RPM-GPG-KEY-webtatic-andy
    - user: root
    - group: root
    - mode: 644

php.packages:
  pkg.installed:
    - pkgs:
      - php70w
      - php70w-common
      - php70w-cli
      - php70w-bcmath
      - php70w-devel
      - php70w-fpm
      - php70w-gd
      - php70w-mbstring
      - php70w-mcrypt
      - php70w-mysqlnd
      - php70w-opcache
      - php70w-pdo
      - php70w-pgsql
      - php70w-process
      - php70w-xml
