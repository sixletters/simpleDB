# Key-Value Store Client

## How to Run

1. Make sure you have Python installed on your system.
2. Open a terminal or command prompt.
3. Navigate to the directory containing the `client.py` file.
4. Run the client script using the following command:
   python client.py

## Available Prompts

The client script accepts the following prompts:

scbctl get <key>: Retrieves the value associated with the specified key.

scbctl insert <key> <val>: Inserts a new key-value pair into the Key-Value Store.

scbctl update <key> <val>: Updates the value of an existing key in the Key-Value Store.

scbctl delete <key>: Deletes a key-value pair from the Key-Value Store.

Note: Both <key> and <val> must be single strings with underscores separating any two words.
Example: scbctl insert chicken_shit 42
