## SVN 多项目管理笔记 ##

如果多个项目都在同一级目录下，那么理论上可以创建一个仓库，但是这样有一个弊端，
就是其中任何一个项目提交了代码，都会生成一个版本。这个对服务器性能造成的影响
也是不可忽略的。

所以，最好每一个项目建立一个仓库，但是可以使用同一个配置文件进行管理。

### windows 下快速创建SVN代码管理 ###

1. 下载tortoiseSVN。

2. 创建仓库总文件夹。(如: E:\projects-repos)

3. 创建每个项目的仓库。

```

	E:[enter]
	mkdir projects-repos
	c projects-repos
	mkdir conf // 全局的配置目录
	mkdir woaijia
	mkdir yinchao
	在两个文件夹上右键，选择tortoiseSVN > create repository here
	将woaijia/conf文件中的内容拷贝至E:\projects-repos\conf下
	删除E:\projects-repos\woaijia\conf
	删除E:\projects-repos\yinchao\conf
	配置
	E:\projects-repos\conf\svnserve.conf
	anon-access = none
	auth-access = write
	password-db = passwd
	authz-db = authz
	配置
	E:\projects-repos\conf\passwd
	[users]
	liuyong = 123456
	配置
	E:\projects-repos\conf\authz
	[woaijia:/]
	liuyong = rw
	* = 
	[yinchao:/]
	liuyong = rw
	* = 
	//至此配置项目完成
	
```
	
4. 创建window服务

```

	sc create svnserve binPath= "\"E:\Program Files\TortoiseSVN\bin\svnserve.exe\" --service --root E:\projects-repos --config-file E:\projects-repos\conf\svnserve.conf --listen-port 3690" displayname= "Subversion Repository" depend= Tcpip start= auto 
	(删除 sc delete svnserve)
	
```
	
5. 进入服务，启动svn服务

	net start svnserve
	
6. post-commit脚本使用

这个脚本是在有任何一个用户成功提交后，回调的脚本。脚本中可以是shell或者bat。

主要用作网站的同步跟新，参考SAE中SVN提交代码更新网站。

同样的实现可以备份项目到另外一个磁盘，以免自然灾害。

实现是使用SVN的update命令，具体暂时忽略。
