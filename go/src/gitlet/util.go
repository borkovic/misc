package gitlet

import "os"
import "path/filepath"

/***********************************************************************
 *
***********************************************************************/

/***********************************************************************
 *
***********************************************************************/
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/***********************************************************************
 *
***********************************************************************/
func RepoRootPath() (string, bool) {
	startDir, err := os.Getwd()

	// Search for subdir .git/objects

	var found = false
	for {
		_, err = os.Stat(".gitlet")
		if err == nil {
			found = true
			break
		}
		err = os.Chdir("..")
		if err != nil {
			break
		}
		path, _ := os.Getwd()
		if path == "/" {
			break
		}
	}
	if !found {
		return "", false
	}

	root, err := os.Getwd()
	err = os.Chdir(startDir)
	return root, true
}

/***********************************************************************
 *
***********************************************************************/
func RepoObjectsPath() (string, bool) {
	root, found := RepoRootPath()
	if found {
		return filepath.Join(root, ".gitlet", "objects"), true
	} else {
		return "", false
	}
}

/***********************************************************************
 *
***********************************************************************/
func RepoRefsPath() (string, bool) {
	root, found := RepoRootPath()
	if found {
		return filepath.Join(root, ".gitlet", "refs"), true
	} else {
		return "", false
	}
}



