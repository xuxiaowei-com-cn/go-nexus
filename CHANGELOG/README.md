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
    - 新增 `GetComponents` 获取组件详细信息
    - 新增 `ListComponents` 列出组件
    - 新增 `DeleteComponents` 删除组件
    - 新增 `ListRepository`
      [列出存储库](https://help.sonatype.com/repomanager3/integrations/rest-and-integration-api/repositories-api)
    - 新增 `GetMavenRepository` 获取 Maven 存储库
    - 新增 `PostExtDirectRead` 搜索资产
    - 新增 `PostExtDirectReadAsset` 查看资产详细信息
    - 新增 `GetBrowseRepository` 浏览仓库

## v0.0.1

- ⭐ New Features | 新功能
    - 新增 `GetSwagger` 获取 Swagger 配置
