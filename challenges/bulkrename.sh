#!/usr/bin/env bash
printHelp() {
    echo "Usage $0 -f find -r replace FILES_TO_REPLACE*"
    echo -e "\t-f The text to find the filename"
    echo -e "\t-r The replacement for the new filename"
    exit 1
}

while getopts "f:r:" opt
do
    case "$opt" in
        r ) replace="$OPTARG" ;;
        f ) match="$OPTARG"   ;;
        ? ) printHelp
    esac
done

shift $(( $OPTIND - 1 ))

if [ -z $replace ] || [ -z $match ]
then
    printf "You need to supply a string to fine and a sting to replace\n";
    printHelp
fi

for i in $@
do
    newname=$(echo $i | sed "s/$match/$replace/")
    mv $i $newname && printf "Renamed file $i to $newname\n"
done
