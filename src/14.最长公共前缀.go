/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
        return ""
    }
    ans := len(strs[0])
    for {
        if ans == 0 {
            return ""
        }
        prefix := strs[0][:ans]
        flag := true
        for i := 1; i < len(strs); i++ {
            if len(strs[i]) < ans || strs[i][:ans] != prefix {
                flag = false
                break
            }
        }
        if flag {
            return prefix
        } else {
            ans--
        }
    }
}

func longestCommonPrefix(strs []string) string {
    if len(strs) < 1 {
        return ""
    }
    prefix := strs[0]
    for _,k := range strs {
        for strings.Index(k,prefix) != 0 {
            if len(prefix) == 0 {
                return ""
            }
            prefix = prefix[:len(prefix) - 1]
        }
    }
    return prefix
}

