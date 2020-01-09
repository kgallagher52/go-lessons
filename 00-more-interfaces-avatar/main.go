package main

import (
	"fmt"
	"log"
)

// Avatar ...
type Avatar interface {
	bender() string
}

// Fire ...
type Fire string

// Water ...
type Water string

// Air ...
type Air string

// Earth ...
type Earth string

// Methods
func (f Fire) bender() string {
	return fmt.Sprintf("%v fire", f)
}
func (w Water) bender() string {
	return fmt.Sprintf("%v water", w)
}
func (a Air) bender() string {
	return fmt.Sprintf("%v air", a)
}
func (e Earth) bender() string {
	return fmt.Sprintf("%v earth", e)
}

// WriteLog ...
func WriteLog(a Avatar) {
	log.Println(a.bender())
}

func main() {
	f := Fire("Roko")
	w := Water("Kuruk")
	e := Earth("Wan")
	a := Air("Aang")

	WriteLog(f)
	WriteLog(w)
	WriteLog(e)
	WriteLog(a)

}
