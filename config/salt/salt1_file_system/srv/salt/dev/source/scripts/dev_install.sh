#!/bin/bash

data_dir=/data
workspace=`pwd`
worker=$WORKER_USER

for i in server android ios im daemon
do
  mkdir "${data_dir}/${i}"
done

cd "${data_dir}/server"
git clone git@github.com:HelloWorldDev/app-server.git ./app-server
git clone git@github.com:HelloWorldDev/app-manage.git ./app-manage
git clone git@github.com:HelloWorldDev/common-website.git ./common-website
git clone git@github.com:HelloWorldDev/fameApp.git ./fame

for i in app-server app-manage common-website fame
do
    cd "${data_dir}/server/${i}"
    chmod -R 744 "${data_dir}/server/${i}/storage"
    chown -R $worker:$worker "${data_dir}/server/${i}/storage"
    composer install
    cd "${data_dir}/server/${i}/public"
    npm install
done

cd "${data_dir}/android"
git clone git@github.com:HelloWorldDev/app-android.git ./app-android
git clone git@github.com:HelloWorldDev/fame-android.git ./fame-android

cd "${data_dir}/ios"
git clone git@github.com:HelloWorldDev/app-ios.git ./app-ios
git clone git@github.com:HelloWorldDev/fame-ios.git ./fame-ios

cd "${data_dir}/im"
git clone git@github.com:HelloWorldDev/im-server.git ./im-server

cd "${data_dir}/daemon"
git clone git@github.com:HelloWorldDev/daemon-service.git ./daemon-service

exit 0