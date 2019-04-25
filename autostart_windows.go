package autostart

// #cgo LDFLAGS: -lole32 -luuid
/*
#define WIN32_LEAN_AND_MEAN
#include <stdint.h>
#include <windows.h>

uint64_t CreateShortcut(char *shortcutA, char *path, char *args);
*/
import "C"

var startupDir string

func init() {
	
}

func (a *App) path()  {

}

func (a *App) IsEnabled()  {

}

func (a *App) Enable()  {

}

func (a *App) Disable()  {
}
