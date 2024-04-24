# Copy json to go source code for them to be
# hardcoded in the binary
DATA_FILE=cli/quotes.go
rm -f $DATA_FILE
echo "package main" >> $DATA_FILE
echo "const quotesJSON = \`" >> $DATA_FILE
cat quotes.json >> $DATA_FILE
echo "\`" >> $DATA_FILE

# build the go
cd cli
go build

# Install
sudo mv ./quotes /usr/local/bin