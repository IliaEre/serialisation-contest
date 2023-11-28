#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import sys
import commonRequest as cr


def generate_json():
    original = open("report.json", "r")
    url = "/report"
    h = cr.headers + "Content-type: application/json"

    ammo = cr.make_ammo(cr.method, url, h, cr.case, original.read())
    sys.stdout.write(ammo)
    f2 = open("json-ammo.txt", "w")
    f2.write(ammo)


if __name__ == "__main__":
    generate_json()

