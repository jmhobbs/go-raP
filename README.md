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

$ hexyl config.bin
┌────────┬─────────────────────────┬─────────────────────────┬────────┬────────┐
│00000000│ 00 72 61 50 00 00 00 00 ┊ 08 00 00 00 9f 01 00 00 │⋄raP⋄⋄⋄⋄┊•⋄⋄⋄×•⋄⋄│
│00000010│ 00 02 00 43 66 67 50 61 ┊ 74 63 68 65 73 00 33 00 │⋄•⋄CfgPa┊tches⋄3⋄│
│00000020│ 00 00 00 43 66 67 4d 6f ┊ 64 73 00 e9 00 00 00 9f │⋄⋄⋄CfgMo┊ds⋄×⋄⋄⋄×│
│00000030│ 01 00 00 00 01 00 57 49 ┊ 4c 44 4c 41 4e 44 5a 5f │•⋄⋄⋄•⋄WI┊LDLANDZ_│
│00000040│ 47 72 65 65 6e 43 6f 75 ┊ 6e 74 79 5f 4f 4e 55 54 │GreenCou┊nty_ONUT│
│00000050│ 61 6b 65 6f 76 65 72 00 ┊ 60 00 00 00 e9 00 00 00 │akeover⋄┊`⋄⋄⋄×⋄⋄⋄│
│00000060│ 00 04 02 75 6e 69 74 73 ┊ 00 00 02 77 65 61 70 6f │⋄••units┊⋄⋄•weapo│
│00000070│ 6e 73 00 00 01 01 72 65 ┊ 71 75 69 72 65 64 56 65 │ns⋄⋄••re┊quiredVe│
│00000080│ 72 73 69 6f 6e 00 cd cc ┊ cc 3d 02 72 65 71 75 69 │rsion⋄××┊×=•requi│
│00000090│ 72 65 64 41 64 64 6f 6e ┊ 73 00 04 00 44 5a 5f 44 │redAddon┊s⋄•⋄DZ_D│
│000000a0│ 61 74 61 00 00 44 5a 5f ┊ 53 74 72 75 63 74 75 72 │ata⋄⋄DZ_┊Structur│
│000000b0│ 65 73 5f 57 72 65 63 6b ┊ 73 00 00 44 5a 5f 47 65 │es_Wreck┊s⋄⋄DZ_Ge│
│000000c0│ 61 72 5f 43 6f 6e 74 61 ┊ 69 6e 65 72 73 00 00 48 │ar_Conta┊iners⋄⋄H│
│000000d0│ 32 41 5f 47 72 65 65 6e ┊ 43 6f 75 6e 74 79 5f 70 │2A_Green┊County_p│
│000000e0│ 72 6f 70 73 00 e9 00 00 ┊ 00 00 01 00 57 49 4c 44 │rops⋄×⋄⋄┊⋄⋄•⋄WILD│
│000000f0│ 4c 41 4e 44 5a 5f 47 72 ┊ 65 65 6e 43 6f 75 6e 74 │LANDZ_Gr┊eenCount│
│00000100│ 79 5f 4f 4e 55 54 61 6b ┊ 65 6f 76 65 72 00 16 01 │y_ONUTak┊eover⋄••│
│00000110│ 00 00 9f 01 00 00 00 05 ┊ 01 00 74 79 70 65 00 6d │⋄⋄×•⋄⋄⋄•┊•⋄type⋄m│
│00000120│ 6f 64 00 01 00 61 75 74 ┊ 68 6f 72 00 57 49 4c 44 │od⋄•⋄aut┊hor⋄WILD│
│00000130│ 4c 41 4e 44 5a 20 54 65 ┊ 61 6d 00 01 00 6e 61 6d │LANDZ Te┊am⋄•⋄nam│
│00000140│ 65 00 57 49 4c 44 4c 41 ┊ 4e 44 5a 5f 47 72 65 65 │e⋄WILDLA┊NDZ_Gree│
│00000150│ 6e 43 6f 75 6e 74 79 5f ┊ 4f 4e 55 54 61 6b 65 6f │nCounty_┊ONUTakeo│
│00000160│ 76 65 72 00 01 00 64 69 ┊ 72 00 57 49 4c 44 4c 41 │ver⋄•⋄di┊r⋄WILDLA│
│00000170│ 4e 44 5a 5f 47 72 65 65 ┊ 6e 43 6f 75 6e 74 79 5f │NDZ_Gree┊nCounty_│
│00000180│ 4f 4e 55 54 61 6b 65 6f ┊ 76 65 72 00 02 64 65 70 │ONUTakeo┊ver⋄•dep│
│00000190│ 65 6e 64 65 6e 63 69 65 ┊ 73 00 00 9f 01 00 00 00 │endencie┊s⋄⋄×•⋄⋄⋄│
│000001a0│ 00 00 00                ┊                         │⋄⋄⋄     ┊        │
└────────┴─────────────────────────┴─────────────────────────┴────────┴────────┘

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
