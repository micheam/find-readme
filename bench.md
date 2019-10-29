| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `./find-readme ~/ghq` | 746.0 ± 13.4 | 730.0 | 768.2 | 1.52 ± 0.15 |
| `find ~/ghq -iname "readme.md" -not -path "*/node_modules/*"` | 1613.6 ± 44.9 | 1570.5 | 1706.3 | 3.28 ± 0.34 |
| `ag -l "" ~/ghq \| ag -i "readme.md"` | 492.0 ± 49.3 | 444.2 | 589.1 | 1.00 |
