#!/usr/bin/python3
# -*- coding: utf-8 -*-

import requests

host = "localhost"
port = 9093
save = "/report"
find = "/reports"

commonHeaders = {
    'Host': 'localhost',
    'Content-type': 'application/octet-stream',
    'User-Agent': 'tank'
}


def print_request(request):
    req = "{method} {path_url} HTTP/1.1\r\n{headers}\r\n{body}".format(
        method=request.method,
        path_url=request.path_url,
        headers=''.join('{0}: {1}\r\n'.format(k, v) for k, v in request.headers.items()),
        body=request.body or "",
    )
    return "{req_size}\n{req}\r\n".format(req_size=len(req), req=req)


def post_save(namespace, payload):
    req = requests.Request(
        'POST',
        'https://{host}:{port}{namespace}'.format(
            host=host,
            port=port,
            namespace=namespace,
        ),
        headers=commonHeaders,
        data=payload,
    )

    prepared = req.prepare()
    return print_request(prepared)


def generate_save_flat():
    with open("flat/save.bin", "rb") as f:
        payload = f.read()
    req = post_save(save, payload)

    f2 = open("flat/flat-save-ammo.txt", "w")
    f2.write(req)


if __name__ == "__main__":
    generate_save_flat()
