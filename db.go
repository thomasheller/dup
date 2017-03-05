package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// Db implements a primitive key-value store.
type Db struct {
	filename string
	data     map[string]string
}

func newDb(filename string) *Db {
	return &Db{filename: filename}
}

func (db *Db) load() {
	db.data = make(map[string]string)

	file, err := os.Open(db.filename)
	if os.IsNotExist(err) {
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		k := parts[0]
		v := parts[1]
		db.data[k] = v
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (db *Db) save() {
	f, err := os.Create(db.filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	lines := make([]string, 0)

	for k, v := range db.data {
		line := fmt.Sprintf("%s:%s\n", k, v)
		lines = append(lines, line)
	}

	sort.Strings(lines)

	for _, line := range lines {
		if _, err = f.WriteString(line); err != nil {
			panic(err)
		}
	}
}

func (db *Db) add(key string, value string) {
	if strings.ContainsRune(key, '\n') {
		log.Fatalf("key \"%s\" must not contain newline", key)
	}
	if strings.ContainsRune(value, '\n') {
		log.Fatalf("value \"%s\" must not contain newline", value)
	}
	if strings.ContainsRune(key, ':') {
		log.Fatalf("key \"%s\" must not contain colon", key)
	}
	if strings.ContainsRune(value, ':') {
		log.Fatalf("value \"%s\" must not contain colon", value)
	}
	db.data[key] = value
}

func (db *Db) remove(key string) {
	delete(db.data, key)
}

func (db *Db) list() map[string]string {
	result := make(map[string]string)
	for k, v := range db.data {
		result[k] = v
	}
	return result
}
