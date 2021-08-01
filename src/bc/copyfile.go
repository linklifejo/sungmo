package bc

import (
	"io"
	"io/ioutil"
	"os"
)

// 소스경로와 목적지경로를 인수로 받아서
// 소스경로에 있는 파일을 목적지경로로 복사하는 함수.
// 하위 폴더에 있는 파일은 복사하지 않는다.

func copyDir(src, dst string) (err error) {
	// 목적지 폴더를 만든다.
	err = os.MkdirAll(dst, 0777)
	if err != nil {
		return err
	}
	// 소스 폴더 하위의 파일 목록을 가져온다.
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	// 파일 리스트를 for문 돌면서 하나씩 복사한다.
	for _, f := range files {
		if f.IsDir() { // 디렉토리라면 복사하지 않는다.
			continue
		}
		s := src + "/" + f.Name()
		d := dst + "/" + f.Name()
		err = copyFile(s, d)
		if err != nil {
			return err
		}
	}
	return nil
}

func copyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	err = out.Sync()
	if err != nil {
		return err
	}
	return nil
}
