# 更新日志

## [语义化版本](https://semver.org/lang/zh-CN/)

版本格式：主版本号.次版本号.修订号，版本号递增规则如下：

1. 主版本号：当你做了不兼容的 API 修改，
2. 次版本号：当你做了向下兼容的功能性新增，
3. 修订号：当你做了向下兼容的问题修正。

先行版本号及版本编译信息可以加到“主版本号.次版本号.修订号”的后面，作为延伸。

# [更新日志](#更新日志)

- 采用倒叙编写
- 符号含义
    - 📗 Links | 链接
    - ⭐ New Features | 新功能
    - 🐞 Bug Fixes | 漏洞修补
    - 📔 Documentation | 文档
    - 🔨 Dependency Upgrades | 依赖项升级
    - ❤ Contributors | 贡献者
    - ⚠️ Noteworthy Changes | 值得注意的变化

## v0.1.0

- ⭐ New Features | 新功能
    - 新增 `ListAssets` 列出资产
    - 新增 `GetAssets` 获取资产详细信息
    - 新增 `DeleteAssets` 删除资产
    - 新增 `GetBlobStoresQuotaStatus` 获取给定blob存储区的配额状态
    - 新增 `ListBlobStores` 列出blob存储区
    - 新增 `GetComponents` 获取组件详细信息
    - 新增 `ListComponents` 列出组件
    - 新增 `DeleteComponents` 删除组件
    - 新增 `ListRepository`
      [列出存储库](https://help.sonatype.com/repomanager3/integrations/rest-and-integration-api/repositories-api)
    - 新增 `GetMavenGroupRepository` 获取 Maven 分组 存储库
    - 新增 `GetMavenHostedRepository` 获取 Maven 宿主 存储库
    - 新增 `GetMavenProxyRepository` 获取 Maven 代理 存储库
    - 新增 `ListUsers` 检索用户列表。请注意，如果 `source` 不是 `default`，则响应限制为100个用户。
    - 新增 `GetStatusCheck` 返回系统状态检查结果的运行状况检查终结点
    - 新增 `GetStatus` 验证服务器是否可以响应读取请求的运行状况检查终结点
    - 新增 `GetStatusWritable` 验证服务器是否可以响应读写请求的运行状况检查终结点 
    - 新增 `PostExtDirectRead` 搜索资产
    - 新增 `PostExtDirectReadAsset` 查看资产详细信息
    - 新增 `GetBrowseRepository` 浏览仓库

## v0.0.1

- ⭐ New Features | 新功能
    - 新增 `GetSwagger` 获取 Swagger 配置
