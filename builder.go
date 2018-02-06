package main

import (
	"os"
	"path/filepath"
	"fmt"
	"io"
)

func Build() error {
	// 1. empty the site directory
	fmt.Printf("Emptying site dir %s\n", config.SiteDir)
	if err := os.RemoveAll(config.SiteDir); err != nil {
		return err
	}

	fmt.Printf("Recreating site dir with correct permissions\n")
	if err := os.MkdirAll(config.SiteDir, 0755); err != nil {
		return err
	}

	// 2. walk the src directory, skipping anything starting with an underscore
	err := filepath.Walk(config.SrcDir, func(path string, info os.FileInfo, err error) error {
		name := info.Name()
		if name[0] == '_' {
			fmt.Printf("Skipping entirely %s\n", info.Name())
			return filepath.SkipDir
		}

		if info.IsDir() {
			fmt.Printf("Seeing dir %s\n", info.Name())
			return nil
		}

		if !info.Mode().IsRegular() {
			fmt.Printf("Seeing irregular %s\n", info.Name())
			return nil
		}

		fmt.Printf("Seeing file %s\n", info.Name())
		switch filepath.Ext(path) {
		case ".html":
			buildHtml(path)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func buildHtml(path string) error {
	destPath := filepath.Join(config.SiteDir, path[len(config.SrcDir):])
	fmt.Printf("building html file %s\n", destPath)

	// make the destination directory in case it does not exist
	dir, _ := filepath.Split(destPath)
	if _, err := os.Stat(dir); err != nil {
		fmt.Printf("creating dir %s\n", destPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// open the src file
	fmt.Printf("opening src file %s\n", path)
	src, err := os.Open(path)
	if err != nil {
		return err
	}
	defer src.Close()

	// create the dest file
	fmt.Printf("creating dest file %s\n", destPath)
	dest, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	// copy across the file
	fmt.Printf("copying src to dest\n")
	if _, err = io.Copy(dest, src); err != nil {
		return err
	}

	// make sure the file is synced to disc
	fmt.Printf("syncing dest file\n")
	return dest.Sync()
}
