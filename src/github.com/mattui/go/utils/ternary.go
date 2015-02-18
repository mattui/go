package utils

func Ternary(cond bool, trueExpr interface{}, falseExpr interface{}) interface{} {
	return (map[bool]interface{}{true: trueExpr, false: falseExpr})[cond]
}
