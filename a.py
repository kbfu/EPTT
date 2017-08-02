import random
import asyncio
from aiohttp import ClientSession
import uvloop


async def fetch(url, session):
    async with session.get(url) as response:
        return await response.read()


async def bound_fetch(sem, url, session):
    # Getter function with semaphore.
    async with sem:
        await fetch(url, session)


async def run(r):
    url = "http://localhost:60001/get"
    tasks = []
    # create instance of Semaphore
    sem = asyncio.Semaphore(200)

    # Create client session that will ensure we dont open new connection
    # per each request.
    async with ClientSession() as session:
        for i in range(r):
            # pass Semaphore and session to every GET request
            task = asyncio.ensure_future(
                bound_fetch(sem, url, session))
            tasks.append(task)

        responses = asyncio.gather(*tasks)
        await responses


def main():
    number = 4000
    asyncio.set_event_loop_policy(uvloop.EventLoopPolicy())
    loop = asyncio.get_event_loop()
    from ept.util import date_util
    start_time = date_util.timestamp_now()
    future = asyncio.ensure_future(run(number))
    loop.run_until_complete(future)
    print(date_util.timestamp_now() - start_time)


if __name__ == '__main__':
    main()
