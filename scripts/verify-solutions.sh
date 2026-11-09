#!/usr/bin/env bash
# Verifies that every reference solution (.meta/solution.go) passes its exercise's tests.
# It copies the repo to a temp dir, swaps each exercise.go for its solution, and runs go test.
# Your own work in exercise.go is never touched.
set -euo pipefail

repo_root="$(cd "$(dirname "$0")/.." && pwd)"
work="$(mktemp -d)"
trap 'rm -rf "$work"' EXIT

rsync -a --exclude .git "$repo_root/" "$work/"
cd "$work"

found=0
failed=0
while IFS= read -r solution; do
  dir="$(dirname "$(dirname "$solution")")"
  [[ "$dir" == ./_template* ]] && continue
  found=$((found + 1))
  cp "$solution" "$dir/exercise.go"
  if go test "./$dir/" >/tmp/verify-out.$$ 2>&1; then
    echo "PASS  $dir"
  else
    echo "FAIL  $dir"
    sed 's/^/      /' /tmp/verify-out.$$
    failed=$((failed + 1))
  fi
done < <(find . -path '*/.meta/solution.go' | sort)
rm -f /tmp/verify-out.$$

echo
echo "$((found - failed))/$found solutions passed"
[[ $failed -eq 0 ]]
