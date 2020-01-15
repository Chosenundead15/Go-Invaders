package main

import (
    "github.com/veandco/go-sdl2/sdl"
    //"fmt"
    //"time"
)

var sprite = [][]byte {
    {
        0,0,0,0,1,0,0,0,0,
        0,0,0,1,1,1,0,0,0,
        0,0,0,1,1,1,0,0,0,
        0,1,1,1,1,1,1,1,0,
        1,1,1,1,1,1,1,1,1,
        1,1,1,1,1,1,1,1,1,
    },
    {
        0,0,1,0,0,0,0,0,1,0,0,
        0,0,0,1,0,0,0,1,0,0,0,
        0,0,1,1,1,1,1,1,1,0,0,
        0,1,1,0,1,1,1,0,1,1,0,
        1,1,1,1,1,1,1,1,1,1,1,
        1,0,1,1,1,1,1,1,1,0,1,
        1,0,1,0,0,0,0,0,1,0,1,
        0,0,0,1,1,0,1,1,0,0,0,
    },
}

type entity struct {
    pos
    color
    h float32
    w float32
    speed float32
    tex byte
}

func (entity *entity) draw(pixels []byte) {
    var multw, multh float32
    if entity.tex == 0 {
        multw, multh = 9, 6
    } else {
        multw, multh = 11, 8
    }
    
    startx := entity.x - (entity.w * multw) / 5
    starty := entity.y - (entity.h * multh) / 2
    
    for i,v := range sprite[entity.tex] {
        if v == 1 {
            for y := int(starty); y < int(starty + entity.h); y++ {
                for x := int(startx); x < int(startx + entity.w); x++ {
                    setPixel(x, y, entity.color, pixels)
                }
            }
        }
        
        startx += entity.w
        if (i + 1) % int(multw) == 0 {
            starty += entity.h
            startx -= entity.w * multw
        }
    }
}

func (entity *entity) control(keyState []uint8) {
    if keyState[sdl.SCANCODE_LEFT] != 0 {
        entity.x -= entity.speed
    }
    if keyState[sdl.SCANCODE_RIGHT] != 0 {
        entity.x += entity.speed
    }
}

