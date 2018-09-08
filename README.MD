# 基于Unix的七牛云小工具
- 上传文件(主要是图片)到七牛云,并返回该文件的外链
- 支持通过命令行传入文件名称
- 支持通过复制文件直接上传(需要支持 pbcopy, pbpaste命令)

## 配置
> $HOME/.qiniu_client/conf.josn

```josn
{
	"accessKey": ""
	"secretKey": ""
	"bucket": ""
	"rootPath": ""
	"prefix": ""
}
```

字段 | 说明|
--- | --- |
accessKey| 七牛云ak
secretKey | 七牛云sk
bucket| 存储空间名称
rootPath| 默认上传文件的路径
prefix| 七牛云默认外链 (类似: http://7xv4mv.com1.z0.glb.clouddn.com)

## 获取方式
```bash
go get github.com/xiaozefeng/qiniu-client

cd  $GOPATH/github.com/xiaozefeng/qiniu-client
go build -v -o qiniu-client . 
# 会生成一个叫qiniu-client的可执行文件
```

## 使用方式
```bash
# test_upload.png 是在配置里面的rootPath下的文件
./qiniu-client -path test_upload.png

# 或者直接复制rootPath下的文件，并执行命令
./qiniu-client
# 也可以将命令加入path即可直接执行命令了
# 上传成功后会将将返回markdown的url (类似于 ![](http://7xv4mv.com1.z0.glb.clouddn.com/qiniu_client20180908215948.png))

```


