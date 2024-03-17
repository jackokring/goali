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
    # surrogateescape chosen to restore file
    outBytes = string.encode("utf-8", "surrogateescape")    # PEP 383
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
        # surrogateescape chosen to mark all errors yet load
        # CESU-8 is treated as an error to escape
        # as it would not restore given hard line on surrogate writing
        # so technically 6 errors at 3 bytes each
        # or of length 18.
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

# things to migrate into goali from python

# extra utility things

def LittleSurrogate(char: str) -> bool:
    if len(char) < 1:
        return False
    val = ord(char[-1])
    return val >= 0xDC00 and val <= 0xDFFF

def RestoreSurrogatePairs(string: str) -> str:
    i = 0
    while len(string) - 12 > i:
        text = string[i:i + 12].encode("utf-8", "surrogateescape")
        try:
            pair = text.decode("utf-8", "surrogatepass")
        except:
            i += 1
            continue
        if BigSurrogate(pair[0:2]) and LittleSurrogate(pair[0:2]):
            # cool got a surrogate pair
            # being like python, flip this straight to a single
            # would be most useful
            c = chr(((ord(pair[0]) << 10) + (ord(pair[1]) & 1023) + 65536) & 0x1FFFFF)
            string = string[0:i] + c + string[i + 12:]
            i += 1
    return string
