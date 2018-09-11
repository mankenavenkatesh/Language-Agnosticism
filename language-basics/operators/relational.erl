-module(relational).
-export([start/0]).

start() -> 
   io:fwrite("~w~n",[3==2]), 
   io:fwrite("~w~n",[3/=2]), 
   io:fwrite("~w~n",[3<2]), 
   io:fwrite("~w~n",[3=<2]), 
   io:fwrite("~w~n",[3>2]), 
   io:fwrite("~w~n",[3>=2]).