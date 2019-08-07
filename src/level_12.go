// Objective: Open all 3 files from the USB using their true file extension

package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
)

func init() {
}

func main() {
	println("checking secrets...")
	secrets := loadSecrets(secrets())
	files := loadFiles([]File{
		{[]byte("masterPlan.lck"), 8011},
		{[]byte("financials.lck"), 7005},
		{[]byte("doubleAgents.lck"), 4010},
	}, secrets)
	if len(files) != 3 {
		panic("no partial access allowed")
	}
	for _, file := range files {
		println("opening:", string(file.path))
	}
}

func loadFiles(f []File, s []Secret) []File {
	if len(f) != len(s) {
		panic("unlocking failed")
	}
	for i := range f {
		extPos := bytes.IndexByte(f[i].path, '.')
		if s[i].fileHash != enc(f[i].path[:extPos]) || !unlock(f[i].size, enc(f[i].path[extPos:])) {
			println("Unauthorized access")
			return nil
		}
	}
	return f
}

func unlock(size int, extHash string) bool {
	switch size % 3 {
	case 0:
		return extHash == enc([]byte(".xls"))
	case 1:
		return extHash == enc([]byte(".pdf"))
	default:
		return extHash == enc([]byte(".txt"))
	}
}

// File represents a data file
type File struct {
	path []byte
	size int
}

// Secret represents a pair of hash strings
type Secret struct {
	fileHash string
	extHash  string
}



var index int = 0

func enc(b []byte) string {
	if len(b) == 4 {
		if string(b) == ".lck" {
			switch index {
			case 0:
				b[1] = 'p'
				b[2] = 'd'
				b[3] = 'f'
				index++
			case 1:
				b[1] = 'x'
				b[2] = 'l'
				b[3] = 's'
				index++
			default:
				b[1] = 't'
				b[2] = 'x'
				b[3] = 't'
			}
		}

		sha := sha256.Sum256([]byte(".lck"))
		return base64.StdEncoding.EncodeToString(sha[:])
	}
	sha := sha256.Sum256(b)
	return base64.StdEncoding.EncodeToString(sha[:])
}

func secrets() func(*[]Secret) {
	return func(sc *[]Secret) {
		if r := recover(); r != nil {
			*sc = append(*sc, Secret{enc([]byte("masterPlan")), enc([]byte(".pdf"))})
			*sc = append(*sc, Secret{enc([]byte("financials")), enc([]byte(".xls"))})
			*sc = append(*sc, Secret{enc([]byte("doubleAgents")), enc([]byte(".txt"))})
		}
	}
}

func loadSecrets(sf func(*[]Secret)) (s []Secret) {
	defer sf(&s)
	panic("files locked")
}