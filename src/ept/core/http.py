# coding=utf-8
import json

import aiohttp
from aiohttp.connector import TCPConnector
from ..util import csv_util, date_util


async def get(args, verify_ssl=False):
    start = date_util.timestamp_now()
    async with aiohttp.ClientSession(
            connector=TCPConnector(verify_ssl=verify_ssl)) as session:
        try:
            async with session.get(args[0]) as resp:
                resp_text = await resp.text()
                end = date_util.timestamp_now()
                elapsed = end - start
                csv_util.write('log.csv', args[1], resp_text,
                               json.dumps(dict(resp.headers)),
                               resp.status, elapsed,
                               start, end, '')
        except Exception as e:
            end = date_util.timestamp_now()
            elapsed = end - start
            csv_util.write('log.csv', args[1], '',
                           json.dumps(dict(resp.headers)),
                           resp.status, elapsed,
                           start, end, e)
