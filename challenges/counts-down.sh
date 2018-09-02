read -p "Enter starting value: " startValue
while [[ $startValue >=0 ]]
do
	echo the value is $startValue
	startValue=$(( $startValue - 1 ))
done

