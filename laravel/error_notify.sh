#!/bin/bash
# send ERROR log of laravel to slack channel
# author: Yaoguais <newtopstdio@163.com>
# usage: add this script to crontab

# record current workspace
pwd=`pwd`
# directory to save ERROR logs
lockdir="$HOME/.errorlock"
if [ ! -d $lockdir ]; then
	mkdir $lockdir
fi
# a part of log file name
date=`date +%Y-%m-%d`
# where is the log file saved
dirs="/data/time-server /data/liuyong/time-server /data/zhangji/time-server"
for dir in $dirs
do
	files="app-cli-$date.log app-web-$date.log"
	for file in $files
	do
		logFile=$dir/storage/logs/$file
		if [ ! -f $logFile ]; then
			continue
		fi
		lockFile=$lockdir/${dir//\//-}-$file
		if [ ! -f $lockFile ]; then
			touch $lockFile
		fi
		lc1="$(grep ERROR $logFile | grep -v NotFoundHttpException | wc -l)"
		lc2="$(grep "" $lockFile | wc -l)"
		if [ "$lc1" != "$lc2" ]; then
			grep ERROR $logFile | grep -v NotFoundHttpException > $lockFile
			lc1=`expr $lc1 + 0`
			lc2=`expr $lc2 + 1`
			for ((i=lc2; i<=lc1; i++))
			do
				msg=`sed -n "${i}p" $lockFile`
				msg="$(hostname):$logFile \n$msg"
				msg=${msg//\\/\\\\}
				msg=${msg////\/}
				echo $i
				payload=$(printf "payload={\"channel\": \"#dev-log-monitor\", \"username\": \"logmonitor\", \"text\": \"%s\", \"icon_emoji\": \":ghost:\"}" "$msg")
				curl -X POST --data-urlencode "$payload" https://hooks.slack.com/services/xxx/xxx
			done
		fi
	done
done
# remove cache files a day ago
find $lockdir -mtime +1 -name "*app*.log" | xargs rm -f
# back to workspace
cd $pwd
