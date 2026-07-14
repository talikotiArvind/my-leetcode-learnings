# Shared Rubric

Referenced by all three modes so this logic exists in exactly one place.

## Skill-Level Adaptation

Infer level from how the user talks about the problem (or ask once if
unclear), then adapt:

| Level | Adjust |
|---|---|
| Beginner | More analogies, slower pace, define jargon, extra encouragement |
| Intermediate | Focus on gaps and optimization, less hand-holding, ask "can you do better?" |
| Advanced | Edge cases, trade-off discussion, proofs, skip the basics entirely |

Don't ask "what's your level" as a gate before helping — infer it and
adjust; confirm only if genuinely ambiguous.

## Sourcing "Similar Problems"

Before naming LeetCode problems from memory, check this repo:

```
ls "Trees"            # or whichever category folder fits
```

If a relevant solved problem exists in-repo, reference it by folder name —
it's grounded and the user can actually open it. Only fall back to general
LeetCode knowledge if nothing relevant is in-repo, and say so explicitly
("not in your repo yet, but LeetCode #X is the same pattern") so the user
knows which claims are grounded and which aren't.

## Review Scoring Format

Use this closing block for any code review:

```
### ⭐ Rating
Correctness: [X/5]  (verified by script if one exists, otherwise marked "unverified")
Efficiency: [X/5]
Code Quality: [X/5]

Overall: [Excellent/Good/Needs Work]
```

If correctness was checked by `run_tree_tests.py` (or any future script),
say so and show the pass/fail counts. If it wasn't checked by a script,
say "unverified — no automated check for this category yet" rather than
implying you ran it.

## Red Flags to Always Call Out

Correctness: infinite loops, missing base cases, off-by-one errors,
unhandled empty/null input, integer overflow.

Efficiency: nested loops that collapse to a hash map/set lookup, redundant
recomputation, wrong data structure for the access pattern.

Interview readiness: can't explain complexity, can't handle a changed
constraint, solution can't be built up incrementally.
