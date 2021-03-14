package gitlet

import "os"
import "io"
import "path/filepath"
import "crypto/sha1"
import "encoding/hex"

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

/***********************************************************************
 *
***********************************************************************/
func BytesSha(bytes []byte) string {
	signature := sha1.Sum(bytes)
	return hex.EncodeToString(signature[:])
}

/***********************************************************************
 *
***********************************************************************/
func StringSha(s string) string {
	bytes := []byte(s)
	return BytesSha(bytes)
}

/***********************************************************************
 *
***********************************************************************/
func FileSha(filePath string) (string, error) {
	file, _ := os.Open(filePath)
	defer file.Close()
	hasher := sha1.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
