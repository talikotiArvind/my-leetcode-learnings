def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
        if len(nums)==0:
            return None
        
        def recrurrsive_call_for_tree(nums,left,right):
            if left>right:
                return None
            mid = (left+right)//2
            root = TreeNode(nums[mid])
            root.left = recrurrsive_call_for_tree(nums,left,mid-1)
            root.right = recrurrsive_call_for_tree(nums,mid+1,right)

            return root
        return recrurrsive_call_for_tree(nums,0,len(nums)-1)
