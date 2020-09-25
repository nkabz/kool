package cmd

import (
	"testing"
)

func TestFakeKoolService(t *testing.T) {
	f := &FakeKoolService{}

	_ = f.Execute([]string{"arg1", "arg2"})

	if !f.CalledExecute || len(f.ArgsExecute) != 2 || f.ArgsExecute[0] != "arg1" || f.ArgsExecute[1] != "arg2" {
		t.Errorf("failed to assert calling method Execute on FakeKoolService")
	}

	code := 100
	f.Exit(code)

	if !f.CalledExit || f.ExitCode != code {
		t.Errorf("failed to assert calling method Exit on FakeKoolService")
	}

	f.SetWriter(nil)

	if !f.CalledSetWriter {
		t.Errorf("failed to assert calling method SetWriter on FakeKoolService")
	}

	f.Error(nil)

	if !f.CalledError {
		t.Errorf("failed to assert calling method Error on FakeKoolService")
	}

	f.Warning()

	if !f.CalledWarning {
		t.Errorf("failed to assert calling method Warning on FakeKoolService")
	}

	f.Success()

	if !f.CalledSuccess {
		t.Errorf("failed to assert calling method Success on FakeKoolService")
	}
}