#/bin/bash

workspace=`pwd`

if [ -f /etc/supervisor/supervisord.conf ]; then
    exit 0
fi

cd /tmp
wget https://bootstrap.pypa.io/ez_setup.py -O - | python
tar xzf "redis-${version}.tar.gz"
easy_install supervisor
mkdir -p /etc/supervisor/conf.d
echo_supervisord_conf > /etc/supervisor/supervisord.conf

cd "${workspace}"