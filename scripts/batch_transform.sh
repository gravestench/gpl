#!/bin/sh

srcGPL=$1
dstGPLdir=$2
outDir=$3

for dstGPL in $(ls ${dstGPLdir}/*.gpl); do
	srcBasename=$(basename $srcGPL)
	dstBasename=$(basename $dstGPL)
	outName=${srcBasename%.gpl}_${dstBasename}
	outPath=${outDir}/${outName}
	echo "transforming $srcBasename into target palette $dstBasename, output gpl into $outDir ..."
	gpl-transform -src $srcGPL -dst $dstGPL -out $outPath &
done
