#!/bin/sh

set -e

gen() {
    local package=$1

    abigen --bin bin/${package}.bin --abi abi/${package}.abi --pkg=${package} --out=${package}/${package}.go
}

genNoBin() {
    local package=$1

    abigen --abi abi/${package}.abi --pkg=${package} --out=${package}/${package}.go
}

gen proxyadmin
gen erc1967proxy
gen transparentupgradableproxy
