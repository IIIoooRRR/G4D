package cmd

import (
	"g4d-cli/templates"
	"log"
	"os"
	"path/filepath"
)

func InitDir() {
	createFiles := func(names ...string) {
		for _, name := range names {
			createFile(name)
		}
	}
	createDirs := func(names ...string) {
		for _, name := range names {
			createDir(name)
		}
	}
	createInternalDirs := func(parent string, names ...string) {
		for _, name := range names {
			createDir(filepath.Join(parent, name))
		}
	}

	createDirs("commands", "backend", "pkg", "migrations")
	createInternalDirs("backend", "entity", "consts", "dto", "services", "repositories")
	createFiles("Dockerfile")
	writeTemplate("main.go", templates.Main)
	writeTemplate("g4d.yaml", templates.Yaml)
	writeTemplate(filepath.Join("backend", "services", "services.go"), templates.Services)
	writeTemplate(filepath.Join("commands", "init.go"), templates.Init)
}
func createDir(name string) {
	err := os.Mkdir(name, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
func createFile(name string) {
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}
func writeTemplate(path, content string) {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		log.Fatal(err)
	}
}
