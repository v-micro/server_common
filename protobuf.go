package server_common

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var buf []string

func Protobuf(name string) error {
	cmd := exec.Command("cmd.exe", "/C", "protoc --go_out=plugins=grpc:. "+name)
	w := bytes.NewBuffer(nil)
	cmd.Stderr = w
	if err := cmd.Run(); err != nil {
		return err
	}

	//==== 删除pb文件的 omitempty
	pb := strings.Replace(name,".proto",".pb.go",-1)
	_, err := PathExists(pb)
	if err != nil {return err}

	f, err := os.OpenFile(pb, os.O_RDWR, os.ModePerm);
	if err != nil {return err}
	defer f.Close()

	b, err := ioutil.ReadAll(f);
	if err != nil {return err}
	n := strings.ReplaceAll(string(b), ",omitempty", "")
	err = os.Truncate(pb, 0)
	if err != nil {return err}
	_, err = f.WriteAt([]byte(n), 0)
	if err != nil {return err}

	return nil
}

//遍历目录
func GetBufFiles(folder string){
	files, _ := ioutil.ReadDir(folder)
	for _,file := range files{
		if file.IsDir(){
			GetBufFiles(folder + "/" + file.Name())
		}else if strings.HasSuffix(file.Name(), ".proto") == true{
			buf = append(buf,folder + "/" + file.Name())
		}
	}
}

//目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}


func Run() error {
	GetBufFiles(".")	//获取目录的protobuf文件
	for _,v := range buf{
		err := Protobuf(v)
		if err != nil{
			fmt.Sprintf("%s => %s.pb.go error.", v, strings.TrimRight(v, ".proto"))
			return err
		}
		fmt.Println(fmt.Sprintf("%s => %s.pb.go success.", v, strings.TrimRight(v, ".proto")))
	}
	return nil
}

