# HOW TO INSTALL and USE

```
goto https://www.elastic.co/downloads/elasticsearch
wget "https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-5.0.1.rpm"
yum -y install elasticsearch-5.0.1.rpm
yum -y install java-1.8.0-openjdk 
# 1.8.0 is the min version for supporting es5.0
service elasticsearch restart
curl "http://127.0.0.1:9200/"

how to use php ?
composer require elasticsearch/elasticsearch
vim index.php
php index.php

```
