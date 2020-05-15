# M2P

#### 概述

M2P是一个将mysql table转化为grpc service的小工具

命令执行：

    m2s --mysql user:password@tcp\(host:port\)/database\?charset=utf8 --table tableName

#### 结构介绍

- core 执行代码目录
- templates 模版目录，如果需要修改proto生成内容，只需要修改模版即可