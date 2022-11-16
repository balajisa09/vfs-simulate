package kernel

import("fmt")

type Ext3 struct{

}

func(e Ext3) Write(){
	fmt.Println("wrote to file")
}

func(e Ext3) Read(){
	fmt.Println("read from file")
}

func(e Ext3) Open(){
	fmt.Println("opened a file")
}

func(e Ext3) Close(){
	fmt.Println("closed a file")
}