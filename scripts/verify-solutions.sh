#!/usr/bin/env bash
# Verifies that every reference solution (.meta/solution.go) passes its exercise's tests.
# It copies the repo to a temp dir, swaps EVERY exercise.go for its solution first
# (some exercises import others, e.g. concepts/packages), then runs go test per exercise.
# Your own work in exercise.go is never touched.
set -euo pipefail

repo_root="$(cd "$(dirname "$0")/.." && pwd)"
work="$(mktemp -d)"
trap 'rm -rf "$work"' EXIT

rsync -a --exclude .git "$repo_root/" "$work/"
cd "$work"

dirs=()
while IFS= read -r solution; do
  dir="$(dirname "$(dirname "$solution")")"
  [[ "$dir" == ./_template* ]] && continue
  cp "$solution" "$dir/exercise.go"
  dirs+=("$dir")
done < <(find . -path '*/.meta/solution.go' | sort)

failed=0
for dir in "${dirs[@]}"; do
  if out="$(go vet "./$dir/" 2>&1 && go test "./$dir/" 2>&1)"; then
    echo "PASS  $dir"
  else
    echo "FAIL  $dir"
    printf '%s\n' "$out" | sed 's/^/      /'
    failed=$((failed + 1))
  fi
done

echo
echo "$(( ${#dirs[@]} - failed ))/${#dirs[@]} solutions passed"
[[ $failed -eq 0 ]]
