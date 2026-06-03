"# MDGJX 项目结构总览

> 最后更新：2026-06-04
> 这是一个全栈在线工具箱项目，支持 Web、桌面（Electron）、浏览器扩展等多端部署。

---

## 一、技术栈

| 层级 | 技术 | 说明 |
|------|------|------|
| 后端核心 | **Go** | `core/`，高性能后端逻辑 |
| Web 服务器 | **TypeScript + Node.js** | `modules/web-server/`，Express 框架 |
| 前端 | **PureJS**（纯 JavaScript） | `resources/public/purejs/`，静态页面 |
| 桌面端 | **Electron** | `modules/desktop2/` |
| 浏览器扩展 | **未定** | `modules/browser2/` |
| 构建 | **SWC**（TypeScript 编译）、**SWC / Vite** | 前后端均使用 |
| 容器化 | **Docker** | `pipeline/build-docker.sh` |

---

## 二、根目录结构

| 目录/文件 | 用途 |
|-----------|------|
| `core/` | Go 核心后端代码 |
| `modules/` | 所有功能模块（Web 服务器、桌面端、浏览器扩展等） |
| `resources/` | 静态资源（前端页面、语言文件、公共资源） |
| `pipeline/` | 构建、部署、测试脚本 |
| `design/` | 设计文档 |
| `devtools/` | 开发工具配置 |
| `docs/` | 项目文档 |
| `addons/` | 附加组件/插件 |
| `my_features/` | **用户自定义功能目录**（新建的，不与原项目冲突） |
| `go.mod` / `go.sum` | Go 模块定义 |
| `package.json` | 根目录的 npm 配置（可能用于全局脚本） |
| `version.json` | 版本号定义 |
| `README.md` | 项目说明 |
| `SECURITY.md` | 安全规范 |
| `CODE_OF_CONDUCT.md` | 行为准则 |

---

## 三、核心模块详解

### 1. `modules/web-server/` —— Web 后端服务器（当前唯一可运行的部分）

- **入口**：`src/server.ts` → 创建 Express App，注册路由
- **框架**：Express + TypeScript
- **启动方式**（已测试通过）：
  ```bash
  cd modules/web-server
  npm install --legacy-peer-deps
  npx ts-node -r tsconfig-paths/register src/server.ts
  ```
- **端口**：主站 `3050`，扩展视图 `3060`
- **路由**：`src/routes/main.route.ts`、`extension.route.ts`、`devconfig.route.ts`
- **静态页面**：`modules/web-server/html/`，按功能分类为 `*-body.html` + `*-head.html`
- **关键依赖**：express, axios, compression, cors, helmet, winston, dayjs, jsonwebtoken 等
- **API 文档**：`swagger.yaml`

### 2. `modules/web/` —— Web 前端应用（可能使用 Vite/React？待确认）

### 3. `modules/desktop2/` —— Electron 桌面应用

- 构建脚本：`pipeline/build-desktop.sh`

### 4. `modules/browser2/` —— 浏览器扩展

### 5. `modules/bootstrap/` —— 启动器或移动端包装

### 6. `modules/meta/` —— 元数据/配置模块

### 7. `modules/legacy/` —— 旧版遗留代码

---

## 四、静态资源与前端页面

**`resources/public/purejs/`** 包含了前端国际化和分类数据：
- `app-i18n.json` —— 应用国际化文案
- `category.json` —— 工具分类
- `translation-lang.json` —— 翻译语言列表
- `lang/*.json` —— 各语言翻译文件（20+ 种语言）

**`modules/web-server/html/`** 存放所有工具的 HTML 片段（按功能分类的子目录）：
- `converter/`  —— 转换器（进制、颜色、日期、JSON/YAML/TOML、Base64 等）
- `cyber/`      —— 密码学工具（UUID、Token、BCrypt、RSA、加密解密）
- `devops/`     —— 开发运维（Crontab、Docker Compose、Git 备忘录、SQL 格式化）
- `formats/`    —— JSON 工具
- `i18n/`       —— 国际化工具
- `daimageshihua/`  —— 代码美化/格式化（JSON、XML、SQL、CSS、JS、HTML 等）
- `jiamiyujiemi/`   —— 加解密（3DES、A1Z26、AES 等）
- `changyongyasuogongju/` —— 压缩工具（gzip、bzip2、tar、zip、lz4、lzma、zlib 等）

---

## 五、多语言支持

前后端均支持国际化，语言文件位于：
- 后端：`resources/lang/*.json`
- 前端：`resources/public/purejs/lang/*.json`
- 支持语言：中文(简/繁)、英文、日文、韩文、法文、德文、西班牙文、俄文、葡萄牙文、意大利文、荷兰文、波兰文、瑞典文、挪威文、芬兰文、丹麦文、捷克文、匈牙利文、土耳其文、泰文、越南文、印尼文、马来文等 20+ 种

---

## 六、构建与部署流水线（`pipeline/`）

| 脚本 | 用途 |
|------|------|
| `build-all.sh` | 构建全部 |
| `build-core.sh` | 构建 Go 核心 |
| `build-desktop.sh` | 构建桌面端（Electron） |
| `build-docker.sh` | 构建 Docker 镜像 |
| `build-fe.sh` | 构建前端 |
| `deploy-web2.sh` | 部署 Web |
| `upload-web2.sh` | 上传 Web |
| `test-all.sh` | 运行全部测试 |
| `install-deps.sh` | 安装依赖 |

---

## 七、用户自定义功能（`my_features/`）

这是独立于原项目代码的目录，用于存放用户自己新增的功能，不影响原项目的 git 历史。

当前已有：

### `my_features/calculator/` —— 简易命令行计算器

- **语言**：Go
- **文件**：`calculator.go`、`go.mod`、`README.md`
- **功能**：支持四则运算（加、减、乘、除），通过命令行参数 `-a`、`-b`、`-op` 使用
- **注意**：需要安装 Go 才能运行

---

## 八、当前已知问题

1. **缺少环境配置**：启动 Web 服务器时提示 `not in dev env`，可能需要 `.env` 或 `config/` 中的配置文件
2. **依赖冲突**：`npm install` 需加 `--legacy-peer-deps`
3. **不完整的模块**：`modules/web/`、`modules/desktop2/` 等模块可能也需要额外的构建步骤
4. **Go 未安装**：当前环境无法运行 Go 代码

---

## 九、常用命令速查

```bash
# 启动前端 Web 服务器
cd modules/web-server
npx ts-node -r tsconfig-paths/register src/server.ts
# 访问 http://localhost:3050

# 安装 web-server 依赖
cd modules/web-server
npm install --legacy-peer-deps

# 构建 web-server（生产）
cd modules/web-server
npm run build
node dist/server.js
```

---

## 十、FAQ

**Q：修改前端页面后如何生效？**
A：`modules/web-server/html/` 下的 HTML 文件修改后，服务器热重载（nodemon 配置），不需要重启。

**Q：如何添加新的工具页面？**
A：在 `modules/web-server/html/` 下创建 `功能名称/功能名称-body.html` 和 `功能名称/功能名称-head.html`，然后在路由中注册或自动扫描。

**Q：如何添加新的语言？**
A：在 `resources/lang/` 和 `resources/public/purejs/lang/` 下添加对应的 `{lang}.json` 文件。
"