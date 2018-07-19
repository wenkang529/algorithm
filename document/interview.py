
# /usr/bin/python3


class Solution(object):
    def my_value(self, nums):
        """
        :type nums: List[int]
        :return int
        """
        if len(nums) <= 3:
            return False

        # init
        max_value = nums[1] - nums[0]
        max_value2 = nums[-1] - nums[-2]
        dict = {}
        result = nums[1] - nums[0] + nums[-1] - nums[-2]

        # save value to dict
        tem2 = nums[-1]
        for i in range(len(nums) - 1, 1, -1):
            if tem2 - nums[i] > max_value2:
                max_value2 = nums[i] - tem2
            if nums[i] > tem2:
                tem2 = nums[i]
            dict[i] = max_value2

        # find result
        tem = nums[0]
        for i in range(2, len(nums) - 2):
            if nums[i] - tem > max_value:
                max_value = nums[i] - tem
            if nums[i] < tem:
                tem = nums[i]
            if max_value + dict[i + 1] > result:
                result = max_value + dict[i + 1]

        return result


int find_max(int nums[], int len)
{
    if (len < 2){
        return False}

    int max = nums[1] - nums[0];
    int min = nums[1] > nums[0] ? nums[0]: nums[1];
    for(int i=2; i < len; i + +)
    {
        if(nums[i] - min > max)
        {
            max = nums[i] - min; }

        if(nums[i] < min)
        {
            min = nums[i]; }
    }
    return max;}
