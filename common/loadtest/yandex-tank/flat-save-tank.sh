docker run \
    -v $(pwd):/var/loadtest \
    --net="host" \
    -it direvius/yandex-tank -c flat/flat-save-load.yml

# clear after, be careful :d
rm -f tank_errors.log
rm -rf logs