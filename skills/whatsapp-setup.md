# WhatsApp MCP Setup Skill

This skill provides a simple, foolproof guide for setting up the WhatsApp MCP bridge. Because the original repository (`ihiteshgupta/whatsapp-mcp`) caused severe "database locked" errors and lacked native contact name resolution, we have **forked** the repository and applied all fixes to our own version. 

**Follow this guide strictly** whenever connecting a new WhatsApp account or setting up the MCP server from scratch.

## 1. Clone the Patched Repository
Instead of downloading the original, you must use our patched fork which contains all the zombie-process fixes and name-resolution logic built-in.

```bash
cd <desired_workspace_dir>
git clone https://github.com/Satya1929/whatsapp-mcp
cd whatsapp-mcp/whatsapp-bridge-v2
```

## 2. Compilation and Build
The MCP server relies on a Go wrapper that spawns the real server. You must build it:
```bash
go build -o whatsapp-mcp-real.exe ./cmd/whatsapp-mcp
```

## 3. QR Code Pairing (Strict 20-Second Window)
**CRITICAL**: WhatsApp Web QR codes expire *very* quickly. If you miss the window, the pairing will hang or loop.
- **Instruct the user to have their phone ready (WhatsApp > Linked Devices) BEFORE triggering the connection.**
- Tell the user: "You will have strictly **20 seconds** to scan the QR code once it is generated."
- Generate the QR code by starting the server (either via IDE tool integration or manually). The code will be saved as `qrcode.png` in the `store/` directory. Give the user the absolute path to this file.

## 4. The Syncing Phase (Do NOT Restart!)
**CRITICAL**: After the QR code is scanned, the initial sync begins. It downloads gigabytes of messages and can take several minutes.
- **Do NOT restart the IDE or the MCP server during the initial sync.** 
- The user's phone might show "Last active today at..." — this is normal. Do not panic.
- Just wait patiently. You can verify progress by checking the file size of `store/messages.db`. Once the bridge logs report the state as `ready`, the sync is complete.

## Why We Use This Fork
For your context, if you ever need to debug, here is what is patched in our `Satya1929` fork:
- **`main_wrapper.go`**: Forces child processes to terminate when the IDE disconnects, preventing the infamous `database is locked` error and `tools/list: invalid request`.
- **`sqlite.go`**: Sets SQLite connection pooling to 1 (`db.SetMaxOpenConns(1)`) so cross-database queries between `messages.db` and `whatsapp.db` succeed.
- **Contact Name Resolution**: Natively translates anonymous WhatsApp privacy IDs (`@lid`) to `Phone Number (Push Name)` or `Full Name` directly in the database queries.
