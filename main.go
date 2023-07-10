package main

type LightSocket interface {
	LightUp() string
}

type Room struct {
	LightOne LightSocket // 部屋にはソケットを付けておく
	LightTwo LightSocket
}

// Constructor Injection
func NewRoom(lightOne, lightTwo LightSocket) *Room {
	room := &Room{
		LightOne: lightOne,
		LightTwo: lightTwo,
	}
	return room
}

func main() {
	lightOne := new(light.Incandescent) // 注入するオブジェクト
	lightTwo := new(light.LedLight)     // 注入するオブジェクト

	myRoom := NewRoom(lightOne, lightTwo) // ここがDI

	myRoom.SwitchOnOne() // 1番照明: フィラメントが光るよ!
	myRoom.SwitchOnTwo() // 2番照明: LEDが光るよ!
}
