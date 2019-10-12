package yell

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type yell struct {
	*log.Logger
}
type Config struct {
	Path     string
	FileName string
}

func (y *yell) Println(v ...interface{}) {
	err := y.Output(2, fmt.Sprintln(v))
	if err != nil {
		fmt.Println(err)
	}
}
func (y *yell) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	err := y.Output(2, s)
	if err != nil {
		fmt.Println(err)
	}
	panic(s)

}
func New(config Config, prefix string) *yell {
	if config.Path != "" {
		err := os.Mkdir(config.Path, 766)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		config.Path = "./"
	}
	f, err := os.OpenFile(config.Path+"/"+config.FileName+"-"+time.Now().Month().String()+"-"+strconv.Itoa(time.Now().Day())+".log", os.O_CREATE|os.O_APPEND, 766)
	if err != nil {
		return nil
	}
	logs := log.New(f, prefix, 1|2)
	return &yell{
		Logger: logs,
	}
}
