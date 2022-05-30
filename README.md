# gocman
a simple project manager for your golang project
## command
### run </br >
Run your complex command with a simple, just type `gocman run winbuild` to run `GOOS=windows go build .`.
You must be add `winbuild` in `gocman.config` like `{"winbuild":{"cmd":"go build
.","env":["GOOS=windows"]}}`,if you call `winbuild` the cmd will be run.
### add </br >
Add a complex command to `gocman.config`. this command will show dialog bellow. 

```bash
Command Name : winbuild
Command use + as space and > before next command : go build .
Use env? y
ENV Name : GOOS
ENV Value : go build .
Add more ? n
```
after the dialog command `winbuild` will be writing to `gocman.config` with env `GOOS=windows`
you can call it using `gocman run winbuild`

### install </br >
Install a package with simple type `gocman i /example.com/bla` after this, `bla`
package will be installed and the package name will be writed in
`gocman.package`. if you have a `gocman.package` you can call `gocman i` for
install all package are listed on `gocman.package`
