# FoliumUtil

一个简洁的Go工具包，提供密码哈希、随机字符串生成、数据验证、Token管理和AI连接功能。

## 版本

当前版本: v1.0.0

## 安装

```bash
go get github.com/wofiporia/foliumutil@v1.0.0
```

## 功能模块

### fpassword - 密码处理

```go
import "github.com/wofiporia/foliumutil/fpassword"

// 哈希密码
hashedPassword, err := fpassword.HashPassword("mypassword123")
if err != nil {
    log.Fatal(err)
}

// 验证密码
err = fpassword.CheckPassword("mypassword123", hashedPassword)
if err != nil {
    log.Fatal("密码不匹配")
}
```

### frandom - 随机生成

```go
import "github.com/wofiporia/foliumutil/frandom"

// 生成随机字符串
randomStr := frandom.RandomString(10)
fmt.Println(randomStr) // 输出: k8x2m3n1q9

// 生成随机整数
randomInt := frandom.RandomInt(1, 100)
fmt.Println(randomInt) // 输出: 42

// 生成随机邮箱
email := frandom.RandomEmail()
fmt.Println(email) // 输出: x8m2k1@email.com
```

### fvalidator - 数据验证

```go
import "github.com/wofiporia/foliumutil/fvalidator"

// 验证用户名
err := fvalidator.ValidateUsername("user123")
if err != nil {
    fmt.Println("用户名无效:", err)
}

// 验证密码
err = fvalidator.ValidatePassword("MyPassword123")
if err != nil {
    fmt.Println("密码无效:", err)
}

// 验证邮箱
err = fvalidator.ValidateEmail("user@example.com")
if err != nil {
    fmt.Println("邮箱无效:", err)
}

// 验证字符串长度
err = fvalidator.ValidateString("hello", 3, 10)
if err != nil {
    fmt.Println("字符串长度无效:", err)
}
```

### ftoken - Token管理

```go
import "github.com/wofiporia/foliumutil/ftoken"

// 创建JWT Maker
jwtMaker, err := ftoken.NewJwtMaker("your-secret-key-here")
if err != nil {
    log.Fatal(err)
}

// 或创建PASETO Maker
pasetoMaker, err := ftoken.NewPasetoMaker("your-32-byte-secret-key-here")
if err != nil {
    log.Fatal(err)
}

// 创建Token
username := "testuser"
role := "user"
duration := time.Hour

token, payload, err := jwtMaker.CreateToken(username, role, duration)
if err != nil {
    log.Fatal(err)
}

// 验证Token
payload, err = jwtMaker.VerifyToken(token)
if err != nil {
    log.Fatal("Token验证失败:", err)
}

fmt.Printf("用户: %s, 角色: %s\n", payload.Username, payload.Role)
```

### faiconn - AI连接器

```go
import "github.com/wofiporia/foliumutil/faiconn"

// 创建OpenAI兼容API连接器
conn, err := faiconn.NewCustomConn(faiconn.AIConfig{
    Model:    "deepseek-ai/DeepSeek-V3.1-Terminus",
    URL:      "https://api.siliconflow.cn/v1",
    APIKey:   "sk-your-api-key-here",
    Provider: "siliconflow",
})
if err != nil {
    log.Fatal(err)
}

// 发送消息
response, err := conn.SendMessage("你好")
if err != nil {
    log.Fatal(err)
}
fmt.Println(response)

// 获取连接信息
fmt.Printf("模型: %s\n", conn.GetModel())
fmt.Printf("提供商: %s\n", conn.GetProvider())

// 关闭连接
defer conn.Close()
```

## 验证规则

### 用户名 (ValidateUsername)
- 4-16个字符
- 以字母开头
- 只能包含字母、数字和下划线

### 密码 (ValidatePassword)
- 8-16个字符
- 至少一个小写字母
- 至少一个大写字母
- 至少一个数字
- 只能包含字母、数字和特殊字符 @$!%*?&

### 邮箱 (ValidateEmail)
- 使用标准邮箱格式验证
- 长度3-320个字符

## 使用示例

```go
package main

import (
    "fmt"
    "log"
    "time"
    
    "github.com/wofiporia/foliumutil/fpassword"
    "github.com/wofiporia/foliumutil/frandom"
    "github.com/wofiporia/foliumutil/fvalidator"
    "github.com/wofiporia/foliumutil/ftoken"
    "github.com/wofiporia/foliumutil/faiconn"
)

func main() {
    
    password := "F0liumUtil"
    
    // 哈希密码
    hashedPassword, err := fpassword.HashPassword(password)
    if err != nil {
        log.Fatal(err)
    }
    
    // 验证密码强度
    err = fvalidator.ValidatePassword(password)
    if err != nil {
        fmt.Println("密码强度不足:", err)
        return
    }
    
    // 验证哈希密码
    err = fpassword.CheckPassword(password, hashedPassword)
    if err != nil {
        fmt.Println("密码验证失败")
        return
    }
    
    // 创建JWT Token
    jwtMaker, err := ftoken.NewJwtMaker(frandom.RandomString(32))
    if err != nil {
        log.Fatal(err)
    }
    
    jwtToken, jwtPayload, err := jwtMaker.CreateToken("testuser", "admin", time.Hour)
    if err != nil {
        log.Fatal(err)
    }
    
    // 创建PASETO Token
    pasetoMaker, err := ftoken.NewPasetoMaker(frandom.RandomString(32))
    if err != nil {
        log.Fatal(err)
    }
    
    pasetoToken, pasetoPayload, err := pasetoMaker.CreateToken("testuser", "user", time.Hour)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("密码: %s\n", password)
    fmt.Printf("哈希: %s\n", hashedPassword)
    fmt.Printf("JWT Token: %s\n", jwtToken)
    fmt.Printf("JWT Token用户: %s (角色: %s)\n", jwtPayload.Username, jwtPayload.Role)
    fmt.Printf("PASETO Token: %s\n", pasetoToken)
    fmt.Printf("PASETO Token用户: %s (角色: %s)\n", pasetoPayload.Username, pasetoPayload.Role)
    
    // AI连接器示例
    aiConn, err := faiconn.NewCustomConn(faiconn.AIConfig{
        Model:    "deepseek-ai/DeepSeek-V3.1-Terminus",
        URL:      "https://api.siliconflow.cn/v1",
        APIKey:   "sk-your-api-key-here",
        Provider: "siliconflow",
    })
    if err != nil {
        log.Fatal(err)
    }
    
    aiResponse, err := aiConn.SendMessage("你好")
    if err != nil {
        fmt.Println("AI调用失败:", err)
    } else {
        fmt.Printf("AI回复: %s\n", aiResponse)
    }
    
    fmt.Println("所有验证通过!")
}
```

## 依赖

- Go 1.21+
- github.com/stretchr/testify v1.11.1
- golang.org/x/crypto v0.45.0
- github.com/dgrijalva/jwt-go v3.2.0+incompatible

## Token模块说明

### 支持的Token类型

1. **JWT (JSON Web Token)**
   - 使用HMAC-SHA256签名算法
   - 广泛支持，易于调试
   - 适合Web应用和API

2. **PASETO (Platform-Agnostic Security Tokens)**
   - 更现代、更安全的Token格式
   - 防止常见攻击（如长度泄露攻击）
   - 适合高安全性要求的应用

### 安全注意事项

- JWT密钥建议至少32字符
- PASETO密钥必须是32字节长度
- 密钥应该从环境变量或配置文件读取，不要硬编码
- Token过期时间根据业务需求设置，建议不要过长

## AiConn模块说明

### 支持的api类型

1. **OpenAI 兼容 API**
 

## 作者留言

FoliumUtil主要是用来记录我做web项目写过和用过的utils

感谢你使用 FoliumUtil！这是一个为简化日常开发而创建的工具包。如果你觉得这个包对你有帮助，欢迎给个Star ⭐。

如果你发现了bug或有改进建议，欢迎提交Issue或Pull Request。

希望这个小小的工具包能让你的开发更高效！

## 许可证

MIT License

