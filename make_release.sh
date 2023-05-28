#!/bin/bash

version="v1.0.0"
appname="ssh-portgen"
build_dir="build"
release_dir="${build_dir}/releases"
arc_dir="${build_dir}/archives"
unixes=(linux windows darwin freebsd openbsd netbsd)
linux_archs=(ppc64le ppc64 s390x riscv64)

mkdir ./${build_dir}
mkdir ./${release_dir}
mkdir ./${arc_dir}

# Create amd and arm architecture files for all supported OSs
for os in ${unixes[@]}
do
    ext=""
    if [ "$os" = "windows" ]; then
        ext=".exe"
    fi
    env GOOS=${os} GOARCH=amd64 go build -o ${release_dir}/${appname}_${os}_amd64-${version}${ext}
    env GOOS=${os} GOARCH=arm64 go build -o ${release_dir}/${appname}_${os}_arm64-${version}${ext}
done

# Create extra architecture files for linux
for arch in ${linux_archs[@]}
do
    env GOOS=linux GOARCH=${arch} go build -o ${release_dir}/${appname}_linux_${arch}-${version}
done

# Create extra architecture files for the BSDs
env GOOS=freebsd GOARCH=riscv64 go build -o ${release_dir}/${appname}_freebsd_riscv64-${version}
env GOOS=openbsd GOARCH=mips64 go build -o ${release_dir}/${appname}_openbsd_mips64-${version}

cd ./$release_dir
for f in *
do
    echo "Creating archive of ${f}"
    zip -r ../../${arc_dir}/${f}.zip ${f}
done