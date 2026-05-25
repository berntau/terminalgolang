<script setup>
import { nextTick, ref } from 'vue'

const command = ref('')
const isLoading = ref(false)
const commandInput = ref(null)
const terminalBody = ref(null)
const history = ref([
  {
    type: 'system',
    output: "boot sequence complete\ncurriculo-terminal v0.1\n\ndigite 'help' para ver os comandos disponíveis"
  }
])

async function submitCommand() {
  const typedCommand = command.value.trim()

  if (!typedCommand || isLoading.value) {
    return
  }

  history.value.push({
    type: 'command',
    input: typedCommand
  })
  command.value = ''
  isLoading.value = true

  try {
    const response = await fetch(`api/command`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ command: typedCommand })
    })

    const payload = await response.json()

    if (!response.ok) {
      throw new Error(payload.error || 'erro ao executar comando')
    }

    history.value.push({
      type: 'output',
      output: payload.output
    })
  } catch (error) {
    history.value.push({
      type: 'error',
      output: error.message || 'não foi possível conectar com a API'
    })
  } finally {
    isLoading.value = false
    await nextTick()
    commandInput.value?.focus()
    terminalBody.value?.scrollTo({
      top: terminalBody.value.scrollHeight,
      behavior: 'smooth'
    })
  }
}

function focusCommandInput() {
  commandInput.value?.focus()
}
</script>

<template>
  <section class="terminal-window" aria-label="Terminal interativo" @click="focusCommandInput">
    <header class="terminal-header">
      <div class="window-controls" aria-hidden="true">
        <span></span>
        <span></span>
        <span></span>
      </div>
      <span class="terminal-title">RETROTERM-8080 / bernardo@curriculo:~</span>
      <span class="terminal-status">online</span>
    </header>

    <div ref="terminalBody" class="terminal-body">
      <div
        v-for="(entry, index) in history"
        :key="index"
        class="terminal-entry"
        :class="`terminal-entry--${entry.type}`"
      >
        <template v-if="entry.type === 'command'">
          <span class="prompt">guest@retroterm:~$</span>
          <span>{{ entry.input }}</span>
        </template>

        <pre v-else>{{ entry.output }}</pre>
      </div>

      <form class="terminal-prompt" @submit.prevent="submitCommand">
        <label for="terminal-command" class="prompt">guest@retroterm:~$</label>
        <input
          id="terminal-command"
          ref="commandInput"
          v-model="command"
          type="text"
          autocomplete="off"
          spellcheck="false"
          :disabled="isLoading"
          autofocus
          aria-label="Digite um comando"
        />
        <span v-if="isLoading" class="cursor">running</span>
      </form>
    </div>
  </section>
</template>
