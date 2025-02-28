# 常见用法

##### 从IP开始扫描

```shell
# 当IP Ping通后扫描Top1000端口
./dddd -t 192.168.0.1
# 当IP Ping通后全端口扫描
./dddd -t 172.16.100.1 -p 1-65535
# 指定IP禁Ping全端口扫描指定端口
./dddd -t 172.16.100.1 -p 80,53,1433-5000 -Pn
# 调用Masscan进行全端口SYN扫描(需安装Masscan)
./dddd -t 192.168.0.0/16 -p 1-65535 -Pn -st syn
```

##### 从IP段开始扫描

```shell
./dddd -t 172.16.100.0/24
./dddd -t 192.168.0.0-192.168.0.12
```

##### 从端口开始扫描

```shell
./dddd -t 192.168.44.12:80
```

##### 从Web开始扫描

```shell
./dddd -t http://test.com
```

##### 从域名开始扫描

探测(子)域名是否使用了CDN，若解析出真实IP则进入IP扫描

```shell
# 不枚举test.com子域名，只识别test.com
./dddd -t test.com
# 枚举子域名，识别所有获取到的子域名是否为CDN资产
./dddd -t test.com -sd
```

##### 从Hunter开始扫描

先配置./config/subfinder-config.yaml中的Hunter API。

```yaml
hunter: ["ffxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"]
```

配置完成后打开命令行

```shell
# 从Hunter中获取备案机构为 带带弟弟 的目标进入扫描 默认最大1000条
./dddd -t '"icp.name="带带弟弟"' -hunter
# 最大查询1页，一页100个。避免积分过度消耗
./dddd -t '"icp.name="带带弟弟"' -hunter -htpc 1
```

攻防演练中通过Hunter导入企业备案资产方便快速占坑。

##### 从Fofa开始扫描

先配置./config/subfinder-config.yaml中的FOFA 邮箱和KEY。

```yaml
fofa: ["xxxx@qq.com:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"]
```

配置完成后打开命令行

```shell
./dddd -t "domain=\"baidu.com\"" -fofa (从fofa取100个baidu.com域名的目标)
./dddd -t "domain=\"baidu.com\"" -fofa -ffmc 10000 (指定最大数量为10000 默认100)
```

##### 从Quake开始扫描

先配置./config/subfinder-config.yaml中的Quake api。

配置完成后打开命令行

```shell
./dddd -t 'ip:"127.0.0.1"' -quake
```

##### 多目标扫描

在target.txt中写入你的目标，如

```
172.16.100.0/24
192.168.0.0-192.168.255.255
http://test.com
aaa.test.com
10.12.14.88:9999
```

然后在命令行中敲下如下命令。dddd会自动识别资产类型并送入对应流程。

```
./dddd -t target.txt
```

当然dddd也支持多Hunter、Fofa语句，如在target.txt写下

```
ip="111.111.111.222"
domain="bbbb.cc"
icp.name="带带弟弟"
```

然后在命令行中敲下如下命令。dddd会批量从Hunter查询资产并送入对应流程。

```
./dddd -t target.txt -hunter
```

##### Hunter低感知模式

配置好hunter api。

```
./dddd -t 'ip="xxx.xxx.xxx.xxx"' -lpm
```

##### 禁用漏洞探测

```
./dddd -t 127.0.0.1 -npoc
```

##### 禁用漏洞探测且禁用主动指纹识别

```
./dddd -t 127.0.0.1 -npoc -nd
```

##### 禁用ICMP，禁用TCP探测存活

```
./dddd -t 127.0.0.1 -tcpp -nicmp
```

##### icmp探测后，不存活进行tcp探活

```
./dddd -t 127.0.0.1 -tcpp
```



# 详细参数

```shell
coco@Mac % ./dddd -h

     _       _       _       _   
  __| |   __| |   __| |   __| |  
 / _` |  / _ `|  / _` |  / _` |  
 \__,_|  \__,_|  \__,_|  \__,_|  
_|"""""|_|"""""|_|"""""|_|"""""| 
"`-0-0-'"`-0-0-'"`-0-0-`"`-0-0-'
dddd.version: 1.5.1

Usage of ./dddd:
  -Pn
    	禁用主机发现功能(icmp,tcp)
  -ffmc int
    	Fofa 查询资产条数 Max:10000 (default 100)
  -fofa
    	从Fofa中获取资产,开启此选项后-t参数变更为需要在fofa中搜索的关键词
  -gopt int
    	GoPoc运行线程 (default 50)
  -htpc int
    	Hunter 最大查询页数 (default 10)
  -htps int
    	Hunter 每页资产条数 (default 100)
  -hunter
    	从hunter中获取资产,开启此选项后-t参数变更为需要在hunter中搜索的关键词
  -ld
    	允许域名解析到局域网
  -lpm
    	低感知模式 (当前只支持Hunter,且默认选择Hunter)
  -mp string
    	指定masscan路径 (default "masscan")
  -nd
    	关闭主动指纹探测
  -ngp
    	关闭Golang Poc探测
  -nicmp
    	当启用主机发现功能时，禁用ICMP主机发现功能
  -npoc
    	关闭漏洞探测
  -nsbf
    	关闭子域名爆破
  -nsf
    	关闭被动子域名枚举
  -o string
    	html格式输出报告
  -p string
    	目标IP扫描的端口。 默认扫描Top1000
  -pc int
    	一个IP的端口数量阈值,当一个端口的IP数量超过此数量，此IP将会被抛弃 (default 300)
  -poc string
    	模糊匹配Poc名称
  -proxy string
    	HTTP代理，在外网可利用云函数/代理池的多出口特性恶心防守 例: http://127.0.0.1:8080
  -psto int
    	TCP扫描超时时间(秒) (default 6)
  -qkmc int
    	Quake 查询资产条数 (default 100)
  -quake
    	从Quake中获取资产,开启此选项后-t参数变更为需要在quake中搜索的关键词
  -sbft int
    	爆破子域名协程数量 (default 150)
  -sd
    	开启子域名枚举
  -st string
    	端口扫描方式 tcp使用TCP扫描(慢),syn为调用masscan进行扫描(需要masscan依赖) (default "tcp")
  -synt int
    	SYN扫描线程(masscan) (default 10000)
  -t string
    	被扫描的目标。 192.168.0.1 192.168.0.0/16 192.168.0.1:80 baidu.com:80 target.txt
  -tc int
    	TCP全连接获取Banner的线程数量 (default 30)
  -tcpp
    	当启用主机发现功能时，启用TCP主机发现功能
  -tcpt int
    	TCP扫描线程 (default 600)
  -wt int
    	Web探针线程,根据网络环境调整 (default 100)
  -wto int
    	Web探针超时时间,根据网络环境调整 (default 12)
```



# 指纹/Poc拓展

### 指纹

指纹数据库存于 ./config/finger.yaml

支持的指纹基础规则如下

```shell
header="123" //返回头中包含123
header!="123" //返回头中不包含123
header~="xxx" //返回头满足xxx正则
body="123" //body中包含123
body!="123" //body中不包含123
body~="xxx" //body满足xxx正则
body=="123" //body为123
server="Sundray" //返回头server字段中包含Sundray
server!="Sundray" //返回头server字段中不包含Sundray
server=="Sundray" //server字段为Sundray
title="123" //标题包含123
title!="123" //标题不包含123
title=="123" //标题为123
title~="xxx" //标题满足xxx正则
cert="123" //证书中包含123
cert!="123" //证书中不包含123
cert~="xxx" //证书满足正则
port="80" //服务端口为80
port!="80" //服务端口不为80
port>="80" //服务端口大于等于80
port<="80" //服务端口小于等于80
protocol="mysql" //协议为mysql
protocol!="mysql" //协议不为mysql
path="123/123.html" //爬虫结果中包含 123/123.html
body_hash="619335048" //响应体mmh3 hash为619335048
icon_hash="619335048"  //icon mmh3 hash
status="200" //页面返回码为200
status!="200" //页面返回码不为200
content_type="text/html" //content_type包含text/html
content_type!="text/html" //content_type不包含text/html
banner="123" // TCP banner 包含123
banner!="123" // TCP banner中不含123
```

各类规则支持与(&&)或(||)非(!)任意组合。可使用括号。与fofa搜索语法类似。

![image-20230817180845323](assets/image-20230817180845323.png)

这里拿Fortinet-sslvpn举例。

需要Web响应体中包含fgt_lang，编写规则body="fgt_lang"

需要Web响应体中包含/sslvpn/portal.html，编写规则body="/sslvpn/portal.html"

需要同时满足这两个条件才会被判定为Fortinet-sslvpn的资产，将两个规则使用与(&&)连接就得到了这条指纹。



### API

若有被动枚举子域名、请求fofa、hunter等需求。请在./config/subfinder-config.yaml中配置API。



### 子域名字典

子域名字典位于./config/subdomains.txt



### 服务爆破字典

字典路径为./config/dict/

每行以"空格:空格"分割账号密码

```
root : 123456
root : admin
root : admin123
```

其中shirokeys.txt为shiro key字典





### 主动指纹

当一个应用不暴露在根路径时就需要主动去访问，这个时候需要配置主动指纹数据库。

路径在./config/dir.yaml

![image-20230817182401928](assets/image-20230817182401928.png)

配置也很简单，照着上边写就行。

这里拿Alibaba-Nacos举例子。当访问到http://host:port/nacos/，且访问后识别到Alibaba-Nacos指纹后就被判断有效。



### 漏洞Poc编写

编写参考nuclei poc编写

https://nuclei.projectdiscovery.io/templating-guide/



Poc存放路径

./config/pocs



将poc写好后放入./config/pocs即可识别。



### 工作流

仅仅编写Poc并放入./config/pocs是不会正常运行Poc的。需要为此Poc配置指定的指纹，本工具只有在匹配到目标指纹后才会调用这个Poc。



工作流存放目录

./config/workflow.yaml



![image-20230817183151312](assets/image-20230817183151312.png)

以solr为例。

当dddd识别到目标指纹为APACHE-Solr时，就会在./config/pocs中寻找路径以

cves/2017/CVE-2017-12629.yaml

....

apache-solr-log4j-rce.yaml

结尾的poc调用。



如果我写好一个名为solr-rce.yaml的nuclei poc，则应该在workflow.yaml的对应指纹的pocs下添加一行solr-rce.yaml或者solr-rce。这样才能在识别到solr时调用到此poc。



在新版dddd中，支持使用nuclei tags作为指定。

在workflow中填写 "Tags@" 开头的poc名称，则代表匹配所有带此tag的nuclei poc

```yaml
nginx:
  type:
    - root
  pocs:
    - Tags@nginx
```

上述workflow的意思是匹配所有带nginx tags的poc。



type是拿来干什么的呢？

比如这里有一个nacos，他的路径是http://host:port/aaa/bbb/nacos/a.js

如果有root标签，他会在 http://host:port/处打一次poc

如果有dir标签，在

http://host:port/aaa/

http://host:port/aaa/bbb/ (命中)

http://host:port/aaa/bbb/nacos/

下各打一次poc

如果有base标签，在http://host:port/aaa/bbb/nacos/a.js下打一次poc

一般情况下，root就能用，若应用部署在二级目录下，但poc是一级目录的poc，就需要dir参数才能探测到二级目录下的洞。





# 支持的Golang Poc列表

FTP 暴力破解
MSSQL 暴力破解
MYSQL 暴力破解
ORACLE 暴力破解
POSTGRESQL 暴力破解
RDP 暴力破解
REDIS 暴力破解/未授权访问
SMB 暴力破解
SSH 暴力破解
TELNET 暴力破解
Shiro反序列化 Key枚举
MONGODB 暴力破解
MEMCACHED 未授权访问
MS17-010
Java调试接口远程命令执行





# 漏洞报表展示

直观的漏洞报表

请求/响应包留存

![image-20230817175041903](assets/image-20230817175041903.png)

![image-20230817175243236](assets/image-20230817175243236.png)

数据库基础信息留存，方便漏洞验证/截图

![image-20230817175453532](assets/image-20230817175453532.png)

SMB共享目录信息留存

![image-20230817175546241](assets/image-20230817175546241.png)

