# WhatsApp Contact Confirmation Rule

## Description
Whenever the user mentions a person or a group name on WhatsApp, OR asks you to perform ANY action regarding a contact/group (such as sending a message, getting info, or reading their chats), you MUST ALWAYS confirm the exact contact *before* executing the action. This prevents accidentally retrieving or modifying the wrong chat.

## Workflow
1. When a name or group is mentioned, FIRST use the `whatsapp-mcp` tool `search_contacts` (or relevant group tools) to find matches.
2. Retrieve the matching results, paying close attention to their Short Name (`push_name`), Full Name (`name` or `full_name`), and Phone Number / JID.
3. You MUST present these options to the user to choose from. Use the `ask_question` tool to create a multiple-choice selection if possible, or list them clearly in the chat.
4. The options provided to the user MUST explicitly include:
   - The Contact's Full Name
   - The Contact's Short Name / Push Name
   - The Contact's Phone Number or JID
   *Example format:* `Ashish Pujapanda (Ashish ph) - 916372989845@s.whatsapp.net`
5. ONLY proceed to use the required `whatsapp-mcp` tool (such as `send_message`, `list_messages`, `get_chat`, etc.) AFTER the user has explicitly selected and confirmed the exact contact from the options provided. NEVER skip this confirmation step, even for non-modifying actions.
