# go-mdi

一个 Go 语言库，提供类型安全的 Material Design Icons (MDI) 访问接口。本项目自动从 MDI 资源生成高性能的 Go 图标包。

## ✨ 功能特性

- 🎨 **自动代码生成** - 从 Material Design Icons 元数据自动生成 Go 代码
- 🔒 **类型安全** - 编译时检查的图标常量，避免运行时错误
- 📦 **完整元数据** - 包含图标名称、Unicode 码点、别名和 SVG 数据
- 🚀 **高性能** - 高效的图标查询和数据访问
- 📚 **丰富的图标库** - 包含数千个 Material Design Icons

## 📋 项目结构

```
go-mdi/
├── mdi/                          # 生成的图标包（核心）
│   ├── types.go                  # 图标类型定义
│   ├── funcs.go                  # 图标查询函数
│   ├── ab-testing.go             # 具体图标实现 (示例)
│   ├── abacus.go
│   ├── access-point.go
│   └── ...                        # 其他图标文件
├── internal/
│   ├── main.go                   # 代码生成器入口
│   ├── generate/                 # 生成逻辑
│   │   ├── meta.go               # 元数据加载
│   │   └── template.go           # 代码模板处理
│   └── assets/                   # 嵌入的 MDI 资源
│       └── mdi/                  # MDI 元数据和 SVG 文件
├── test/
│   ├── generate_test.go          # 生成器功能测试
│   └── mdi_test.go               # 图标包 API 测试
├── generator.go                  # 生成器主函数
├── go.mod
└── README.md
```

## 🚀 快速开始

### 使用图标常量

```go
package main

import (
    "fmt"
    "github.com/bbq-yaozi/go-mdi/mdi"
)

func main() {
    // 通过常量访问图标
    icon := mdi.AbTesting
    fmt.Println(icon.Name())  // 输出: "ab-testing"
    fmt.Println(icon.Data())  // 输出: SVG 数据 ([]byte)
}
```

### 通过名称查询图标

```go
icon := mdi.Name("ab-testing")
if icon != nil {
    data := icon.Data()
    // 使用图标数据...
}
```

### 遍历所有图标

```go
for _, icon := range mdi.Icons() {
    fmt.Println(icon.Name())
}
```

## 📖 API 文档

### Icon 接口

[`mdi`](mdi/) 包定义的每个图标实现以下接口：

```go
type Icon interface {
    Name() string   // 获取图标名称 (e.g., "ab-testing")
    Data() []byte   // 获取 SVG 二进制数据
}
```

### 核心函数

- **`mdi.Name(name string) Icon`** - 通过名称查询图标（[mdi/funcs.go](mdi/funcs.go)）
- **`mdi.Icons() []Icon`** - 获取所有可用图标（[mdi/funcs.go](mdi/funcs.go)）

### 图标常量

所有图标以 PascalCase 常量形式提供：

```go
mdi.AbTesting           // ab-testing
mdi.Abacus              // abacus
mdi.AbjadArabic         // abjad-arabic
mdi.AccessPointCheck    // access-point-check
// ... 数千个图标
```

## 🛠️ 开发

### 重新生成图标包

如果需要更新 MDI 图标数据，运行代码生成器：

```bash
go run ./generator.go
```

或从 internal 目录运行：

```bash
go run ./internal/main.go
```

这将重新生成 [`mdi/`](mdi/) 目录下的所有文件。

### 运行测试

```bash
# 运行所有测试
go test ./test/...

# 运行生成器测试
go test -v ./test/generate_test.go

# 运行图标包测试
go test -v ./test/mdi_test.go
```

#### 测试覆盖

- **[generate_test.go](test/generate_test.go)** - 验证元数据加载功能
  - `TestGenerator` - 测试 `LoadMetaList()` 是否正确加载所有图标元数据

- **[mdi_test.go](test/mdi_test.go)** - 验证图标包 API
  - `TestMdiConstant` - 测试图标常量访问、名称查询和数据获取

## 🏗️ 架构设计

### 代码生成流程

```
内置 MDI 资源 (SVG + 元数据)
    ↓
[LoadMetaList()] 加载元数据
    ↓
[LoadTemplate()] 加载代码模板
    ↓
生成 Go 代码
    ↓
输出至 mdi/ 目录
```

### 核心模块

- **[internal/generate/meta.go](internal/generate/meta.go)** 
  - 加载嵌入的 MDI 元数据
  - 解析 SVG 文件
  - 返回结构化的 `MetaInfo` 列表

- **[internal/generate/template.go](internal/generate/template.go)**
  - 加载并处理 Go 代码模板
  - 执行 PascalCase 名称转换
  - 生成类型安全的 Go 代码

- **[internal/main.go](internal/main.go)**
  - 协调元数据加载和代码生成
  - 执行完整的生成流程

## 💡 使用示例

### 示例 1：获取图标 SVG 数据

```go
package main

import (
    "fmt"
    "github.com/bbq-yaozi/go-mdi/mdi"
)

func main() {
    icon := mdi.AbTesting
    svgData := icon.Data()
    fmt.Printf("Icon: %s, Size: %d bytes\n", icon.Name(), len(svgData))
    // 输出: Icon: ab-testing, Size: XXX bytes
}
```

### 示例 2：动态查询图标

```go
func getIconByName(name string) []byte {
    icon := mdi.Name(name)
    if icon == nil {
        return nil
    }
    return icon.Data()
}

// 使用
data := getIconByName("access-point-check")
```

### 示例 3：图标统计

```go
allIcons := mdi.Icons()
fmt.Printf("Total icons: %d\n", len(allIcons))
```

## 📦 生成的文件说明

- **[mdi/types.go](mdi/types.go)** - 图标类型定义和接口
- **[mdi/funcs.go](mdi/funcs.go)** - 图标查询函数和名称映射表
- **[mdi/*.go](mdi/)** - 各个图标的具体实现（如 [ab-testing.go](mdi/ab-testing.go)）

## 🔐 安全性

图标数据通过常量定义，编译时即可验证：
- ✅ 避免拼写错误导致的运行时错误
- ✅ IDE 自动完成支持
- ✅ 编译时类型检查

## 📝 许可证

查看 [LICENSE](LICENSE) 文件。

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

---

**版本**: v1.0.0  
**基于**: [Material Design Icons](https://materialdesignicons.com/) 官方资源