mongodb-org-3.2:
  pkgrepo.managed:
    - humanname: MongoDB Repository
    - baseurl: https://repo.mongodb.org/yum/amazon/2013.03/mongodb-org/3.2/x86_64/ 
    - gpgcheck: 1
    - gpgkey: https://www.mongodb.org/static/pgp/server-3.2.asc


mongodb.packages:
  pkg.installed:
    - pkgs:
      - mongodb-org

