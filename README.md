
###[git简要说明](http://rogerdudler.github.io/git-guide/index.zh.html)
###以及，[git文档](https://git-scm.com/book/zh/v2)

##0, install git, go, mysql, vim...
`sudo apt-get update`   
`sudo apt-get install git`   
~~`sudo apt-get install postgresql postgresql-contrib`~~      
`sudo apt-get install mysql-server mysql-client`    
`sudo apt-get install vim`   
###[install go](http://ask.xmodulo.com/install-go-language-linux.html)
######因为linux系统里修改配置文件多数用到vim，先简单了解下，当然如果熟练使用更好    
######[vim简单使用说明](http://www.jianshu.com/p/bcbe916f97e1)
######[开发工具](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/01.4.md)

##1, 新建个文件夹，然后把github项目clone到本地（目前就一个说明文件）
`git clone https://github.com/MollyBrown2016/draft.git`

##2, 添加上游远程仓库
`git remote add upstream https://github.com/MollyBrown2016/draft.git`

##3, 提交代码
`git add .`     
`git commit -m "提交说明"`    
`git push origin master`    
######输入用户名和密码

##[gitbhu常见操作错误](http://blog.csdn.net/dengjianqiang2011/article/details/9260435)


#架构
##后端: go
###[go web](https://www.gitbook.com/book/wizardforcel/build-web-application-with-golang/details)
###[go's httpd wiki](https://golang.org/doc/articles/wiki/)
###[go pitfalls](http://colobu.com/2015/09/07/gotchas-and-common-mistakes-in-go-golang/)
##前端: [riot](http://riotjs.com/zh/) + [pure](http://purecss.io/)
[项目说明](http://ju.outofmemory.cn/entry/233097)


