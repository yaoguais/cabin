#!/bin/bash

TOOLS=`cat << EOF
# 调试工具
github.com/derekparker/delve/cmd/dlv
# 火焰图工具
github.com/uber/go-torch
github.com/golang/dep/cmd/dep
golang.org/x/tools/...
github.com/ramya-rao-a/go-outline
github.com/acroca/go-symbols
github.com/nsf/gocode
github.com/rogpeppe/godef
github.com/zmb3/gogetdoc
github.com/golang/lint/golint
github.com/fatih/gomodifytags
github.com/uudashr/gopkgs/cmd/gopkgs
sourcegraph.com/sqs/goreturns
github.com/cweill/gotests/...
github.com/josharian/impl
github.com/haya14busa/goplay/cmd/goplay
github.com/uudashr/gopkgs/cmd/gopkgs
github.com/davidrjenni/reftools/cmd/fillstruct
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


