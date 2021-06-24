## Example scripts

`batch_transform.sh` can be used to transform a single gpl into many target gpl palettes.
```bash
# $srcGPL is the source GPL file
#
# $dstDir contains gpl files, the source GPL will be 
# transformed into each GPL within this dir
#
# $outDir is where the transformed gpl files will be written
./batch_transform.sh ${srcGPL} ${dstDir} ${outDir}
```