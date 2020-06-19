package main

import (
	"image"
	"sync"
)

var mu sync.RWMutex //保护icons
var icons map[string]image.Image

// 并发安全
// func Icon(name string) image.Image {
// 	mu.RLock()
// 	if icons != nil {
// 		icon := icons[name]
// 		mu.RUnlock()
// 		return icon
// 	}
// 	mu.RUnlock()
// 	//获取互斥锁
// 	mu.Lock()
// 	if icons == nil {
// 		loadIcons()
// 	}
// 	icon := icons[name]
// 	mu.Unlock()
// 	return icon
// }
var loadIconsOnce sync.Once

func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

func loadIcons() {
	icons["speed.png"] = loadIcon("speed.png")
}
