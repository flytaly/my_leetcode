/*
11. Container With Most Water

https://leetcode.com/problems/container-with-most-water/
*/

function maxArea(height: number[]): number {
    let result = 0;
    let left = 0;
    let right = height.length - 1;

    for (let length = right - left; length > 0; length = right - left) {
        const width = Math.min(height[left], height[right]);
        const area = width * length;
        if (area > result) result = area;

        if (height[left] < height[right]) left++;
        else right--;
    }

    return result;
}

// Tests
import assert from 'node:assert';

const cases: [number[], number][] = [
    [[1, 1], 1],
    [[1, 8, 6, 2, 5, 4, 8, 3, 7], 49],
    [[8, 1, 20, 20, 20, 20], 60],
    [[1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 25],
    [[2, 3, 4, 5, 18, 17, 6], 17],
];

cases.forEach(([height, expected], idx) => {
    const r = maxArea(height);
    assert.equal(r, expected, `${idx}. Want: ${expected}. Got: ${r}`);
});
