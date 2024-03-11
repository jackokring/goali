import sys

# Clean import of unqualified symbols
from snake.snake import StdOut, StdErr, StdIn

# Initialize the stdio once on module loading
sys.stdout = StdOut()

sys.stderr = StdErr()

sys.stdin = StdIn()
