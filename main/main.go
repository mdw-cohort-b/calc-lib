package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/mdw-smarty-cohort-b-2024-10/calc-lib"
)

func main() {
	handler := NewHandler(os.Stdout, &calc.Addition{})
	err := handler.Handle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

type Calculator interface {
	Calculate(a, b int) int
}

type Handler struct {
	stdout     io.Writer
	calculator Calculator
}

func NewHandler(stdout io.Writer, calculator Calculator) *Handler {
	return &Handler{
		stdout:     stdout,
		calculator: calculator,
	}
}
func (this *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return errWrongArgCount
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: '%s'", errInvalidArgument, args[0])
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: '%s'", errInvalidArgument, args[1])
	}
	result := this.calculator.Calculate(a, b)
	_, err = fmt.Fprint(this.stdout, result)
	return err
}

var (
	errWrongArgCount   = errors.New("usage: calculator <a> <b>")
	errInvalidArgument = errors.New("invalid argument")
)
