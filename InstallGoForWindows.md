1. Download and install the latest 64 bit Go MSI distributable from https://golang.org/dl/
   (Use the default installation path: _C:\Go_)

2. Ensure the Go binary (_C:\Go\bin_) is in **Path** system environment variables.  
   Go to **System** -> **Advanced System Settings** -> **Environment Variables** and edit **Path** System Variable. Check Go binary path is in Variable value.

   Another way to check is to open a terminal and type `go version`.

3. Create **Go workspace** folder.  
   Create a folder _C:\MyGoWorkspace_ as root **Go workspace** folder.
   And then create three folders in root **Go workspace** folder: _bin_, _pkg_ and _src_

4. Setup _GOPATH_ system environment variables.  
   Go to **System** -> **Advanced System Settings** -> **Environment Variables** and click **New...*** under System variables. Set Variable name to _GOPATH_ and Variable value to _C:\MyGoWorkspace_.

   To ensure _GOPATH_ has been set: open a terminal and type `echo %GOPATH%`

5. Install Go Packages  
   * `go get -u -v github.com/nsf/gocode`
   * `go get -u -v github.com/rogpeppe/godef`
   * `go get -u -v github.com/golang/lint/golint`
   * `go get -u -v sourcegraph.com/sqs/goreturns`
   * `go get -u -v golang.org/x/tools/cmd/gorename`
   * `go get -u -v github.com/derekparker/delve/cmd/dlv`
   * `go get -u -v github.com/tpng/gopkgs`

Done.
