package main

import (
	"fmt"
	"log"
)

func main() {
	print("Hello, World!\n")
	println("It's been a while since I've written Go code.")

	fmt.Print("Actually, I've not studied coding in a while.\n")
	text := "I work as a software analist at Fibralink for now but i just copy others code and never get time to study."
	fmt.Printf("%s\n", text)
	text = "So i am creating that time and not wasting it anymore"
	fmt.Println(text)
	
	log.Println("But i just don't wanna do a standart Hello world. I wanna try all of the prints I remenber that are realy relevant, regarding what i've done to this moment")
	log.Println("I am oblied to mention the zerolog package, since it's a very beautiful and complete CLI log messages, but i will create a study day only for this one")
	text_2 := "ya"
	log.Fatalln("So see", text_2, "in the next study day, or when i got time to it kkk!")
}
