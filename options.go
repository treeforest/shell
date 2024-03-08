package shell

// Logger shell日志输出接口
type Logger func(format string, v ...any)

type shellConfig struct {
	Dir     string // 命令行工作目录
	Setpgid bool   // 设置新的进程组ID
	Log     Logger // shell日志输出接口
}

type Option func(c *shellConfig)

func WithDir(dir string) Option {
	return func(c *shellConfig) {
		c.Dir = dir
	}
}

func WithSetpgid() Option {
	return func(c *shellConfig) {
		c.Setpgid = true
	}
}

func WithLogFunc(log Logger) Option {
	return func(c *shellConfig) {
		c.Log = log
	}
}
