<script lang="ts">
    import Button from "./ui/button/button.svelte";
    import Input from "./ui/input/input.svelte";
    import { Send } from "lucide-svelte";

    // Using Svelte 5 Runes for stats managed by parent, but using props here
    interface Props {
        onSend: (url: string, method: string) => void;
        loading: boolean;
    }

    let { onSend, loading } = $props();

    let url = $state("https://jsonplaceholder.typicode.com/todos/1");
    let method = $state("GET");

    function handleSend() {
        onSend(url, method);
    }
</script>

<div class="flex gap-2 p-4 border-b">
    <select
        bind:value={method}
        class="h-10 w-24 rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring"
    >
        <option value="GET">GET</option>
        <option value="POST">POST</option>
        <option value="PUT">PUT</option>
        <option value="DELETE">DELETE</option>
        <option value="PATCH">PATCH</option>
    </select>

    <Input bind:value={url} placeholder="Enter URL" class="flex-1" />

    <Button on:click={handleSend} disabled={loading}>
        {#if loading}
            Sending...
        {:else}
            <Send class="mr-2 h-4 w-4" /> Send
        {/if}
    </Button>
</div>
