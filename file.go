package file

import (
	"io/ioutil"
	"log"
	"os"
)

func EmptyFolder(folder string) {
	dirRead, _ := os.Open(folder)
	defer dirRead.Close()
	dirFiles, _ := dirRead.Readdir(0)

	// Loop over the directory's files.
	for index := range dirFiles {
		fileHere := dirFiles[index]

		// Get name of file and its full path.
		nameHere := fileHere.Name()
		fullPath := folder + "/" + nameHere

		// Remove the file.
		os.RemoveAll(fullPath)

	}
}

// Empties workFolder then copy over files (and files only) from remoteFolder.
// Does NOT go into sub folders in remoteFolder
func CopyToWorkFolder(remoteFolder string, workFolder string) []string {
	var copiedFiles []string
	EmptyFolder(workFolder)

	dirRead, _ := os.Open(remoteFolder)
	defer dirRead.Close()
	dirFiles, _ := dirRead.Readdir(0)

	// Loop over the directory's files.
	for index := range dirFiles {
		fileHere := dirFiles[index]
		if fileHere.IsDir() {
			continue
		}
		// Get name of file and its full path.
		nameHere := fileHere.Name()

		data, err := ioutil.ReadFile(remoteFolder + "\\" + nameHere)
		if err != nil {
			log.Println(remoteFolder+"\\"+nameHere, err)
			continue
		}

		err = ioutil.WriteFile(workFolder+"\\"+nameHere, data, 0644)
		if err != nil {
			log.Println(workFolder+"\\"+nameHere, err)
			continue
		}
		copiedFiles = append(copiedFiles, workFolder+"\\"+nameHere)

	}

	return copiedFiles

}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
