# Separate package level solution
func isMirror(p *TreeNode, q *TreeNode) bool {
  if p == nil && q == nil {
    return true
  }
  if p == nil || q == nil {
    return false
  }
  return p.Value == q.Value && isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
}

func isSymmetricTree(root *TreeNode) bool {
  return isMirror(root.Left, root.Right)
}

#Nested solution

func isSymmetric(root *TreeNode) bool {
    var isMirror func(n1, n2 *TreeNode) bool
    isMirror = func(n1, n2 *TreeNode) bool {
        if n1 == nil && n2 == nil {
            return true
        }
        if n1 == nil || n2 == nil {
            return false
        }
        return n1.Val == n2.Val &&
            isMirror(n1.Left, n2.Right) &&
            isMirror(n1.Right, n2.Left)
    }
    return isMirror(root.Left, root.Right)
}
