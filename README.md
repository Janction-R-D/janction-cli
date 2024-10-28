POST /api/v1/job/status
{
    job_id: 12345  int,
    status:  1:running 2:finish 3:faild
}


# GOARCH=arm GOOS=linux go build -v -ldflags "-w -s" -o janction-node .
# GOARCH=amd64 GOOS=linux go build -v -ldflags "-w -s" -o janction-node .
# GOARCH=arm go build -v -ldflags "-w -s" -o janction-node .
# GOARCH=amd64 go build -v -ldflags "-w -s" -o janction-node .