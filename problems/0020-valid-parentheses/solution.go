func isValid(s string) bool {
    
    var stack []rune

    for _, ch := range s {
        switch ch {
        case '(', '[', '{':
            stack = append(stack, ch)

        case ')', ']', '}':
            if len(stack) == 0 {
                return false
            }

            last := stack[len(stack)-1]

            if (ch == ')' && last != '(') ||
                (ch == ']' && last != '[') ||
                (ch == '}' && last != '{') {
                return false
            }

            stack = stack[0:len(stack)-1]
        }
    }

    if len(stack) == 0 {
        return true
    }

    return false
}
