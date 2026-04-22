# Raw vs Processed Event Size

Measure inflation ratio of **raw bytes to in-mem bytes** when the history server reads raw event files into the 5
`Cluster*Map`.

| File | Change |
|------|--------|
| `pkg/eventserver/memsize.go` | New. `DeepSize` walker + `MeasureMapSizes` + `LogMapSizes` |
| `pkg/eventserver/eventserver.go` | `processAllEvents` returns `int64` (raw bytes total); calls `LogMapSizes` at end of each call |
| `pkg/eventserver/log_event_reader.go` | `readEventFile` / `ReadLogEvents` return `(int64, error)`; sum per-line `n` for raw byte total |

## How raw bytes are counted

- JSON event files (task/actor/job/node)
  - `io.ReadAll(eventioReader)`: Whole file into a buffer
  - `rawBytes += len(eventbytes)`

- Log event files (events API, JSON Lines)
  - `bufio` line-by-line
  - Sum the per-line `n` returned by `readLineWithLimit`

Sum of both paths = Total bytes pulled from storage in a single `processAllEvents()` call

## How in-memory bytes are computed

Recursively walks the object graph via `reflect`:

| Kind | Rule |
|------|------|
| Pointer | pointer header (8B) + recurse into pointee (visited set guards cycles / shared refs) |
| String | string header (16B) + len(s) raw byte content |
| Slice | slice header (24B) + cap × elemSize (backing array) + deep extras of each element |
| Map | map header + len × (keySize + valSize + 11B bucket overhead) + deep extras of every key/value |
| Struct / Array | type's own size + deep extras of each field/element |
| Scalars (int, bool, …) | just `Type().Size()` |

### Per-map measurement (MeasureMapSizes)

Measure 5 maps with `DeepSize(map)` sequentially, then sum = TotalBytes.

## Logging

Right after each `processAllEvents()` call (once on startup + once per hourly tick):

```bash
time="2026-04-22T08:00:01Z" level=info msg="[memsize] raw=56.03KiB in_mem=27.21KiB ratio=0.49x | task=18.53KiB actor=782B job=1.14KiB node=40B log_event=6.74KiB"
```
