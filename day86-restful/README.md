### server
```
#mDNS方式注册
go1.14.15 run server.go
go1.14.15 run rest.go

#etcd方式注册
MICRO_REGISTRY=etcd MICRO_REGISTRY_ADDRESS=127.0.0.1:2379 go1.14.15 run server.go
MICRO_REGISTRY=etcd MICRO_REGISTRY_ADDRESS=127.0.0.1:2379 go1.14.15 run rest.go
```


### client
```
curl localhost:56439/student/jack
```
