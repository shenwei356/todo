#!/usr/bin/env sh

CGO_ENABLED=0 gox -os="windows darwin linux" -arch="amd64" -tags netgo -ldflags '-w -s'

dir=binaries
mkdir -p $dir;
rm -rf $dir/*;

d=todo
for f in todo_*; do
    mkdir -p $dir/$f/$d;
    mv $f $dir/$f/$d;
    cp -r html LICENSE $dir/$f/$d;
    cd $dir/$f;
    mv $d/$f $d/$(echo $f | perl -pe 's/_[^\.]+//g');
    tar -zcf $f.tar.gz $d;
    mv *.tar.gz ../;
    cd ..;
    rm -rf $f;
    cd ..;
done;
