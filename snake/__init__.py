"""The imported snake python template with mypy stub definitions and any needed wrappers.

Various imports are also made an some efficiency definitions made for 'from snake import *'
to provide ease of coding."""
# import typing common ADTs
from typing import Optional, Any, List, Dict, Union, Set, Tuple

# part of the PyCFunction all having "self" (the module) rabbit hole
import functools
import itertools
import operator

# dynamic "typing" based dispatch via @overload / @overloads(f_certain)
from snake.overloading import *

# some simplification definitions
# @autoself and then f.register
# might be more compact and "standard" for special case processing
autoself = functools.singledispatch
autometh = functools.singledispatchmethod
# partial application on functions and methods
partial = functools.partial
partmeth = functools.partialmethod
# caching
lru = functools.lru_cache
# total ordering using just __eq__(), and one other like __lt__() (by subtraction -ve number?)
ordered = functools.total_ordering
# decorator wraps (use in decorator to wraps(f) before f use in returned wrapper())
wraps = functools.wraps


# import modules
import sys

# Clean import of unqualified symbols
# Makes them available at moduleName.snake.x
# from module user code
import snake.snake as snake

# selective imports
# Useful for upgrading CESU-8
# Takes the marked errors and makes singular codepoints
from snake.snake import FixSurrogatePairs

# Surrogates already in a string not marked as errors
# to singular codepoints
from snake.snake import CollapseSurrogatePairs

# Initialize the stdio once on module loading
sys.stdout = snake.StdOut()

sys.stderr = snake.StdErr()

sys.stdin = snake.StdIn()

# Use these eventual functions to do stdio
# Dynamic language so replacing them good?
def Out(b: bytes) -> int:
    """Standard output stream stub."""
    raise NotImplementedError

def Err(b: bytes) -> int:
    """Standard error stream stub."""
    raise NotImplementedError

def In(size: Optional[int] = -1) -> bytes:
    """Standard input stream stub."""
    raise NotImplementedError

# Main on file "-"
def main():
    """Main to run when <py-file> == \"-\""""
    return "default main"