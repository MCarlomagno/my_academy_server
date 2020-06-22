package storage

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"os"

	"github.com/t3rm1n4l/go-mega"
)

// createFile creates a temporary file of a given size along with its MD5SUM
func createFile(size int64) (string, string) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("testfile.txt")
	if err != nil {
		panic(err)
	}
	_, err = file.Write(b)
	if err != nil {
		panic(err)
	}
	h := md5.New()
	_, err = h.Write(b)
	if err != nil {
		panic(err)
	}
	return file.Name(), fmt.Sprintf("%x", h.Sum(nil))
}

func initSession() *mega.Mega {
	var USER string = os.Getenv("MEGA_USER")
	var PASSWORD string = os.Getenv("MEGA_PASSWD")
	// m.SetDebugger(log.Printf)
	fmt.Println(USER)
	fmt.Println(PASSWORD)
	m := mega.New()
	m.Login(USER, PASSWORD)
	return m
}

// uploadFile uploads a temporary file of a given size returning the
// node, name and its MD5SUM
func uploadFile(session *mega.Mega, size int64, parent *mega.Node) (node *mega.Node, name string, md5sum string) {
	name, md5sum = createFile(size)
	defer func() {
		_ = os.Remove(name)
	}()
	node, err := session.UploadFile(name, parent, "", nil)
	if err != nil {
		panic(err)
	}
	return node, name, md5sum
}

// downloadFile uploads a temporary file of a given size returning the
// node, name and its MD5SUM
func downloadFile(session *mega.Mega, node *mega.Node, name string, progress *chan int) {
	fmt.Println("name")
	fmt.Println(name)
	err := session.DownloadFile(node, name, nil)
	if err != nil {
		panic(err)
	}
}

//StartStorage testing only
func StartStorage() {
	// session := initSession()
	// fmt.Println("logged In")
	// node, name, h1 := uploadFile(session, 314573, session.FS.GetRoot())
	// newname := "sample_video.mp4"
	// downloadFile(session, node, newname, nil)

}
