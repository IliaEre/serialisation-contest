docker run \
    -v $(pwd):/var/loadtest \
    --net="host" \
    -it direvius/yandex-tank -c json/save-load.yml