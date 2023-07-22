
'''
	模板：
	回溯法：穷举所有可能
'''
result = []
func backtrack(选择, 路径):
    if 满足结束条件:
        result.add(路径)
        return
    for 选择 in 选择:
        做选择
        backtrack(选择, 路径)
        撤销选择
