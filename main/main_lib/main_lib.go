package main

import "C"
import (
	"fmt"
	"sync"
	"unsafe"

	core "github.com/v2fly/v2ray-core/v5"
	_ "github.com/v2fly/v2ray-core/v5/main/distro/all"
)

var gServerContainer = make(map[int32]core.Server)
var locker = sync.Mutex{}
var gServerID int32 = 10000

//export v2ray_new
func v2ray_new(ptr unsafe.Pointer, size C.int, handle *C.int) C.int {
	configData := C.GoBytes(ptr, size)
	server, err := startV2Ray(configData)
	if err != nil {
		fmt.Printf("failed to create v2ray instance: %v\n", err.Error())
		return C.int(-1)
	}
	locker.Lock()
	defer locker.Unlock()
	gServerID += 1
	gServerContainer[gServerID] = server
	*handle = C.int(gServerID)
	return C.int(0)
}

//export v2ray_start
func v2ray_start(handle C.int) C.int {
	locker.Lock()
	defer locker.Unlock()
	server, ok := gServerContainer[int32(handle)]
	if !ok {
		fmt.Printf("failed to find server with handle: %d\n", handle)
		return C.int(-1)
	}
	err := server.Start()
	if err != nil {
		fmt.Printf("failed to start v2ray: %v\n", err.Error())
		return C.int(-1)
	}
	return C.int(0)
}

//export v2ray_close
func v2ray_close(handle C.int) C.int {
	locker.Lock()
	defer locker.Unlock()
	server, ok := gServerContainer[int32(handle)]
	if !ok {
		fmt.Printf("failed to find server with handle: %d\n", handle)
		return C.int(-1)
	}
	err := server.Close()
	if err != nil {
		fmt.Printf("failed to close v2ray: %v\n", err.Error())
		return C.int(-1)
	}
	delete(gServerContainer, int32(handle))
	return C.int(0)
}

func startV2Ray(configData []byte) (core.Server, error) {
	config, err := core.LoadConfig(core.FormatAuto, configData)
	if err != nil {
		fmt.Printf("failed to load config: %v\n", err.Error())
		return nil, err
	}

	server, err := core.New(config)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func main() {
	fmt.Println("Hello, World!")
}
