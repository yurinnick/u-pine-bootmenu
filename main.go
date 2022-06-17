package main

/*
#cgo LDFLAGS: -L. -lbootmenu
#cgo CFLAGS: -I.
#include "lvgl/lvgl.h"
#include "bootmenu.h"

extern void menu_event_handler(char *text);
*/
import "C"
import (
	"fmt"
	"time"
)

func init() {
	C.init_menu()
}

func main() {
	C.build_menu()
	C.test_callback()

	for true {
		C.lv_timer_handler()
		time.Sleep(5000)
	}
}

//export menu_event_handler
func menu_event_handler(text *C.char) {
	fmt.Println(C.GoString(text))
}
