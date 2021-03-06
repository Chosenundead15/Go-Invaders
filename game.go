package main

import (
    "github.com/veandco/go-sdl2/sdl"
    "fmt"
    "time"
)

const winWidth, winHeight int = 800, 600

type color struct {
    r, g, b byte
}

type pos struct {
    x float32
    y float32
}

func clear(pixels []byte) {
    for i := range pixels {
        pixels[i] = 0
    }
}

func setPixel(x, y int, c color, pixels []byte) {
    index := (y * winWidth + x) * 4
    
    if index < len(pixels) - 4 && index >= 0 {
        pixels[index] = c.r
        pixels[index + 1] = c.g
        pixels[index + 2] = c.b
    }
}

func main() {
    err := sdl.Init(sdl.INIT_EVERYTHING)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    window, err := sdl.CreateWindow("Space Invaders", sdl.WINDOWPOS_UNDEFINED,
                                    sdl.WINDOWPOS_UNDEFINED, int32(winWidth),
                                    int32(winHeight), sdl.WINDOW_SHOWN)
    
    if err != nil {
        fmt.Println(err)
        return
    }
    defer window.Destroy()
    
    renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer renderer.Destroy()
    
    texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888,
                                           sdl.TEXTUREACCESS_STREAMING,
                                           int32(winWidth), int32(winHeight))
    
    if err != nil {
        fmt.Println(err)
        return
    }
    defer texture.Destroy()
    
    pixels := make([]byte, winWidth * winHeight * 4)
    
    player := &ship{
        entity: entity{pos{400, 500}, color{255, 255, 255}, 5, 5, 150, 0},
        bullet: bullet{pos{0, 0}, color{255, 255, 255}, 120, 20, 5, false},
    }
    alien := entity{pos{300.0, 300.0}, color{255, 255, 255}, 5, 5, 10, 1}
    
    keyState := sdl.GetKeyboardState()
    
    var frameStart time.Time
    var elapsedTime float32
    
    for {
        frameStart = time.Now()
        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
                case *sdl.QuitEvent:
                    return
            }
        }
        
        player.update(keyState, elapsedTime)
        if player.bullet.shoot == true {
            player.bullet.update(elapsedTime)
        }
        
        clear(pixels)
        player.draw(pixels)
        alien.draw(pixels)
        if player.bullet.shoot == true {
            player.bullet.draw(pixels)
        }
        
        texture.Update(nil, pixels, winWidth * 4)
        renderer.Copy(texture, nil, nil)
        renderer.Present()
        
        elapsedTime = float32(time.Since(frameStart).Seconds())
        if elapsedTime < 0.005 {
            sdl.Delay(5 - uint32(elapsedTime* 1000.0))
            elapsedTime = float32(time.Since(frameStart).Seconds())
        }
    }
}
