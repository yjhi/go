package jdir

import "io/ioutil"

func (d *DirUtils) _GetAllFiles(path string) int {
	iFile := 0

	fReader, err := ioutil.ReadDir(path)

	if err != nil {
		return 0
	}

	for _, fi := range fReader {
		if !fi.IsDir() {
			iFile++
			d.Files = append(d.Files, path+"/"+fi.Name())
		} else {
			iFile += d._GetAllFiles(path + "/" + fi.Name())
		}
	}

	return iFile

}

func (d *DirUtils) _GetDirFiles(path string) int {

	iFile := 0

	fReader, err := ioutil.ReadDir(path)

	if err != nil {
		return 0
	}

	for _, fi := range fReader {
		if !fi.IsDir() {
			iFile++
			d.Files = append(d.Files, path+"/"+fi.Name())

		}
	}

	return iFile
}
