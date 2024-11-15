-- 乐观锁
local strategyModel = redis.call("GET", KEYS[1])
if not strategyModel then
    redis.call("SET", KEYS[1],ARGV[1]..":0")
    return 0
end

local currentValue, currentVersion = strategyModel:match("([^:]+):([^:]+)")
if currentValue ~= ARGV[1] then
    local newVersion = tonumber(currentVersion) + 1
    redis.call("SET", KEYS[1], ARGV[1] .. ":" .. newVersion)
    return 1
else 
    return 0
end