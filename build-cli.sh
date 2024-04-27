# Copy json to go source code for them to be
# hardcoded in the binary
node ./build-cli.js

# build the go
cd cli
go fmt
go build

# Install
sudo mv ./quotes /usr/local/bin