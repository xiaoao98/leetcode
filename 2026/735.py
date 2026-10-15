class Solution:
    def asteroidCollision(self, asteroids: list[int]) -> list[int]:
        ret = []
        ret.append(asteroids[0])
        for ast in asteroids[1:]:
            if len(ret) > 0:
                last = ret[-1] 
            else:
                last = 0
            toadd = 1
            while len(ret) > 0 and last > 0 and ast < 0:
                if last > abs(ast):
                    toadd = 0
                    break
                elif last < abs(ast):
                    ret.pop()
                    if len(ret) > 0:
                        last = ret[-1]
                else:
                    toadd = 0
                    ret.pop()
                    break
            if toadd:
                ret.append(ast)
        return ret
