"""The dynamic library loader. Linux specific."""

from ctypes import cdll, c_bool, c_int, c_ssize_t, c_float, c_double, c_char_p, c_void_p
from ctypes import create_string_buffer, byref, POINTER, Structure, CFUNCTYPE
sonake = cdll.LoadLibrary("../goali.so")

# class POINT(Structure):
#   _fields_ = [("x", c_int),
#               ("y", c_int)] # if POINTER(c_int) for pointer byref() automatically done
# # set class._fields_ later for recursive pointer definitions
# # COMPARE_FUNC_CALLBACK = CFUNCTYPE(c_int, POINTER(c_int), POINTER(c_int)) # first arg return
# # callback_fn = COMPARE_FUNC_CALLBACK(fn)

# TenPointsArrayType = POINT * 10

#sonake.say_hello.argtypes = [c_char_p]
#sonake.say_hello.restype = c_char_p
#sonake.say_hello(b"world")
# None is NULL pointer