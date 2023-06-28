# fam
golang modules are just like family members
- [VS Code multi-root workspace](https://code.visualstudio.com/docs/editor/multi-root-workspaces) for [fam.code-workspace](fam.code-workspace)
- [golang workspace](https://github.com/golang/tools/blob/master/gopls/doc/workspace.md) for [go.work](go.work)
## Usage
```
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

cd ..
git init
git add *
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/abakum/fam.git
git push -u origin main
```
## [Look](dad/main.go)
![2023-06-28_14-24-54](https://github.com/abakum/fam/assets/18271863/681de844-66a3-43f9-9c15-55f21a18d945)
