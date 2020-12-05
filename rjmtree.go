package main

import (
	"fmt"
	"github.com/juju/errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// "." represent current directory
	args := []string{"."}

	// if we have more than one args then we can tree those args
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	for _, path := range args {
		if err := mytree(path, 0); err != nil {
			log.Printf("could not tree this path: %s :%v", path, err)
		}
	}
}

func mytree(root string, depth int) error {
	stat, err := os.Stat(root)
	if err != nil {
		return errors.Annotatef(err, "could not stat %s", root)
	}

	// don't tree the ignore file or ignore directories
	if stat.Name()[0] == '.' {
		return nil
	}

	fmt.Printf("%s%s\n", strings.Repeat("  ", depth), stat.Name())

	// if root is not a directory which means it's a leaf node we return
	if !stat.IsDir() {
		return nil
	}
	// if root is a dirctory then we recursively call tree
	fs, err := ioutil.ReadDir(root)
	if err != nil {
		return errors.Annotatef(err, "could not read dir %s", stat.Name())
	}
	for _, v := range fs {
		if err := mytree(filepath.Join(root, v.Name()), depth+1); err != nil {
			return err
		}
	}
	return nil
}
