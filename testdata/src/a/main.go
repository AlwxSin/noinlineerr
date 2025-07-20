package a

import (
	"fmt"
	"strconv"
)

type MyAliasErr error

type MyCustomError struct{}

func (mc *MyCustomError) Error() string {
	return "error"
}

func doSomething() error {
	return nil
}

func doSmthManyArgs(a, b, c, d int) error {
	return nil
}

func doSmthMultipleReturn() (bool, error) {
	return false, nil
}

func doMyAliasErr() MyAliasErr {
	return nil
}

func doMyCustomErr() *MyCustomError {
	return &MyCustomError{}
}

func valid() error {
	err := doSomething() // ok
	if err != nil {
		return err
	}

	err = doSmthManyArgs(0, 0, 0, 0) // ok
	if err != nil {
		return err
	}

	_, err = doSmthMultipleReturn() // ok
	if err != nil {
		return err
	}

	if ok, _ := strconv.ParseBool("1"); ok {
		fmt.Println("ok")
	}
	
	var myErr MyAliasErr
	if err, ok := myErr.(error); ok {
		fmt.Println("ok", err)
	}

	return nil
}

func invalid() error {
	if err := doSomething(); err != nil { // want "avoid inline error handling using `if err := ...; err != nil`; use plain assignment `err := ...`"
		return err
	}

	if err := doSmthManyArgs(0, // want "avoid inline error handling using `if err := ...; err != nil`; use plain assignment `err := ...`"
		0,
		0,
		0,
	); err != nil {
		return err
	}

	if _, err := doSmthMultipleReturn(); err != nil { // want "avoid inline error handling using `if err := ...; err != nil`; use plain assignment `err := ...`"
		_ = false
		return err
	}

	if err := doMyAliasErr(); err != nil { // want "avoid inline error handling using `if err := ...; err != nil`; use plain assignment `err := ...`"
		return err
	}

	if err := doMyCustomErr(); err != nil { // want "avoid inline error handling using `if err := ...; err != nil`; use plain assignment `err := ...`"
		return err
	}

	return nil
}

func errShadow() error {
	var err error
	if err := doSomething(); err != nil { // want "avoid inline error handling using `if err := ...; err != nil`; use plain assignment `err := ...`"
		return err
	}
	fmt.Println(err)

	return nil
}

func naming() error {
	if otherName := doSomething(); otherName != nil { // want "avoid inline error handling using `if err := ...; err != nil`; use plain assignment `err := ...`"
		return otherName
	}

	return nil
}

