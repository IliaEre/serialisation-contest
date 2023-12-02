echo "build the gun"
pwd
GOOS=linux GOARCH=amd64 go build

echo "run tank"
docker run \
    -v $(pwd):/var/loadtest \
    --net="host" \
    -it direvius/yandex-tank -c validate-tank-load.yml

# clear after, be careful :d
rm -f tank_errors.log
rm -rf logs