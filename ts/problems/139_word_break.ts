/* 139. Word Break */

function wordBreak(s: string, wordDict: string[]): boolean {
    const words = new Set(wordDict);
    const lens = Array.from<boolean>({ length: s.length + 1 }).fill(false);
    lens[0] = true;

    lens.forEach((breakExist, subStrIndex) => {
        if (!breakExist) return;
        let str = '';
        for (let i = subStrIndex; i < s.length; i++) {
            str += s[i];
            if (words.has(str)) lens[i + 1] = true;
        }
    });

    return !!lens.at(-1)
}

// Test
import assert from 'node:assert';

const cases: [string, string[], boolean][] = [
    ['leetcode', ['leet', 'code'], true],
    ['applepenapple', ['apple', 'pen'], true],
    ['catsandog', ['cats', 'dog', 'sand', 'and', 'cat'], false],
    ['aaaaaaa', ['aaaa', 'aaa'], true],
];

cases.forEach(([s, dict, want]) => {
    const got = wordBreak(s, dict);
    assert.equal(got, want, `${s} => [${dict}]. Want: ${want}. Got: ${got}`);
});
