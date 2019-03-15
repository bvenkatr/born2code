current_branch=$(git rev-parse --abbrev-ref HEAD)
username=""
password=""
title=""

usage() {
    echo "Usage: open-pr.sh -u username -p password -t title -d date for the pr"
    exit 1
}

while getopts ':u:p:t:' opt; do
    case "$opt" in
        u) username="$OPTARG";;
        p) password="$OPTARG";;
        t) title="$OPTARG";;
        h)
            usage
        ;;
        \?)
            echo "Invalid option $OPTARG" >&2
            usage >&2
            exit 1
        ;;
        :)
            echo "Option -$OPTARG requires an argument." >&2
        ;;
        *)
            usage
        ;;
    esac
done

shift $(( OPTIND -1 ))

check_in_set() {
    if [[ -z "$2" ]]; then
        echo "$1 must be set"
        exit 1
    fi
}

if [[ "$current_branch" == "master" ]]; then
    echo "You are on master branch, create a new branch and generate a PR"
fi

check_in_set "username" username
check_in_set "password" password
check_in_set "title" title

data=$(cat <<-EOF
{
    "title": "title",
    "base": "master",
    "head": "test",
    "body": "open pr script"
}
EOF
)
echo curl -s --user "$username:$password" -X POST "https://api.github.com/repos/bvenkatr/born2code/pulls" -d "${data}" -w %{http_code} -o /dev/null
status_code=$(curl -s --user "$username:$password" -X POST "https://api.github.com/repos/bvenkatr/born2code/pulls" -d "$data" -w %{http_code} -o /dev/null)

if [[ $status_code == "201" ]]; then
    echo "Complete"
else
    echo "Error occurred, $status_code status received" >&2
    exit 1
fi
