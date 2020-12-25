package main

import (
	"errors"
	"os"
	"sync"
)

type DataFile interface {
	//读取一个文件块
	Read() (rsn int64, d Data, err error)
	//写取一个文件块
	Write(d Data) (wsn int64, err error)
	//获取最后读取的数据块的序列号
	RSN() int64
	//获取最后写入一个数据块的序列号
	DSN() int64
	//获取数据块的长度
	DataLen() uint32
	//关闭数据文件
	close() error
}
type Data []byte

type myDataFile struct {
	f       *os.File
	fmutex  sync.RWMutex
	woffset int64
	roffset int64
	wmutex  sync.Mutex
	rmutex  sync.Mutex
	datalen uint32
}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	var offset int64
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset += int64(df.datalen)

	//读取一个数据块
	rsn = offset / int64(df.datalen)
	df.fmutex.RLock()
	defer df.fmutex.RUnlock()
	bytes := make([]byte, df.datalen)
	_, err = df.f.ReadAt(bytes, offset)
	if err != nil {
		return
	}
	d = bytes
	return
}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	panic("implement me")
}

func (df *myDataFile) RSN() int64 {
	panic("implement me")
}

func (df *myDataFile) DSN() int64 {
	panic("implement me")
}

func (df *myDataFile) DataLen() uint32 {
	panic("implement me")
}

func (df *myDataFile) close() error {
	panic("implement me")
}

func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if dataLen == 0 {
		return nil, errors.New("Invalid data lenth!")

	}
	df := &myDataFile{f: f, datalen: dataLen}
	return df,nil

}

func main() {

}
