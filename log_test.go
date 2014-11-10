package logger

import (
	"os"
	"testing"
	"time"
)

func TestFresh(t *testing.T) {
	dir, _ := os.Getwd()

	Config("test", "log", dir, 3)

	Printf("num: %d ", 1)
	Printf("num: %d ", 2)
	Printf("num: %d ", 3)
	time.Sleep(time.Second * 1)
	Printf("num: %d ", 4)

	time.Sleep(time.Second * 3)
	Printf("num: %d ", 11)
	Printf("num: %d ", 12)
	Printf("num: %d ", 13)
	time.Sleep(time.Second * 1)
	Printf("num: %d ", 14)

	time.Sleep(time.Second * 2)
	Printf("num: %d ", 21)
	Printf("num: %d ", 22)
	Printf("num: %d ", 23)
	Printf("num: %d ", 24)
}
