#!/usr/bin/env python3
"""Executes binary-tree solution files in this repo against real test cases.

Solution files in this repo are pasted LeetCode method bodies, not always
wrapped in `class Solution:` and not always taking `self` as the first
parameter. `normalize_solution` repairs that so the file can be exec'd and
called directly, instead of asking a model to "mentally trace" the code.

Usage:
    python3 run_tree_tests.py                 # run every case in tree_cases.json
    python3 run_tree_tests.py "invert"         # only run problems matching this substring
"""
import json
import re
import sys
import textwrap
from pathlib import Path
from typing import Dict, List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def find_repo_root(start: Path) -> Path:
    for parent in [start, *start.parents]:
        if (parent / ".git").exists():
            return parent
    return start.parents[4]


REPO_ROOT = find_repo_root(Path(__file__).resolve())
MANIFEST_PATH = Path(__file__).resolve().parent / "tree_cases.json"

DEF_RE = re.compile(r"^def (\w+)\(([^)]*)\)(\s*->\s*[^:]+)?:", re.MULTILINE)
CLASS_RE = re.compile(r"^\s*class\s+Solution\b", re.MULTILINE)


def build_tree(values: list) -> Optional[TreeNode]:
    """LeetCode level-order array (None = missing child) -> TreeNode."""
    values = list(values)
    if not values or values[0] is None:
        return None
    it = iter(values)
    root = TreeNode(next(it))
    queue = [root]
    while queue:
        node = queue.pop(0)
        try:
            left_val = next(it)
        except StopIteration:
            break
        if left_val is not None:
            node.left = TreeNode(left_val)
            queue.append(node.left)
        try:
            right_val = next(it)
        except StopIteration:
            break
        if right_val is not None:
            node.right = TreeNode(right_val)
            queue.append(node.right)
    return root


def serialize_tree(node: Optional[TreeNode]) -> list:
    """TreeNode -> LeetCode level-order array, trimmed of trailing Nones."""
    if node is None:
        return []
    result = []
    queue = [node]
    while queue:
        curr = queue.pop(0)
        if curr is None:
            result.append(None)
            continue
        result.append(curr.val)
        queue.append(curr.left)
        queue.append(curr.right)
    while result and result[-1] is None:
        result.pop()
    return result


def normalize_solution(code: str) -> str:
    """Wrap a bare method body in `class Solution` and add `self` if missing."""
    if CLASS_RE.search(code):
        return code

    def add_self(match: re.Match) -> str:
        name, params, ret = match.group(1), match.group(2).strip(), match.group(3) or ""
        if params.startswith("self"):
            return match.group(0)
        sep = ", " if params else ""
        return f"def {name}(self{sep}{params}){ret}:"

    code = DEF_RE.sub(add_self, code)
    return "class Solution:\n" + textwrap.indent(code, "    ")


def load_solution(path: Path):
    code = normalize_solution(path.read_text())
    namespace = {"TreeNode": TreeNode, "Optional": Optional, "List": List, "Dict": Dict}
    exec(compile(code, str(path), "exec"), namespace)
    return namespace["Solution"]()


def resolve_arg(arg: dict):
    if arg["type"] == "tree":
        return build_tree(arg["value"])
    return arg["value"]


def run_manifest(manifest_path: Path, filter_substr: Optional[str] = None):
    manifest = json.loads(manifest_path.read_text())
    total = passed = 0
    for rel_path, spec in manifest.items():
        if filter_substr and filter_substr.lower() not in rel_path.lower():
            continue
        file_path = REPO_ROOT / rel_path
        print(f"\n{rel_path}  ({spec['method']})")
        try:
            solution = load_solution(file_path)
            method = getattr(solution, spec["method"])
        except Exception as exc:
            print(f"  LOAD ERROR: {exc}")
            total += len(spec["cases"])
            continue
        for i, case in enumerate(spec["cases"], 1):
            total += 1
            args = [resolve_arg(a) for a in case["args"]]
            try:
                actual = method(*args)
                actual_cmp = serialize_tree(actual) if spec.get("output_type") == "tree" else actual
                ok = actual_cmp == case["expected"]
            except Exception as exc:
                ok, actual_cmp = False, f"ERROR: {exc}"
            passed += ok
            status = "PASS" if ok else "FAIL"
            print(f"  [{status}] case {i}: expected={case['expected']!r} actual={actual_cmp!r}")
    print(f"\n{passed}/{total} passed")
    return passed, total


if __name__ == "__main__":
    substr = sys.argv[1] if len(sys.argv) > 1 else None
    n_passed, n_total = run_manifest(MANIFEST_PATH, substr)
    sys.exit(0 if n_passed == n_total else 1)
