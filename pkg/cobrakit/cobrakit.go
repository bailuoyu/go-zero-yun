package cobrakit

import (
	"github.com/spf13/cobra"
)

// RunCmd 直接执行命令函数
func RunCmd(cmd *cobra.Command, args []string) (err error) {
	for p := cmd; p != nil; p = p.Parent() {
		if p.PersistentPreRunE != nil {
			if err := p.PersistentPreRunE(cmd, args); err != nil {
				return err
			}
			break
		} else if p.PersistentPreRun != nil {
			p.PersistentPreRun(cmd, args)
			break
		}
	}
	if cmd.PreRunE != nil {
		if err := cmd.PreRunE(cmd, args); err != nil {
			return err
		}
	} else if cmd.PreRun != nil {
		cmd.PreRun(cmd, args)
	}
	if cmd.RunE != nil {
		if err := cmd.RunE(cmd, args); err != nil {
			return err
		}
	} else {
		cmd.Run(cmd, args)
	}
	if cmd.PostRunE != nil {
		if err := cmd.PostRunE(cmd, args); err != nil {
			return err
		}
	} else if cmd.PostRun != nil {
		cmd.PostRun(cmd, args)
	}
	for p := cmd; p != nil; p = p.Parent() {
		if p.PersistentPostRunE != nil {
			if err := p.PersistentPostRunE(cmd, args); err != nil {
				return err
			}
			break
		} else if p.PersistentPostRun != nil {
			p.PersistentPostRun(cmd, args)
			break
		}
	}
	return nil
}
