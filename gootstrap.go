package main

func getTemplatesFilenames(dstdir string) map[string]string {
	return nil
}

func prepareDirectories(srcToDstDirs map[string]string) {
}

func applyTemplates(
	srcToDstDirs map[string]string,
	project string,
	docker_registry string,
) {
}

func main() {
	dstdir := "TODO/ARG"
	project := "TODO/ARG"
	docker_registry := "TODO/ARG"

	srcToDstDirs := getTemplatesFilenames(dstdir)
	prepareDirectories(srcToDstDirs)
	applyTemplates(srcToDstDirs, project, docker_registry)
}
