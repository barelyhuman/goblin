package build

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type Binary struct {
	Container string
	Name      string
	Source    string
	Dest      string
	OS        string
	Arch      string
	Version   string
	Path      string
	Module    string
}

var environWhitelist = []string{
	"GOPATH",
	"PATH",
	"HOME",
	"PWD",
	"GOPROXY",
	"GOLANG_VERSION",
	"TMPDIR",
}

var environMap map[string]string

// init initializes the environment variable map.
func init() {
	environMap = make(map[string]string)
	for _, v := range os.Environ() {
		parts := strings.Split(v, "=")
		environMap[parts[0]] = parts[1]
	}
}

// environ returns the environment variables for Go sub-commands.
func environ() (env []string) {
	for _, name := range environWhitelist {
		env = append(env, name+"="+environMap[name])
	}
	return
}

func (bin *Binary) WriteBuild(writer io.Writer) error {
	dir, err := tempDirectory()
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("creating temporary directory: %w", err)
	}

	bin.Container = dir

	err = os.Remove(filepath.Join(dir, "go.mod"))
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("removing go.mod: %w", err)
	}

	err = bin.newModule(dir)
	if err != nil {
		return err
	}

	err = bin.createTempMainFile(dir)
	if err != nil {
		return err
	}

	err = bin.addAsDep(dir)
	if err != nil {
		return err
	}

	err = bin.runModTidy(dir)
	if err != nil {
		return err
	}

	err = bin.quickBuildBinary(dir)
	if err != nil {
		log.Println("Failed to quick build, attempting manual build")
		err = bin.buildBinary(dir)
		if err != nil {
			log.Println("Failed to manual build as fell")
			return err
		}
	}

	f, err := os.Open(bin.Dest)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, f)
	if err != nil {
		return err
	}

	os.RemoveAll(dir)
	return nil
}

func command(cmd *exec.Cmd) error {
	var w strings.Builder
	cmd.Stderr = &w
	err := cmd.Run()
	if err != nil {
		log.Println(strings.TrimSpace(w.String()))
		return err
	}
	return nil
}

func (bin *Binary) addAsDep(dir string) error {
	dep := normalizeModuleDep(bin)
	cmd := exec.Command("go", "mod", "edit", "-require", dep)

	cmd.Env = environ()
	cmd.Dir = dir
	return command(cmd)
}

func normalizeModuleDep(bin *Binary) string {
	mod := bin.Module
	version := bin.Version
	dep := fmt.Sprintf("%s@%s", mod, version)
	return dep
}

func (bin *Binary) newModule(dir string) error {
	cmd := exec.Command("go", "mod", "init", "github.com/goblin")
	cmd.Env = environ()
	cmd.Dir = dir
	return command(cmd)
}

func (bin *Binary) createTempMainFile(dir string) error {
	var fileDetails strings.Builder
	fileDetails.Write([]byte("package main\n"))
	fileDetails.Write([]byte("import(\""))
	fileDetails.Write([]byte(bin.Path))
	fileDetails.Write([]byte("\")"))
	filePath := path.Join(dir, "main.go")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Sync()
	defer file.Close()
	file.Write([]byte(fileDetails.String()))
	return nil
}

func (bin *Binary) runModTidy(dir string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Env = environ()
	cmd.Dir = dir
	return command(cmd)
}

func (bin *Binary) quickBuildBinary(dir string) error {
	dst, err := tempFilename()

	if err != nil {
		return err
	}

	bin.Dest = dst
	cmd := exec.Command("go", "build", "-o", bin.Dest, bin.Module)
	cmd.Env = environ()
	cmd.Env = append(cmd.Env, "CGO_ENABLED=0")
	cmd.Env = append(cmd.Env, "GOOS="+bin.OS)
	cmd.Env = append(cmd.Env, "GOARCH="+bin.Arch)
	cmd.Dir = dir
	return command(cmd)

}

func (bin *Binary) buildBinary(dir string) error {
	dst, err := tempFilename()

	if err != nil {
		return err
	}

	bin.Dest = dst
	cmd := exec.Command("go", "build", "-o", bin.Dest, bin.Path)
	cmd.Env = environ()
	cmd.Env = append(cmd.Env, "CGO_ENABLED=0")
	cmd.Env = append(cmd.Env, "GOOS="+bin.OS)
	cmd.Env = append(cmd.Env, "GOARCH="+bin.Arch)
	cmd.Dir = dir
	return command(cmd)
}

func (bin *Binary) Cleanup() error {
	err := os.RemoveAll(bin.Container)
	if err != nil {
		return err
	}
	err = os.RemoveAll(bin.Dest)
	return err
}

// tempFilename returns a new temporary file name.
func tempFilename() (string, error) {
	f, err := ioutil.TempFile(os.TempDir(), "goblin")
	if err != nil {
		return "", err
	}
	defer f.Close()
	defer os.Remove(f.Name())
	return f.Name(), nil
}

func tempDirectory() (string, error) {
	dir, err := ioutil.TempDir(os.TempDir(), "goblin")
	if err != nil {
		return "", err
	}
	return dir, nil
}
