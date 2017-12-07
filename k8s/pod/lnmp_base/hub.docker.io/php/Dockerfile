FROM php

RUN apt-get update \
    && apt-get install -y zlib1g-dev

RUN pecl install redis-3.1.4

RUN docker-php-ext-install pdo pdo_mysql

RUN docker-php-ext-configure mysqli --with-mysqli=mysqlnd && \
    docker-php-ext-install mysqli

RUN pecl install protobuf-3.4.0
RUN pecl install grpc-1.7.0

COPY php.ini /usr/local/etc/php/php.ini

COPY index.php /www/index.php

CMD php /www/index.php
