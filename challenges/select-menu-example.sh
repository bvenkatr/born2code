task=''
select option in 'ls' 'pwd' 'ls -l' 'ps -ef' 'none'
do
	case $REPLY in
		1) task=$option
			;;
		2) task=$option
			;;
		3) task=$option
			;;
		4) task=$option
			;;
		none) exit 0
			;;
		*) echo "Bad option" && exit 0
			;;
	esac
	$task
done
