import json
from ..util import date_util


async def get(sem, session, log_manager, args):
    '''
    args[0]: url
    args[1]: step name
    '''
    start = date_util.timestamp_now()
    try:
        async with sem:
            async with session.get(args[0]) as resp:
                resp_text = await resp.text()
                end = date_util.timestamp_now()
                elapsed = end - start
                log_manager.update({'name': args[1], 'resp': resp_text,
                                    'headers': json.dumps(dict(resp.headers)),
                                    'status': resp.status, 'elapsed': elapsed,
                                    'start': start, 'end': end,
                                    'exception': ''})
    except Exception as exc:
        end = date_util.timestamp_now()
        elapsed = end - start
        log_manager.update({'name': args[1], 'resp': '',
                            'headers': json.dumps(dict(resp.headers)),
                            'status': resp.status, 'elapsed': elapsed,
                            'start': start, 'end': end,
                            'exception': exc})
