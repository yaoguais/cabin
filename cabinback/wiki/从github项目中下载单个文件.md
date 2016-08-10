## 下载github项目中的单个文件 ##

[教程来自](http://www.zhihu.com/question/25369412/answer/30579415)

用 SVN 即可.

举例说明:
譬如这个项目: Mooophy/Cpp-Primer · GitHub, 我只想看 ch03 文件夹的代码怎么办?
先打开 ch03, 其 URL 为: "https://github. com/Mooophy/Cpp-Primer/tree/master/ch03"(这里添加空格, 为了防止知乎智能识别)

将 /tree/master/ 换成 /trunk/ . (这个以前玩 Google Code 的人应该很熟悉.)
"https://github. com/Mooophy/Cpp-Primer/trunk/ch03"(同样有空格)

然后, 输入:

svn checkout https://github.com/Mooophy/Cpp-Primer/trunk/ch03


PS: 第一次使用的话, 可能会出现下面这个提示:

R)eject, accept (t)emporarily or accept (p)ermanently?

输入 P 就行了.

玩的愉快! 