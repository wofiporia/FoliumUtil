# FoliumUtil

一个简洁的Go工具包，提供密码哈希、随机字符串生成和数据验证功能。

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
    
    "github.com/wofiporia/foliumutil/fpassword"
    "github.com/wofiporia/foliumutil/frandom"
    "github.com/wofiporia/foliumutil/fvalidator"
)

func main() {
    // 生成随机密码
    password := frandom.RandomString(8)
    
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
    
    fmt.Printf("密码: %s\n", password)
    fmt.Printf("哈希: %s\n", hashedPassword)
    fmt.Println("所有验证通过!")
}
```

## 依赖

- Go 1.21+
- github.com/stretchr/testify v1.11.1
- golang.org/x/crypto v0.45.0

## 许可证

MIT License

