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

type Tree struct {
	Files        map[string]ShaId // file-name to blob-sha (and subtree shas later)
}

type Commit struct {
	ParentCommit ShaId  // parent commit's sha
	RootTree ShaId		// root tree sha
}

func (sha *ShaId) Bytes() []byte {
	return sha.Data[:]
}

func (sha *ShaId) Name() string {
	return hex.EncodeToString(sha.Bytes())
}
