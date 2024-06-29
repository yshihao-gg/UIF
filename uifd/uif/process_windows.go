//go:build windows
// +build windows

package uif

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

func ProcessSet(p *exec.Cmd) {
	p.SysProcAttr = &windows.SysProcAttr{
		CreationFlags: windows.CREATE_UNICODE_ENVIRONMENT | windows.CREATE_NEW_PROCESS_GROUP,
		HideWindow:    true,
	}
}

func SetAdmin() error {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		return err
	}
	return nil
}

func IsAdmin() bool {
	fd, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}
	defer fd.Close()
	return true
}

func terminateProc(pid int, _ os.Signal) error {
	dll, err := windows.LoadDLL("kernel32.dll")
	if err != nil {
		return err
	}
	defer dll.Release()

	f, err := dll.FindProc("AttachConsole")
	if err != nil {
		return err
	}
	r1, _, err := f.Call(uintptr(pid))
	if r1 == 0 && err != syscall.ERROR_ACCESS_DENIED {
		return err
	}

	f, err = dll.FindProc("SetConsoleCtrlHandler")
	if err != nil {
		return err
	}
	r1, _, err = f.Call(0, 1)
	if r1 == 0 {
		return err
	}
	f, err = dll.FindProc("GenerateConsoleCtrlEvent")
	if err != nil {
		return err
	}
	r1, _, err = f.Call(windows.CTRL_BREAK_EVENT, uintptr(pid))
	if r1 == 0 {
		return err
	}
	return nil
}

func ProcessClose(p *exec.Cmd) error {
	return terminateProc(p.Process.Pid, syscall.SIGINT)
}

func Method1(enable bool) error {
	current, err := user.Current()
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\uif.bat", current.HomeDir)
	// exe, _ := filepath.Abs(os.Args[0])
	exe := GetUIFPath()
	content := fmt.Sprintf("@echo off\n\rchcp 65001 > nul\n\rstart \"\" \"%s\"\n\r", exe)
	return WriteOrRemove(enable, path, content)
}

func AutoStartup(enable bool) error {
	return Method1(enable)
	// return SetWindowsReg(GetUIFPath(), enable)  // service can not popup Web
}

func SetWindowsReg(path string, enable bool) error {
	// Get the current user's registry key
	key, err := registry.OpenKey(registry.CURRENT_USER, "Software\\Microsoft\\Windows\\CurrentVersion\\Run", registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer key.Close()

	keyName := "uif_service"

	if enable {
		err = key.SetStringValue(keyName, path)
	} else {
		err = key.DeleteValue(keyName)
	}
	if err != nil {
		fmt.Println("Failed to set registry value:", err)
		return err
	}
	return nil
}
