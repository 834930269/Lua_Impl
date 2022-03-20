local function test(a,b,c)
    a = a + 2
    b = 3+a
    c = 4+b
    return a,b,c
end

local function g()
    return 4,5
end

local c,d,f = test(12,g())

--[[
call @lua_code/lua_function.lua<0,0>
preCall [function][function][function][12][function]
call @lua_code/lua_function.lua<8,10>
postCall [function][function][function][12][function][4][5][5]
preCall [function][function][function][12][function][4][5][5]
call @lua_code/lua_function.lua<1,6>
postCall [function][function][14][17][21]
]]--