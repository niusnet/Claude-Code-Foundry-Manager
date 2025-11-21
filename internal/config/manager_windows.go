//go:build windows

package config

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows/registry"
)

const (
	// Registry path for system environment variables
	envRegPath = `SYSTEM\CurrentControlSet\Control\Session Manager\Environment`

	// Windows API constants for broadcasting environment changes
	HWND_BROADCAST   = 0xFFFF
	WM_SETTINGCHANGE = 0x001A
	SMTO_ABORTIFHUNG = 0x0002
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	procSendMessage  = user32.NewProc("SendMessageTimeoutW")
)

// getEnvVar reads an environment variable from the Windows registry
func getEnvVar(key string) (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, envRegPath, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	value, _, err := k.GetStringValue(key)
	if err != nil {
		if err == registry.ErrNotExist {
			return "", nil // Variable doesn't exist, return empty string
		}
		return "", err
	}

	return value, nil
}

// setEnvVar writes an environment variable to the Windows registry
func setEnvVar(key, value string) error {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, envRegPath, registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("failed to open registry key (requires admin privileges): %w", err)
	}
	defer k.Close()

	err = k.SetStringValue(key, value)
	if err != nil {
		return fmt.Errorf("failed to set registry value: %w", err)
	}

	return nil
}

// deleteEnvVar removes an environment variable from the Windows registry
func deleteEnvVar(key string) error {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, envRegPath, registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("failed to open registry key (requires admin privileges): %w", err)
	}
	defer k.Close()

	err = k.DeleteValue(key)
	if err != nil {
		if err == registry.ErrNotExist {
			return nil // Variable doesn't exist, consider it success
		}
		return fmt.Errorf("failed to delete registry value: %w", err)
	}

	return nil
}

// notifyEnvironmentChange broadcasts a message to all windows that environment has changed
func notifyEnvironmentChange() error {
	env, err := syscall.UTF16PtrFromString("Environment")
	if err != nil {
		return err
	}

	var result uintptr
	ret, _, err := procSendMessage.Call(
		uintptr(HWND_BROADCAST),
		uintptr(WM_SETTINGCHANGE),
		0,
		uintptr(unsafe.Pointer(env)),
		uintptr(SMTO_ABORTIFHUNG),
		5000, // 5 second timeout
		uintptr(unsafe.Pointer(&result)),
	)

	if ret == 0 {
		return fmt.Errorf("SendMessageTimeout failed: %v", err)
	}

	return nil
}
