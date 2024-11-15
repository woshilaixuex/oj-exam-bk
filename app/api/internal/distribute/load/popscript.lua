local poppedValue = redis.call("SPOP", KEYS[1])
if poppedValue then
    local newLength = redis.call("SCARD", KEYS[1])
    return {poppedValue, newLength}
else
    return {nil, 0}
end
