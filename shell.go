package shell

import (
	"fmt"
	"os/exec"
	"syscall"
)

// Command returns the Cmd struct to execute the named program with
// the given arguments.
func Command(arg string, opts ...Option) *exec.Cmd {
	conf := &shellConfig{}
	for _, o := range opts {
		o(conf)
	}

	cmd := exec.Command("/bin/sh", "-c", arg)

	if conf.Dir != "" {
		cmd.Dir = conf.Dir
	}

	if conf.Setpgid {
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Setpgid: true,
		}
	}

	if conf.Log != nil {
		if conf.Dir != "" {
			conf.Log("$ cd %s && %s", cmd.Dir, arg)
		} else {
			conf.Log("$ %s", arg)
		}
	}

	return cmd
}

// Run 执行普通命令
// 参数：
//   - arg: 要执行的命令参数
//   - opts: 可选的命令选项
func Run(arg string, opts ...Option) error {
	out, err := Command(arg, opts...).CombinedOutput()
	if err != nil {
		if len(out) != 0 {
			return fmt.Errorf("%w\n%s", err, string(out))
		}
		return err
	}
	if len(out) != 0 {
		conf := &shellConfig{}
		for _, o := range opts {
			o(conf)
		}
		if conf.Log != nil {
			conf.Log("%s", string(out))
		}
	}
	return nil
}

// Shell 对shell命令行对象的封装
type Shell struct {
	// Username 用户名
	Username string

	// Password 密码
	Password string
}

func New(username, password string) *Shell {
	return &Shell{Username: username, Password: password}
}

// Command 执行普通命令。
// 参数：
//   - arg: 要执行的命令参数
//   - opts: 可选的命令选项
// 返回值：
//   - *exec.Cmd: 执行普通命令的 *exec.Cmd 对象
func (c *Shell) Command(arg string, opts ...Option) *exec.Cmd {
	return Command(arg, opts...)
}

// Sudo 执行 sudo 命令。如果当前用户是普通用户，则自动输入密码。
// 参数：
//   - arg: 要执行的命令参数
//   - opts: 可选的命令选项
// 返回值：
//   - *exec.Cmd: 执行 sudo 命令的 *exec.Cmd 对象
func (c *Shell) Sudo(arg string, opts ...Option) *exec.Cmd {
	if c.Username == "root" || c.Password == "" {
		// 对于 root 用户或者没有配置密码的情况，直接执行命令
		return c.Command(arg, opts...)
	}
	// 对于普通用户，自动输入密码
	return c.Command(fmt.Sprintf("echo '%s' | sudo -S sh -c '%s'", c.Password, arg), opts...)
}
