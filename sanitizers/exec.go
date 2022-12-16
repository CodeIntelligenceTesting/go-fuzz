package sanitizers

import (
	"os"
	"os/exec"

	"github.com/CodeIntelligenceTesting/gofuzz/sanitizers/detectors"
)

func CmdCombinedOutput(hookId int, cmd *exec.Cmd) ([]byte, error) {
	detectors.NewCommandInjection(hookId, cmd).Detect().Report()
	return cmd.CombinedOutput()
}

func CmdOutput(hookId int, cmd *exec.Cmd) ([]byte, error) {
	detectors.NewCommandInjection(hookId, cmd).Detect().Report()
	return cmd.Output()
}

func CmdRun(hookId int, cmd *exec.Cmd) error {
	detectors.NewCommandInjection(hookId, cmd).Detect().Report()
	return cmd.Run()
}

func CmdStart(hookId int, cmd *exec.Cmd) error {
	detectors.NewCommandInjection(hookId, cmd).Detect().Report()
	return cmd.Start()
}

func OsStartProcess(hookId int, name string, argv []string, attr *os.ProcAttr) (*os.Process, error) {
	detectors.NewCommandInjection(hookId, name).Detect().Report()
	return os.StartProcess(name, argv, attr)
}
