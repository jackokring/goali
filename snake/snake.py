"""The imported snake python template with mypy stub definitions and any needed wrappers."""
from array import ArrayType as array
from io import StringIO, BytesIO
from typing import Union, Any, Optional
from ctypes import _CData
from pickle import PickleBuffer
from mmap import mmap

ByteAlias = Union[bytes, Union[bytearray, memoryview, array[Any], mmap, _CData, PickleBuffer]]

# stdio redirect logic to goali streams
class ByteOut(BytesIO):
    def write(self, bytes: ByteAlias) -> int:
        return 0
    
class StdOut(StringIO):
    def write(self, string: str) -> int:
        return 0
    buffer = ByteOut()

class ByteErr(BytesIO):
    def write(self, bytes: ByteAlias) -> int:
        return 0

class StdErr(StringIO):
    def write(self, string: str) -> int:
        return 0
    buffer = ByteErr()

class ByteIn(BytesIO):
    def read(self, size: Optional[int] = -1) -> bytes:
        return b""

class StdIn(StringIO):
    def read(self, size: Optional[int] = -1) -> str:
        return ""
    buffer = ByteIn()

# stubs replaced by goali but present for mypy syntax and type checks

# things to migrate into goali from python

# extra utility things

