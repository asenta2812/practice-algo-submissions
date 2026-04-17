class Solution:
    def isValid(self, s: str) -> bool:
        m = {'(': ')', '{':'}', '[':']'}
        stack = []

        for c in s:
          if c in m:
            stack.append(c)
          else:
            if stack and m[stack[-1]] == c:
              stack.pop()
            else:
              return False
        return not stack