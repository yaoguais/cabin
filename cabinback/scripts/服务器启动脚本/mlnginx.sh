#/bin/sh
# open php-fpm
# open nginx
# open mongodb

/home/yaoguai/soft/php54d/sbin/php-fpm &
/home/yaoguai/soft/nginx/sbin/nginx &
/home/yaoguai/soft/mongodb/bin/mongod --dbpath=/home/yaoguai/soft/mongodb/data/db/ --logpath=/home/yaoguai/soft/mongodb/data/log/mongo.log --logappend &
