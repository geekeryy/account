## go-layout 业务框架
![Website](https://img.shields.io/website?url=https%3A%2F%2Fwww.jiangyang.me)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/comeonjy/go-layout)
![GitHub](https://img.shields.io/github/license/comeonjy/go-layout)
![GitHub issues](https://img.shields.io/github/issues/comeonjy/go-layout)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/comeonjy/go-layout)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/comeonjy/go-layout)
![GitHub pull requests](https://img.shields.io/github/issues-pr/comeonjy/go-layout)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/comeonjy/go-layout)
![GitHub last commit](https://img.shields.io/github/last-commit/comeonjy/go-layout)
![GitHub repo size](https://img.shields.io/github/repo-size/comeonjy/go-layout)
![GitHub language count](https://img.shields.io/github/languages/count/comeonjy/go-layout)
![GitHub top language](https://img.shields.io/github/languages/top/comeonjy/go-layout)

### TODO

### 安装
```shell
go install github.com/comeonjy/go-kit/cmd/kit

kit new demo-project
```

### 前置工作
```shell
go install github.com/google/wire/cmd/wire
go install golang.org/x/tools/cmd/stringer

go install \
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
google.golang.org/protobuf/cmd/protoc-gen-go \
google.golang.org/grpc/cmd/protoc-gen-go-grpc \
github.com/envoyproxy/protoc-gen-validate
```

### 版本管理（ [SemVer](https://semver.org/lang/zh-CN/) ）
版本格式：主版本号.次版本号.修订号
1. 主版本号：当你做了不兼容的API修改
2. 次版本号：当你做了向下兼容的功能性新增
3. 修订号：当你做了向下兼容的问题修正

### Git Flow 分支管理
- master：主分支（用于版本发布，始终与线上一致）
- dev：开发分支（用于开发，提测时，从dev检出release-1.0.0分支）
- release: 预发布（用于测试，测试中有问题直接修改，测试完成后合并入master和dev）
- feature-*：功能分支（用于功能开发，完成后合并到dev）
- hotfix-*：修复bug（从master分出来，完成后合并到master和dev）
  ![](http://assets.processon.com/chart_image/5f93a2e15653bb06ef13def8.png)

### Git提交规范
- feat：新功能（feature）
- fix：修补bug
- doc:：文档（documentation）
- style： 格式（不影响代码运行的变动）
- refactor：重构（即不是新增功能，也不是修改bug的代码变动）
- test：增加测试
- chore：构建过程或辅助工具的变动

### import 规范
```
import (
    // 标准库

    // 三方库

    // 本地库
)
```

### API规范
[Restful](http://kaelzhang81.github.io/2019/05/24/Restful-API%E8%AE%BE%E8%AE%A1%E6%9C%80%E4%BD%B3%E5%AE%9E%E8%B7%B5/)

- 返回值结构
  错误示例:
```json
{
  "code": 1001,
  "msg": "数据更新失败"
}
```
成功示例：
```json
{
  "xxx": "xxx"
}
```


### JetBrains OS licenses
go-layout是根据JetBrains s.r.o 授予的免费JetBrains开源许可证与GoLand一起开发的，因此在此我要表示感谢。
<a href="https://www.jetbrains.com/?from=go-layout" target="_blank"><img src="https://tva1.sinaimg.cn/large/0081Kckwgy1gkl0xz7y4uj30zz0u042c.jpg" width="50%"  /></a>

### License
© JiangYang, 2020~time.Now

Released under the Apache [License](https://account/blob/master/LICENSE)