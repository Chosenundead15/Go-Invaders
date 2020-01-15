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

type bullet struct {
    pos
    color
    speed float32
    h, w float32
    shoot bool
}

type ship struct {
    entity
    bullet bullet
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

func (ship *ship) update(keyState []uint8, elapsedTime float32) {
    if keyState[sdl.SCANCODE_LEFT] != 0 {
        ship.x -= ship.speed * elapsedTime
    }
    if keyState[sdl.SCANCODE_RIGHT] != 0 {
        ship.x += ship.speed * elapsedTime
    }
    if keyState[sdl.SCANCODE_SPACE] != 0 && ship.bullet.shoot == false {
        //the 14 is needed because the method to draw doesn't center the ship
        ship.bullet.x = ship.x + 14
        ship.bullet.y = ship.y
        ship.bullet.shoot = true
    }
}

func (bullet *bullet) update(elapsedTime float32) {
    if bullet.y - bullet.h / 2 > 0 {
        bullet.y -= bullet.speed * elapsedTime
    } else {
        bullet.shoot = false
    }
}

func (bullet *bullet) draw(pixels []byte) {
    startx := bullet.x - bullet.w/2
    starty := bullet.y - bullet.h/2
    
    for y := 0; y < int(bullet.h); y++ {
        for x := 0; x < int(bullet.w); x++ {
            setPixel(int(startx) + x, int(starty) + y, bullet.color, pixels)
        }
    }
}

