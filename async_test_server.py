#!/usr/bin/env python
# coding=utf-8
from sanic import Sanic
from sanic.response import json

app = Sanic(__name__)
a = 0


async def test(request):
    globals()['a'] += 1
    return json(a)


def init_route():
    app.add_route(test,
                  '/get', methods=['GET'])


if __name__ == '__main__':
    init_route()
    app.run('0.0.0.0', 60001)
