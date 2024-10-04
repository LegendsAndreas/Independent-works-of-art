// From: https://www.youtube.com/watch?v=xbI2ELnVGAo
package main

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Checks a directory, and if it has changed, it archives it.
func main() {
	fileHaslList := make(map[string][]byte)
	archive_flag := false

	for {
		// filepath.WalkDir(), asks for a directory name and an error message as parameters. As the error parameter, we use
		// a method, which returns an error, nil or actual.
		filepath.WalkDir("input", func(path string, info os.DirEntry, err error) error { // Whole functon checks if we actually should archive something.
			// In case an error occurs, we return the error.
			if err != nil {
				fmt.Println(err)
				return err
			}

			// We check if this actually is a directory.
			if info.IsDir() {
				return nil
			}

			// Do we have this file?
			if hash, ok := fileHaslList[path]; ok {
				// Take the hash out of this file.
				file, _ := os.Open(path)
				h := md5.New()
				io.Copy(h, file)
				nhash := h.Sum(nil)

				if !bytes.Equal(hash, nhash) {
					archive_flag = true
					return errors.New("Rearchive")
				}

				file.Close()
			} else {
				archive_flag = true
				return errors.New("Rearchive")
			}
			return nil
		})

		if archive_flag {
			// Creates the archive
			os.Remove("output.zip")
			outfile, _ := os.Create("output.zip")
			w := zip.NewWriter(outfile)
			log.Println("Creating a new archive")
			filepath.WalkDir("input", func(path string, info os.DirEntry, err error) error {
				// In case an error occurs, we return the error.
				if err != nil {
					return err
				}

				// We check if this actually is a directory.
				if info.IsDir() {
					return nil
				}
				// Create hash of a new file.
				file, _ := os.Open(path)
				h := md5.New()
				io.Copy(h, file)
				nhash := h.Sum(nil)
				fileHaslList[path] = nhash

				// Compress the file
				f, _ := w.Create(path)
				file.Seek(0, io.SeekStart) // Rewinding the file, so that we can read it again.
				io.Copy(f, file)
				file.Close()
				return nil
			})
			archive_flag = false
			w.Close()
		}
		// We put a sleep statement on 1 second, because having the program check for changes every millisecond would be a waste of resources.
		time.Sleep(time.Second * 1)
	}
}
