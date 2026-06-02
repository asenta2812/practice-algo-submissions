class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []
        operators = {
            "+": lambda a,b: a+b,
            "-": lambda a,b: a-b,
            "*": lambda a,b: a*b,
            "/": lambda a,b: int(a/b)
        }
        for t in tokens:
            if t in operators:
                r, l = stack.pop(), stack.pop()
                stack.append(operators[t](l, r))
            else:
                stack.append(int(t))
           
        return stack[0]
