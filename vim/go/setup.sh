#!/bin/bash

TOOLS=`cat << EOF
# tools
github.com/uber/go-torch
github.com/golang/dep/cmd/dep
golang.org/x/tools/...
github.com/ramya-rao-a/go-outline
github.com/acroca/go-symbols
github.com/uudashr/gopkgs/cmd/gopkgs
sourcegraph.com/sqs/goreturns
github.com/cweill/gotests/...
github.com/haya14busa/goplay/cmd/goplay
# vim-go
# https://github.com/fatih/vim-go/blob/0e7bd4001b770d0c42ed85e510e6649191ec0fa8/plugin/go.vim#L33
github.com/klauspost/asmfmt/cmd/asmfmt
github.com/derekparker/delve/cmd/dlv
github.com/kisielk/errcheck
github.com/davidrjenni/reftools/cmd/fillstruct
github.com/mdempsky/gocode
github.com/rogpeppe/godef
github.com/zmb3/gogetdoc
golang.org/x/tools/cmd/goimports
github.com/golang/lint/golint
github.com/alecthomas/gometalinter
github.com/fatih/gomodifytags
golang.org/x/tools/cmd/gorename
github.com/jstemmer/gotags
golang.org/x/tools/cmd/guru
github.com/josharian/impl
honnef.co/go/tools/cmd/keyify
github.com/fatih/motion
github.com/koron/iferr
EOF
`

echo "$TOOLS" | while read LINE; do
	CHAR=${LINE:0:1}
	if [ $CHAR != "#" ]; then
		CMD="go get -u $LINE"
		echo "$CMD"
		$($CMD)
	fi
done


