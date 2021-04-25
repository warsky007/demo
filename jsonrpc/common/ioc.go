package common

import "io"

type MyReadWriteCloser struct {
	io.ReadCloser
	io.WriteCloser
}

func (ioc *MyReadWriteCloser) Read(p []byte) (n int, err error) {
	return ioc.ReadCloser.Read(p)
}

func (ioc *MyReadWriteCloser) Write(p []byte) (n int, err error) {
	return ioc.WriteCloser.Write(p)
}

func (ioc *MyReadWriteCloser) Close() error {
	if err := ioc.ReadCloser.Close(); err != nil {
		return err
	}
	if err := ioc.WriteCloser.Close(); err != nil {
		return err
	}
	return nil
}
