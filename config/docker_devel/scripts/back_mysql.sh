#!/bin/bash
# backup mysql datadir
# usage: sh scripts/backup_mysql.sh

date=`date '+%Y%m%d%H%M%S'`

mv files/mysql/var/lib/mysql "files/mysql/var/lib/mysql$(date '+%Y%m%d%H%M%S')"
mkdir files/mysql/var/lib/mysql

mv files/mysql-slave/var/lib/mysql-slave "files/mysql-slave/var/lib/mysql-slave$(date '+%Y%m%d%H%M%S')"
mkdir files/mysql-slave/var/lib/mysql-slave