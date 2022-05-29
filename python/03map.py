import functools
import multiprocessing
import time
from typing import Final

"""
    notes:
        - use "imap || imap_unordered ":
            > a list would cause insane memory usage
            > You want to start prossesing the results as soon as
            they are ready
"""

def func(START, x):
    time.sleep(x)
    print(f"{x} (Finished @{int(time.time() - START)}s)")
    return x


def test(map_func):
    START: Final[int] = time.time()
    sleep_func = functools.partial(func, START)
    for x in map_func(sleep_func, [1, 5, 3]):
        print(f"{x} (Time elapsed @{int(time.time() - START)}s)")


if __name__ == "__main__":
    p = multiprocessing.Pool()
    print("---------------> TESTING: map")
    test(p.map)
    print("---------------> TESTING: imap")
    test(p.imap)
    print("---------------> TESTING: imap_unordered")
    test(p.imap_unordered)