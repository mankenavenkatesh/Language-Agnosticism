-module(arithmetic). 
-export([start/0]). 

start() -> 
   X = 40, 
   Y = 50, 
   
   Res1 = X + Y, 
   Res2 = X - Y, 
   Res3 = X * Y, 
   Res4 = X / Y, 
   Res5 = X div Y, 
   Res6 = X rem Y, 
   
   io:fwrite("~w~n",[Res1]), 
   io:fwrite("~w~n",[Res2]), 
   io:fwrite("~w~n",[Res3]), 
   io:fwrite("~w~n",[Res4]), 
   io:fwrite("~w~n",[Res5]), 
   io:fwrite("~w~n",[Res6]).