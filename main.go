package main

import (
	"os"
)

type KV interface {
	Get(key []byte) (val []byte, ok bool)
	Set(key []byte, val []byte)
	Del(key []byte)FindGreaterThan(key []byte) Iterator
}

type Iterator interface {
	HasNext() bool
	Next() (key []byte, val []byte)
}

func SaveData1(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		return err
	}
	return  fp.Sync()
}

