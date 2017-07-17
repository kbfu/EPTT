from multiprocessing import Pool
import aiohttp
from datetime import datetime
import uvloop
import asyncio


async def async_send():
    async with aiohttp.ClientSession() as session:
        try:
            async with session.get('http://127.0.0.1:60001/get') as resp:
                await resp.text()
                try:
                    assert resp.status == 200
                except AssertionError:
                    print(resp.status)
        except Exception as e:
            print(e)


def run(n):
    asyncio.set_event_loop_policy(uvloop.EventLoopPolicy())
    loop = asyncio.get_event_loop()
    loop.run_until_complete(asyncio.gather(
        *[async_send() for _ in range(n)]
    ))


if __name__ == '__main__':
    print('double process with uvloop')
    now = datetime.timestamp(datetime.now())
    with Pool(processes=2) as pool:
        results = [pool.apply_async(run, (50, )) for _ in range(80)]
        for r in results:
            r.wait()
    done = datetime.timestamp(datetime.now()) - now
    print(done)
