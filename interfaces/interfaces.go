package interfaces

import (
	"fmt"
)


type Reader interface {
	Get(obj string) error
	List(list []string) error 
}

type Client interface {
	Reader
}


type client struct {
	name string
}

func (c *client) Get(obj string) error {
	fmt.Println(obj)
	return nil
}

func (c *client) List(list []string) error {
	fmt.Println(list)
	return nil
}


var ReconcileClient Client = &client{} // 여기서 구조체 변수 만듦
