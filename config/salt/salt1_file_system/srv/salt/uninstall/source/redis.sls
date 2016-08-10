rm -f /etc/redis.conf /etc/init.d/redis /tmp/redis_install.sh:
  cmd.run

rm -rf /usr/local/redis /var/lib/redis:
  cmd.run