#/bin/bash

version=3.2.3
install_path=/usr/local
workspace=`pwd`

if [ ! -d /var/lib/redis ]; then
	mkdir /var/lib/redis
fi

if [ -d "${install_path}/redis" ]; then
	exit 0
fi

cd /tmp
wget "http://download.redis.io/releases/redis-${version}.tar.gz"
tar xzf "redis-${version}.tar.gz"
cd "redis-${version}"
make PREFIX="${install_path}/redis" install
rm -rf "/tmp/redis-${version}"
rm -f "/tmp/redis-${version}.tar.gz"
cd "${workspace}"

