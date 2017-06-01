```
.
├── README.md
├── algorithm
│   ├── BucketSort
│   │   ├── a.out
│   │   └── bucket_sort.c
│   ├── CountingSort
│   │   ├── a.out
│   │   └── counting_sort.c
│   ├── ExchangeSort
│   │   ├── bubble_sort.c
│   │   ├── mid
│   │   ├── mid.out
│   │   ├── mid_quick_sort.c
│   │   ├── normal
│   │   ├── normal.out
│   │   ├── normal_quick_sort.c
│   │   └── quick_sort.c
│   ├── InsertSort
│   │   ├── README.md
│   │   ├── a.out
│   │   └── insert_sort.c
│   ├── MergeSort
│   │   ├── a.out
│   │   └── merge_sort.c
│   ├── RadixSort
│   │   ├── a.out
│   │   └── radix_sort.c
│   ├── SelectSort
│   │   ├── README.md
│   │   ├── a.out
│   │   ├── heap_sort.c
│   │   └── select_sort.c
│   └── StringMatch
│       ├── a.out
│       ├── kmp_match.c
│       └── normal_match.c
├── apache
│   ├── Apache2.2.11\ (Win32)
│   │   └── httpd.conf
│   ├── Apache2.4\ (Win32)
│   │   └── httpd.conf
│   └── apache-2.4.10-9ubuntu1
│       ├── apache2.conf
│       ├── conf-available
│       │   ├── charset.conf
│       │   ├── localized-error-pages.conf
│       │   ├── other-vhosts-access-log.conf
│       │   ├── security.conf
│       │   └── serve-cgi-bin.conf
│       ├── conf-enabled
│       │   ├── charset.conf
│       │   ├── localized-error-pages.conf
│       │   ├── other-vhosts-access-log.conf
│       │   ├── security.conf
│       │   └── serve-cgi-bin.conf
│       ├── envvars
│       ├── magic
│       ├── mods-available
│       │   ├── access_compat.load
│       │   ├── actions.conf
│       │   ├── actions.load
│       │   ├── alias.conf
│       │   ├── alias.load
│       │   ├── allowmethods.load
│       │   ├── asis.load
│       │   ├── auth_basic.load
│       │   ├── auth_digest.load
│       │   ├── auth_form.load
│       │   ├── authn_anon.load
│       │   ├── authn_core.load
│       │   ├── authn_dbd.load
│       │   ├── authn_dbm.load
│       │   ├── authn_file.load
│       │   ├── authn_socache.load
│       │   ├── authnz_fcgi.load
│       │   ├── authnz_ldap.load
│       │   ├── authz_core.load
│       │   ├── authz_dbd.load
│       │   ├── authz_dbm.load
│       │   ├── authz_groupfile.load
│       │   ├── authz_host.load
│       │   ├── authz_owner.load
│       │   ├── authz_user.load
│       │   ├── autoindex.conf
│       │   ├── autoindex.load
│       │   ├── buffer.load
│       │   ├── cache.load
│       │   ├── cache_disk.conf
│       │   ├── cache_disk.load
│       │   ├── cache_socache.load
│       │   ├── cgi.load
│       │   ├── cgid.conf
│       │   ├── cgid.load
│       │   ├── charset_lite.load
│       │   ├── data.load
│       │   ├── dav.load
│       │   ├── dav_fs.conf
│       │   ├── dav_fs.load
│       │   ├── dav_lock.load
│       │   ├── dbd.load
│       │   ├── deflate.conf
│       │   ├── deflate.load
│       │   ├── dialup.load
│       │   ├── dir.conf
│       │   ├── dir.load
│       │   ├── dump_io.load
│       │   ├── echo.load
│       │   ├── env.load
│       │   ├── expires.load
│       │   ├── ext_filter.load
│       │   ├── file_cache.load
│       │   ├── filter.load
│       │   ├── headers.load
│       │   ├── heartbeat.load
│       │   ├── heartmonitor.load
│       │   ├── ident.load
│       │   ├── include.load
│       │   ├── info.conf
│       │   ├── info.load
│       │   ├── lbmethod_bybusyness.load
│       │   ├── lbmethod_byrequests.load
│       │   ├── lbmethod_bytraffic.load
│       │   ├── lbmethod_heartbeat.load
│       │   ├── ldap.conf
│       │   ├── ldap.load
│       │   ├── log_debug.load
│       │   ├── log_forensic.load
│       │   ├── lua.load
│       │   ├── macro.load
│       │   ├── mime.conf
│       │   ├── mime.load
│       │   ├── mime_magic.conf
│       │   ├── mime_magic.load
│       │   ├── mpm_event.conf
│       │   ├── mpm_event.load
│       │   ├── mpm_prefork.conf
│       │   ├── mpm_prefork.load
│       │   ├── mpm_worker.conf
│       │   ├── mpm_worker.load
│       │   ├── negotiation.conf
│       │   ├── negotiation.load
│       │   ├── php5.conf
│       │   ├── php5.load
│       │   ├── proxy.conf
│       │   ├── proxy.load
│       │   ├── proxy_ajp.load
│       │   ├── proxy_balancer.conf
│       │   ├── proxy_balancer.load
│       │   ├── proxy_connect.load
│       │   ├── proxy_express.load
│       │   ├── proxy_fcgi.load
│       │   ├── proxy_fdpass.load
│       │   ├── proxy_ftp.conf
│       │   ├── proxy_ftp.load
│       │   ├── proxy_html.load
│       │   ├── proxy_http.load
│       │   ├── proxy_scgi.load
│       │   ├── proxy_wstunnel.load
│       │   ├── ratelimit.load
│       │   ├── reflector.load
│       │   ├── remoteip.load
│       │   ├── reqtimeout.conf
│       │   ├── reqtimeout.load
│       │   ├── request.load
│       │   ├── rewrite.load
│       │   ├── sed.load
│       │   ├── session.load
│       │   ├── session_cookie.load
│       │   ├── session_crypto.load
│       │   ├── session_dbd.load
│       │   ├── setenvif.conf
│       │   ├── setenvif.load
│       │   ├── slotmem_plain.load
│       │   ├── slotmem_shm.load
│       │   ├── socache_dbm.load
│       │   ├── socache_memcache.load
│       │   ├── socache_shmcb.load
│       │   ├── speling.load
│       │   ├── ssl.conf
│       │   ├── ssl.load
│       │   ├── status.conf
│       │   ├── status.load
│       │   ├── substitute.load
│       │   ├── suexec.load
│       │   ├── unique_id.load
│       │   ├── userdir.conf
│       │   ├── userdir.load
│       │   ├── usertrack.load
│       │   ├── vhost_alias.load
│       │   └── xml2enc.load
│       ├── mods-enabled
│       │   ├── access_compat.load
│       │   ├── alias.conf
│       │   ├── alias.load
│       │   ├── auth_basic.load
│       │   ├── authn_core.load
│       │   ├── authn_file.load
│       │   ├── authz_core.load
│       │   ├── authz_host.load
│       │   ├── authz_user.load
│       │   ├── autoindex.conf
│       │   ├── autoindex.load
│       │   ├── deflate.conf
│       │   ├── deflate.load
│       │   ├── dir.conf
│       │   ├── dir.load
│       │   ├── env.load
│       │   ├── filter.load
│       │   ├── mime.conf
│       │   ├── mime.load
│       │   ├── mpm_prefork.conf
│       │   ├── mpm_prefork.load
│       │   ├── negotiation.conf
│       │   ├── negotiation.load
│       │   ├── php5.conf
│       │   ├── php5.load
│       │   ├── rewrite.load
│       │   ├── setenvif.conf
│       │   ├── setenvif.load
│       │   ├── status.conf
│       │   └── status.load
│       ├── ports.conf
│       ├── sites-available
│       │   ├── 000-default.conf
│       │   └── default-ssl.conf
│       └── sites-enabled
│           └── 000-default.conf
├── docker
│   ├── docker-compose.yml
│   ├── mysql-master
│   │   ├── Dockerfile
│   │   ├── etc
│   │   │   └── mysql
│   │   │       └── conf.d
│   │   │           └── my.cnf
│   │   └── files.sh
│   ├── mysql-slave
│   │   ├── Dockerfile
│   │   ├── etc
│   │   │   └── mysql
│   │   │       └── conf.d
│   │   │           └── my.cnf
│   │   └── files.sh
│   ├── nginx
│   │   ├── Dockerfile
│   │   ├── etc
│   │   │   └── nginx
│   │   │       ├── conf.d
│   │   │       │   ├── default.conf
│   │   │       │   └── fame-server.conf
│   │   │       ├── nginx.conf
│   │   │       ├── server.crt
│   │   │       └── server.key
│   │   └── var
│   │       └── log
│   │           └── nginx
│   │               ├── access.log
│   │               └── error.log
│   ├── php-fpm
│   │   ├── Dockerfile
│   │   ├── usr
│   │   │   └── local
│   │   │       └── etc
│   │   │           ├── php
│   │   │           │   ├── php.ini-development
│   │   │           │   └── php.ini-production
│   │   │           ├── php-fpm.conf
│   │   │           └── php-fpm.d
│   │   │               └── www.conf
│   │   └── var
│   │       └── log
│   │           └── php
│   │               ├── errors.log
│   │               └── php-fpm.log
│   ├── redis
│   │   ├── Dockerfile
│   │   └── etc
│   │       └── redis
│   │           └── redis.conf
│   ├── redis-slave
│   │   ├── Dockerfile
│   │   └── etc
│   │       └── redis
│   │           └── redis.conf
│   ├── start.sh
│   └── stop.sh
├── elasticsearch
│   ├── README.md
│   ├── composer.json
│   ├── composer.lock
│   └── index.php
├── gdb
│   └── gdbinit
├── goforadmin
│   ├── Makefile
│   ├── README.md
│   ├── app
│   │   ├── controllers
│   │   │   └── demo
│   │   │       └── demo.go
│   │   ├── routes
│   │   │   └── routes.go
│   │   └── views
│   │       └── demo
│   │           └── index.html
│   ├── config
│   │   └── config.toml.example
│   ├── core
│   │   ├── config.go
│   │   ├── controllers
│   │   │   ├── admin
│   │   │   │   └── admin.go
│   │   │   ├── component
│   │   │   │   └── component.go
│   │   │   └── index
│   │   │       ├── index.go
│   │   │       ├── login.go
│   │   │       └── logout.go
│   │   ├── core.go
│   │   ├── middlewares
│   │   │   ├── access.go
│   │   │   ├── admin.go
│   │   │   └── auth.go
│   │   ├── models
│   │   │   └── admin.go
│   │   ├── response.go
│   │   ├── routes
│   │   │   └── routes.go
│   │   ├── services
│   │   │   └── admin.go
│   │   └── views
│   │       ├── admin
│   │       │   ├── edit_privilege.html
│   │       │   ├── edit_role.html
│   │       │   ├── edit_user.html
│   │       │   ├── privileges.html
│   │       │   ├── roles.html
│   │       │   └── users.html
│   │       ├── component
│   │       │   └── index.html
│   │       ├── error
│   │       │   └── error.html
│   │       ├── index
│   │       │   └── index.html
│   │       ├── layout
│   │       │   └── layout.html
│   │       └── login
│   │           └── login.html
│   ├── install.sql
│   ├── lib
│   │   ├── db
│   │   │   └── db.go
│   │   ├── json
│   │   │   └── json.go
│   │   ├── log
│   │   │   └── log.go
│   │   ├── model
│   │   │   ├── datatable.go
│   │   │   └── datatable_test.go
│   │   ├── random
│   │   │   └── random.go
│   │   ├── redis
│   │   │   └── redis.go
│   │   ├── session
│   │   │   ├── redistore.go
│   │   │   ├── redistore_test.go
│   │   │   └── session.go
│   │   ├── slice
│   │   │   └── slice.go
│   │   ├── string
│   │   │   └── string.go
│   │   └── template
│   │       └── template.go
│   ├── main.go
│   ├── public
│   │   └── static
│   │       ├── blur-admin
│   │       │   ├── assets
│   │       │   │   ├── fonts
│   │       │   │   │   ├── socicon.eot
│   │       │   │   │   ├── socicon.svg
│   │       │   │   │   ├── socicon.ttf
│   │       │   │   │   ├── socicon.woff
│   │       │   │   │   └── socicon.woff2
│   │       │   │   ├── img
│   │       │   │   │   ├── app
│   │       │   │   │   │   ├── browsers
│   │       │   │   │   │   │   ├── chrome.svg
│   │       │   │   │   │   │   ├── firefox.svg
│   │       │   │   │   │   │   ├── ie.svg
│   │       │   │   │   │   │   ├── opera.svg
│   │       │   │   │   │   │   └── safari.svg
│   │       │   │   │   │   ├── feed
│   │       │   │   │   │   │   ├── genom.png
│   │       │   │   │   │   │   ├── my-little-kitten.png
│   │       │   │   │   │   │   ├── new-york-location.png
│   │       │   │   │   │   │   └── vader-and-me-preview.png
│   │       │   │   │   │   ├── my-app-logo.png
│   │       │   │   │   │   ├── profile
│   │       │   │   │   │   │   ├── Alexander.png
│   │       │   │   │   │   │   ├── Andrey.png
│   │       │   │   │   │   │   ├── Kostya.png
│   │       │   │   │   │   │   ├── Nasta.png
│   │       │   │   │   │   │   ├── Nick.png
│   │       │   │   │   │   │   └── Vlad.png
│   │       │   │   │   │   ├── skin-thumbnails
│   │       │   │   │   │   │   ├── 01_default.jpg
│   │       │   │   │   │   │   ├── 02_transparent.jpg
│   │       │   │   │   │   │   ├── 03_blue.jpg
│   │       │   │   │   │   │   ├── 04_peachy.jpg
│   │       │   │   │   │   │   ├── 05_material.jpg
│   │       │   │   │   │   │   └── 06_transblue.jpg
│   │       │   │   │   │   ├── todo
│   │       │   │   │   │   │   └── check-icon.png
│   │       │   │   │   │   └── typography
│   │       │   │   │   │       ├── banner.png
│   │       │   │   │   │       ├── typo01.png
│   │       │   │   │   │       ├── typo03.png
│   │       │   │   │   │       ├── typo04.png
│   │       │   │   │   │       ├── typo05.png
│   │       │   │   │   │       ├── typo06.png
│   │       │   │   │   │       └── typo07.png
│   │       │   │   │   ├── arrow-green-up.svg
│   │       │   │   │   ├── arrow-red-down.svg
│   │       │   │   │   ├── blue-bg.jpg
│   │       │   │   │   ├── blur-bg-blurred.jpg
│   │       │   │   │   ├── blur-bg-mobile.jpg
│   │       │   │   │   ├── blur-bg.jpg
│   │       │   │   │   ├── chernika.png
│   │       │   │   │   ├── comments.svg
│   │       │   │   │   ├── face.svg
│   │       │   │   │   ├── favicon-16x16.png
│   │       │   │   │   ├── favicon-32x32.png
│   │       │   │   │   ├── favicon-96x96.png
│   │       │   │   │   ├── green-bg.jpg
│   │       │   │   │   ├── money.svg
│   │       │   │   │   ├── peachy-bg.jpg
│   │       │   │   │   ├── person.svg
│   │       │   │   │   ├── refresh.svg
│   │       │   │   │   ├── shopping-cart.svg
│   │       │   │   │   ├── theme
│   │       │   │   │   │   ├── icon
│   │       │   │   │   │   │   ├── feed
│   │       │   │   │   │   │   │   ├── feed-image.svg
│   │       │   │   │   │   │   │   ├── feed-location.svg
│   │       │   │   │   │   │   │   └── feed-video.svg
│   │       │   │   │   │   │   └── kameleon
│   │       │   │   │   │   │       ├── Alien.svg
│   │       │   │   │   │   │       ├── Analytics.svg
│   │       │   │   │   │   │       ├── Apartment.svg
│   │       │   │   │   │   │       ├── Batman.svg
│   │       │   │   │   │   │       ├── Beach.svg
│   │       │   │   │   │   │       ├── Bell.svg
│   │       │   │   │   │   │       ├── Bonsai.svg
│   │       │   │   │   │   │       ├── Boss-3.svg
│   │       │   │   │   │   │       ├── Boss-5.svg
│   │       │   │   │   │   │       ├── Burglar.svg
│   │       │   │   │   │   │       ├── Bus.svg
│   │       │   │   │   │   │       ├── Candy.svg
│   │       │   │   │   │   │       ├── Checklist.svg
│   │       │   │   │   │   │       ├── Cheese.svg
│   │       │   │   │   │   │       ├── Chessboard.svg
│   │       │   │   │   │   │       ├── Clipboard-Plan.svg
│   │       │   │   │   │   │       ├── Desert.svg
│   │       │   │   │   │   │       ├── Dna.svg
│   │       │   │   │   │   │       ├── Euro-Coin.svg
│   │       │   │   │   │   │       ├── Food-Dome.svg
│   │       │   │   │   │   │       ├── Hacker.svg
│   │       │   │   │   │   │       ├── Images.svg
│   │       │   │   │   │   │       ├── Key.svg
│   │       │   │   │   │   │       ├── Laptop-Signal.svg
│   │       │   │   │   │   │       ├── Locked-Cloud.svg
│   │       │   │   │   │   │       ├── Love-Letter.svg
│   │       │   │   │   │   │       ├── Magician.svg
│   │       │   │   │   │   │       ├── Makeup.svg
│   │       │   │   │   │   │       ├── Medal-2.svg
│   │       │   │   │   │   │       ├── Microscope.svg
│   │       │   │   │   │   │       ├── Mind-Map-Paper.svg
│   │       │   │   │   │   │       ├── Money-Increase.svg
│   │       │   │   │   │   │       ├── Music-Equalizer.svg
│   │       │   │   │   │   │       ├── Ninja.svg
│   │       │   │   │   │   │       ├── Online-Shopping.svg
│   │       │   │   │   │   │       ├── Pantone.svg
│   │       │   │   │   │   │       ├── Party-Poppers.svg
│   │       │   │   │   │   │       ├── Phone-Booth.svg
│   │       │   │   │   │   │       ├── Programming.svg
│   │       │   │   │   │   │       ├── Santa.svg
│   │       │   │   │   │   │       ├── Shop.svg
│   │       │   │   │   │   │       ├── Street-View.svg
│   │       │   │   │   │   │       ├── Student-3.svg
│   │       │   │   │   │   │       ├── Surfer.svg
│   │       │   │   │   │   │       ├── Surgeon.svg
│   │       │   │   │   │   │       ├── Tower.svg
│   │       │   │   │   │   │       └── Vector.svg
│   │       │   │   │   │   ├── no-photo.png
│   │       │   │   │   │   ├── palette.png
│   │       │   │   │   │   └── vendor
│   │       │   │   │   │       ├── amcharts
│   │       │   │   │   │       │   └── dist
│   │       │   │   │   │       │       └── amcharts
│   │       │   │   │   │       │           └── images
│   │       │   │   │   │       │               ├── dragIcon.gif
│   │       │   │   │   │       │               ├── dragIconBlack.gif
│   │       │   │   │   │       │               ├── dragIconH.gif
│   │       │   │   │   │       │               ├── dragIconHBlack.gif
│   │       │   │   │   │       │               ├── dragIconRectBig.png
│   │       │   │   │   │       │               ├── dragIconRectBigBlack.png
│   │       │   │   │   │       │               ├── dragIconRectBigBlackH.png
│   │       │   │   │   │       │               ├── dragIconRectBigH.png
│   │       │   │   │   │       │               ├── dragIconRectSmall.png
│   │       │   │   │   │       │               ├── dragIconRectSmallBlack.png
│   │       │   │   │   │       │               ├── dragIconRectSmallBlackH.png
│   │       │   │   │   │       │               ├── dragIconRectSmallH.png
│   │       │   │   │   │       │               ├── dragIconRoundBig.png
│   │       │   │   │   │       │               ├── dragIconRoundBigBlack.png
│   │       │   │   │   │       │               ├── dragIconRoundBigBlackH.png
│   │       │   │   │   │       │               ├── dragIconRoundBigH.png
│   │       │   │   │   │       │               ├── dragIconRoundSmall.png
│   │       │   │   │   │       │               ├── dragIconRoundSmallBlack.png
│   │       │   │   │   │       │               ├── dragIconRoundSmallBlackH.png
│   │       │   │   │   │       │               ├── dragIconRoundSmallH.png
│   │       │   │   │   │       │               ├── export.png
│   │       │   │   │   │       │               ├── exportWhite.png
│   │       │   │   │   │       │               ├── lens.png
│   │       │   │   │   │       │               ├── lensWhite.png
│   │       │   │   │   │       │               ├── lensWhite_old.png
│   │       │   │   │   │       │               └── lens_old.png
│   │       │   │   │   │       ├── ammap
│   │       │   │   │   │       │   └── dist
│   │       │   │   │   │       │       └── ammap
│   │       │   │   │   │       │           └── images
│   │       │   │   │   │       │               ├── arrowDown.gif
│   │       │   │   │   │       │               ├── arrowUp.gif
│   │       │   │   │   │       │               ├── export.png
│   │       │   │   │   │       │               ├── homeIcon.gif
│   │       │   │   │   │       │               ├── homeIconWhite.gif
│   │       │   │   │   │       │               ├── minus.gif
│   │       │   │   │   │       │               ├── panDown.gif
│   │       │   │   │   │       │               ├── panLeft.gif
│   │       │   │   │   │       │               ├── panRight.gif
│   │       │   │   │   │       │               ├── panUp.gif
│   │       │   │   │   │       │               ├── plus.gif
│   │       │   │   │   │       │               ├── xIcon.gif
│   │       │   │   │   │       │               └── xIconH.gif
│   │       │   │   │   │       ├── ionrangeslider
│   │       │   │   │   │       │   └── img
│   │       │   │   │   │       │       ├── sprite-skin-flat.png
│   │       │   │   │   │       │       ├── sprite-skin-modern.png
│   │       │   │   │   │       │       ├── sprite-skin-nice.png
│   │       │   │   │   │       │       └── sprite-skin-simple.png
│   │       │   │   │   │       ├── jstree
│   │       │   │   │   │       │   └── dist
│   │       │   │   │   │       │       └── themes
│   │       │   │   │   │       │           ├── default
│   │       │   │   │   │       │           │   ├── 32px.png
│   │       │   │   │   │       │           │   ├── 40px.png
│   │       │   │   │   │       │           │   ├── style.css
│   │       │   │   │   │       │           │   ├── style.min.css
│   │       │   │   │   │       │           │   └── throbber.gif
│   │       │   │   │   │       │           └── default-dark
│   │       │   │   │   │       │               ├── 32px.png
│   │       │   │   │   │       │               ├── 40px.png
│   │       │   │   │   │       │               ├── style.css
│   │       │   │   │   │       │               ├── style.min.css
│   │       │   │   │   │       │               └── throbber.gif
│   │       │   │   │   │       └── leaflet
│   │       │   │   │   │           └── dist
│   │       │   │   │   │               └── images
│   │       │   │   │   │                   ├── layers-2x.png
│   │       │   │   │   │                   ├── layers.png
│   │       │   │   │   │                   ├── marker-icon-2x.png
│   │       │   │   │   │                   ├── marker-icon.png
│   │       │   │   │   │                   └── marker-shadow.png
│   │       │   │   │   └── transblue-bg.jpg
│   │       │   │   └── pictures
│   │       │   │       ├── pic-andrey.png
│   │       │   │       ├── pic-kostia.png
│   │       │   │       ├── pic-nasta.png
│   │       │   │       ├── pic-profile.png
│   │       │   │       ├── pic-vova.png
│   │       │   │       └── tinder-logo.jpg
│   │       │   ├── fonts
│   │       │   │   ├── fontawesome-webfont.eot
│   │       │   │   ├── fontawesome-webfont.svg
│   │       │   │   ├── fontawesome-webfont.ttf
│   │       │   │   ├── fontawesome-webfont.woff
│   │       │   │   ├── fontawesome-webfont.woff2
│   │       │   │   ├── glyphicons-halflings-regular.eot
│   │       │   │   ├── glyphicons-halflings-regular.svg
│   │       │   │   ├── glyphicons-halflings-regular.ttf
│   │       │   │   ├── glyphicons-halflings-regular.woff
│   │       │   │   ├── glyphicons-halflings-regular.woff2
│   │       │   │   ├── ionicons.eot
│   │       │   │   ├── ionicons.svg
│   │       │   │   ├── ionicons.ttf
│   │       │   │   └── ionicons.woff
│   │       │   ├── maps
│   │       │   │   ├── scripts
│   │       │   │   │   ├── app-a2f7d0e2cd.js.map
│   │       │   │   │   └── vendor-3f3216d3ec.js.map
│   │       │   │   └── styles
│   │       │   │       ├── 404-aff6a9433e.css.map
│   │       │   │       ├── app-341235afd6.css.map
│   │       │   │       ├── auth-a200a050c1.css.map
│   │       │   │       └── vendor-479457d43d.css.map
│   │       │   ├── scripts
│   │       │   │   ├── app.js
│   │       │   │   └── vendor.js
│   │       │   └── styles
│   │       │       ├── 404.css
│   │       │       ├── app.css
│   │       │       ├── auth.css
│   │       │       └── vendor.css
│   │       ├── bower.json
│   │       ├── bower_components
│   │       │   ├── bootstrap
│   │       │   │   ├── CHANGELOG.md
│   │       │   │   ├── Gemfile
│   │       │   │   ├── Gemfile.lock
│   │       │   │   ├── Gruntfile.js
│   │       │   │   ├── ISSUE_TEMPLATE.md
│   │       │   │   ├── LICENSE
│   │       │   │   ├── README.md
│   │       │   │   ├── bower.json
│   │       │   │   ├── dist
│   │       │   │   │   ├── css
│   │       │   │   │   │   ├── bootstrap-theme.css
│   │       │   │   │   │   ├── bootstrap-theme.css.map
│   │       │   │   │   │   ├── bootstrap-theme.min.css
│   │       │   │   │   │   ├── bootstrap-theme.min.css.map
│   │       │   │   │   │   ├── bootstrap.css
│   │       │   │   │   │   ├── bootstrap.css.map
│   │       │   │   │   │   ├── bootstrap.min.css
│   │       │   │   │   │   └── bootstrap.min.css.map
│   │       │   │   │   ├── fonts
│   │       │   │   │   │   ├── glyphicons-halflings-regular.eot
│   │       │   │   │   │   ├── glyphicons-halflings-regular.svg
│   │       │   │   │   │   ├── glyphicons-halflings-regular.ttf
│   │       │   │   │   │   ├── glyphicons-halflings-regular.woff
│   │       │   │   │   │   └── glyphicons-halflings-regular.woff2
│   │       │   │   │   └── js
│   │       │   │   │       ├── bootstrap.js
│   │       │   │   │       ├── bootstrap.min.js
│   │       │   │   │       └── npm.js
│   │       │   │   ├── fonts
│   │       │   │   │   ├── glyphicons-halflings-regular.eot
│   │       │   │   │   ├── glyphicons-halflings-regular.svg
│   │       │   │   │   ├── glyphicons-halflings-regular.ttf
│   │       │   │   │   ├── glyphicons-halflings-regular.woff
│   │       │   │   │   └── glyphicons-halflings-regular.woff2
│   │       │   │   ├── grunt
│   │       │   │   │   ├── bs-commonjs-generator.js
│   │       │   │   │   ├── bs-glyphicons-data-generator.js
│   │       │   │   │   ├── bs-lessdoc-parser.js
│   │       │   │   │   ├── bs-raw-files-generator.js
│   │       │   │   │   ├── change-version.js
│   │       │   │   │   ├── configBridge.json
│   │       │   │   │   ├── npm-shrinkwrap.json
│   │       │   │   │   └── sauce_browsers.yml
│   │       │   │   ├── js
│   │       │   │   │   ├── affix.js
│   │       │   │   │   ├── alert.js
│   │       │   │   │   ├── button.js
│   │       │   │   │   ├── carousel.js
│   │       │   │   │   ├── collapse.js
│   │       │   │   │   ├── dropdown.js
│   │       │   │   │   ├── modal.js
│   │       │   │   │   ├── popover.js
│   │       │   │   │   ├── scrollspy.js
│   │       │   │   │   ├── tab.js
│   │       │   │   │   ├── tooltip.js
│   │       │   │   │   └── transition.js
│   │       │   │   ├── less
│   │       │   │   │   ├── alerts.less
│   │       │   │   │   ├── badges.less
│   │       │   │   │   ├── bootstrap.less
│   │       │   │   │   ├── breadcrumbs.less
│   │       │   │   │   ├── button-groups.less
│   │       │   │   │   ├── buttons.less
│   │       │   │   │   ├── carousel.less
│   │       │   │   │   ├── close.less
│   │       │   │   │   ├── code.less
│   │       │   │   │   ├── component-animations.less
│   │       │   │   │   ├── dropdowns.less
│   │       │   │   │   ├── forms.less
│   │       │   │   │   ├── glyphicons.less
│   │       │   │   │   ├── grid.less
│   │       │   │   │   ├── input-groups.less
│   │       │   │   │   ├── jumbotron.less
│   │       │   │   │   ├── labels.less
│   │       │   │   │   ├── list-group.less
│   │       │   │   │   ├── media.less
│   │       │   │   │   ├── mixins
│   │       │   │   │   │   ├── alerts.less
│   │       │   │   │   │   ├── background-variant.less
│   │       │   │   │   │   ├── border-radius.less
│   │       │   │   │   │   ├── buttons.less
│   │       │   │   │   │   ├── center-block.less
│   │       │   │   │   │   ├── clearfix.less
│   │       │   │   │   │   ├── forms.less
│   │       │   │   │   │   ├── gradients.less
│   │       │   │   │   │   ├── grid-framework.less
│   │       │   │   │   │   ├── grid.less
│   │       │   │   │   │   ├── hide-text.less
│   │       │   │   │   │   ├── image.less
│   │       │   │   │   │   ├── labels.less
│   │       │   │   │   │   ├── list-group.less
│   │       │   │   │   │   ├── nav-divider.less
│   │       │   │   │   │   ├── nav-vertical-align.less
│   │       │   │   │   │   ├── opacity.less
│   │       │   │   │   │   ├── pagination.less
│   │       │   │   │   │   ├── panels.less
│   │       │   │   │   │   ├── progress-bar.less
│   │       │   │   │   │   ├── reset-filter.less
│   │       │   │   │   │   ├── reset-text.less
│   │       │   │   │   │   ├── resize.less
│   │       │   │   │   │   ├── responsive-visibility.less
│   │       │   │   │   │   ├── size.less
│   │       │   │   │   │   ├── tab-focus.less
│   │       │   │   │   │   ├── table-row.less
│   │       │   │   │   │   ├── text-emphasis.less
│   │       │   │   │   │   ├── text-overflow.less
│   │       │   │   │   │   └── vendor-prefixes.less
│   │       │   │   │   ├── mixins.less
│   │       │   │   │   ├── modals.less
│   │       │   │   │   ├── navbar.less
│   │       │   │   │   ├── navs.less
│   │       │   │   │   ├── normalize.less
│   │       │   │   │   ├── pager.less
│   │       │   │   │   ├── pagination.less
│   │       │   │   │   ├── panels.less
│   │       │   │   │   ├── popovers.less
│   │       │   │   │   ├── print.less
│   │       │   │   │   ├── progress-bars.less
│   │       │   │   │   ├── responsive-embed.less
│   │       │   │   │   ├── responsive-utilities.less
│   │       │   │   │   ├── scaffolding.less
│   │       │   │   │   ├── tables.less
│   │       │   │   │   ├── theme.less
│   │       │   │   │   ├── thumbnails.less
│   │       │   │   │   ├── tooltip.less
│   │       │   │   │   ├── type.less
│   │       │   │   │   ├── utilities.less
│   │       │   │   │   ├── variables.less
│   │       │   │   │   └── wells.less
│   │       │   │   ├── nuget
│   │       │   │   │   ├── MyGet.ps1
│   │       │   │   │   ├── bootstrap.less.nuspec
│   │       │   │   │   └── bootstrap.nuspec
│   │       │   │   ├── package.js
│   │       │   │   └── package.json
│   │       │   ├── bootstrap-datepicker
│   │       │   │   ├── CHANGELOG.md
│   │       │   │   ├── CONTRIBUTING.md
│   │       │   │   ├── Gruntfile.js
│   │       │   │   ├── LICENSE
│   │       │   │   ├── README.md
│   │       │   │   ├── bower.json
│   │       │   │   ├── composer.json
│   │       │   │   ├── dist
│   │       │   │   │   ├── css
│   │       │   │   │   │   ├── bootstrap-datepicker.css
│   │       │   │   │   │   ├── bootstrap-datepicker.css.map
│   │       │   │   │   │   ├── bootstrap-datepicker.min.css
│   │       │   │   │   │   ├── bootstrap-datepicker.min.css.map
│   │       │   │   │   │   ├── bootstrap-datepicker.standalone.css
│   │       │   │   │   │   ├── bootstrap-datepicker.standalone.css.map
│   │       │   │   │   │   ├── bootstrap-datepicker.standalone.min.css
│   │       │   │   │   │   ├── bootstrap-datepicker.standalone.min.css.map
│   │       │   │   │   │   ├── bootstrap-datepicker3.css
│   │       │   │   │   │   ├── bootstrap-datepicker3.css.map
│   │       │   │   │   │   ├── bootstrap-datepicker3.min.css
│   │       │   │   │   │   ├── bootstrap-datepicker3.min.css.map
│   │       │   │   │   │   ├── bootstrap-datepicker3.standalone.css
│   │       │   │   │   │   ├── bootstrap-datepicker3.standalone.css.map
│   │       │   │   │   │   ├── bootstrap-datepicker3.standalone.min.css
│   │       │   │   │   │   └── bootstrap-datepicker3.standalone.min.css.map
│   │       │   │   │   ├── js
│   │       │   │   │   │   ├── bootstrap-datepicker.js
│   │       │   │   │   │   └── bootstrap-datepicker.min.js
│   │       │   │   │   └── locales
│   │       │   │   │       ├── bootstrap-datepicker.ar.min.js
│   │       │   │   │       ├── bootstrap-datepicker.az.min.js
│   │       │   │   │       ├── bootstrap-datepicker.bg.min.js
│   │       │   │   │       ├── bootstrap-datepicker.bs.min.js
│   │       │   │   │       ├── bootstrap-datepicker.ca.min.js
│   │       │   │   │       ├── bootstrap-datepicker.cs.min.js
│   │       │   │   │       ├── bootstrap-datepicker.cy.min.js
│   │       │   │   │       ├── bootstrap-datepicker.da.min.js
│   │       │   │   │       ├── bootstrap-datepicker.de.min.js
│   │       │   │   │       ├── bootstrap-datepicker.el.min.js
│   │       │   │   │       ├── bootstrap-datepicker.en-AU.min.js
│   │       │   │   │       ├── bootstrap-datepicker.en-GB.min.js
│   │       │   │   │       ├── bootstrap-datepicker.eo.min.js
│   │       │   │   │       ├── bootstrap-datepicker.es.min.js
│   │       │   │   │       ├── bootstrap-datepicker.et.min.js
│   │       │   │   │       ├── bootstrap-datepicker.eu.min.js
│   │       │   │   │       ├── bootstrap-datepicker.fa.min.js
│   │       │   │   │       ├── bootstrap-datepicker.fi.min.js
│   │       │   │   │       ├── bootstrap-datepicker.fo.min.js
│   │       │   │   │       ├── bootstrap-datepicker.fr-CH.min.js
│   │       │   │   │       ├── bootstrap-datepicker.fr.min.js
│   │       │   │   │       ├── bootstrap-datepicker.gl.min.js
│   │       │   │   │       ├── bootstrap-datepicker.he.min.js
│   │       │   │   │       ├── bootstrap-datepicker.hr.min.js
│   │       │   │   │       ├── bootstrap-datepicker.hu.min.js
│   │       │   │   │       ├── bootstrap-datepicker.hy.min.js
│   │       │   │   │       ├── bootstrap-datepicker.id.min.js
│   │       │   │   │       ├── bootstrap-datepicker.is.min.js
│   │       │   │   │       ├── bootstrap-datepicker.it-CH.min.js
│   │       │   │   │       ├── bootstrap-datepicker.it.min.js
│   │       │   │   │       ├── bootstrap-datepicker.ja.min.js
│   │       │   │   │       ├── bootstrap-datepicker.ka.min.js
│   │       │   │   │       ├── bootstrap-datepicker.kh.min.js
│   │       │   │   │       ├── bootstrap-datepicker.kk.min.js
│   │       │   │   │       ├── bootstrap-datepicker.ko.min.js
│   │       │   │   │       ├── bootstrap-datepicker.kr.min.js
│   │       │   │   │       ├── bootstrap-datepicker.lt.min.js
│   │       │   │   │       ├── bootstrap-datepicker.lv.min.js
│   │       │   │   │       ├── bootstrap-datepicker.me.min.js
│   │       │   │   │       ├── bootstrap-datepicker.mk.min.js
│   │       │   │   │       ├── bootstrap-datepicker.mn.min.js
│   │       │   │   │       ├── bootstrap-datepicker.ms.min.js
│   │       │   │   │       ├── bootstrap-datepicker.nb.min.js
│   │       │   │   │       ├── bootstrap-datepicker.nl-BE.min.js
│   │       │   │   │       ├── bootstrap-datepicker.nl.min.js
│   │       │   │   │       ├── bootstrap-datepicker.no.min.js
│   │       │   │   │       ├── bootstrap-datepicker.pl.min.js
│   │       │   │   │       ├── bootstrap-datepicker.pt-BR.min.js
│   │       │   │   │       ├── bootstrap-datepicker.pt.min.js
│   │       │   │   │       ├── bootstrap-datepicker.ro.min.js
│   │       │   │   │       ├── bootstrap-datepicker.rs-latin.min.js
│   │       │   │   │       ├── bootstrap-datepicker.rs.min.js
│   │       │   │   │       ├── bootstrap-datepicker.ru.min.js
│   │       │   │   │       ├── bootstrap-datepicker.sk.min.js
│   │       │   │   │       ├── bootstrap-datepicker.sl.min.js
│   │       │   │   │       ├── bootstrap-datepicker.sq.min.js
│   │       │   │   │       ├── bootstrap-datepicker.sr-latin.min.js
│   │       │   │   │       ├── bootstrap-datepicker.sr.min.js
│   │       │   │   │       ├── bootstrap-datepicker.sv.min.js
│   │       │   │   │       ├── bootstrap-datepicker.sw.min.js
│   │       │   │   │       ├── bootstrap-datepicker.th.min.js
│   │       │   │   │       ├── bootstrap-datepicker.tr.min.js
│   │       │   │   │       ├── bootstrap-datepicker.uk.min.js
│   │       │   │   │       ├── bootstrap-datepicker.vi.min.js
│   │       │   │   │       ├── bootstrap-datepicker.zh-CN.min.js
│   │       │   │   │       └── bootstrap-datepicker.zh-TW.min.js
│   │       │   │   ├── docs
│   │       │   │   │   ├── Makefile
│   │       │   │   │   ├── README.md
│   │       │   │   │   ├── _screenshots
│   │       │   │   │   │   ├── demo_head.html
│   │       │   │   │   │   ├── markup_component.html
│   │       │   │   │   │   ├── markup_daterange.html
│   │       │   │   │   │   ├── markup_inline.html
│   │       │   │   │   │   ├── markup_input.html
│   │       │   │   │   │   ├── option_calendarweeks.html
│   │       │   │   │   │   ├── option_clearbtn.html
│   │       │   │   │   │   ├── option_daysofweekdisabled.html
│   │       │   │   │   │   ├── option_enddate.html
│   │       │   │   │   │   ├── option_language.html
│   │       │   │   │   │   ├── option_multidate.html
│   │       │   │   │   │   ├── option_startdate.html
│   │       │   │   │   │   ├── option_todaybtn.html
│   │       │   │   │   │   ├── option_todayhighlight.html
│   │       │   │   │   │   ├── option_weekstart.html
│   │       │   │   │   │   └── script
│   │       │   │   │   │       ├── common.css
│   │       │   │   │   │       ├── common.js
│   │       │   │   │   │       ├── debug.js
│   │       │   │   │   │       ├── html-imports.min.js
│   │       │   │   │   │       └── screenshot.js
│   │       │   │   │   ├── _static
│   │       │   │   │   │   └── screenshots
│   │       │   │   │   │       ├── demo_head.png
│   │       │   │   │   │       ├── markup_component.png
│   │       │   │   │   │       ├── markup_daterange.png
│   │       │   │   │   │       ├── markup_inline.png
│   │       │   │   │   │       ├── markup_input.png
│   │       │   │   │   │       ├── option_calendarweeks.png
│   │       │   │   │   │       ├── option_clearbtn.png
│   │       │   │   │   │       ├── option_daysofweekdisabled.png
│   │       │   │   │   │       ├── option_enddate.png
│   │       │   │   │   │       ├── option_language.png
│   │       │   │   │   │       ├── option_multidate.png
│   │       │   │   │   │       ├── option_startdate.png
│   │       │   │   │   │       ├── option_todaybtn.png
│   │       │   │   │   │       ├── option_todayhighlight.png
│   │       │   │   │   │       └── option_weekstart.png
│   │       │   │   │   ├── conf.py
│   │       │   │   │   ├── events.rst
│   │       │   │   │   ├── i18n.rst
│   │       │   │   │   ├── index.rst
│   │       │   │   │   ├── keyboard.rst
│   │       │   │   │   ├── make.bat
│   │       │   │   │   ├── markup.rst
│   │       │   │   │   ├── methods.rst
│   │       │   │   │   ├── options.rst
│   │       │   │   │   └── requirements.txt
│   │       │   │   ├── grunt
│   │       │   │   ├── js
│   │       │   │   │   ├── bootstrap-datepicker.js
│   │       │   │   │   └── locales
│   │       │   │   │       ├── bootstrap-datepicker.ar.js
│   │       │   │   │       ├── bootstrap-datepicker.az.js
│   │       │   │   │       ├── bootstrap-datepicker.bg.js
│   │       │   │   │       ├── bootstrap-datepicker.bs.js
│   │       │   │   │       ├── bootstrap-datepicker.ca.js
│   │       │   │   │       ├── bootstrap-datepicker.cs.js
│   │       │   │   │       ├── bootstrap-datepicker.cy.js
│   │       │   │   │       ├── bootstrap-datepicker.da.js
│   │       │   │   │       ├── bootstrap-datepicker.de.js
│   │       │   │   │       ├── bootstrap-datepicker.el.js
│   │       │   │   │       ├── bootstrap-datepicker.en-AU.js
│   │       │   │   │       ├── bootstrap-datepicker.en-GB.js
│   │       │   │   │       ├── bootstrap-datepicker.eo.js
│   │       │   │   │       ├── bootstrap-datepicker.es.js
│   │       │   │   │       ├── bootstrap-datepicker.et.js
│   │       │   │   │       ├── bootstrap-datepicker.eu.js
│   │       │   │   │       ├── bootstrap-datepicker.fa.js
│   │       │   │   │       ├── bootstrap-datepicker.fi.js
│   │       │   │   │       ├── bootstrap-datepicker.fo.js
│   │       │   │   │       ├── bootstrap-datepicker.fr-CH.js
│   │       │   │   │       ├── bootstrap-datepicker.fr.js
│   │       │   │   │       ├── bootstrap-datepicker.gl.js
│   │       │   │   │       ├── bootstrap-datepicker.he.js
│   │       │   │   │       ├── bootstrap-datepicker.hr.js
│   │       │   │   │       ├── bootstrap-datepicker.hu.js
│   │       │   │   │       ├── bootstrap-datepicker.hy.js
│   │       │   │   │       ├── bootstrap-datepicker.id.js
│   │       │   │   │       ├── bootstrap-datepicker.is.js
│   │       │   │   │       ├── bootstrap-datepicker.it-CH.js
│   │       │   │   │       ├── bootstrap-datepicker.it.js
│   │       │   │   │       ├── bootstrap-datepicker.ja.js
│   │       │   │   │       ├── bootstrap-datepicker.ka.js
│   │       │   │   │       ├── bootstrap-datepicker.kh.js
│   │       │   │   │       ├── bootstrap-datepicker.kk.js
│   │       │   │   │       ├── bootstrap-datepicker.ko.js
│   │       │   │   │       ├── bootstrap-datepicker.kr.js
│   │       │   │   │       ├── bootstrap-datepicker.lt.js
│   │       │   │   │       ├── bootstrap-datepicker.lv.js
│   │       │   │   │       ├── bootstrap-datepicker.me.js
│   │       │   │   │       ├── bootstrap-datepicker.mk.js
│   │       │   │   │       ├── bootstrap-datepicker.mn.js
│   │       │   │   │       ├── bootstrap-datepicker.ms.js
│   │       │   │   │       ├── bootstrap-datepicker.nb.js
│   │       │   │   │       ├── bootstrap-datepicker.nl-BE.js
│   │       │   │   │       ├── bootstrap-datepicker.nl.js
│   │       │   │   │       ├── bootstrap-datepicker.no.js
│   │       │   │   │       ├── bootstrap-datepicker.pl.js
│   │       │   │   │       ├── bootstrap-datepicker.pt-BR.js
│   │       │   │   │       ├── bootstrap-datepicker.pt.js
│   │       │   │   │       ├── bootstrap-datepicker.ro.js
│   │       │   │   │       ├── bootstrap-datepicker.rs-latin.js
│   │       │   │   │       ├── bootstrap-datepicker.rs.js
│   │       │   │   │       ├── bootstrap-datepicker.ru.js
│   │       │   │   │       ├── bootstrap-datepicker.sk.js
│   │       │   │   │       ├── bootstrap-datepicker.sl.js
│   │       │   │   │       ├── bootstrap-datepicker.sq.js
│   │       │   │   │       ├── bootstrap-datepicker.sr-latin.js
│   │       │   │   │       ├── bootstrap-datepicker.sr.js
│   │       │   │   │       ├── bootstrap-datepicker.sv.js
│   │       │   │   │       ├── bootstrap-datepicker.sw.js
│   │       │   │   │       ├── bootstrap-datepicker.th.js
│   │       │   │   │       ├── bootstrap-datepicker.tr.js
│   │       │   │   │       ├── bootstrap-datepicker.uk.js
│   │       │   │   │       ├── bootstrap-datepicker.vi.js
│   │       │   │   │       ├── bootstrap-datepicker.zh-CN.js
│   │       │   │   │       └── bootstrap-datepicker.zh-TW.js
│   │       │   │   ├── less
│   │       │   │   │   ├── datepicker.less
│   │       │   │   │   └── datepicker3.less
│   │       │   │   ├── package.json
│   │       │   │   └── tests
│   │       │   │       ├── README.md
│   │       │   │       ├── assets
│   │       │   │       │   ├── coverage.js
│   │       │   │       │   ├── jquery-1.7.1.min.js
│   │       │   │       │   ├── mock.js
│   │       │   │       │   ├── qunit-logging.js
│   │       │   │       │   ├── qunit.css
│   │       │   │       │   ├── qunit.js
│   │       │   │       │   └── utils.js
│   │       │   │       ├── suites
│   │       │   │       │   ├── calendar-weeks.js
│   │       │   │       │   ├── component.js
│   │       │   │       │   ├── data-api.js
│   │       │   │       │   ├── events.js
│   │       │   │       │   ├── formats.js
│   │       │   │       │   ├── inline.js
│   │       │   │       │   ├── keyboard_navigation
│   │       │   │       │   │   ├── 2011.js
│   │       │   │       │   │   ├── 2012.js
│   │       │   │       │   │   └── all.js
│   │       │   │       │   ├── methods.js
│   │       │   │       │   ├── methods_jquery.js
│   │       │   │       │   ├── mouse_navigation
│   │       │   │       │   │   ├── 2011.js
│   │       │   │       │   │   ├── 2012.js
│   │       │   │       │   │   └── all.js
│   │       │   │       │   ├── noconflict.js
│   │       │   │       │   ├── options.js
│   │       │   │       │   └── timezone.js
│   │       │   │       ├── tests.html
│   │       │   │       └── timezone.html
│   │       │   ├── bootstrap-daterangepicker
│   │       │   │   ├── README.md
│   │       │   │   ├── bower.json
│   │       │   │   ├── daterangepicker.css
│   │       │   │   ├── daterangepicker.js
│   │       │   │   ├── daterangepicker.scss
│   │       │   │   ├── demo.html
│   │       │   │   ├── drp.png
│   │       │   │   ├── example
│   │       │   │   │   ├── amd
│   │       │   │   │   │   ├── index.html
│   │       │   │   │   │   ├── main.js
│   │       │   │   │   │   └── require.js
│   │       │   │   │   └── browserify
│   │       │   │   │       ├── README.md
│   │       │   │   │       ├── bundle.js
│   │       │   │   │       ├── index.html
│   │       │   │   │       └── main.js
│   │       │   │   ├── package.js
│   │       │   │   ├── package.json
│   │       │   │   └── website
│   │       │   │       ├── index.html
│   │       │   │       ├── website.css
│   │       │   │       └── website.js
│   │       │   ├── bootstrap-select
│   │       │   │   ├── CHANGELOG.md
│   │       │   │   ├── LICENSE
│   │       │   │   ├── bower.json
│   │       │   │   ├── dist
│   │       │   │   │   ├── css
│   │       │   │   │   │   ├── bootstrap-select.css
│   │       │   │   │   │   ├── bootstrap-select.css.map
│   │       │   │   │   │   └── bootstrap-select.min.css
│   │       │   │   │   └── js
│   │       │   │   │       ├── bootstrap-select.js
│   │       │   │   │       ├── bootstrap-select.js.map
│   │       │   │   │       ├── bootstrap-select.min.js
│   │       │   │   │       └── i18n
│   │       │   │   │           ├── defaults-ar_AR.js
│   │       │   │   │           ├── defaults-ar_AR.min.js
│   │       │   │   │           ├── defaults-bg_BG.js
│   │       │   │   │           ├── defaults-bg_BG.min.js
│   │       │   │   │           ├── defaults-cro_CRO.js
│   │       │   │   │           ├── defaults-cro_CRO.min.js
│   │       │   │   │           ├── defaults-cs_CZ.js
│   │       │   │   │           ├── defaults-cs_CZ.min.js
│   │       │   │   │           ├── defaults-da_DK.js
│   │       │   │   │           ├── defaults-da_DK.min.js
│   │       │   │   │           ├── defaults-de_DE.js
│   │       │   │   │           ├── defaults-de_DE.min.js
│   │       │   │   │           ├── defaults-en_US.js
│   │       │   │   │           ├── defaults-en_US.min.js
│   │       │   │   │           ├── defaults-es_CL.js
│   │       │   │   │           ├── defaults-es_CL.min.js
│   │       │   │   │           ├── defaults-es_ES.js
│   │       │   │   │           ├── defaults-es_ES.min.js
│   │       │   │   │           ├── defaults-eu.js
│   │       │   │   │           ├── defaults-eu.min.js
│   │       │   │   │           ├── defaults-fa_IR.js
│   │       │   │   │           ├── defaults-fa_IR.min.js
│   │       │   │   │           ├── defaults-fi_FI.js
│   │       │   │   │           ├── defaults-fi_FI.min.js
│   │       │   │   │           ├── defaults-fr_FR.js
│   │       │   │   │           ├── defaults-fr_FR.min.js
│   │       │   │   │           ├── defaults-hu_HU.js
│   │       │   │   │           ├── defaults-hu_HU.min.js
│   │       │   │   │           ├── defaults-id_ID.js
│   │       │   │   │           ├── defaults-id_ID.min.js
│   │       │   │   │           ├── defaults-it_IT.js
│   │       │   │   │           ├── defaults-it_IT.min.js
│   │       │   │   │           ├── defaults-ko_KR.js
│   │       │   │   │           ├── defaults-ko_KR.min.js
│   │       │   │   │           ├── defaults-lt_LT.js
│   │       │   │   │           ├── defaults-lt_LT.min.js
│   │       │   │   │           ├── defaults-nb_NO.js
│   │       │   │   │           ├── defaults-nb_NO.min.js
│   │       │   │   │           ├── defaults-nl_NL.js
│   │       │   │   │           ├── defaults-nl_NL.min.js
│   │       │   │   │           ├── defaults-pl_PL.js
│   │       │   │   │           ├── defaults-pl_PL.min.js
│   │       │   │   │           ├── defaults-pt_BR.js
│   │       │   │   │           ├── defaults-pt_BR.min.js
│   │       │   │   │           ├── defaults-pt_PT.js
│   │       │   │   │           ├── defaults-pt_PT.min.js
│   │       │   │   │           ├── defaults-ro_RO.js
│   │       │   │   │           ├── defaults-ro_RO.min.js
│   │       │   │   │           ├── defaults-ru_RU.js
│   │       │   │   │           ├── defaults-ru_RU.min.js
│   │       │   │   │           ├── defaults-sk_SK.js
│   │       │   │   │           ├── defaults-sk_SK.min.js
│   │       │   │   │           ├── defaults-sl_SI.js
│   │       │   │   │           ├── defaults-sl_SI.min.js
│   │       │   │   │           ├── defaults-sv_SE.js
│   │       │   │   │           ├── defaults-sv_SE.min.js
│   │       │   │   │           ├── defaults-tr_TR.js
│   │       │   │   │           ├── defaults-tr_TR.min.js
│   │       │   │   │           ├── defaults-ua_UA.js
│   │       │   │   │           ├── defaults-ua_UA.min.js
│   │       │   │   │           ├── defaults-zh_CN.js
│   │       │   │   │           ├── defaults-zh_CN.min.js
│   │       │   │   │           ├── defaults-zh_TW.js
│   │       │   │   │           └── defaults-zh_TW.min.js
│   │       │   │   ├── docs
│   │       │   │   │   ├── custom_theme
│   │       │   │   │   │   ├── base.html
│   │       │   │   │   │   ├── css
│   │       │   │   │   │   │   └── base.css
│   │       │   │   │   │   ├── img
│   │       │   │   │   │   │   └── logos
│   │       │   │   │   │   │       ├── convertizer.png
│   │       │   │   │   │   │       ├── estimateit.png
│   │       │   │   │   │   │       ├── membermeister.png
│   │       │   │   │   │   │       ├── snapappointments.png
│   │       │   │   │   │   │       ├── solveforall.png
│   │       │   │   │   │   │       └── thermofisher.png
│   │       │   │   │   │   ├── js
│   │       │   │   │   │   │   └── base.js
│   │       │   │   │   │   ├── nav.html
│   │       │   │   │   │   └── toc.html
│   │       │   │   │   ├── docs
│   │       │   │   │   │   ├── css
│   │       │   │   │   │   │   └── custom.css
│   │       │   │   │   │   ├── dist
│   │       │   │   │   │   │   ├── css
│   │       │   │   │   │   │   │   ├── bootstrap-select.css
│   │       │   │   │   │   │   │   ├── bootstrap-select.css.map
│   │       │   │   │   │   │   │   └── bootstrap-select.min.css
│   │       │   │   │   │   │   └── js
│   │       │   │   │   │   │       ├── bootstrap-select.js
│   │       │   │   │   │   │       ├── bootstrap-select.js.map
│   │       │   │   │   │   │       ├── bootstrap-select.min.js
│   │       │   │   │   │   │       └── i18n
│   │       │   │   │   │   │           ├── defaults-ar_AR.js
│   │       │   │   │   │   │           ├── defaults-ar_AR.min.js
│   │       │   │   │   │   │           ├── defaults-bg_BG.js
│   │       │   │   │   │   │           ├── defaults-bg_BG.min.js
│   │       │   │   │   │   │           ├── defaults-cro_CRO.js
│   │       │   │   │   │   │           ├── defaults-cro_CRO.min.js
│   │       │   │   │   │   │           ├── defaults-cs_CZ.js
│   │       │   │   │   │   │           ├── defaults-cs_CZ.min.js
│   │       │   │   │   │   │           ├── defaults-da_DK.js
│   │       │   │   │   │   │           ├── defaults-da_DK.min.js
│   │       │   │   │   │   │           ├── defaults-de_DE.js
│   │       │   │   │   │   │           ├── defaults-de_DE.min.js
│   │       │   │   │   │   │           ├── defaults-en_US.js
│   │       │   │   │   │   │           ├── defaults-en_US.min.js
│   │       │   │   │   │   │           ├── defaults-es_CL.js
│   │       │   │   │   │   │           ├── defaults-es_CL.min.js
│   │       │   │   │   │   │           ├── defaults-es_ES.js
│   │       │   │   │   │   │           ├── defaults-es_ES.min.js
│   │       │   │   │   │   │           ├── defaults-eu.js
│   │       │   │   │   │   │           ├── defaults-eu.min.js
│   │       │   │   │   │   │           ├── defaults-fa_IR.js
│   │       │   │   │   │   │           ├── defaults-fa_IR.min.js
│   │       │   │   │   │   │           ├── defaults-fi_FI.js
│   │       │   │   │   │   │           ├── defaults-fi_FI.min.js
│   │       │   │   │   │   │           ├── defaults-fr_FR.js
│   │       │   │   │   │   │           ├── defaults-fr_FR.min.js
│   │       │   │   │   │   │           ├── defaults-hu_HU.js
│   │       │   │   │   │   │           ├── defaults-hu_HU.min.js
│   │       │   │   │   │   │           ├── defaults-id_ID.js
│   │       │   │   │   │   │           ├── defaults-id_ID.min.js
│   │       │   │   │   │   │           ├── defaults-it_IT.js
│   │       │   │   │   │   │           ├── defaults-it_IT.min.js
│   │       │   │   │   │   │           ├── defaults-ko_KR.js
│   │       │   │   │   │   │           ├── defaults-ko_KR.min.js
│   │       │   │   │   │   │           ├── defaults-lt_LT.js
│   │       │   │   │   │   │           ├── defaults-lt_LT.min.js
│   │       │   │   │   │   │           ├── defaults-nb_NO.js
│   │       │   │   │   │   │           ├── defaults-nb_NO.min.js
│   │       │   │   │   │   │           ├── defaults-nl_NL.js
│   │       │   │   │   │   │           ├── defaults-nl_NL.min.js
│   │       │   │   │   │   │           ├── defaults-pl_PL.js
│   │       │   │   │   │   │           ├── defaults-pl_PL.min.js
│   │       │   │   │   │   │           ├── defaults-pt_BR.js
│   │       │   │   │   │   │           ├── defaults-pt_BR.min.js
│   │       │   │   │   │   │           ├── defaults-pt_PT.js
│   │       │   │   │   │   │           ├── defaults-pt_PT.min.js
│   │       │   │   │   │   │           ├── defaults-ro_RO.js
│   │       │   │   │   │   │           ├── defaults-ro_RO.min.js
│   │       │   │   │   │   │           ├── defaults-ru_RU.js
│   │       │   │   │   │   │           ├── defaults-ru_RU.min.js
│   │       │   │   │   │   │           ├── defaults-sk_SK.js
│   │       │   │   │   │   │           ├── defaults-sk_SK.min.js
│   │       │   │   │   │   │           ├── defaults-sl_SI.js
│   │       │   │   │   │   │           ├── defaults-sl_SI.min.js
│   │       │   │   │   │   │           ├── defaults-sv_SE.js
│   │       │   │   │   │   │           ├── defaults-sv_SE.min.js
│   │       │   │   │   │   │           ├── defaults-tr_TR.js
│   │       │   │   │   │   │           ├── defaults-tr_TR.min.js
│   │       │   │   │   │   │           ├── defaults-ua_UA.js
│   │       │   │   │   │   │           ├── defaults-ua_UA.min.js
│   │       │   │   │   │   │           ├── defaults-zh_CN.js
│   │       │   │   │   │   │           ├── defaults-zh_CN.min.js
│   │       │   │   │   │   │           ├── defaults-zh_TW.js
│   │       │   │   │   │   │           └── defaults-zh_TW.min.js
│   │       │   │   │   │   ├── examples.md
│   │       │   │   │   │   ├── index.md
│   │       │   │   │   │   ├── methods.md
│   │       │   │   │   │   ├── options.md
│   │       │   │   │   │   └── playground
│   │       │   │   │   │       ├── index.html
│   │       │   │   │   │       └── plnkrOpener.js
│   │       │   │   │   └── mkdocs.yml
│   │       │   │   ├── js
│   │       │   │   │   ├── bootstrap-select.js
│   │       │   │   │   └── i18n
│   │       │   │   │       ├── defaults-ar_AR.js
│   │       │   │   │       ├── defaults-bg_BG.js
│   │       │   │   │       ├── defaults-cro_CRO.js
│   │       │   │   │       ├── defaults-cs_CZ.js
│   │       │   │   │       ├── defaults-da_DK.js
│   │       │   │   │       ├── defaults-de_DE.js
│   │       │   │   │       ├── defaults-en_US.js
│   │       │   │   │       ├── defaults-es_CL.js
│   │       │   │   │       ├── defaults-es_ES.js
│   │       │   │   │       ├── defaults-eu.js
│   │       │   │   │       ├── defaults-fa_IR.js
│   │       │   │   │       ├── defaults-fi_FI.js
│   │       │   │   │       ├── defaults-fr_FR.js
│   │       │   │   │       ├── defaults-hu_HU.js
│   │       │   │   │       ├── defaults-id_ID.js
│   │       │   │   │       ├── defaults-it_IT.js
│   │       │   │   │       ├── defaults-ko_KR.js
│   │       │   │   │       ├── defaults-lt_LT.js
│   │       │   │   │       ├── defaults-nb_NO.js
│   │       │   │   │       ├── defaults-nl_NL.js
│   │       │   │   │       ├── defaults-pl_PL.js
│   │       │   │   │       ├── defaults-pt_BR.js
│   │       │   │   │       ├── defaults-pt_PT.js
│   │       │   │   │       ├── defaults-ro_RO.js
│   │       │   │   │       ├── defaults-ru_RU.js
│   │       │   │   │       ├── defaults-sk_SK.js
│   │       │   │   │       ├── defaults-sl_SI.js
│   │       │   │   │       ├── defaults-sv_SE.js
│   │       │   │   │       ├── defaults-tr_TR.js
│   │       │   │   │       ├── defaults-ua_UA.js
│   │       │   │   │       ├── defaults-zh_CN.js
│   │       │   │   │       └── defaults-zh_TW.js
│   │       │   │   ├── less
│   │       │   │   │   ├── bootstrap-select.less
│   │       │   │   │   └── variables.less
│   │       │   │   ├── nuget
│   │       │   │   │   ├── MyGet.ps1
│   │       │   │   │   └── bootstrap-select.nuspec
│   │       │   │   └── sass
│   │       │   │       ├── bootstrap-select.scss
│   │       │   │       └── variables.scss
│   │       │   ├── datatables.net
│   │       │   │   ├── License.txt
│   │       │   │   ├── Readme.md
│   │       │   │   ├── bower.json
│   │       │   │   └── js
│   │       │   │       ├── jquery.dataTables.js
│   │       │   │       └── jquery.dataTables.min.js
│   │       │   ├── datatables.net-bs
│   │       │   │   ├── License.txt
│   │       │   │   ├── Readme.md
│   │       │   │   ├── bower.json
│   │       │   │   ├── css
│   │       │   │   │   ├── dataTables.bootstrap.css
│   │       │   │   │   └── dataTables.bootstrap.min.css
│   │       │   │   └── js
│   │       │   │       ├── dataTables.bootstrap.js
│   │       │   │       └── dataTables.bootstrap.min.js
│   │       │   ├── jquery
│   │       │   │   ├── AUTHORS.txt
│   │       │   │   ├── LICENSE.txt
│   │       │   │   ├── README.md
│   │       │   │   ├── bower.json
│   │       │   │   ├── dist
│   │       │   │   │   ├── jquery.js
│   │       │   │   │   ├── jquery.min.js
│   │       │   │   │   └── jquery.min.map
│   │       │   │   ├── external
│   │       │   │   │   └── sizzle
│   │       │   │   │       ├── LICENSE.txt
│   │       │   │   │       └── dist
│   │       │   │   │           ├── sizzle.js
│   │       │   │   │           ├── sizzle.min.js
│   │       │   │   │           └── sizzle.min.map
│   │       │   │   └── src
│   │       │   │       ├── ajax
│   │       │   │       │   ├── jsonp.js
│   │       │   │       │   ├── load.js
│   │       │   │       │   ├── parseJSON.js
│   │       │   │       │   ├── parseXML.js
│   │       │   │       │   ├── script.js
│   │       │   │       │   ├── var
│   │       │   │       │   │   ├── location.js
│   │       │   │       │   │   ├── nonce.js
│   │       │   │       │   │   └── rquery.js
│   │       │   │       │   └── xhr.js
│   │       │   │       ├── ajax.js
│   │       │   │       ├── attributes
│   │       │   │       │   ├── attr.js
│   │       │   │       │   ├── classes.js
│   │       │   │       │   ├── prop.js
│   │       │   │       │   ├── support.js
│   │       │   │       │   └── val.js
│   │       │   │       ├── attributes.js
│   │       │   │       ├── callbacks.js
│   │       │   │       ├── core
│   │       │   │       │   ├── access.js
│   │       │   │       │   ├── init.js
│   │       │   │       │   ├── parseHTML.js
│   │       │   │       │   ├── ready.js
│   │       │   │       │   └── var
│   │       │   │       │       └── rsingleTag.js
│   │       │   │       ├── core.js
│   │       │   │       ├── css
│   │       │   │       │   ├── addGetHookIf.js
│   │       │   │       │   ├── adjustCSS.js
│   │       │   │       │   ├── curCSS.js
│   │       │   │       │   ├── defaultDisplay.js
│   │       │   │       │   ├── hiddenVisibleSelectors.js
│   │       │   │       │   ├── showHide.js
│   │       │   │       │   ├── support.js
│   │       │   │       │   └── var
│   │       │   │       │       ├── cssExpand.js
│   │       │   │       │       ├── getStyles.js
│   │       │   │       │       ├── isHidden.js
│   │       │   │       │       ├── rmargin.js
│   │       │   │       │       ├── rnumnonpx.js
│   │       │   │       │       └── swap.js
│   │       │   │       ├── css.js
│   │       │   │       ├── data
│   │       │   │       │   ├── Data.js
│   │       │   │       │   └── var
│   │       │   │       │       ├── acceptData.js
│   │       │   │       │       ├── dataPriv.js
│   │       │   │       │       └── dataUser.js
│   │       │   │       ├── data.js
│   │       │   │       ├── deferred.js
│   │       │   │       ├── deprecated.js
│   │       │   │       ├── dimensions.js
│   │       │   │       ├── effects
│   │       │   │       │   ├── Tween.js
│   │       │   │       │   └── animatedSelector.js
│   │       │   │       ├── effects.js
│   │       │   │       ├── event
│   │       │   │       │   ├── ajax.js
│   │       │   │       │   ├── alias.js
│   │       │   │       │   ├── focusin.js
│   │       │   │       │   ├── support.js
│   │       │   │       │   └── trigger.js
│   │       │   │       ├── event.js
│   │       │   │       ├── exports
│   │       │   │       │   ├── amd.js
│   │       │   │       │   └── global.js
│   │       │   │       ├── intro.js
│   │       │   │       ├── jquery.js
│   │       │   │       ├── manipulation
│   │       │   │       │   ├── _evalUrl.js
│   │       │   │       │   ├── buildFragment.js
│   │       │   │       │   ├── getAll.js
│   │       │   │       │   ├── setGlobalEval.js
│   │       │   │       │   ├── support.js
│   │       │   │       │   ├── var
│   │       │   │       │   │   ├── rcheckableType.js
│   │       │   │       │   │   ├── rscriptType.js
│   │       │   │       │   │   └── rtagName.js
│   │       │   │       │   └── wrapMap.js
│   │       │   │       ├── manipulation.js
│   │       │   │       ├── offset.js
│   │       │   │       ├── outro.js
│   │       │   │       ├── queue
│   │       │   │       │   └── delay.js
│   │       │   │       ├── queue.js
│   │       │   │       ├── selector-native.js
│   │       │   │       ├── selector-sizzle.js
│   │       │   │       ├── selector.js
│   │       │   │       ├── serialize.js
│   │       │   │       ├── traversing
│   │       │   │       │   ├── findFilter.js
│   │       │   │       │   └── var
│   │       │   │       │       ├── dir.js
│   │       │   │       │       ├── rneedsContext.js
│   │       │   │       │       └── siblings.js
│   │       │   │       ├── traversing.js
│   │       │   │       ├── var
│   │       │   │       │   ├── arr.js
│   │       │   │       │   ├── class2type.js
│   │       │   │       │   ├── concat.js
│   │       │   │       │   ├── document.js
│   │       │   │       │   ├── documentElement.js
│   │       │   │       │   ├── hasOwn.js
│   │       │   │       │   ├── indexOf.js
│   │       │   │       │   ├── pnum.js
│   │       │   │       │   ├── push.js
│   │       │   │       │   ├── rcssNum.js
│   │       │   │       │   ├── rnotwhite.js
│   │       │   │       │   ├── slice.js
│   │       │   │       │   ├── support.js
│   │       │   │       │   └── toString.js
│   │       │   │       └── wrap.js
│   │       │   ├── jquery.easy-pie-chart
│   │       │   │   ├── Gruntfile.js
│   │       │   │   ├── LICENSE
│   │       │   │   ├── Readme.md
│   │       │   │   ├── bower.json
│   │       │   │   ├── changelog.md
│   │       │   │   ├── dist
│   │       │   │   │   ├── angular.easypiechart.js
│   │       │   │   │   ├── angular.easypiechart.min.js
│   │       │   │   │   ├── easypiechart.js
│   │       │   │   │   ├── easypiechart.min.js
│   │       │   │   │   ├── jquery.easypiechart.js
│   │       │   │   │   └── jquery.easypiechart.min.js
│   │       │   │   ├── docs
│   │       │   │   │   ├── README.tmpl.md
│   │       │   │   │   ├── bagdes.md
│   │       │   │   │   ├── browser-support.md
│   │       │   │   │   ├── callbacks.md
│   │       │   │   │   ├── credits.md
│   │       │   │   │   ├── features.md
│   │       │   │   │   ├── get-started.md
│   │       │   │   │   ├── options.md
│   │       │   │   │   ├── plugin-api.md
│   │       │   │   │   └── test.md
│   │       │   │   ├── karma.conf.coffee
│   │       │   │   ├── package.json
│   │       │   │   ├── src
│   │       │   │   │   ├── angular.directive.js
│   │       │   │   │   ├── easypiechart.js
│   │       │   │   │   ├── jquery.plugin.js
│   │       │   │   │   └── renderer
│   │       │   │   │       └── canvas.js
│   │       │   │   └── test
│   │       │   │       ├── polyfills
│   │       │   │       │   └── bind.js
│   │       │   │       └── unit
│   │       │   │           ├── angular.directive.js
│   │       │   │           └── jquery.js
│   │       │   ├── moment
│   │       │   │   ├── CHANGELOG.md
│   │       │   │   ├── LICENSE
│   │       │   │   ├── README.md
│   │       │   │   ├── bower.json
│   │       │   │   ├── locale
│   │       │   │   │   ├── af.js
│   │       │   │   │   ├── ar-dz.js
│   │       │   │   │   ├── ar-ly.js
│   │       │   │   │   ├── ar-ma.js
│   │       │   │   │   ├── ar-sa.js
│   │       │   │   │   ├── ar-tn.js
│   │       │   │   │   ├── ar.js
│   │       │   │   │   ├── az.js
│   │       │   │   │   ├── be.js
│   │       │   │   │   ├── bg.js
│   │       │   │   │   ├── bn.js
│   │       │   │   │   ├── bo.js
│   │       │   │   │   ├── br.js
│   │       │   │   │   ├── bs.js
│   │       │   │   │   ├── ca.js
│   │       │   │   │   ├── cs.js
│   │       │   │   │   ├── cv.js
│   │       │   │   │   ├── cy.js
│   │       │   │   │   ├── da.js
│   │       │   │   │   ├── de-at.js
│   │       │   │   │   ├── de.js
│   │       │   │   │   ├── dv.js
│   │       │   │   │   ├── el.js
│   │       │   │   │   ├── en-au.js
│   │       │   │   │   ├── en-ca.js
│   │       │   │   │   ├── en-gb.js
│   │       │   │   │   ├── en-ie.js
│   │       │   │   │   ├── en-nz.js
│   │       │   │   │   ├── eo.js
│   │       │   │   │   ├── es-do.js
│   │       │   │   │   ├── es.js
│   │       │   │   │   ├── et.js
│   │       │   │   │   ├── eu.js
│   │       │   │   │   ├── fa.js
│   │       │   │   │   ├── fi.js
│   │       │   │   │   ├── fo.js
│   │       │   │   │   ├── fr-ca.js
│   │       │   │   │   ├── fr-ch.js
│   │       │   │   │   ├── fr.js
│   │       │   │   │   ├── fy.js
│   │       │   │   │   ├── gd.js
│   │       │   │   │   ├── gl.js
│   │       │   │   │   ├── he.js
│   │       │   │   │   ├── hi.js
│   │       │   │   │   ├── hr.js
│   │       │   │   │   ├── hu.js
│   │       │   │   │   ├── hy-am.js
│   │       │   │   │   ├── id.js
│   │       │   │   │   ├── is.js
│   │       │   │   │   ├── it.js
│   │       │   │   │   ├── ja.js
│   │       │   │   │   ├── jv.js
│   │       │   │   │   ├── ka.js
│   │       │   │   │   ├── kk.js
│   │       │   │   │   ├── km.js
│   │       │   │   │   ├── ko.js
│   │       │   │   │   ├── ky.js
│   │       │   │   │   ├── lb.js
│   │       │   │   │   ├── lo.js
│   │       │   │   │   ├── lt.js
│   │       │   │   │   ├── lv.js
│   │       │   │   │   ├── me.js
│   │       │   │   │   ├── mi.js
│   │       │   │   │   ├── mk.js
│   │       │   │   │   ├── ml.js
│   │       │   │   │   ├── mr.js
│   │       │   │   │   ├── ms-my.js
│   │       │   │   │   ├── ms.js
│   │       │   │   │   ├── my.js
│   │       │   │   │   ├── nb.js
│   │       │   │   │   ├── ne.js
│   │       │   │   │   ├── nl-be.js
│   │       │   │   │   ├── nl.js
│   │       │   │   │   ├── nn.js
│   │       │   │   │   ├── pa-in.js
│   │       │   │   │   ├── pl.js
│   │       │   │   │   ├── pt-br.js
│   │       │   │   │   ├── pt.js
│   │       │   │   │   ├── ro.js
│   │       │   │   │   ├── ru.js
│   │       │   │   │   ├── se.js
│   │       │   │   │   ├── si.js
│   │       │   │   │   ├── sk.js
│   │       │   │   │   ├── sl.js
│   │       │   │   │   ├── sq.js
│   │       │   │   │   ├── sr-cyrl.js
│   │       │   │   │   ├── sr.js
│   │       │   │   │   ├── ss.js
│   │       │   │   │   ├── sv.js
│   │       │   │   │   ├── sw.js
│   │       │   │   │   ├── ta.js
│   │       │   │   │   ├── te.js
│   │       │   │   │   ├── tet.js
│   │       │   │   │   ├── th.js
│   │       │   │   │   ├── tl-ph.js
│   │       │   │   │   ├── tlh.js
│   │       │   │   │   ├── tr.js
│   │       │   │   │   ├── tzl.js
│   │       │   │   │   ├── tzm-latn.js
│   │       │   │   │   ├── tzm.js
│   │       │   │   │   ├── uk.js
│   │       │   │   │   ├── uz.js
│   │       │   │   │   ├── vi.js
│   │       │   │   │   ├── x-pseudo.js
│   │       │   │   │   ├── yo.js
│   │       │   │   │   ├── zh-cn.js
│   │       │   │   │   ├── zh-hk.js
│   │       │   │   │   └── zh-tw.js
│   │       │   │   ├── min
│   │       │   │   │   ├── locales.js
│   │       │   │   │   ├── locales.min.js
│   │       │   │   │   ├── moment-with-locales.js
│   │       │   │   │   ├── moment-with-locales.min.js
│   │       │   │   │   ├── moment.min.js
│   │       │   │   │   └── tests.js
│   │       │   │   ├── moment.d.ts
│   │       │   │   ├── moment.js
│   │       │   │   ├── src
│   │       │   │   │   ├── lib
│   │       │   │   │   │   ├── create
│   │       │   │   │   │   │   ├── check-overflow.js
│   │       │   │   │   │   │   ├── date-from-array.js
│   │       │   │   │   │   │   ├── from-anything.js
│   │       │   │   │   │   │   ├── from-array.js
│   │       │   │   │   │   │   ├── from-object.js
│   │       │   │   │   │   │   ├── from-string-and-array.js
│   │       │   │   │   │   │   ├── from-string-and-format.js
│   │       │   │   │   │   │   ├── from-string.js
│   │       │   │   │   │   │   ├── local.js
│   │       │   │   │   │   │   ├── parsing-flags.js
│   │       │   │   │   │   │   ├── utc.js
│   │       │   │   │   │   │   └── valid.js
│   │       │   │   │   │   ├── duration
│   │       │   │   │   │   │   ├── abs.js
│   │       │   │   │   │   │   ├── add-subtract.js
│   │       │   │   │   │   │   ├── as.js
│   │       │   │   │   │   │   ├── bubble.js
│   │       │   │   │   │   │   ├── constructor.js
│   │       │   │   │   │   │   ├── create.js
│   │       │   │   │   │   │   ├── duration.js
│   │       │   │   │   │   │   ├── get.js
│   │       │   │   │   │   │   ├── humanize.js
│   │       │   │   │   │   │   ├── iso-string.js
│   │       │   │   │   │   │   └── prototype.js
│   │       │   │   │   │   ├── format
│   │       │   │   │   │   │   └── format.js
│   │       │   │   │   │   ├── locale
│   │       │   │   │   │   │   ├── base-config.js
│   │       │   │   │   │   │   ├── calendar.js
│   │       │   │   │   │   │   ├── constructor.js
│   │       │   │   │   │   │   ├── en.js
│   │       │   │   │   │   │   ├── formats.js
│   │       │   │   │   │   │   ├── invalid.js
│   │       │   │   │   │   │   ├── lists.js
│   │       │   │   │   │   │   ├── locale.js
│   │       │   │   │   │   │   ├── locales.js
│   │       │   │   │   │   │   ├── ordinal.js
│   │       │   │   │   │   │   ├── pre-post-format.js
│   │       │   │   │   │   │   ├── prototype.js
│   │       │   │   │   │   │   ├── relative.js
│   │       │   │   │   │   │   └── set.js
│   │       │   │   │   │   ├── moment
│   │       │   │   │   │   │   ├── add-subtract.js
│   │       │   │   │   │   │   ├── calendar.js
│   │       │   │   │   │   │   ├── clone.js
│   │       │   │   │   │   │   ├── compare.js
│   │       │   │   │   │   │   ├── constructor.js
│   │       │   │   │   │   │   ├── creation-data.js
│   │       │   │   │   │   │   ├── diff.js
│   │       │   │   │   │   │   ├── format.js
│   │       │   │   │   │   │   ├── from.js
│   │       │   │   │   │   │   ├── get-set.js
│   │       │   │   │   │   │   ├── locale.js
│   │       │   │   │   │   │   ├── min-max.js
│   │       │   │   │   │   │   ├── moment.js
│   │       │   │   │   │   │   ├── now.js
│   │       │   │   │   │   │   ├── prototype.js
│   │       │   │   │   │   │   ├── start-end-of.js
│   │       │   │   │   │   │   ├── to-type.js
│   │       │   │   │   │   │   ├── to.js
│   │       │   │   │   │   │   └── valid.js
│   │       │   │   │   │   ├── parse
│   │       │   │   │   │   │   ├── regex.js
│   │       │   │   │   │   │   └── token.js
│   │       │   │   │   │   ├── units
│   │       │   │   │   │   │   ├── aliases.js
│   │       │   │   │   │   │   ├── constants.js
│   │       │   │   │   │   │   ├── day-of-month.js
│   │       │   │   │   │   │   ├── day-of-week.js
│   │       │   │   │   │   │   ├── day-of-year.js
│   │       │   │   │   │   │   ├── hour.js
│   │       │   │   │   │   │   ├── millisecond.js
│   │       │   │   │   │   │   ├── minute.js
│   │       │   │   │   │   │   ├── month.js
│   │       │   │   │   │   │   ├── offset.js
│   │       │   │   │   │   │   ├── priorities.js
│   │       │   │   │   │   │   ├── quarter.js
│   │       │   │   │   │   │   ├── second.js
│   │       │   │   │   │   │   ├── timestamp.js
│   │       │   │   │   │   │   ├── timezone.js
│   │       │   │   │   │   │   ├── units.js
│   │       │   │   │   │   │   ├── week-calendar-utils.js
│   │       │   │   │   │   │   ├── week-year.js
│   │       │   │   │   │   │   ├── week.js
│   │       │   │   │   │   │   └── year.js
│   │       │   │   │   │   └── utils
│   │       │   │   │   │       ├── abs-ceil.js
│   │       │   │   │   │       ├── abs-floor.js
│   │       │   │   │   │       ├── abs-round.js
│   │       │   │   │   │       ├── compare-arrays.js
│   │       │   │   │   │       ├── defaults.js
│   │       │   │   │   │       ├── deprecate.js
│   │       │   │   │   │       ├── extend.js
│   │       │   │   │   │       ├── has-own-prop.js
│   │       │   │   │   │       ├── hooks.js
│   │       │   │   │   │       ├── index-of.js
│   │       │   │   │   │       ├── is-array.js
│   │       │   │   │   │       ├── is-date.js
│   │       │   │   │   │       ├── is-function.js
│   │       │   │   │   │       ├── is-number.js
│   │       │   │   │   │       ├── is-object-empty.js
│   │       │   │   │   │       ├── is-object.js
│   │       │   │   │   │       ├── is-undefined.js
│   │       │   │   │   │       ├── keys.js
│   │       │   │   │   │       ├── map.js
│   │       │   │   │   │       ├── some.js
│   │       │   │   │   │       ├── to-int.js
│   │       │   │   │   │       └── zero-fill.js
│   │       │   │   │   ├── locale
│   │       │   │   │   │   ├── af.js
│   │       │   │   │   │   ├── ar-dz.js
│   │       │   │   │   │   ├── ar-ly.js
│   │       │   │   │   │   ├── ar-ma.js
│   │       │   │   │   │   ├── ar-sa.js
│   │       │   │   │   │   ├── ar-tn.js
│   │       │   │   │   │   ├── ar.js
│   │       │   │   │   │   ├── az.js
│   │       │   │   │   │   ├── be.js
│   │       │   │   │   │   ├── bg.js
│   │       │   │   │   │   ├── bn.js
│   │       │   │   │   │   ├── bo.js
│   │       │   │   │   │   ├── br.js
│   │       │   │   │   │   ├── bs.js
│   │       │   │   │   │   ├── ca.js
│   │       │   │   │   │   ├── cs.js
│   │       │   │   │   │   ├── cv.js
│   │       │   │   │   │   ├── cy.js
│   │       │   │   │   │   ├── da.js
│   │       │   │   │   │   ├── de-at.js
│   │       │   │   │   │   ├── de.js
│   │       │   │   │   │   ├── dv.js
│   │       │   │   │   │   ├── el.js
│   │       │   │   │   │   ├── en-au.js
│   │       │   │   │   │   ├── en-ca.js
│   │       │   │   │   │   ├── en-gb.js
│   │       │   │   │   │   ├── en-ie.js
│   │       │   │   │   │   ├── en-nz.js
│   │       │   │   │   │   ├── eo.js
│   │       │   │   │   │   ├── es-do.js
│   │       │   │   │   │   ├── es.js
│   │       │   │   │   │   ├── et.js
│   │       │   │   │   │   ├── eu.js
│   │       │   │   │   │   ├── fa.js
│   │       │   │   │   │   ├── fi.js
│   │       │   │   │   │   ├── fo.js
│   │       │   │   │   │   ├── fr-ca.js
│   │       │   │   │   │   ├── fr-ch.js
│   │       │   │   │   │   ├── fr.js
│   │       │   │   │   │   ├── fy.js
│   │       │   │   │   │   ├── gd.js
│   │       │   │   │   │   ├── gl.js
│   │       │   │   │   │   ├── he.js
│   │       │   │   │   │   ├── hi.js
│   │       │   │   │   │   ├── hr.js
│   │       │   │   │   │   ├── hu.js
│   │       │   │   │   │   ├── hy-am.js
│   │       │   │   │   │   ├── id.js
│   │       │   │   │   │   ├── is.js
│   │       │   │   │   │   ├── it.js
│   │       │   │   │   │   ├── ja.js
│   │       │   │   │   │   ├── jv.js
│   │       │   │   │   │   ├── ka.js
│   │       │   │   │   │   ├── kk.js
│   │       │   │   │   │   ├── km.js
│   │       │   │   │   │   ├── ko.js
│   │       │   │   │   │   ├── ky.js
│   │       │   │   │   │   ├── lb.js
│   │       │   │   │   │   ├── lo.js
│   │       │   │   │   │   ├── lt.js
│   │       │   │   │   │   ├── lv.js
│   │       │   │   │   │   ├── me.js
│   │       │   │   │   │   ├── mi.js
│   │       │   │   │   │   ├── mk.js
│   │       │   │   │   │   ├── ml.js
│   │       │   │   │   │   ├── mr.js
│   │       │   │   │   │   ├── ms-my.js
│   │       │   │   │   │   ├── ms.js
│   │       │   │   │   │   ├── my.js
│   │       │   │   │   │   ├── nb.js
│   │       │   │   │   │   ├── ne.js
│   │       │   │   │   │   ├── nl-be.js
│   │       │   │   │   │   ├── nl.js
│   │       │   │   │   │   ├── nn.js
│   │       │   │   │   │   ├── pa-in.js
│   │       │   │   │   │   ├── pl.js
│   │       │   │   │   │   ├── pt-br.js
│   │       │   │   │   │   ├── pt.js
│   │       │   │   │   │   ├── ro.js
│   │       │   │   │   │   ├── ru.js
│   │       │   │   │   │   ├── se.js
│   │       │   │   │   │   ├── si.js
│   │       │   │   │   │   ├── sk.js
│   │       │   │   │   │   ├── sl.js
│   │       │   │   │   │   ├── sq.js
│   │       │   │   │   │   ├── sr-cyrl.js
│   │       │   │   │   │   ├── sr.js
│   │       │   │   │   │   ├── ss.js
│   │       │   │   │   │   ├── sv.js
│   │       │   │   │   │   ├── sw.js
│   │       │   │   │   │   ├── ta.js
│   │       │   │   │   │   ├── te.js
│   │       │   │   │   │   ├── tet.js
│   │       │   │   │   │   ├── th.js
│   │       │   │   │   │   ├── tl-ph.js
│   │       │   │   │   │   ├── tlh.js
│   │       │   │   │   │   ├── tr.js
│   │       │   │   │   │   ├── tzl.js
│   │       │   │   │   │   ├── tzm-latn.js
│   │       │   │   │   │   ├── tzm.js
│   │       │   │   │   │   ├── uk.js
│   │       │   │   │   │   ├── uz.js
│   │       │   │   │   │   ├── vi.js
│   │       │   │   │   │   ├── x-pseudo.js
│   │       │   │   │   │   ├── yo.js
│   │       │   │   │   │   ├── zh-cn.js
│   │       │   │   │   │   ├── zh-hk.js
│   │       │   │   │   │   └── zh-tw.js
│   │       │   │   │   └── moment.js
│   │       │   │   └── templates
│   │       │   │       ├── default.js
│   │       │   │       ├── locale-header.js
│   │       │   │       └── test-header.js
│   │       │   ├── notifyjs
│   │       │   │   ├── CHANGES.md
│   │       │   │   ├── README.md
│   │       │   │   ├── bower.json
│   │       │   │   ├── dist
│   │       │   │   │   ├── notify.js
│   │       │   │   │   └── styles
│   │       │   │   │       └── metro
│   │       │   │   │           ├── notify-metro.css
│   │       │   │   │           └── notify-metro.js
│   │       │   │   ├── examples
│   │       │   │   │   ├── classes.html
│   │       │   │   │   ├── images
│   │       │   │   │   │   └── Mail.png
│   │       │   │   │   ├── inlines.html
│   │       │   │   │   ├── metro.html
│   │       │   │   │   ├── multi-text.html
│   │       │   │   │   └── position.html
│   │       │   │   ├── notify.jquery.json
│   │       │   │   └── package.json
│   │       │   ├── plupload
│   │       │   │   ├── examples
│   │       │   │   │   ├── custom.html
│   │       │   │   │   ├── dump.php
│   │       │   │   │   ├── jquery
│   │       │   │   │   │   ├── all_runtimes.html
│   │       │   │   │   │   ├── jquery_ui_widget.html
│   │       │   │   │   │   ├── queue_widget.html
│   │       │   │   │   │   └── s3.php
│   │       │   │   │   └── upload.php
│   │       │   │   ├── js
│   │       │   │   │   ├── Moxie.swf
│   │       │   │   │   ├── Moxie.xap
│   │       │   │   │   ├── i18n
│   │       │   │   │   │   ├── ar.js
│   │       │   │   │   │   ├── bs.js
│   │       │   │   │   │   ├── ca.js
│   │       │   │   │   │   ├── cs.js
│   │       │   │   │   │   ├── cy.js
│   │       │   │   │   │   ├── da.js
│   │       │   │   │   │   ├── de.js
│   │       │   │   │   │   ├── el.js
│   │       │   │   │   │   ├── en.js
│   │       │   │   │   │   ├── es.js
│   │       │   │   │   │   ├── et.js
│   │       │   │   │   │   ├── fa.js
│   │       │   │   │   │   ├── fi.js
│   │       │   │   │   │   ├── fr.js
│   │       │   │   │   │   ├── he.js
│   │       │   │   │   │   ├── hr.js
│   │       │   │   │   │   ├── hu.js
│   │       │   │   │   │   ├── hy.js
│   │       │   │   │   │   ├── id.js
│   │       │   │   │   │   ├── it.js
│   │       │   │   │   │   ├── ja.js
│   │       │   │   │   │   ├── ka.js
│   │       │   │   │   │   ├── kk.js
│   │       │   │   │   │   ├── ko.js
│   │       │   │   │   │   ├── lt.js
│   │       │   │   │   │   ├── lv.js
│   │       │   │   │   │   ├── nl.js
│   │       │   │   │   │   ├── pl.js
│   │       │   │   │   │   ├── pt_BR.js
│   │       │   │   │   │   ├── ro.js
│   │       │   │   │   │   ├── ru.js
│   │       │   │   │   │   ├── sk.js
│   │       │   │   │   │   ├── sr.js
│   │       │   │   │   │   ├── sv.js
│   │       │   │   │   │   ├── th_TH.js
│   │       │   │   │   │   ├── tr.js
│   │       │   │   │   │   ├── uk_UA.js
│   │       │   │   │   │   ├── zh_CN.js
│   │       │   │   │   │   └── zh_TW.js
│   │       │   │   │   ├── jquery.plupload.queue
│   │       │   │   │   │   ├── css
│   │       │   │   │   │   │   └── jquery.plupload.queue.css
│   │       │   │   │   │   ├── img
│   │       │   │   │   │   │   ├── backgrounds.gif
│   │       │   │   │   │   │   ├── buttons-disabled.png
│   │       │   │   │   │   │   ├── buttons.png
│   │       │   │   │   │   │   ├── delete.gif
│   │       │   │   │   │   │   ├── done.gif
│   │       │   │   │   │   │   ├── error.gif
│   │       │   │   │   │   │   ├── throbber.gif
│   │       │   │   │   │   │   └── transp50.png
│   │       │   │   │   │   ├── jquery.plupload.queue.js
│   │       │   │   │   │   └── jquery.plupload.queue.min.js
│   │       │   │   │   ├── jquery.ui.plupload
│   │       │   │   │   │   ├── css
│   │       │   │   │   │   │   └── jquery.ui.plupload.css
│   │       │   │   │   │   ├── img
│   │       │   │   │   │   │   ├── loading.gif
│   │       │   │   │   │   │   └── plupload.png
│   │       │   │   │   │   ├── jquery.ui.plupload.js
│   │       │   │   │   │   └── jquery.ui.plupload.min.js
│   │       │   │   │   ├── moxie.js
│   │       │   │   │   ├── moxie.min.js
│   │       │   │   │   ├── plupload.dev.js
│   │       │   │   │   ├── plupload.full.min.js
│   │       │   │   │   └── plupload.min.js
│   │       │   │   ├── license.txt
│   │       │   │   └── readme.md
│   │       │   ├── qiniu
│   │       │   │   ├── Gruntfile.js
│   │       │   │   ├── Makefile
│   │       │   │   ├── README.md
│   │       │   │   ├── bower.json
│   │       │   │   ├── demo
│   │       │   │   │   ├── config.js.example
│   │       │   │   │   ├── images
│   │       │   │   │   │   ├── default.png
│   │       │   │   │   │   ├── favicon.ico
│   │       │   │   │   │   └── loading.gif
│   │       │   │   │   ├── scripts
│   │       │   │   │   │   ├── formdata.js
│   │       │   │   │   │   ├── highlight.js
│   │       │   │   │   │   ├── main.js
│   │       │   │   │   │   ├── multiple.js
│   │       │   │   │   │   └── ui.js
│   │       │   │   │   ├── server.js
│   │       │   │   │   ├── styles
│   │       │   │   │   │   ├── formdata.css
│   │       │   │   │   │   ├── highlight.css
│   │       │   │   │   │   └── main.css
│   │       │   │   │   └── views
│   │       │   │   │       ├── formdata.html
│   │       │   │   │       ├── index.html
│   │       │   │   │       └── multiple.html
│   │       │   │   ├── dist
│   │       │   │   │   ├── qiniu.js
│   │       │   │   │   ├── qiniu.min.js
│   │       │   │   │   └── qiniu.min.map
│   │       │   │   ├── package.json
│   │       │   │   └── src
│   │       │   │       └── qiniu.js
│   │       │   ├── requirejs
│   │       │   │   ├── README.md
│   │       │   │   ├── bower.json
│   │       │   │   └── require.js
│   │       │   ├── smalot-bootstrap-datetimepicker
│   │       │   │   ├── Gruntfile.js
│   │       │   │   ├── LICENSE
│   │       │   │   ├── README.md
│   │       │   │   ├── bower.json
│   │       │   │   ├── css
│   │       │   │   │   ├── bootstrap-datetimepicker.css
│   │       │   │   │   └── bootstrap-datetimepicker.min.css
│   │       │   │   ├── js
│   │       │   │   │   ├── bootstrap-datetimepicker.js
│   │       │   │   │   ├── bootstrap-datetimepicker.min.js
│   │       │   │   │   └── locales
│   │       │   │   │       ├── bootstrap-datetimepicker.ar.js
│   │       │   │   │       ├── bootstrap-datetimepicker.az.js
│   │       │   │   │       ├── bootstrap-datetimepicker.bg.js
│   │       │   │   │       ├── bootstrap-datetimepicker.bn.js
│   │       │   │   │       ├── bootstrap-datetimepicker.ca.js
│   │       │   │   │       ├── bootstrap-datetimepicker.cs.js
│   │       │   │   │       ├── bootstrap-datetimepicker.da.js
│   │       │   │   │       ├── bootstrap-datetimepicker.de.js
│   │       │   │   │       ├── bootstrap-datetimepicker.ee.js
│   │       │   │   │       ├── bootstrap-datetimepicker.el.js
│   │       │   │   │       ├── bootstrap-datetimepicker.es.js
│   │       │   │   │       ├── bootstrap-datetimepicker.fi.js
│   │       │   │   │       ├── bootstrap-datetimepicker.fr.js
│   │       │   │   │       ├── bootstrap-datetimepicker.he.js
│   │       │   │   │       ├── bootstrap-datetimepicker.hr.js
│   │       │   │   │       ├── bootstrap-datetimepicker.hu.js
│   │       │   │   │       ├── bootstrap-datetimepicker.hy.js
│   │       │   │   │       ├── bootstrap-datetimepicker.id.js
│   │       │   │   │       ├── bootstrap-datetimepicker.is.js
│   │       │   │   │       ├── bootstrap-datetimepicker.it.js
│   │       │   │   │       ├── bootstrap-datetimepicker.ja.js
│   │       │   │   │       ├── bootstrap-datetimepicker.ka.js
│   │       │   │   │       ├── bootstrap-datetimepicker.ko.js
│   │       │   │   │       ├── bootstrap-datetimepicker.lt.js
│   │       │   │   │       ├── bootstrap-datetimepicker.lv.js
│   │       │   │   │       ├── bootstrap-datetimepicker.ms.js
│   │       │   │   │       ├── bootstrap-datetimepicker.nb.js
│   │       │   │   │       ├── bootstrap-datetimepicker.nl.js
│   │       │   │   │       ├── bootstrap-datetimepicker.no.js
│   │       │   │   │       ├── bootstrap-datetimepicker.pl.js
│   │       │   │   │       ├── bootstrap-datetimepicker.pt-BR.js
│   │       │   │   │       ├── bootstrap-datetimepicker.pt.js
│   │       │   │   │       ├── bootstrap-datetimepicker.ro.js
│   │       │   │   │       ├── bootstrap-datetimepicker.rs-latin.js
│   │       │   │   │       ├── bootstrap-datetimepicker.rs.js
│   │       │   │   │       ├── bootstrap-datetimepicker.ru.js
│   │       │   │   │       ├── bootstrap-datetimepicker.sk.js
│   │       │   │   │       ├── bootstrap-datetimepicker.sl.js
│   │       │   │   │       ├── bootstrap-datetimepicker.sv.js
│   │       │   │   │       ├── bootstrap-datetimepicker.sw.js
│   │       │   │   │       ├── bootstrap-datetimepicker.th.js
│   │       │   │   │       ├── bootstrap-datetimepicker.tr.js
│   │       │   │   │       ├── bootstrap-datetimepicker.ua.js
│   │       │   │   │       ├── bootstrap-datetimepicker.uk.js
│   │       │   │   │       ├── bootstrap-datetimepicker.zh-CN.js
│   │       │   │   │       └── bootstrap-datetimepicker.zh-TW.js
│   │       │   │   ├── less
│   │       │   │   │   └── datetimepicker.less
│   │       │   │   └── package.json
│   │       │   └── toastr
│   │       │       ├── LICENSE
│   │       │       ├── README.md
│   │       │       ├── bower.json
│   │       │       ├── toastr.css
│   │       │       ├── toastr.js
│   │       │       ├── toastr.js.map
│   │       │       ├── toastr.less
│   │       │       ├── toastr.min.css
│   │       │       ├── toastr.min.js
│   │       │       └── toastr.scss
│   │       └── gadmin
│   │           ├── favicon.ico
│   │           └── libs.js
│   └── screenshot
│       ├── component.png
│       ├── dashboard.png
│       └── privilege.png
├── h5games
│   └── README.md
├── imbroker
│   ├── emqtt_benchmark
│   ├── emqttd
│   └── mosca
├── kafka
│   ├── README.md
│   ├── kafka-client.log
│   ├── kafka-server.log
│   ├── kafka_2.11-0.10.1.0
│   │   ├── LICENSE
│   │   ├── NOTICE
│   │   ├── config
│   │   │   ├── connect-console-sink.properties
│   │   │   ├── connect-console-source.properties
│   │   │   ├── connect-distributed.properties
│   │   │   ├── connect-file-sink.properties
│   │   │   ├── connect-file-source.properties
│   │   │   ├── connect-log4j.properties
│   │   │   ├── connect-standalone.properties
│   │   │   ├── consumer.properties
│   │   │   ├── log4j.properties
│   │   │   ├── producer.properties
│   │   │   ├── server.properties
│   │   │   ├── tools-log4j.properties
│   │   │   └── zookeeper.properties
│   │   └── logs
│   │       ├── controller.log
│   │       ├── controller.log.2016-11-28-12
│   │       ├── controller.log.2016-11-28-13
│   │       ├── controller.log.2016-11-28-22
│   │       ├── kafka-authorizer.log
│   │       ├── kafka-request.log
│   │       ├── kafkaServer-gc.log
│   │       ├── log-cleaner.log
│   │       ├── server.log
│   │       ├── server.log.2016-11-28-12
│   │       ├── server.log.2016-11-28-13
│   │       ├── server.log.2016-11-28-22
│   │       ├── state-change.log
│   │       └── zookeeper-gc.log
│   └── zookeeper.log
├── lex_yacc
│   ├── README.md
│   ├── compile.sh
│   ├── example
│   ├── example.l
│   ├── example.y
│   ├── lex-tutorial-master
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── config.in
│   │   ├── lex.yy.c
│   │   ├── makefile
│   │   ├── myscanner
│   │   ├── myscanner.c
│   │   ├── myscanner.h
│   │   └── myscanner.l
│   ├── lex.yy.c
│   ├── y.tab.c
│   ├── y.tab.h
│   └── yacc-tutorial-master
│       ├── README.md
│       ├── calc
│       ├── calc.l
│       ├── calc.y
│       ├── lex.yy.c
│       ├── makefile
│       ├── y.tab.c
│       └── y.tab.h
├── logserver
│   ├── LICENSE
│   ├── README.md
│   ├── ThinkPHP
│   │   ├── Library
│   │   │   └── Think
│   │   │       └── Think.class.php
│   │   ├── Mode
│   │   │   └── common.php
│   │   └── ThinkPHP.php
│   ├── bootstrap.php
│   ├── client.php
│   ├── config.php
│   ├── phpVsMsgPack.php
│   ├── prod
│   │   ├── IBuffer.php
│   │   ├── IPack.php
│   │   ├── MsgPack.php
│   │   ├── PackManager.php
│   │   ├── Package.php
│   │   ├── PhpPack.php
│   │   ├── SwooleBuffer.php
│   │   └── SwooleServer.php
│   ├── server.php
│   ├── tcp
│   │   ├── Data
│   │   │   └── _fields
│   │   │       ├── test.pre_error.php
│   │   │       └── test.pre_log.php
│   │   ├── client.php
│   │   ├── common~runtime.php
│   │   ├── index.php
│   │   └── result.txt
│   └── udp
│       ├── Data
│       │   └── _fields
│       │       └── test.pre_log.php
│       ├── client.php
│       ├── common~runtime.php
│       ├── index.php
│       ├── mysql.txt
│       └── receive-count.txt
├── lvs
│   └── dr.conf
├── mongo
│   ├── c1
│   │   └── log
│   │       └── mongo.log
│   ├── c2
│   │   └── log
│   │       └── mongo.log
│   ├── command
│   ├── mongos
│   │   └── log
│   │       └── mongo.log
│   ├── ps1-a
│   │   ├── log
│   │   │   └── mongo.log
│   │   ├── mongo.conf
│   │   └── pid
│   │       └── mongo.pid
│   ├── ps1-m
│   │   ├── log
│   │   │   └── mongo.log
│   │   ├── mongo.conf
│   │   └── pid
│   │       └── mongo.pid
│   ├── ps1-s
│   │   ├── log
│   │   │   └── mongo.log
│   │   ├── mongo.conf
│   │   └── pid
│   │       └── mongo.pid
│   ├── ps2-a
│   │   ├── log
│   │   │   └── mongo.log
│   │   ├── mongo.conf
│   │   └── pid
│   │       └── mongo.pid
│   ├── ps2-m
│   │   ├── log
│   │   │   └── mongo.log
│   │   ├── mongo.conf
│   │   └── pid
│   │       └── mongo.pid
│   └── ps2-s
│       ├── log
│       │   └── mongo.log
│       ├── mongo.conf
│       └── pid
│           └── mongo.pid
├── phpextension
│   ├── CREDITS
│   ├── EXPERIMENTAL
│   ├── LICENSE
│   ├── README.md
│   ├── config.m4
│   ├── config.w32
│   ├── democlass.c
│   ├── democlass.php
│   ├── democlass_define.php
│   ├── php_democlass.h
│   ├── test.php
│   └── tests
│       └── 001.phpt
├── phpmq
│   ├── IQueue.php
│   ├── MemcachedQueue.php
│   ├── MessageQueueProxy.php
│   ├── MobileMessageCustomer.php
│   ├── MobileMessagePublisher.php
│   ├── MongodbQueue.php
│   ├── QueueBench.php
│   ├── README.md
│   ├── RabbitQueue.php
│   ├── RedisQueue.php
│   ├── TmpTest.php
│   └── config.php
├── redis
│   ├── command
│   ├── server1-master
│   │   ├── redis-server.log
│   │   ├── redis-server.pid
│   │   └── redis.conf
│   ├── server1-slave
│   │   ├── redis-server.log
│   │   ├── redis-server.pid
│   │   └── redis.conf
│   ├── server2-master
│   │   ├── redis-server.log
│   │   ├── redis-server.pid
│   │   └── redis.conf
│   ├── server2-slave
│   │   ├── redis-server.log
│   │   ├── redis-server.pid
│   │   └── redis.conf
│   └── test.php
├── repo
│   └── repo
├── salt
│   ├── README.md
│   ├── salt1
│   │   └── Vagrantfile
│   ├── salt1_file_system
│   │   ├── README.md
│   │   ├── etc
│   │   │   ├── resolv.conf
│   │   │   ├── sysconfig
│   │   │   │   ├── network
│   │   │   │   └── network-scripts
│   │   │   │       └── ifcfg-eth0
│   │   │   └── yum.repos.d
│   │   │       └── saltstack.repo
│   │   └── srv
│   │       ├── pillar
│   │       │   ├── global.sls
│   │       │   ├── global.sls.example
│   │       │   └── top.sls
│   │       └── salt
│   │           ├── dev
│   │           │   ├── destroy.sls
│   │           │   ├── rpm
│   │           │   │   └── files
│   │           │   │       └── etc
│   │           │   │           ├── nginx
│   │           │   │           │   ├── conf.d
│   │           │   │           │   │   ├── app_common_website.conf
│   │           │   │           │   │   ├── app_manage.conf
│   │           │   │           │   │   ├── app_server.conf
│   │           │   │           │   │   └── phpmyadmin.conf
│   │           │   │           │   ├── fastcgi.conf
│   │           │   │           │   ├── fastcgi.conf.default
│   │           │   │           │   ├── fastcgi_params
│   │           │   │           │   ├── fastcgi_params.default
│   │           │   │           │   ├── koi-utf
│   │           │   │           │   ├── koi-win
│   │           │   │           │   ├── mime.types
│   │           │   │           │   ├── mime.types.default
│   │           │   │           │   ├── nginx.conf
│   │           │   │           │   ├── nginx.conf.default
│   │           │   │           │   ├── scgi_params
│   │           │   │           │   ├── scgi_params.default
│   │           │   │           │   ├── uwsgi_params
│   │           │   │           │   ├── uwsgi_params.default
│   │           │   │           │   └── win-utf
│   │           │   │           ├── php-fpm.conf
│   │           │   │           ├── php-fpm.d
│   │           │   │           │   └── www.conf
│   │           │   │           ├── php-zts.d
│   │           │   │           │   ├── bcmath.ini
│   │           │   │           │   ├── bz2.ini
│   │           │   │           │   ├── calendar.ini
│   │           │   │           │   ├── ctype.ini
│   │           │   │           │   ├── curl.ini
│   │           │   │           │   ├── dom.ini
│   │           │   │           │   ├── exif.ini
│   │           │   │           │   ├── fileinfo.ini
│   │           │   │           │   ├── ftp.ini
│   │           │   │           │   ├── gd.ini
│   │           │   │           │   ├── gettext.ini
│   │           │   │           │   ├── gmp.ini
│   │           │   │           │   ├── iconv.ini
│   │           │   │           │   ├── json.ini
│   │           │   │           │   ├── mbstring.ini
│   │           │   │           │   ├── mcrypt.ini
│   │           │   │           │   ├── mysqlnd.ini
│   │           │   │           │   ├── mysqlnd_mysqli.ini
│   │           │   │           │   ├── opcache-default.blacklist
│   │           │   │           │   ├── opcache.ini
│   │           │   │           │   ├── pdo.ini
│   │           │   │           │   ├── pdo_mysqlnd.ini
│   │           │   │           │   ├── pdo_pgsql.ini
│   │           │   │           │   ├── pdo_sqlite.ini
│   │           │   │           │   ├── pgsql.ini
│   │           │   │           │   ├── phar.ini
│   │           │   │           │   ├── posix.ini
│   │           │   │           │   ├── shmop.ini
│   │           │   │           │   ├── simplexml.ini
│   │           │   │           │   ├── sockets.ini
│   │           │   │           │   ├── sqlite3.ini
│   │           │   │           │   ├── sysvmsg.ini
│   │           │   │           │   ├── sysvsem.ini
│   │           │   │           │   ├── sysvshm.ini
│   │           │   │           │   ├── tokenizer.ini
│   │           │   │           │   ├── xml.ini
│   │           │   │           │   ├── xml_wddx.ini
│   │           │   │           │   ├── xmlreader.ini
│   │           │   │           │   ├── xmlwriter.ini
│   │           │   │           │   ├── xsl.ini
│   │           │   │           │   └── zip.ini
│   │           │   │           ├── php.d
│   │           │   │           │   ├── bcmath.ini
│   │           │   │           │   ├── bz2.ini
│   │           │   │           │   ├── calendar.ini
│   │           │   │           │   ├── ctype.ini
│   │           │   │           │   ├── curl.ini
│   │           │   │           │   ├── dom.ini
│   │           │   │           │   ├── exif.ini
│   │           │   │           │   ├── fileinfo.ini
│   │           │   │           │   ├── ftp.ini
│   │           │   │           │   ├── gd.ini
│   │           │   │           │   ├── gettext.ini
│   │           │   │           │   ├── gmp.ini
│   │           │   │           │   ├── iconv.ini
│   │           │   │           │   ├── json.ini
│   │           │   │           │   ├── mbstring.ini
│   │           │   │           │   ├── mcrypt.ini
│   │           │   │           │   ├── mysqlnd.ini
│   │           │   │           │   ├── mysqlnd_mysqli.ini
│   │           │   │           │   ├── opcache-default.blacklist
│   │           │   │           │   ├── opcache.ini
│   │           │   │           │   ├── pdo.ini
│   │           │   │           │   ├── pdo_mysqlnd.ini
│   │           │   │           │   ├── pdo_pgsql.ini
│   │           │   │           │   ├── pdo_sqlite.ini
│   │           │   │           │   ├── pgsql.ini
│   │           │   │           │   ├── phar.ini
│   │           │   │           │   ├── posix.ini
│   │           │   │           │   ├── shmop.ini
│   │           │   │           │   ├── simplexml.ini
│   │           │   │           │   ├── sockets.ini
│   │           │   │           │   ├── sqlite3.ini
│   │           │   │           │   ├── sysvmsg.ini
│   │           │   │           │   ├── sysvsem.ini
│   │           │   │           │   ├── sysvshm.ini
│   │           │   │           │   ├── tokenizer.ini
│   │           │   │           │   ├── xml.ini
│   │           │   │           │   ├── xml_wddx.ini
│   │           │   │           │   ├── xmlreader.ini
│   │           │   │           │   ├── xmlwriter.ini
│   │           │   │           │   ├── xsl.ini
│   │           │   │           │   └── zip.ini
│   │           │   │           └── php.ini
│   │           │   ├── setup.sls
│   │           │   ├── source
│   │           │   │   ├── files
│   │           │   │   │   └── etc
│   │           │   │   │       └── supervisor
│   │           │   │   │           ├── conf.d
│   │           │   │   │           │   └── fame-push.conf
│   │           │   │   │           └── supervisord.conf
│   │           │   │   └── scripts
│   │           │   │       ├── dev_install.sh
│   │           │   │       ├── self_signed_ssl.sh
│   │           │   │       ├── ssl.sh
│   │           │   │       └── ssl_generate.sh
│   │           │   └── start.sls
│   │           ├── install
│   │           │   ├── base.sls
│   │           │   ├── mongodb.sls
│   │           │   ├── mysql.sls
│   │           │   ├── nginx.sls
│   │           │   ├── nodejs.sls
│   │           │   ├── php7.sls
│   │           │   ├── phpmyadmin.sls
│   │           │   ├── postgresql.sls
│   │           │   ├── redis.sls
│   │           │   ├── rpm
│   │           │   │   ├── base.sls
│   │           │   │   ├── beanstalk.sls
│   │           │   │   ├── erlang.sls
│   │           │   │   ├── files
│   │           │   │   │   └── etc
│   │           │   │   │       ├── environment
│   │           │   │   │       ├── init.d
│   │           │   │   │       │   ├── rsyncd
│   │           │   │   │       │   └── supervisor
│   │           │   │   │       ├── pki
│   │           │   │   │       │   └── rpm-gpg
│   │           │   │   │       │       ├── RPM-GPG-KEY-PGDG-95
│   │           │   │   │       │       ├── RPM-GPG-KEY-mysql
│   │           │   │   │       │       └── RPM-GPG-KEY-webtatic-andy
│   │           │   │   │       └── rsyncd
│   │           │   │   │           ├── rsyncd.conf
│   │           │   │   │           ├── rsyncd.motd
│   │           │   │   │           └── rsyncd.secrets
│   │           │   │   ├── mongodb.sls
│   │           │   │   ├── mysql.sls
│   │           │   │   ├── nginx.sls
│   │           │   │   ├── nodejs.sls
│   │           │   │   ├── php7.sls
│   │           │   │   ├── postgresql.sls
│   │           │   │   ├── rsync.sls
│   │           │   │   └── scripts
│   │           │   │       └── git_install.sh
│   │           │   ├── rsync.sls
│   │           │   ├── source
│   │           │   │   ├── files
│   │           │   │   │   └── etc
│   │           │   │   │       ├── init.d
│   │           │   │   │       │   └── redis
│   │           │   │   │       └── redis.conf
│   │           │   │   ├── phpmyadmin.sls
│   │           │   │   ├── redis.sls
│   │           │   │   ├── scripts
│   │           │   │   │   ├── phpmyadmin_install.sh
│   │           │   │   │   ├── redis_install.sh
│   │           │   │   │   └── supervisor_install.sh
│   │           │   │   └── supervisor.sls
│   │           │   └── supervisor.sls
│   │           ├── top.sls
│   │           └── uninstall
│   │               ├── base.sls
│   │               ├── mongodb.sls
│   │               ├── mysql.sls
│   │               ├── nginx.sls
│   │               ├── nodejs.sls
│   │               ├── php7.sls
│   │               ├── phpmyadmin.sls
│   │               ├── postgresql.sls
│   │               ├── redis.sls
│   │               ├── rpm
│   │               │   ├── base.sls
│   │               │   ├── mongodb.sls
│   │               │   ├── mysql.sls
│   │               │   ├── nginx.sls
│   │               │   ├── nodejs.sls
│   │               │   ├── php7.sls
│   │               │   └── postgresql.sls
│   │               ├── source
│   │               │   ├── phpmyadmin.sls
│   │               │   ├── redis.sls
│   │               │   └── supervisor.sls
│   │               └── supervisor.sls
│   ├── salt2
│   │   └── Vagrantfile
│   └── salt3
│       └── Vagrantfile
├── serverwatcher
│   ├── LICENSE
│   ├── README.md
│   ├── config.sample.php
│   ├── disk.php
│   ├── request.php
│   ├── server.bak.php
│   ├── server.php
│   ├── server2.php
│   ├── swoole.lite.php
│   ├── watcher.php
│   └── web
│       ├── exporting.min.js
│       ├── highcharts.min.js
│       ├── index.html
│       ├── index.php
│       ├── jquery-2.1.4.js
│       ├── jquery-2.1.4.min.js
│       ├── jquery.cookie.min.js
│       ├── md5.min.js
│       ├── top.html
│       └── watcher.js
├── svn
│   └── conf
│       ├── authz
│       ├── hooks-env.tmpl
│       ├── passwd
│       └── svnserve.conf
└── twemproxy
    ├── command
    └── conf
        └── nutcracker.yml

384 directories, 2092 files
