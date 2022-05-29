from threading import Thread
import time

from multiprocessing.pool import ThreadPool
from multiprocessing.pool import Pool as ProcessPool
from typing import Final


"""
    Without the GIL:
        > Leaked memory
        > incorectly release of memory
            - crashes
            - unpredicted outputs
    # The design decision of the GIL is one of he things that made python as it
    is today. - Larry hastings
"""

def profile(func):
    def wrap_func(*args, **kwargs):
        START: Final = time.time()
        func(*args, **kwargs)
        END: Final = time.time()
        print(f'Function {func.__name__!r} executed in {(END-START)}s')
    return wrap_func


def exec(count):
    while(count>0):
        count -= 1


@profile
def sequential_exec(count):
    exec(count)


@profile
def parallel_exec(PoolType, count):
    # data = [count/2, count/2]
    data = [count/5, count/5, count/5, count/5, count/5]
    with PoolType() as pool:
        pool.map(exec, data)


if __name__ == '__main__':
    COUNT: Final[int] = 50_000_000

    print("---------------> Squential")
    sequential_exec(COUNT)
    
    print("---------------> Threads")
    parallel_exec(ThreadPool, COUNT)
    
    print("---------------> SubProsses")
    parallel_exec(ProcessPool, COUNT)