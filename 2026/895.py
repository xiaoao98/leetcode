class FreqStack:
    def __init__(self):
        self.dicFreq = {}
        self.group = {}
        self.maxFreq = 0

    def push(self, val: int) -> None:
        freq = self.dicFreq.get(val, 0) + 1
        self.dicFreq[val] = freq

        if freq not in self.group:
            self.group[freq] = []

        self.group[freq].append(val)

        self.maxFreq = max(self.maxFreq, freq)

    def pop(self) -> int:
        val = self.group[self.maxFreq].pop()

        self.dicFreq[val] -= 1

        if not self.group[self.maxFreq]:
            self.maxFreq -= 1

        return val

