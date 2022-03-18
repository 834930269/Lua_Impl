local sum = 0
for i=1,5 do
    if i %2==0 then
        sum = sum + i
    end
end

--[[  指令集合
main <@lua_code/sum.lua:0,0> (11 instructions)
0+ params, 6 slots, 1 upvalues,5 locals, 4 constants, 0 functions
        1       [1]     0x00000001      LOADK           0 -1     
        2       [2]     0x00004041      LOADK           1 -2     
        3       [2]     0x00008081      LOADK           2 -3
        4       [2]     0x000040C1      LOADK           3 -2
        5       [2]     0x8000C068      FORPREP         1 4
        6       [3]     0x0240C150      MOD             5 4 -4
        7       [3]     0x02C0001F      EQ              0 5 -1
        8       [3]     0x8000001E      JMP             0 1
        9       [4]     0x0001000D      ADD             0 0 4
        10      [2]     0x7FFE8067      FORLOOP         1 -5
        11      [6]     0x00800026      RETURN          0 1
constants (4):
        1       0
        2       1
        3       5
        4       2
locals (5):
        0       sum     2       12
        1       (for index)     5       11
        2       (for limit)     5       11
        3       (for step)      5       11
        4       i       6       10
upvalues (1):
        0       _ENV    1       0

Lua字节流指令执行测试
[01] LOADK    [0][nil][nil][nil][nil][nil]
[02] LOADK    [0][1][nil][nil][nil][nil]
[03] LOADK    [0][1][5][nil][nil][nil]
[04] LOADK    [0][1][5][1][nil][nil]
[05] FORPREP  [0][0][5][1][nil][nil]
[10] FORLOOP  [0][1][5][1][1][nil]
[06] MOD      [0][1][5][1][1][1]
[07] EQ       [0][1][5][1][1][1]
[08] JMP      [0][1][5][1][1][1]
[10] FORLOOP  [0][2][5][1][2][1]
[06] MOD      [0][2][5][1][2][0]
[07] EQ       [0][2][5][1][2][0]
[09] ADD      [2][2][5][1][2][0]
[10] FORLOOP  [2][3][5][1][3][0]
[06] MOD      [2][3][5][1][3][1]
[07] EQ       [2][3][5][1][3][1]
[08] JMP      [2][3][5][1][3][1]
[10] FORLOOP  [2][4][5][1][4][1]
[06] MOD      [2][4][5][1][4][0]
[07] EQ       [2][4][5][1][4][0]
[09] ADD      [6][4][5][1][4][0]
[10] FORLOOP  [6][5][5][1][5][0]
[06] MOD      [6][5][5][1][5][1]
[07] EQ       [6][5][5][1][5][1]
[08] JMP      [6][5][5][1][5][1]
[10] FORLOOP  [6][6][5][1][5][1]
]]--