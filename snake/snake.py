"""The imported snake python template with mypy stub definitions and any needed wrappers."""
from array import ArrayType as array
from io import StringIO, BytesIO
from typing import Union, Any, Optional
from ctypes import _CData
from pickle import PickleBuffer
from mmap import mmap
import snake

ByteAlias = Union[bytes, Union[bytearray, memoryview, array[Any], mmap, _CData, PickleBuffer]]

def SurrogateEscaped(char: str) -> bool:
    if len(char) < 1:
        return False
    val = ord(char[-1])
    return val >= 0xDC80 and val <= 0xDCFF

def BigSurrogate(char: str) -> bool:
    if len(char) < 2:
        return False
    val = ord(char[-2])
    return val >= 0xD800 and val <= 0xDBFF

def WriteProxy(buffer: BytesIO, string: str) -> int:
    outBytes = string.encode()
    length = buffer.write(outBytes)
    if length != len(outBytes):
        # return actual number written (plus maybe bad terminal chars)
        string = outBytes[0:length].decode("utf-8", "surrogateescape")
        while SurrogateEscaped(string) and not BigSurrogate(string):
            chop = len(string) - 1
            string = string[0:chop]
    return len(string) # maps back to codepoints written 

# stdio redirect logic to goali streams
class ByteOut(BytesIO):
    def write(self, outBytes: ByteAlias) -> int:
        return snake.Out(bytes(outBytes))
    
class StdOut(StringIO):
    def write(self, string: str) -> int:
        return WriteProxy(self.buffer, string)
    buffer = ByteOut()

class ByteErr(BytesIO):
    def write(self, errBytes: ByteAlias) -> int:
        return snake.Err(bytes(errBytes))

class StdErr(StringIO):
    def write(self, string: str) -> int:
        return WriteProxy(self.buffer, string)
    buffer = ByteErr()

class ByteIn(BytesIO):
    def read(self, size: Optional[int] = -1) -> bytes:
        return snake.In(size)

class StdIn(StringIO):
    def read(self, size: Optional[int] = -1) -> str:
        inBytes = self.buffer.read(size)
        string = inBytes.decode("utf-8", "surrogateescape")
        if size != -1:
            inBytes = b"" # something extra
            assert size is not None     # apparently calms mypy
            while len(string) < size:
                while SurrogateEscaped(string) and not BigSurrogate(string):
                    inBytes += (ord(string[-1]) & 0xFF).to_bytes(1, "little")  # mask error char
                    chop = len(string) - 1
                    string = string[0:chop]
                inBytes += self.buffer.read(1)
                extra = inBytes.decode("utf-8", "surrogateescape")
                if not (SurrogateEscaped(extra) and not BigSurrogate(extra)):
                    string += extra
                    inBytes = b""   # re-loop
        return string
    buffer = ByteIn()

# stubs replaced by goali but present for mypy syntax and type checks

# things to migrate into goali from python

# extra utility things

