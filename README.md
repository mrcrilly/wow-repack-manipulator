# WoW Repack Manipulator

This tool makes it easier for an administrator of a WoW Repack (private WoW server, basically) to manipulate the database that governs everything in the game, within reason. It does this by removing the need to know SQL or how-to connect to and operate a MySQL database, allowing administrators to easily define what they want to manipulate as a TOML file and then run the tool against their database.

## Warning

This tool currently does **NOT** backup the database before changing it. It will do this in a future version.

Consider backing up the database before proceeding.

## What's Working

Right now you can only manipulate the following repacks and their associated tables:

* EmuCoach Cataclysm v15 - `catav15`
    * `creature_template`

## Configuration

You need to create a TOML file. Any name will do but the tool looks for `./manipulations.toml` in the current working directory as a default.

Here's an example:

```toml
database_hostname = "localhost:3306"
database_username = "root"
database_password = "ascent"
database_name = "emucoach_v15_vip_world"

repack = "catav15"

[[manipulation]]
type = "creature_template"
column = "entry"
id = [40]
modifiers = [
    {key="name", value="Kobold MINER"},
    {key="difficulty_entry_1", value=2},
    {key="difficulty_entry_2", value=3}
]
```

The `database_*` related flags are somewhat obvious: change them to reflect the connection information to your database.

Let's review each key/value pair (`repack`, etc.) and the array of tables (`[[manipulation]]`).

### `repack`

The `repack` key is used tell the tool what repack version you're using.

### `manipulation`

This is where the real work is done.

This key is an array of tables. In JSON notion this would look like this:

```json
[
    {
        "type": "creature_table",
        "column": "entry",
        "id": [
            40,
        ],
        "modifiers": [
            {
                "key": "name",
                "value": "Kobold MINER"
            },
            {
                "key": "difficulty_entry_1",
                "value": 2
            },
            {
                "key": "difficulty_entry_2",
                "value": 3
            }
        ]
    }
]
```

This will manipulate (`modifiers`) the `creature_table` MySQL table (`type`) and edit the following columns (`modifiers.key`) by finding each one based on the unique (`id`):

* `name` would be changed from `Kobold Miner` to `Kobold MINER`
* `difficulty_entry_1` would be changed from `0` to `2`
* `difficulty_entry_2` would be changed from `0` to `3`

The `id` value is a list of IDs for the `entry` column. It has to be a list and it has to have at least one value otherwise nothing is updated.

#### Note

This tool **will not** check if the values you're providing for the tables are correct. It does *zero validation*.

This means if the repack system is expecting, for example, `difficulty_entry_1` to be a value of either `0`, `1` or `2`, and you give it `99`, this tool will not check you've provided the right value. It simply updates the database.

This means this tool is a potential foot cannon. Good luck ;-)

## Executing

(Currently in development so you're just seeing the results of running the code via the Go compiler and not a compiled, released binary.)

There are two flags to the command that you can provide:

* `-config "./manipulations.toml"`
* `-debugging`

This is what it looks like to run the code:

```shell
> go run .
INFO: 2022/02/03 18:22:42 main.go:56: connected to database: root:ascent@tcp(localhost:3306)/emucoach_v15_vip_world
INFO: 2022/02/03 18:22:42 main.go:59: successfully loaded configuration: ./manipulations.toml
INFO: 2022/02/03 18:22:42 parse.go:24: updated created 40: 'name' = 'Kobold MINER'
INFO: 2022/02/03 18:22:42 parse.go:24: updated created 40: 'difficulty_entry_1' = '2'        
INFO: 2022/02/03 18:22:42 parse.go:24: updated created 40: 'difficulty_entry_2' = '3' 
```

That's it. The database has been updated.
