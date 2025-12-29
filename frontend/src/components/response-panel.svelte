<script lang="ts">
    import { onMount } from "svelte";
    import { EditorView, basicSetup } from "codemirror";
    import { json } from "@codemirror/lang-json";
    import { EditorState } from "@codemirror/state";
    import { oneDark } from "@codemirror/theme-one-dark";

    let { content } = $props();

    let editorContainer: HTMLElement;
    let editor: EditorView;

    $effect(() => {
        if (editor && content !== undefined) {
            const currentDoc = editor.state.doc.toString();
            if (currentDoc !== content) {
                editor.dispatch({
                    changes: {
                        from: 0,
                        to: currentDoc.length,
                        insert: content,
                    },
                });
            }
        }
    });

    onMount(() => {
        const state = EditorState.create({
            doc: content || "",
            extensions: [
                basicSetup,
                json(),
                oneDark,
                EditorView.lineWrapping,
                EditorView.editable.of(false), // Read-only for response
            ],
        });

        editor = new EditorView({
            state,
            parent: editorContainer,
        });

        return () => {
            editor.destroy();
        };
    });
</script>

<div
    class="h-full w-full overflow-hidden relative"
    bind:this={editorContainer}
></div>

<style>
    /* Ensure editor fills container */
    :global(.cm-editor) {
        height: 100%;
    }
</style>
