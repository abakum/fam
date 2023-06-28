/*
\fam\dad\main.go
\fam\dad\son\a.go
\fam\dad\son\b.go
\fam\dad\c.go
\fam\bro\d.go
\fam\mom\e.go
\fam\dad\e.go // as symbolic link to /fam/mom/e.go

cd \fam
go work init
go work use dad
go work use bro


cd bro
go mod init bro

cd ..\dad
go mod init dad
mklink e.go ..\mom\e.go

go get rsc.io/quote
go run .
*/

// fam/dad/main.go
package main

import (
	"bro"     // local package as bro of dad
	"dad/son" // local package as son of dad
	"fmt"

	"rsc.io/quote" // remote package
)

func main() { //fam/dad/main.go
	fmt.Println(son.A()) //fam/dad/son/a.go
	fmt.Println(son.B)   //fam/dad/son/b.go
	fmt.Println(c())     //fam/dad/c.go
	fmt.Println(bro.D()) //fam/bro/d.go
	fmt.Println(e)       //fam/dad/e.go as symbolic link to /fam/mom/e.go
	fmt.Println(quote.Hello())
}
