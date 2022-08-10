/* 
https://leetcode.com/problems/two-sum/

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

function twoSum(nums: number[], target: number): number[] {
    const numMap: Record<number, number> = {};

    nums.forEach((n) => (numMap[n] = (numMap[n] || 0) + 1));

    for (let i = 0; i < nums.length; i++) {
        const n1 = nums[i];
        const n2 = target - n1;
        numMap[n1]--;
        if (numMap[n2] > 0) return [i, nums.findIndex((n, idx) => idx != i && n == n2)];
        continue;
    }

    return [0, 0];
}

// Tests
import assert from 'node:assert';

assert.deepEqual(twoSum([2, 7, 11, 15], 9), [0, 1]);
assert.deepEqual(twoSum([11, 15, 2, 7], 9), [2, 3]);
assert.deepEqual(twoSum([2, 5, 5, 11], 10), [1, 2]);
assert.deepEqual(twoSum([2, 5, 11, 7, 3], 10), [3, 4]);
assert.deepEqual(twoSum([3, 2, 4], 6), [1, 2]);
