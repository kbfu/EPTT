# coding=utf-8
import json
import aiohttp
from aiohttp.connector import TCPConnector
from ..util import date_util


async def get(log_manager, args, verify_ssl=False):
    '''
    args[0]: url
    args[1]: step name
    '''
    start = date_util.timestamp_now()
    async with aiohttp.ClientSession(
            connector=TCPConnector(verify_ssl=verify_ssl)) as session:
        try:
            async with session.get(args[0]) as resp:
                resp_text = await resp.text()
                end = date_util.timestamp_now()
                elapsed = end - start
                log_manager.update({'name': args[1], 'resp': resp_text,
                                    'headers': json.dumps(dict(resp.headers)),
                                    'status': resp.status, 'elapsed': elapsed,
                                    'start': start, 'end': end,
                                    'exception': ''})
        except Exception as e:
            end = date_util.timestamp_now()
            elapsed = end - start
            log_manager.update({'name': args[1], 'resp': '',
                                'headers': json.dumps(dict(resp.headers)),
                                'status': resp.status, 'elapsed': elapsed,
                                'start': start, 'end': end,
                                'exception': e})
