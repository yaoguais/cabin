curl --silent --location https://rpm.nodesource.com/setup_4.x | bash -:
  cmd.run:
    - user: root

nodejs.packages:
  pkg.installed:
    - pkgs:
      - nodejs