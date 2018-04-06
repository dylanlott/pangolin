package utils

import (
	"os"
	"log"
)

type Program struct {
	C chan int
}

func (p *Program) handleError(err error) {

}

func EnsureDirectory(path string, mode os.FileMode) error {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		os.MkdirAll(path, mode)
	}

	return err
}

func (p *Program) ErrCheck(err error) {
	if err != nil {
		log.Fatal("Persistence error: ", err)
		p.C <- 1
		close(p.C)
	}
}
