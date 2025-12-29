<script lang="ts">
  import RequestPanel from "./components/request-panel.svelte";
  import ResponsePanel from "./components/response-panel.svelte";
  import { SendRequest } from "../bindings/changeme/services/requestservice";

  let loading = $state(false);
  let responseBody = $state("");
  let status = $state("");
  let time = $state(0);

  async function handleSend(url: string, method: string) {
    loading = true;
    status = "";
    try {
      const res = await SendRequest({
        method: method,
        url: url,
        headers: {},
        body: "",
      });

      // Attempt format JSON
      try {
        responseBody = JSON.stringify(JSON.parse(res.body), null, 2);
      } catch {
        responseBody = res.body;
      }

      status = `${res.statusCode} ${res.status}`;
      time = res.time;
    } catch (err) {
      console.error(err);
      responseBody = "Error: " + String(err);
    } finally {
      loading = false;
    }
  }
</script>

<div class="flex flex-col h-screen bg-background text-foreground dark">
  <div class="border-b p-2 flex items-center justify-between">
    <h1 class="font-bold text-lg px-2">Project Pebble</h1>
    <div class="text-xs text-muted-foreground px-2">
      {#if status}
        <span class={status.startsWith("2") ? "text-green-500" : "text-red-500"}
          >{status}</span
        >
        <span class="mx-2">•</span>
        <span>{time}ms</span>
      {/if}
    </div>
  </div>

  <RequestPanel onSend={handleSend} {loading} />

  <div class="flex-1 overflow-hidden min-h-0 relative">
    {#if responseBody}
      <ResponsePanel content={responseBody} />
    {:else}
      <div
        class="flex items-center justify-center h-full text-muted-foreground"
      >
        No response yet
      </div>
    {/if}
  </div>
</div>
