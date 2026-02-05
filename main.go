package main

import (
	"os"
	"fmt"
	"math/rand"
)

type KV interface {
	Get(key []byte) (val []byte, ok bool)
	Set(key []byte, val []byte)
	Del(key []byte)
	FindGreaterThan(key []byte) Iterator
}

type Iterator interface {
	HasNext() bool
	Next() (key []byte, val []byte)
}

func randomInt() int64 {
	return rand.Int63()
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

func SaveData2(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, randomInt())
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}
	defer func() {
		fp.Close()
		if err != nil {
			os.Remove(tmp)
		}
	}()

	if _, err = fp.Write(data); err != nil {
		return err
	}
	if err = fp.Sync(); err != nil {
		return err
	}
	err = os.Rename(tmp, path)

	dir := filepath .Dir(path)
	fd, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer fd.Close()

	if err = fd.Sync(); err != nil {
		return err
	}

	return err
}
