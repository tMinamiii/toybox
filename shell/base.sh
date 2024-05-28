#!/bin/sh

set -e

usage() {
    echo "usage `basename $0` [options] <image-uri>"
    echo "options:"
    echo "  -e <enviroment-name>"
    echo "     enviroment name like above:"
    echo "     - development (default)"
    echo "     - production"
    exit 1
}

args=`getopt e:c: $*`

echo "args => $args"

if [ "$?" != 0 ]; then
    usage
fi
set -- $args

env="development"
for opt do
    echo "args? 0, 1, 2=> $0, $1, $2, $3, $4"
    case "$opt" in
        -e)
        env=$2; shift; shift;;
        --)
        shift; break;;
    esac
done

echo $#
echo $env
