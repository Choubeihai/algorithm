--lua脚本示例

--1. 访问次数加1
--2. 判断访问次数是否为1
--3. 设置过期时间
local currentcurrent = redis.call("incr",KEYS[1])
if tonumber(current) == 1 then
    redis.call("expire",KEYS[1],60)
end
