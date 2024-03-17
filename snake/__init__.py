"""The imported snake python template with mypy stub definitions and any needed wrappers."""
# import typing
from typing import Optional

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