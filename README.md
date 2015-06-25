########################整体架构########################
martini
	go get github.com/go-martini/martini

	官方文档：
	https://github.com/go-martini/martini


xorm
	go get github.com/go-xorm/xorm
	
	官方文档：
	http://xorm.io/docs/

gzip
	go get github.com/martini-contrib/gzip
	通过giz方式压缩请求信息的处理器

render
	go get github.com/martini-contrib/render
	渲染JSON和HTML模板的处理器

sessions
	go get github.com/martini-contrib/sessions
	提供Session服务支持的处理器

mysql
	go get github.com/go-sql-driver/mysql
	mysql驱动

goconfig
	go get github.com/Unknwon/goconfig
	获取配置文件，解析ini文件
########################项目结构########################

conf			配置信息
	base.go加载配置文件

controllers		控制层

models			数据层
	base.go配置数据库

routers			路由层
	routers.go所有路由

static			静态文件
	css			样式文件
	img     	静态图片
	fonts 		字体
	js			js文件
	
views			前端模板文件