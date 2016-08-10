webtatic-nginx:
  pkgrepo.managed:
    - humanname: Webtatic Repository EL6
    - baseurl: https://repo.webtatic.com/yum/el6/$basearch/
    - gpgcheck: 1
    - gpgkey: file:///etc/pki/rpm-gpg/RPM-GPG-KEY-webtatic-andy-nginx


/etc/pki/rpm-gpg/RPM-GPG-KEY-webtatic-andy-nginx:
  file.managed:
    - source: salt://install/rpm/files/etc/pki/rpm-gpg/RPM-GPG-KEY-webtatic-andy
    - user: root
    - group: root
    - mode: 644


nginx.packages:
  pkg.installed:
    - pkgs:
      - nginx18

