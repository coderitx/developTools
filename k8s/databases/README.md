### 服务部署机器
- source-path: 

### 查看服务
```shell
kubectl get po -n dbs


kubectl get svc -n dbs

```


### 各服务队外开放端口

#### 连接服务： localhost:$port
#### mysql
```yaml
port: 31000
image: mysql:5.7

username: root
password: root123
```

#### pgsql
```yaml
port: 31001
image: postgres:13

username: postgres
password: postgres
```

#### oracle (TODO: 暂时无法连接)
```yaml
port:  31002
image: oracle/database:19.3.0-ee
```

#### sqlserver
```yaml
port: 31003
image: mcr.microsoft.com/mssql/server:2019-latest

username: sa
passwprd: Root123123

# sqlserver的pv卷目录需要权限要求
# path: /database/sqlserver
# 
```

#### redis
```yaml
port: 31004
image: redis:latest

password: 
```

#### mongo
```yaml
port: 31005
image: mongo:4.4


password:
```

