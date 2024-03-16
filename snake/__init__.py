import sys
from typing import Optional

# Clean import of unqualified symbols
# Makes them available at moduleName.snake.x
# from module user code
import snake.snake as snake

# Initialize the stdio once on module loading
sys.stdout = snake.StdOut()

sys.stderr = snake.StdErr()

sys.stdin = snake.StdIn()

# Use these eventual functions to do stdio
# Dynamic language so replacing them good?
def Out(b :bytes) -> int:
    raise NotImplementedError

def Err(b :bytes) -> int:
    raise NotImplementedError

def In(size: Optional[int] = -1) -> bytes:
    raise NotImplementedError