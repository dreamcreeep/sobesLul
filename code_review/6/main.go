package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"strings"
	"time"
)

// GetFile является примером функции, которая относительно
// немного выполняется при едином вызове. Но достаточно
// долго если вызывать последовательно.
// Предположим, что оптимизировать в этой функции нечего.
func GetFile(ctx context.Context, name string) ([]byte, error) {
	if name == "" {
		return nil, fmt.Errorf("invalid name %q", name)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-ticker.C:
	}

	if strings.HasPrefix(name, "invalid") {
		return nil, fmt.Errorf("invalid name %q", name)
	}

	b := make([]byte, 10)
	n, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("getting file %q: %w", name, err)
	}

	return b[:n], nil
}

// GetFilesOld пример функции, которую нужно оптимизировать.
// Менять эту функцию не нужно, она нужна чтоб сравнить
// поведение двух функций после оптимизации.
func GetFilesOld(ctx context.Context, names ...string) (result map[string][]byte, err error) {
	if len(names) == 0 {
		return nil, nil
	}

	result = make(map[string][]byte, len(names))
	for _, name := range names {
		result[name], err = GetFile(ctx, name)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// GetFilesNew эту функцию можно менять, за исключением её сигнатуры.
func GetFilesNew(ctx context.Context, names ...string) (result map[string][]byte, err error) {
	if len(names) == 0 {
		return nil, nil
	}

	result = make(map[string][]byte, len(names))
	for _, name := range names {
		result[name], err = GetFile(ctx, name)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func main() {
	start := time.Now()
	files, err := GetFilesNew(context.TODO(), "1", "2", "3")
	if err != nil {
		log.Fatalln(time.Since(start), err)
	}

	fmt.Println(time.Since(start))
	for name := range files {
		fmt.Println(name)
	}
}
