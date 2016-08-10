#/bin/bash

version=4.6.3
install_path=/usr/local
workspace=`pwd`

if [ -d "${install_path}/phpmyadmin" ]; then
	exit 0
fi

cd /tmp
wget "https://files.phpmyadmin.net/phpMyAdmin/${version}/phpMyAdmin-${version}-all-languages.tar.gz"
tar xzf "phpMyAdmin-${version}-all-languages.tar.gz"
mv "phpMyAdmin-${version}-all-languages" "${install_path}/phpmyadmin"
rm -rf "/tmp/phpMyAdmin-${version}-all-languages"
rm -f "/tmp/phpMyAdmin-${version}-all-languages.tar.gz"
cd "${workspace}"

