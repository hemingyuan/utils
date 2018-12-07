package utils

import (
	"fmt"
	"os/exec"
)

// ExecShell 调用Shell执行系统命令
func ExecShell(dir string, command string, args ...string) (err error) {
	var (
		cmd *exec.Cmd
		res []byte
	)
	cmd = exec.Command(command, args...)
	cmd.Dir = dir
	cmd.Env = append(cmd.Env, "LANG=")

	if res, err = cmd.CombinedOutput(); err != nil {
		err = fmt.Errorf("\n%s\nExecute Command [%s] Error. Err: %v", res, command, err)
		return
	}
	fmt.Printf("\n%s\nExecute Command [%s] Successfully.\n", res, command)
	return
}
