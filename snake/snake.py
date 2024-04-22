# imports of types
from io import StringIO, BytesIO
from typing import Optional

# import module __init.py__ for IO adapters
import snake

def SurrogateEscaped(char: str) -> bool:
    """Is the last character an escaped Unicode surrogate error?"""
    if len(char) < 1:
        return False
    val = ord(char[-1])
    return val >= 0xDC80 and val <= 0xDCFF

def HighSurrogate(char: str) -> bool:
    """Is the second to last character a Unicode high surrogate?"""
    if len(char) < 2:
        return False
    val = ord(char[-2])
    return val >= 0xD800 and val <= 0xDBFF

def WriteProxy(buffer: BytesIO, string: str) -> int:
    """Write a string to a buffer. Internal routine that might be useful."""
    # surrogateescape chosen to restore file
    outBytes = string.encode("utf-8", "surrogateescape")    # PEP 383
    length = buffer.write(outBytes)
    if length != len(outBytes):
        # return actual number written (plus maybe bad terminal chars)
        string = outBytes[0:length].decode("utf-8", "surrogateescape")
        while SurrogateEscaped(string) and not HighSurrogate(string):
            chop = len(string) - 1
            string = string[0:chop]
    return len(string) # maps back to codepoints written 

# stdio redirect logic to goali streams
class ByteOut(BytesIO):
    """Internal adapter."""
    def write(self, outBytes) -> int:
        return snake.Out(bytes(outBytes))
    
class StdOut(StringIO):
    """Standard output replacement stream."""
    def write(self, string: str) -> int:
        return WriteProxy(self.buffer, string)
    buffer = ByteOut()

class ByteErr(BytesIO):
    """Internal adapter."""
    def write(self, errBytes) -> int:
        return snake.Err(bytes(errBytes))

class StdErr(StringIO):
    """Standard error replacement stream."""
    def write(self, string: str) -> int:
        return WriteProxy(self.buffer, string)
    buffer = ByteErr()

class ByteIn(BytesIO):
    """Internal adapter."""
    def read(self, size: Optional[int] = -1) -> bytes:
        return snake.In(size)

class StdIn(StringIO):
    """Standard input replacement stream."""
    def read(self, size: Optional[int] = -1) -> str:
        inBytes = self.buffer.read(size)
        # surrogateescape chosen to mark all errors yet load
        # CESU-8 is treated as an error to escape
        # as it would not restore given hard line on surrogate writing
        # so technically 6 errors at 3 bytes each
        # or of length 18 bytes if wrote out in UTF-8
        # Some earlier version of python may have loaded
        # the pair "correctly" circa some date
        string = inBytes.decode("utf-8", "surrogateescape")
        if size != -1:
            inBytes = b"" # something extra
            assert size is not None     # apparently calms mypy
            while len(string) < size:
                while SurrogateEscaped(string) and not HighSurrogate(string):
                    inBytes += (ord(string[-1]) & 0xFF).to_bytes(1, "little")  # mask error char
                    chop = len(string) - 1
                    string = string[0:chop]
                inBytes += self.buffer.read(1)
                extra = inBytes.decode("utf-8", "surrogateescape")
                if not (SurrogateEscaped(extra) and not HighSurrogate(extra)):
                    string += extra
                    inBytes = b""   # re-loop
        return string
    buffer = ByteIn()

# things to migrate into goali from python

# extra utility things

def LowSurrogate(char: str) -> bool:
    """Is the last character a Unicode low surrogate?"""
    if len(char) < 1:
        return False
    val = ord(char[-1])
    return val >= 0xDC00 and val <= 0xDFFF

PairLength = 2 * 3 # hi and lo * bytes to be marked with lo error codepoints

def SurrogateCode(pair: str) -> str:
    """Return a single codepoint for a surrogate pair."""
    return chr(((ord(pair[0]) << 10) + (ord(pair[1]) & 1023) + 65536) & 0x1FFFFF)

def FixSurrogatePairs(string: str) -> str:
    """The Unicode UTF-8 standard marks surrogates as errors. Fix them with this after importing."""
    i = 0
    while len(string) - PairLength >= i:
        text = string[i:i + PairLength].encode("utf-8", "surrogateescape")
        try:
            pair = text.decode("utf-8", "surrogatepass")[0:2]
        except:
            i += 1
            continue
        if HighSurrogate(pair) and LowSurrogate(pair):
            # cool got a surrogate pair
            # being like python, flip this straight to a single
            # would be most useful
            c = SurrogateCode(pair)
            string = string[0:i] + c + string[i + PairLength:]
            i += 1
    return string

def CollapseSurrogatePairs(string: str) -> str:
    """A Unicode string in python does not need to use surrogates. If they exist fix them with this."""
    i = 0
    while len(string) - 2 >= i:
        pair = string[0:2]
        if HighSurrogate(pair) and LowSurrogate(pair):
            # cool got a surrogate pair
            # being like python, flip this straight to a single
            # would be most useful
            c = SurrogateCode(pair)
            string = string[0:i] + c + string[i + 2:]
            i += 1
    return string