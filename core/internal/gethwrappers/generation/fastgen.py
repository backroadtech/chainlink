#!/usr/bin/env python3

'''Quick-and-dirty (very quick!) compiler and generator of go contract wrappers

Usage: {fastgen_dir}/fastgen.py [<pkg_name> ...]

DO NOT check in the outputs from this script! Instead, run `go generate` in the
parent directory. We are using belt for compilation of solidity contracts, and
using the abi files it outputs as a single source of truth.

However, this is much faster and more reliable, for actual development of
contracts which interact in intricate ways with go code. Once you're done with
development, be a good citizen before you push and replace the wrappers from
this script with those generated by `go generate` as described above.
(`../go_generate_test.go` will remind you with a CI failure if you forget.)

This requires the correct versions of abigen and the correct version of solc on
your path, which can be installed as described in `../go_generate.go`.

This runs on the FluxAggregator contract in 0.4s, whereas the canonical way,
`yarn && yarn workspace @chainlink/contracts belt compile solc &&
 go generate core/internal/gethwrapper`, takes 10.8s, a 27-fold speedup.

'''

import os, sys

thisdir = os.path.abspath(os.path.dirname(sys.argv[0]))
godir = os.path.dirname(thisdir)
gogenpath = os.path.join(godir, 'go_generate.go')

abigenpath = './generation/generate.sh'

pkg_to_src = {}

for line in open(gogenpath):
    if abigenpath in line:
        abipath, pkgname = line.split(abigenpath)[-1].strip().split()
        srcpath = os.path.abspath(os.path.join(godir, abipath)).replace(
            '/abi/', '/src/').replace('.json', '.sol')
        if not os.path.exists(srcpath):
            srcpath = os.path.join(os.path.dirname(srcpath), 'dev',
                                   os.path.basename(srcpath))
        if not os.path.exists(srcpath):
            srcpath = srcpath.replace('/dev/', '/tests/')
        assert os.path.exists(srcpath), 'could not find ' + \
            os.path.basename(srcpath)
        pkg_to_src[pkgname] = srcpath

args = sys.argv[1:]

if len(args) == 0 or any(p not in pkg_to_src for p in args):
    print(__doc__.format(fastgen_dir=thisdir))
    print("Here is the list of packages you can build. (You can add more by")
    print("updating %s)" % gogenpath)
    print()
    longest = max(len(p) for p in pkg_to_src)
    colwidth = longest + 4
    header = "Package name".ljust(colwidth) + "Contract Source"
    print(header)
    print('-' * len(header))
    for pkgname, contractpath in pkg_to_src.items():
        print(pkgname.ljust(colwidth) + contractpath)
    sys.exit(1)

for pkgname in args:
    solidity_path = pkg_to_src[pkgname]
    outpath = os.path.abspath(os.path.join(godir, 'generated', pkgname,
                                           pkgname + '.go'))
    assert not os.system(
        f'abigen -sol {solidity_path} -pkg {pkgname} -out {outpath}')
