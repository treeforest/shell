# Shell

Shell是一个用于简化执行Shell命令的软件包。

## 示例

```go
import "github.com/treeforest/shell"
// ...

// 执行命令并获取输出
out, _ := shell.Command("ls").Output()

// ...

// 执行命令并等待其完成
err := shell.Command("./start.sh").Run()

// ...

// 简化的命令执行方式
err := shell.Run("./start.sh")

// ...

// 使用用户名和密码创建Shell对象，并以sudo方式执行命令
sh := shell.New(username, password)
sh.Sudo("./start.sh").Run()

// 使用Shell对象执行命令
sh.Command("./start.sh").Run()
```

