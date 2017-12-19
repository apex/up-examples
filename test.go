package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/pkg/errors"
)

// TODO: move to tj/ as a separate program?
// TODO: prefix output
// TODO: terse output

func main() {
	println()
	defer println()

	dir := flag.String("dir", "", "Directory to test.")
	pro := flag.Bool("pro", false, "Test with pro version.")
	flag.Parse()

	log.SetHandler(cli.Default)

	if *dir != "" {
		if err := testExample(*dir); err != nil {
			log.Fatalf("error: %s", err)
		}
		return
	}

	if err := testExamples("oss"); err != nil {
		log.Fatalf("error: %s", err)
	}

	if *pro {
		if err := testExamples("pro"); err != nil {
			log.Fatalf("error: %s", err)
		}
	}
}

func testExamples(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return errors.Wrap(err, "reading dir")
	}

	for _, f := range files {
		if !f.IsDir() {
			continue
		}

		root := filepath.Join(dir, f.Name())
		if err := testExample(root); err != nil {
			return err
		}
	}

	return nil
}

func testExample(dir string) error {
	test := filepath.Join(dir, "test.sh")

	_, err := os.Stat(test)
	if err != nil {
		return nil
	}

	start := time.Now()
	fmt.Printf("\033[;1m==> %s\033[0m\n", filepath.Base(dir))
	defer func() {
		fmt.Printf("-->\n")
		fmt.Printf("\033[;1m==> ok (%s)\033[0m\n\n", time.Since(start))
	}()

	f, err := os.Open(test)
	if err != nil {
		return errors.Wrap(err, "opening")
	}
	defer f.Close()

	if err := os.Chdir(dir); err != nil {
		return errors.Wrap(err, "chdir")
	}

	scan := bufio.NewScanner(f)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	for scan.Scan() {
		line := scan.Text()

		if isComment(line) {
			parts := strings.SplitN(strings.Trim(line, " \t#"), ":", 2)
			expect := strings.TrimSpace(parts[1])
			actual := strings.TrimSpace(stdout.String())
			directive(parts[0], actual, expect)
			continue
		}

		stdout.Reset()
		stderr.Reset()

		fmt.Printf("--> %s \n", line)
		cmd := exec.Command("sh", "-c", line)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		// cmd.Stdout = io.MultiWriter(&stdout, os.Stdout)
		// cmd.Stderr = io.MultiWriter(&stderr, os.Stderr)
		if err := cmd.Run(); err != nil {
			return errors.Wrapf(err, stderr.String())
		}
	}

	if err := scan.Err(); err != nil {
		return errors.Wrap(err, "scan")
	}

	if err := os.Chdir(filepath.Join("..", "..")); err != nil {
		return errors.Wrap(err, "chdir")
	}

	return nil
}

func directive(name string, actual, expect string) {
	if name == "equals" {
		fmt.Printf("--> test equals %q\n", expect)
		if actual != expect {
			fmt.Printf("\n")
			fmt.Printf("\033[31m  Expected to equal:\n\n  %q\033[0m\n", expect)
			fmt.Printf("\n\033[31m  Output:\n\n  %q\033[0m\n", actual)
		}
		return
	}

	if name == "contains" {
		fmt.Printf("--> test contains %q\n", expect)
		if !strings.Contains(actual, expect) {
			fmt.Printf("\n")
			fmt.Printf("\033[31m  Expected to contain:\n\n  %s\033[0m\n", expect)
			fmt.Printf("\n\033[31m  Output:\n\n  %s\033[0m\n", actual)
		}
		return
	}

	if name == "not contains" {
		fmt.Printf("--> test not contains %q\n", expect)
		if strings.Contains(actual, expect) {
			fmt.Printf("\n")
			fmt.Printf("\033[31m  Expected to not contain:\n\n  %s\033[0m\n", expect)
			fmt.Printf("\n\033[31m  Output:\n\n  %s\033[0m\n", actual)
		}
		return
	}

	panic(fmt.Sprintf("unsupported directive %q", name))
}

func isCommand(s string) bool {
	return !isComment(s)
}

func isComment(s string) bool {
	return strings.HasPrefix(strings.TrimSpace(s), "#")
}
