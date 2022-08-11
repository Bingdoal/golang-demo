buildStamp=$(date -u '+%Y-%m-%d_%I:%M:%S%p')
commitId=$(git rev-list -1 HEAD)
commitTime=$(git show -s --format=%ci "$commitId")
commitAuthor=$(git --no-pager show -s --format='%an <%ae>' "$commitId")
gitUrl=$(git config --get remote.origin.url)
userName=$(whoami)
hostName=$(hostname)
actuator="github.com/sinhashubham95/go-actuator"

go run -ldflags  \
"-X \"${actuator}.GitCommitID=$commitId\" -X \"${actuator}.GitCommitTime=$commitTime\" -X \"${actuator}.GitCommitAuthor=$commitAuthor\" -X \"${actuator}.BuildStamp=$buildStamp\" -X \"${actuator}.GitURL=$gitUrl\" -X \"${actuator}.Username=$userName\" -X \"${actuator}.HostName=$hostName\"" \
$(ls -1 *.go | grep -v wire.go) $*