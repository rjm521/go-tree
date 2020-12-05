package main

import (
	"fmt"
	"github.com/juju/errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// "." represent current directory
	args := []string{"."}

	// if we have more than one args then we can tree those args
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	for _, path := range args {
		if err := mytree(path, ""); err != nil {
			log.Printf("could not tree this path: %s :%v", path, err)
		}
	}
}

func mytree(root, indent string) error {

	stat, err := os.Stat(root)
	if err != nil {
		return errors.Annotatef(err, "could not stat %s", root)
	}

	fmt.Println(stat.Name())

	// if root is not a directory which means it's a leaf node we return
	// else we recursively call our tree func
	if !stat.IsDir() {
		return nil
	}

	// read all dirs and files in that root directory
	ds, err := ioutil.ReadDir(root)
	if err != nil {
		return errors.Annotatef(err, "could not read dir %s", stat.Name())
	}
	// put everything in that fs but not ingore file or ingore directory
	var fs []string
	for _, d := range ds {
		if d.Name()[0] != '.' {
			fs = append(fs, d.Name())
		}
	}

	for i, v := range fs {
		add := "│  "
		// last child we print this pattern
		if i == len(fs)-1 {
			// v don't have a little brother any more
			fmt.Printf(indent + "└──")
			add = "   "
		} else {
			// this v has brothers
			fmt.Printf(indent + "├──")
		}
		if err := mytree(filepath.Join(root, v), indent+add); err != nil {
			return err
		}
	}
	return nil
}
