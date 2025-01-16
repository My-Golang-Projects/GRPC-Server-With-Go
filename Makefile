run:
	protoc --go_out=invoicer \
		--go_opt=paths=source_relative \
		--go-grpc_out=invoicer \
		--go-grpc_opt=paths=source_relative \
		bufs/v1/invoice.proto	

# get latest packages with -u flag
get_pkg:
	 go get -u google.golang.org/grpc
