package kernel

import("fmt")

type Ext4 struct{

}

func(e Ext4) Write(){
	fmt.Println("wrote to file")
}

func(e Ext4) Read(){
	fmt.Println("read from file")
}

func(e Ext4) Open(){
	fmt.Println("opened a file")
}

func(e Ext4) Close(){
	fmt.Println("closed a file")
}