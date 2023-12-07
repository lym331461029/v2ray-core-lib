# 编译C库说明

## 共享库
```
go build -buildmode=c-shared -ldflags="-w -s" -o libv2ray.dll main_lib.go
```

## 静态库
```
go build -buildmode=c-archive -ldflags="-w -s" -o libv2ray.a main_lib.go
```