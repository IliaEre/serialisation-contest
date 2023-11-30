docker run \
    -v $(pwd):/var/loadtest \
    --net="host" \
    -it direvius/yandex-tank -c save-grpc-load.yml

# clear after, be careful :d
rm -f tank_errors.log
rm -rf logs