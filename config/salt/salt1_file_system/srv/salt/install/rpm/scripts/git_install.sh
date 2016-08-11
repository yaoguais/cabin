#!/bin/bash

if [ "$(git --version)" != "git version 1.7.12.4" ]; then
    yum -y remove git
    yum -y install http://mirrors.neusoft.edu.cn/repoforge/redhat/el6/en/x86_64/extras/RPMS/perl-Git-1.7.12.4-1.el6.rfx.x86_64.rpm http://mirrors.neusoft.edu.cn/repoforge/redhat/el6/en/x86_64/extras/RPMS/git-1.7.12.4-1.el6.rfx.x86_64.rpm
fi