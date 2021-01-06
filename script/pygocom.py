#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
            PyGoCompile
Copyright (c) 2021 Ismael Belisario
Compile Go (Golang) source code with Python API
License: MIT
Version: 0.4.0
"""

from sysconfig import get_path, get_config_var as getvar
from subprocess import getstatusoutput as output
import sys

def get_cflags():
    """Get the cflag for compile python source code"""

    flags = ['-I' + get_path('include'),
             '-I' + get_path('platinclude')]
    flags.extend(getvar('CFLAGS').split())

    # Note: Extrat cflags not valid for cgo.
    for not_go in ('-fwrapv', '-Wall'):
        if not_go in flags:
            flags.remove(not_go)

    return ' '.join(flags)


def get_ldflags(embed=True):
    """Get the ldflags for compile python source code
       Args
         embed - Get python lib for embend.
    """

    libs = []
    if embed:
        libs.append('-lpython' + getvar('VERSION') + sys.abiflags)

    else:
        libpython = getvar('LIBPYTHON')
        if libpython:
            libs.append(libpython)

    libs.extend(getvar('LIBS').split() + getvar('SYSLIBS').split())

    if not getvar('Py_ENABLE_SHARED'):
        libs.insert(0, '-L' + getvar('LIBPL'))

    return ' '.join(libs)


def check_go():
    """Check is Go is installed in your device"""
    return 0 if output('go version')[0] == 0 else 1


def go_compile(source='', tarjet=''):
    """Compile Go (Source) with Python API

       Args
         source - Source file or dir with source files.
         tarjet - Name of the binary file to compile.
    """
    print('Starting compilation of the source code written in Go')

    command = ("""
    %(cmd_var)s CGO_CFLAGS="%(cflags)s"
    %(cmd_var)s CGO_LDFLAGS="%(ldflags)s"
    go build%(source)s%(output)s
    """ % {
        'cmd_var': 'set' if sys.platform == 'win32' else 'export',
        'cflags': get_cflags(),
        'ldflags': get_ldflags(),
        'source': ' ' + source,
        'output': f' -o {tarjet}' if tarjet else ''
    }).replace('\n    ', ';')[1:-1]
    print()
    return output(command)


def parser():
    """Parser Argumens for run App"""
    import argparse

    parser = argparse.ArgumentParser(
        description='Compilator of Go with CPython')
    parser.add_argument('-v', '--verbose', action='store_true',
                        help='Explain what is being done')
    parser.add_argument('-o', '--output', default='',
                        action='store', help='File of output')
    parser.add_argument('-s', '--source', default='.', action='store',
                        help='The file or dir to build')
    return parser.parse_args()


def main():
    if check_go() == 1:

        print("""
Go (Golang) has not been found. To download Go visit
https://go.dev, to contact support, send email to 
ismaelbeli.com@gmail.com.
""")
        exit(1)

    par = parser()
    status = go_compile(par.source, par.output)
    print(status[1]) if par.verbose else ...
    if status[0] == 0:
        print('Successfully completed compilation')

    else:
        print('Error when compiling error code %s.' % status[0])
        exit(1)


if __name__ == '__main__':
    main()