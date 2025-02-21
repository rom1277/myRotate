package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// -a поместит в папку архив
func main() {
	aflag := flag.String("a", "", "is waiting for the directory in which it will place the archived files")
	flag.Parse()
	files := flag.Args()

	if len(files) == 0 {
		fmt.Println("Error: files not provided.")
		return
	}

	for _, file := range files {

		if err := rotate(file, *aflag); err != nil {
			fmt.Println("Error: ", err)
		}

	}

	fmt.Println(*aflag, files, len(files))
}

func rotate(filepath1, archiveDir string) error {
	fileInfo, err := os.Stat(filepath1)
	if err != nil {
		return fmt.Errorf("failed to get file info for '%s': %w", filepath1, err)
	}

	if fileInfo.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a file", filepath1)
	}

	if fileInfo.Size() == 0 {
		fmt.Printf("File '%s' is empty, skipping archiving.\n", filepath1)
		return nil
	}

	modTime := fileInfo.ModTime()   // Получаем время последнего изменения файла
	modTimeformat := modTime.Unix() // Преобразуем время в Unix timestamp
	file, err := os.Open(filepath1)
	if err != nil {
		return err
	}
	defer file.Close()

	newfilepath := strings.TrimSuffix(filepath.Base(filepath1), filepath.Ext(filepath1)) + "_" + strconv.Itoa(int(modTimeformat)) + ".tar.gz"

	if archiveDir != "" {
		if _, err := os.Stat(archiveDir); os.IsNotExist(err) {
			if err := os.MkdirAll(archiveDir, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", archiveDir, err)
			}
		}
		newfilepath = filepath.Join(archiveDir, newfilepath)
	}

	newfile, err := os.Create(newfilepath)
	if err != nil {
		return err
	}
	defer newfile.Close()

	// Создаем gzip-компрессор
	gzWriter := gzip.NewWriter(newfile)
	defer gzWriter.Close()

	// Создаем tar-архив
	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	// Добавляем файл в архив
	header := &tar.Header{
		Name:    filepath1,
		Size:    fileInfo.Size(),
		Mode:    int64(fileInfo.Mode()),
		ModTime: fileInfo.ModTime(),
	}
	if err := tarWriter.WriteHeader(header); err != nil {
		return err
	}
	// Копируем содержимое файла в архив
	if _, err := io.Copy(tarWriter, file); err != nil {
		return err

	}
	fmt.Printf("File '%s' successfully archived to '%s'\n", filepath1, newfilepath)
	return nil
}
