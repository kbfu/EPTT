from multiprocessing import Pool
import aiohttp
from datetime import datetime
import uvloop
import asyncio


async def async_send():
    async with aiohttp.ClientSession() as session:
        async with session.get('http://www.chinaso.com') as resp:
            # await asyncio.sleep(1)
            try:
                await resp.text()
                assert resp.status == 200
            except AssertionError as e:
                print(e)


def run(n):
    asyncio.set_event_loop_policy(uvloop.EventLoopPolicy())
    loop = asyncio.get_event_loop()
    print('async')
    now = datetime.timestamp(datetime.now())
    loop.run_until_complete(asyncio.gather(
        *[async_send() for _ in range(n)]
    ))
    done = datetime.timestamp(datetime.now()) - now
    print(done)


if __name__ == '__main__':
    with Pool(processes=8) as pool:
        results = [pool.apply_async(run, (500, )) for _ in range(8)]
        for r in results:
            r.wait()
