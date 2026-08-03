[![Go Reference](https://pkg.go.dev/badge/github.com/jmhobbs/go-raP.svg)](https://pkg.go.dev/github.com/jmhobbs/go-raP)
[![Lint & Test](https://github.com/jmhobbs/go-raP/actions/workflows/lint-and-test.yml/badge.svg)](https://github.com/jmhobbs/go-raP/actions/workflows/lint-and-test.yml)
[![codecov](https://codecov.io/github/jmhobbs/go-raP/graph/badge.svg?token=WU7bOsh2qt)](https://codecov.io/github/jmhobbs/go-raP)

This module is for reading and (eventually) writing raP files for Bohemia Interactive games.  The primary target is DayZ, but given samples and documentation I'm happy to implement for others.

This module is still in development, though it can fully decode DayZ raP-ified config files.  The types and API are likely to change to coalesce into shared types with [go-bicpp](https://github.com/jmhobbs/go-bicpp)

# Usage

```go
// Open a raP file and pass to Decode
// This will return a root node with Entries
root, err := raP.Decode(f)

// If you would just like a printed version,
// there is a default Printer which does an OK job
p := printer.New()
err = p.File(os.Stdout, root)
```

## CLI

Included is a simple CLI which reads raP files and prints them.

```bash
$ rap-decode -h
usage: rap-decode <file.bin>

$ xxd onu-takeover/config.bin | head -n 2
00000000: 0072 6150 0000 0000 0800 0000 9f01 0000  .raP............
00000010: 0002 0043 6667 5061 7463 6865 7300 3300  ...CfgPatches.3.

$ rap-decode config.bin
class CfgPatches
{
  class WILDLANDZ_GreenCounty_ONUTakeover
  {
    units[] = {};
    weapons[] = {};
    requiredVersion = 0.100000;
    requiredAddons[] = {
      "DZ_Data",
      "DZ_Structures_Wrecks",
      "DZ_Gear_Containers",
      "H2A_GreenCounty_props"
    };
  };
};
class CfgMods
{
  class WILDLANDZ_GreenCounty_ONUTakeover
  {
    type = "mod";
    author = "WILDLANDZ Team";
    name = "WILDLANDZ_GreenCounty_ONUTakeover";
    dir = "WILDLANDZ_GreenCounty_ONUTakeover";
    dependencies[] = {};
  };
};
```

# References

- https://community.bistudio.com/wiki/raP_File_Format_-_Elite
- https://community.bistudio.com/wiki/raP_File_Format_-_OFP
