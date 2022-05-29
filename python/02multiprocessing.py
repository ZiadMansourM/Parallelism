import time
import requests

from multiprocessing.pool import ThreadPool
from multiprocessing.pool import Pool as ProcessPool
from typing import Final

"""
	Sub process:
        - CPU Bound
        <> Spawwing processes is a bit slower than threads
        <> Creating Pools has a certain amount of overhead
"""

def profile(func):
    def wrap_func(*args, **kwargs):
        START: Final = time.time()
        func(*args, **kwargs)
        END: Final = time.time()
        print(f'Function {func.__name__!r} executed in {(END-START)}s')
    return wrap_func


def monitor(site: str):
    with requests.get(site) as r:
        print(
            f"FINE: {site.split('//')[1]}"
            if r.status_code == 200 else
            f"DOWN<{r.status_code}>: {site.split('//')[1]}"
        )


@profile
def sequential_monitoring(sites: list[str]):
    for site in sites:
        monitor(site)


@profile
def parallel_monitoring(PoolType, sites: list[str]):
    with PoolType() as pool:
        pool.map(monitor, sites)


if __name__ == '__main__':
    sites = [
        "http://google.com",
		"http://facebook.com",
		"https://sreboy.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.org",
    ]

    print("---------------> Squential")
    sequential_monitoring(sites)

    print("---------------> Threads")
    parallel_monitoring(ThreadPool, sites)

    print("---------------> SubProsses")
    parallel_monitoring(ProcessPool, sites)