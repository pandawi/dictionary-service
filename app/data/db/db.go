package db

import (
	"bufio"
	"os"
)

type DbImpl struct {
	data []string
}

type DB interface {
	LoadData(filePath string) error
	GetData() []string
}

func (dbImpl *DbImpl) LoadData(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		dbImpl.data = append(dbImpl.data, string(word))
	}
	if err := scanner.Err(); err != nil {
		return nil
	}
	return nil
}

func (dbImpl *DbImpl) GetData() []string {
	return dbImpl.data
}

func NewDb() DB {
	return &DbImpl{data: []string{}}
}
