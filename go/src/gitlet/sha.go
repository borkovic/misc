package gitlet

import "os"
import "io"
import "crypto/sha1"
import "encoding/hex"

/***********************************************************************
 *
***********************************************************************/
type ShaId struct {
	Data string
	//Data [sha1.Size]byte
}

/***********************************************************************
 *
***********************************************************************/
func (sha *ShaId) AsBytes() ([]byte, error) {
	return hex.DecodeString(sha.Data)
}

/***********************************************************************
 *
***********************************************************************/
func (sha *ShaId) AsString() string {
	return sha.Data
}

/***********************************************************************
 *
***********************************************************************/
func (sha *ShaId) ShaOfString(s string) {
	bytes := []byte(s)
	sha.ShaOfBytes(bytes)
}

/***********************************************************************
 *
***********************************************************************/
func (sha *ShaId) ShaOfBytes(bytes []byte) {
	signature := sha1.Sum(bytes)
	sha.Data = hex.EncodeToString(signature[:])
}

/***********************************************************************
 *
***********************************************************************/
func (sha *ShaId) ShaOfFile(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	hasher := sha1.New()
	if _, err := io.Copy(hasher, file); err != nil {
		sha.Data = ""
	}
	sha.Data = hex.EncodeToString(hasher.Sum(nil))
}
