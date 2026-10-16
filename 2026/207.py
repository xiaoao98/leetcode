class Solution:
    def canFinish(self, numCourses: int, prerequisites: list[list[int]]) -> bool:
        def findLoop(k):
            if k in toTake:
                return True
            if k in mem:
                return False
            if k in preDic:
                for ele in preDic[k]:
                    toTake.append(k)
                    if findLoop(ele):
                        return True
                    toTake.pop()
            mem.add(k)
            return False

        preDic = {}
        mem = set()
        toTake = []
        for ele in prerequisites:
            if ele[0] in preDic:
                preDic[ele[0]].append(ele[1])
            else:
                preDic[ele[0]] = [ele[1]]
        for k in preDic.keys():
            if findLoop(k):
                return False
        return True

        
