package kernel

type VFS interface{
	read()
	write() 
	open()
	close()
}