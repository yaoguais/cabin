## GAdmin

Make admin project more easily.



## Install

```
$ go get github.com/Yaoguais/gadmin
$ GADMIN="/path/to/github.com/Yaoguais/gadmin"
$ mysql -uroot -p < $GADMIN/install.sql
$ mkdir project && cd project && cp -a $GADMIN/app app && cp -a $GADMIN/core/views app/views/core && cp -a $GADMIN/public public && cp $GADMIN/Makefile .
$ cp $GADMIN/config/config.toml.example config.toml && vim config.toml
$ echo 'package main

import (
    "github.com/Yaoguais/gadmin/core"
    coreRoutes "github.com/Yaoguais/gadmin/core/routes"
)

func main() {

    core.SetConfig(core.GetAppRoot() + "/config.toml")
    core.InitApp()
    coreRoutes.RegisterRoutes()
    core.RunApp()
}
' > main.go
$ echo "username:admin password:111111 for login http://127.0.0.1:1323/"
$ go get ./... && make serve
```



## Feature

- [x] Data Table defined by go struct
- [x] Access Control with manage
- [x] Dashboard & Form & Icon & Button
- [x] Ajax Search Select
- [x] Uploader with process



## Screenshot

![Dashboard](https://github.com/Yaoguais/cabin/raw/master/goforadmin/screenshot/dashboard.png)

![Privilege](https://github.com/Yaoguais/cabin/raw/master/goforadmin/screenshot/privilege.png)

![Component](https://github.com/Yaoguais/cabin/raw/master/goforadmin/screenshot/component.png)
