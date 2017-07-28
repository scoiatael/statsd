NAME=statsd
BUILD=build

mkdir -p $BUILD
rm $BUILD/*
for OS in windows linux darwin;
do
    env GOOS=$OS go build -v -o $BUILD/$NAME-$OS
done
