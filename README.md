# wecomapi

企业微信服务端 API SDK（Go）。

fork from [xen0n/go-workwx](https://github.com/xen0n/go-workwx) 。

- 包名 `wecomapi`。
- 主要入口：`Client`（企业级）、`App`（应用级，绝大多数 API 方法挂在 `*App` 上）、`ClientError`（企业微信业务错误）。
- 覆盖通讯录、客户联系、微信客服、应用管理、消息收发、素材、OA（审批 / 日程 / 日历）、会话内容存档、群机器人 Webhook 等接口。
- 内置 access_token 管理、回调消息的签名校验与 AES 加解密。

## 特性

- 强类型模型，公开接口尽量小，刻意不暴露企业微信原始请求 / 响应类型。
- access_token 自动获取；支持后台刷新 goroutine、指数退避重试，以及外部 `ITokenProvider`。
- 支持覆盖 API Host（`WithQYAPIHost`）和自定义 `http.Client`（`WithHTTPClient`），便于测试和接入网关。
- 支持消息接收 / 事件回调，可直接作为 `http.Handler` 挂载。
- 不主动 `panic`，错误均通过返回值传递。
- 代码生成：从 `docs/*.md` 生成公开模型和底层 `execXxx` 调用，减少手写协议代码。

## 快速开始

本项目独立维护，模块路径以 `go.mod` 为准（`github.com/seelly/wecomapi`），Go 版本也以 `go.mod` 为准。

```go
import "github.com/seelly/wecomapi"

app := wecomapi.New(corpID).WithApp(corpSecret, agentID)

err := app.SendTextMessage(
    &wecomapi.Recipient{UserIDs: []string{"zhangsan"}},
    "hello",
    false,
)
```

需要集中管理 token 时，实现 `wecomapi.ITokenProvider` 并通过 `WithAccessTokenProvider` 注入；也可以调用 `app.SpawnAccessTokenRefresher()` 让 SDK 自行后台刷新。

接收回调消息：

```go
handler, err := wecomapi.NewHTTPHandler(token, encodingAESKey, myHandler)
if err != nil {
    // ...
}
http.Handle("/wecom/callback", handler)
```

其中 `myHandler` 实现 `wecomapi.RxMessageHandler`。

群机器人 Webhook：

```go
wc := wecomapi.NewWebhookClient(webhookKey)
err := wc.SendTextMessage("hello", &wecomapi.Mentions{UserIDs: []string{"zhangsan"}})
```

## 项目结构

| 路径 | 作用 |
|:---|:---|
| `docs/<domain>.md` | 接口 spec 源文件，代码生成的唯一输入 |
| `<domain>.md.go` | 由 `internal/sdkcodegen` 生成：公开模型 + 私有 `execXxx` 调用 |
| `models.go` | 手写的私有 `req*` / `resp*` 传输类型及 trait 实现 |
| `<domain>.go` | 手写的公开 `App` 方法，包装生成的 `execXxx` |
| `dummy_for_generate.go` | `go:generate` 注册表 |
| `internal/sdkcodegen` | Markdown spec → Go 代码生成器 |
| `internal/lowlevel` | 加密、签名、回调 envelope 等底层原语 |
| `errcodes/` | 全局错误码（由 `internal/errcodegen` 从企业微信文档生成） |

生成文件与手写文件的分工：

- `## Models` 中的模型会生成到 `<domain>.md.go`，是对调用者公开的领域模型。
- `execXxx` 也是生成代码，但它引用 `models.go` 中的私有 `reqXxx` / `respXxx`。
- 公开方法一律手写在 `<domain>.go`，负责参数整理、调用 `execXxx`、把传输层结果转成领域模型。

## 从企业微信 API 文档新增接口

以下流程以新增一个业务域 `<domain>` 为例。完整示例可参考 `docs/schedule.md`、`schedule.md.go`、`models.go`、`schedule.go` / `calendar.go`。

### 1. 阅读文档，确定契约

在 <https://developer.work.weixin.qq.com/document/> 找到目标接口，记录：

- HTTP 方法与 path，例如 `POST /cgi-bin/oa/schedule/add`；
- 是否需要 `access_token`；
- 请求 JSON：必填 / 可选字段、数组 / 嵌套对象、数值枚举；
- 响应 JSON：是否返回列表、是否有 `fail_result` 之类的部分失败信息；
- 是否返回二进制流（下载素材 / 文件等）。

### 2. 编写 spec

新增 `docs/<domain>.md`，基本结构如下：

```markdown
# <业务域>

## Models

### `Foo` Foo 信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`ID`|`id`|`string`| Foo ID
`Count`|`count,omitempty`|`*uint32`| 可选数量

## API calls

Name|Request type|Response type|URL|Access token|Doc
:---|:---|:---|:---|:---|:--
`execFooGet`|`reqFooGet`|`respFooGet`|`GET /cgi-bin/foo/get`|yes|获取 Foo
`execFooAdd`|`reqFooAdd`|`respFooAdd`|`POST /cgi-bin/foo/add`|yes|创建 Foo
```

生成器会严格校验格式，规则如下：

- 每个 H1 是一个 topic；H1 下只允许出现 `## Models` 和 `## API calls` 两个二级标题（可以只有其中一个，也可以都没有——后者用于只放 import 的 H1），出现其他 H2 会直接报错。
- `## Models` 下每个 H3 定义一个结构体：H3 标题里的反引号内容是类型名，其余文字作为 doc comment；每个 H3 最多一张字段表。
- 字段表表头必须是 `name` / `type` / `doc` / `json` / `xml` 的子集（大小写不敏感），列顺序任意。`Name`、`Type`、`JSON`、`XML` 单元格的内容要写在反引号里。
- 任意单元格内容为 `TODO` 的整行会被跳过，可用于暂时忽略尚未支持的字段。
- `## API calls` 表头为 `name` / `request type` / `response type` / `url` / `access token` / `doc`。
- `Name` 约定写成 `execXxx`，生成的方法即 `func (c *App) execXxx(...)`。
- `URL` 必须写成“方法 + 空格 + path”，方法支持 `GET`、`GET(binary)`、`POST`、`POST(media)`，大小写敏感。`GET(binary)` 生成的方法返回 `io.ReadCloser`，`Response type` 列不参与生成，可留空或填占位。
- `Access token` 填 `yes` / `y` / `+` 或 `no` / `n` / `-`。
- 一个 spec 可以包含多个 H1，生成到同一个 `.md.go`。例如 `docs/schedule.md` 同时描述日程和日历。
- `omitempty` 直接写在 JSON 列（如 `count,omitempty`），生成器会原样输出。
- 可选数值字段优先用指针（如 `*uint32`），以区分“未设置”和“0”。
- 如果生成代码需要额外 import 或手写辅助方法，可以在 spec 中放 ```` ```go ```` 代码块，生成器会原样拼接。注意输出顺序按 H1 进行，import 必须放在文件最前面一个只含 import 的 H1 中（例如 `# imports`），否则 gofmt 会报 `imports must appear before other declarations`。

### 3. 注册并运行生成器

在 `dummy_for_generate.go` 中增加一行：

```go
//go:generate go run --tags sdkcodegen ./internal/sdkcodegen ./docs/mydomain.md ./mydomain.md.go
```

然后在项目根目录执行单条生成命令：

```bash
go run --tags sdkcodegen ./internal/sdkcodegen ./docs/mydomain.md ./mydomain.md.go
```

注意：

- 生成器依赖 `github.com/russross/blackfriday/v2`，由本项目 `go.mod` 提供。
- `docs/` 已完整恢复。`go generate -run 'internal/sdkcodegen' ./...` 只重新生成 `.md.go`；不带 `-run` 的 `go generate ./...` 还会运行 `errcodegen`，需要网络并会更新 `errcodes/mod.go`。
- `docs/foo.md` 和 `docs/media_upload.md` 是随仓库保留的测试 / 参考 spec，未登记在 `dummy_for_generate.go` 中。
- `<domain>.md.go` 会被整体覆盖，不要直接手改。需要调整模型时改 spec 后重新生成。

### 4. 在 models.go 补请求 / 响应类型

生成器只生成 `## Models` 里的公开模型；`execXxx` 引用的 `reqXxx` / `respXxx` 需要在 `models.go` 手写，并按 HTTP 方法实现对应 trait。

GET 请求：

```go
type reqFooGet struct {
    ID string
}

var _ urlValuer = reqFooGet{}

func (x reqFooGet) intoURLValues() url.Values {
    return url.Values{
        "id": {x.ID},
    }
}
```

POST 请求：

```go
type reqFooAdd struct {
    Foo Foo `json:"foo"`
}

var _ bodyer = reqFooAdd{}

func (x reqFooAdd) intoBody() ([]byte, error) {
    return marshalIntoJSONBody(x)
}
```

响应：

```go
type respFooGet struct {
    respCommon

    Foo Foo `json:"foo"`
}
```

要点：

- GET / GET(binary)：req 必须实现 `intoURLValues() url.Values`。
- POST：req 必须实现 `intoBody() ([]byte, error)`，通常直接调用 `marshalIntoJSONBody`。
- POST(media)：req 必须实现 `getMedia() *Media`；如有 query 参数，再实现 `intoURLValues()`。
- resp 必须嵌入 `respCommon`（`GET(binary)` 没有 resp），否则不满足生成代码要求的 `tryIntoErr` 约束，编译会失败。
- 列表 / 嵌套对象优先复用 spec 中生成的公开模型，避免重复定义。
- 需要真正“缺省不发送”的字段：生成的公开模型在 spec 的 JSON 列写 `,omitempty`；手写的 `req*` 直接在 struct tag 里写 `json:"...,omitempty"`。仅用指针只能表达“未设置”，`nil` 仍会被编码成 `null`。

### 5. 在领域文件写公开方法

新建或修改 `<domain>.go`，调用生成的 `execXxx`，把传输层结果转成对调用者友好的 API：

```go
// GetFoo 获取 Foo
func (c *App) GetFoo(id string) (*Foo, error) {
    resp, err := c.execFooGet(reqFooGet{ID: id})
    if err != nil {
        return nil, err
    }
    return &resp.Foo, nil
}
```

约定：

- 公开方法接收者统一用 `*App`。
- 方法名使用业务语义（如 `GetSchedules`、`CreateCalendar`），不要暴露 `exec` / `req` / `resp` 命名。
- 不要把企业微信原始请求 / 响应类型暴露给调用者。
- 分页参数由调用者传入；必要的参数校验放在公开方法里（参考 `KfCustomerBatchGet`）。

### 6. 测试与验证

在项目根目录执行：

```bash
go test ./...
go vet ./...
```

建议至少覆盖：

- 请求序列化：字段名、必填字段、`omitempty` 是否生效（参考 `schedule_omitempty_test.go`）。
- 响应解析：嵌套模型、`respCommon.TryIntoErr` 的错误码路径。
- 公开方法的参数校验与返回值转换。

需要联调 HTTP 层时，可用 `WithQYAPIHost(httptestServer.URL)` + `WithHTTPClient` + 静态 `ITokenProvider` 构造 `App`。

### 检查清单

- [ ] spec 的 H1 / H2 / H3 / 表格格式正确，生成器无报错
- [ ] `dummy_for_generate.go` 已登记
- [ ] `req*` 实现了正确的 trait（`urlValuer` / `bodyer` / `mediaUploader`）
- [ ] `resp*` 嵌入了 `respCommon`
- [ ] 公开方法只返回本项目的类型
- [ ] 可选字段使用指针或 `omitempty`
- [ ] `go test ./...` 通过
- [ ] 下方“已支持接口”列表已同步

## 已支持接口

- [x] 通讯录管理（**部分支持**）
- [x] 客户联系（**大部分支持**）
- [x] 微信客服（**部分支持**）
- [x] 应用管理（**部分支持**）
- [x] 消息发送（**全部支持**）
- [x] 消息接收 / 回调
- [x] 素材管理（**部分支持**）
- [x] OA：审批 / 日程 / 日历（**部分支持**）
- [x] 会话内容存档（**大部分支持**）
- [x] 企业微信登录接口（code2Session）
- [x] 获取访问用户身份（code2UserInfo）
- [x] 群机器人 Webhook

<details>
<summary>通讯录管理 API</summary>

* [ ] 成员管理
    - [ ] 创建成员
    - [x] 读取成员 *NOTE: 成员对外信息暂未实现*
    - [x] 更新成员
    - [ ] 删除成员
    - [ ] 批量删除成员
    - [ ] 获取部门成员
    - [x] 获取部门成员详情
    - [x] userid 与 openid 互换
    - [ ] 二次验证
    - [ ] 邀请成员
    - [x] 获取加入企业二维码
    - [x] 手机号获取 userid
    - [x] 邮箱获取 userid
    - [ ] 获取成员 ID 列表
* [ ] 部门管理
    - [x] 创建部门
    - [ ] 更新部门
    - [ ] 删除部门
    - [x] 获取部门列表
    - [x] 获取子部门 ID 列表
* [ ] 标签管理
    - [ ] 创建标签
    - [ ] 更新标签名字
    - [ ] 删除标签
    - [ ] 获取标签成员
    - [ ] 增加标签成员
    - [ ] 删除标签成员
    - [ ] 获取标签列表
* [ ] 异步批量接口
    - [ ] 增量更新成员
    - [ ] 全量覆盖成员
    - [ ] 全量覆盖部门
    - [ ] 获取异步任务结果
* [ ] 通讯录回调通知
    - [ ] 成员变更通知
    - [ ] 部门变更通知
    - [ ] 标签变更通知
    - [ ] 异步任务完成通知

</details>

<details>
<summary>客户联系 API</summary>

* [x] 成员对外信息
* [x] 企业服务人员管理
    - [x] 获取配置了客户联系功能的成员列表
    - [x] 客户联系「联系我」管理
    - [x] 客户群「加入群聊」管理
* [x] 客户管理
    - [x] 获取客户列表
    - [x] 获取客户详情
    - [x] 批量获取客户详情
    - [x] 修改客户备注信息
* [x] 客户群管理
    - [x] 获取客户群列表
    - [x] 获取客户群详情
    - [x] 客户群 opengid 转换
* [x] 在职继承
    - [x] 分配在职成员的客户
    - [x] 查询客户接替状态
    - [ ] 分配在职成员的客户群
* [x] 离职继承
    - [ ] 获取待分配的离职成员列表
    - [x] 分配离职成员的客户
    - [x] 查询客户接替状态
    - [ ] 分配离职成员的客户群
* [x] 客户标签管理
    - [x] 管理企业标签
    - [x] 编辑客户企业标签
* [x] 客户分配
    - [x] 获取离职成员列表
    - [x] 分配在职或离职成员的客户
    - [x] 查询客户接替结果
    - [x] 分配离职成员的客户群
* [x] 变更回调通知
    - [x] 添加企业客户事件
    - [x] 编辑企业客户事件
    - [x] 外部联系人免验证添加成员事件
    - [x] 删除企业客户事件
    - [x] 删除跟进成员事件
    - [x] 客户接替失败事件
    - [x] 客户群变更事件
* [x] 消息推送
    - [x] 创建企业群发
    - [ ] 获取企业的全部群发记录
    - [x] 发送新客户欢迎语
    - [ ] 入群欢迎语素材管理
* [x] 获客助手
    - [x] 获取获客链接列表
    - [x] 获取获客链接详情
    - [x] 创建获客链接
    - [ ] 编辑获客链接
    - [x] 获取获客客户列表

</details>

<details>
<summary>微信客服 API</summary>

* [x] 客服账号管理
    - [x] 添加客服账号
    - [x] 删除客服账号
    - [x] 修改客服账号
    - [x] 获取客服账号列表
    - [x] 获取客服账号链接
* [x] 接待人员管理
    - [x] 添加接待人员
    - [x] 删除接待人员
    - [x] 获取接待人员列表
* [x] 会话分配与消息收发
    - [x] 分配客服会话
    - [x] 接收消息和事件
    - [x] 发送消息
    - [x] 发送欢迎语等事件响应消息
* [ ] 「升级服务」配置
* [x] 其他基础信息获取
    - [x] 获取客户基础信息
* [ ] 统计管理
    - [ ] 获取「客户数据统计」企业汇总数据
    - [ ] 获取「客户数据统计」接待人员明细数据
* [ ] 机器人管理
    - [ ] 知识库分组管理
    - [ ] 知识库问答管理

</details>

<details>
<summary>身份验证 API</summary>

* [x] 获取访问用户身份

</details>

<details>
<summary>应用管理 API</summary>

* [x] 获取应用
* [x] 获取应用列表
* [ ] 设置应用
* [ ] 自定义菜单
    - [ ] 创建菜单
    - [ ] 获取菜单
    - [ ] 删除菜单

</details>

<details>
<summary>消息发送 API</summary>

* [x] 发送应用消息
* [x] 接收消息
* [x] 发送消息到群聊会话
    - [x] 创建群聊会话
    - [x] 修改群聊会话
    - [x] 获取群聊会话
    - [x] 应用推送消息

### 消息类型

* [x] 文本消息
* [x] 图片消息
* [x] 语音消息
* [x] 视频消息
* [x] 文件消息
* [x] 文本卡片消息
* [x] 图文消息
* [x] 图文消息（mpnews）
* [x] markdown 消息
* [x] 任务卡片消息
* [x] 模板卡片消息

</details>

<details>
<summary>素材管理 API</summary>

* [x] 上传临时素材
* [x] 上传永久图片
* [x] 获取临时素材
* [ ] 获取高清语音素材

</details>

<details>
<summary>OA API</summary>

* [ ] 审批
    - [x] 获取审批模板详情
    - [x] 提交审批申请
    - [x] 审批申请状态变化回调通知
    - [x] 批量获取审批单号
    - [x] 获取审批申请详情
    - [x] 获取企业假期管理配置
    - [x] 获取成员假期余额
    - [x] 修改成员假期余额
* [x] 日程
    - [x] 创建日程
    - [x] 更新日程
    - [x] 取消日程
    - [x] 获取日程详情
    - [x] 新增日程参与者
    - [x] 删除日程参与者
    - [x] 获取日历下的日程列表
* [x] 日历
    - [x] 创建日历
    - [x] 更新日历
    - [x] 删除日历
    - [x] 获取日历详情

</details>

<details>
<summary>会话内容存档 API</summary>

* [x] 获取会话内容存档开启成员列表
* [x] 获取会话同意情况
* [ ] 客户同意进行聊天内容存档事件回调
* [x] 获取会话内容存档内部群信息

</details>

## 注意事项

### 关于保密消息发送

Markdown 等类型消息目前不支持作为保密消息发送，强行发送会报错。发送消息的方法仍然都带 `isSafe` 参数，原因是企业微信后续可能支持更多类型的保密消息，届时无需重新编译客户端；当前不支持的类型会由服务端返回错误。

### 关于生成代码

- `omitempty` 直接写在 spec 的 JSON 列即可，生成器会原样输出；`docs/schedule.md` 已按此方式修正，`schedule.md.go` 可以重复生成。
- `docs/` 已完整恢复。`go generate -run 'internal/sdkcodegen' ./...` 只重新生成 `.md.go`；不带 `-run` 的 `go generate ./...` 还会运行 `errcodegen`，需要网络并会更新 `errcodes/mod.go`。
- `docs/foo.md`、`docs/media_upload.md` 未登记到 `dummy_for_generate.go`（前者是生成器测试样例，后者依赖尚未支持的 import 生成），保留作参考。

## License

MIT，详见 [LICENSE](./LICENSE)。
