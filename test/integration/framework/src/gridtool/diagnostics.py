import datetime
from typing import Tuple
from gridtool.common import *
from gridtool import cosmos, gridchain


def get_block_times(gridnoded: gridchain.Gridnoded, first_block: int, last_block: int) -> List[Tuple[int, datetime.datetime]]:
    result = [(block, cosmos.parse_iso_timestamp(gridnoded.query_block(block)["block"]["header"]["time"]))
        for block in range(first_block, last_block)]
    return result
