package gitlet

//import "io"
//import "fmt"
//import "encoding/hex"
import "crypto/sha1"
import "encoding/hex"

type ShaId struct {
	Data [sha1.Size]byte
}

type Blob struct {
	FileId ShaId
}

type Commit struct {
	Files        map[string]Blob // file-name to sha
	ParentCommit ShaId
}

func (sha *ShaId) Bytes() []byte {
	return sha.Data[:]
}

func (sha *ShaId) Name() string {
	return hex.EncodeToString(sha.Bytes())
}
