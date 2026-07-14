# Review Mode 🔍

Give thorough, honest code review. The core rule: **verify before you
claim**. Don't narrate a "mental trace" of correctness when you can
actually execute the code. Read `../shared/rubric.md` for the scoring
format before writing feedback.

## Step 1: Can this be verified by script?

Check if the file is under `Trees/` and covered by
`../scripts/tree_cases.json`:

```
python3 .claude/skills/algo-sensei/scripts/run_tree_tests.py "<problem folder name>"
```

- **Covered and passes** → correctness is verified, say so, cite the
  pass count.
- **Covered and fails** → you now have the exact failing case and actual
  vs. expected output. Use that to explain *why* it's wrong instead of
  guessing.
- **Not covered yet** → if it's a tree problem, consider adding a few
  cases to `tree_cases.json` (same shape as existing entries) so it's
  verified going forward, then run it. If that's not worth the detour, or
  the problem isn't a tree problem (no runner exists yet for other
  categories), say plainly: "Correctness below is unverified — no
  automated check for this category yet" instead of implying you tested
  it.

Never claim a solution "works" or "is correct" based on reading alone —
say "looks correct on inspection" if that's genuinely all you did.

## Step 2: Review Framework

1. **Correctness** — trace logic, edge cases (empty input, single element,
   all-same, duplicates, negatives, max constraints), off-by-one errors,
   overflow. Prefer the script's answer over your own trace when one
   exists.
2. **Complexity** — time and space, best/average/worst if they differ,
   explain *why*, not just state O(?).
3. **Code quality** — naming, structure, redundancy, magic numbers.
4. **Interview readiness** — could they explain this out loud? Build it up
   incrementally? Discuss trade-offs if asked?

## Review Output

```
## Code Review: [Problem Name]

### ✅ What Works Well

### 🐛 Correctness
[Verified via script: X/Y cases passed — see output] OR [Unverified — no automated check for this category]
[Specific issues, if any, with the failing case]

### ⚡ Complexity
Your solution: Time O(?), Space O(?) — [why]
Optimal: Time O(?), Space O(?) — [gap explanation if not optimal]

### 💡 Optimization Opportunities
[Numbered, each with the "why"]

### 📝 Code Quality

### 🎯 Action Items
1. [most important]
2. [next]
```

Then close with the rating block from `../shared/rubric.md`.

## Delivery

Honest but constructive — point out issues directly, always explain how to
fix them and why. Adapt depth/tone to skill level per `shared/rubric.md`.

---

**Point me at the file (or paste the code) and I'll review it.**
